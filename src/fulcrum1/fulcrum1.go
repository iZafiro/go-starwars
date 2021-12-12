package main

import (
	"context"
	"fmt"
	"go-starwars/api/fulcrumpb"
	"go-starwars/src/concerns"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
)

type server struct{}

var planetVectors map[string][3]int32
var folder string
var node int32

var cf2 fulcrumpb.FulcrumServiceClient
var cf3 fulcrumpb.FulcrumServiceClient
var s *grpc.Server

func main() {
	planetVectors = make(map[string][3]int32)
	folder = "src/fulcrum1/out/"
	node = int32(1)
	concerns.CRemoveContents("src/fulcrum1/out")
	/*
		Local test
		concerns.CRemoveContents("out")
		_, planetVectors = concerns.CAddCity("planeta", "uwu", int32(100), planetVectors, folder, node)
		_, planetVectors = concerns.CUpdateName("planeta", "uwu", "awa", planetVectors, folder, node)
		_, planetVectors = concerns.CUpdateNumber("planeta", "awa", int32(50), planetVectors, folder, node)
		_, planetVectors = concerns.CAddCity("planeta", "owo", int32(100), planetVectors, folder, node)
		_, planetVectors = concerns.CAddCity("planeta", "iwi", int32(100), planetVectors, folder, node)
		_, planetVectors = concerns.CAddCity("planeta2", "xd", int32(100), planetVectors, folder, node)
		_, planetVectors = concerns.CDeleteCity("planeta", "awa", planetVectors, folder, node)
		planetVectors = concerns.CMerge([]string{"planeta\n1, 2, 3\nlinea1\nlinea2\n..."}, planetVectors, folder)
		log.Println(planetVectors)
	*/
	// Connect to fulcrum2 server
	fmt.Println("Starting Client...")
	cc, err := grpc.Dial("0.0.0.0:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()
	cf2 = fulcrumpb.NewFulcrumServiceClient(cc)

	// Connect to fulcrum3 server
	fmt.Println("Starting Client...")
	cc, err = grpc.Dial("0.0.0.0:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()
	cf3 = fulcrumpb.NewFulcrumServiceClient(cc)

	// Se crea una subrutina para hacer un merge cronjob cada dos minutos
	go mergeCronjobs()

	// Start server
	fmt.Println("Starting server...")
	l, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	s = grpc.NewServer()
	fulcrumpb.RegisterFulcrumServiceServer(s, &server{})
	if err := s.Serve(l); err != nil {
		log.Fatalf("Failed to server %v", err)
	}
}

func mergeCronjobs() {
	ticker := time.NewTicker(120 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				go merge()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func merge() {
	fmt.Println("Started merge.")

	// Se define un arreglo que guardará los planetas que han sido cambiados
	// en algún fulcrum desde el último merge
	touchedPlanets := []string{}

	// Se borran los logs de registro actuales
	os.Remove(folder + "Logs.txt")

	// Se piden los logs de registro de los otros dos fulcrum
	// Pack request
	req := &fulcrumpb.GetLogsRequest{
		Value: 1,
	}

	// Send request
	res2, err := cf2.GetLogs(context.Background(), req)
	if err != nil {
		log.Fatalf("Error Call RPC: %v", err)
	}

	// Send request
	res3, err := cf3.GetLogs(context.Background(), req)
	if err != nil {
		log.Fatalf("Error Call RPC: %v", err)
	}

	// Se unen los logs de registro de los otros dos fulcrum
	logs := append(res2.Logs, res3.Logs...)

	// Se itera en los logs de registro
	for _, log := range logs {
		// Se separa cada log del fulcrum X del reloj de vector más reciente
		// del fulcrum X del planeta asociado al log
		// En efecto, un ejemplo de log es
		// AddCity Tattoine Mos_Eisley 1\n1, 1, 2
		logArray := strings.Split(log, "\n")

		// Si el log es válido, se continúa
		if len(logArray) > 1 {
			// Se obtiene la operación
			opArray := strings.Split(logArray[0], " ")

			// Se obtienen los parámetros de la operación
			planet := opArray[1]
			city := opArray[2]
			var valueString string
			if len(opArray) > 3 {
				valueString = opArray[3]
			} else {
				valueString = "0"
			}

			// Se ejecuta la operación
			var valueInt int
			switch opArray[0] {
			case "AddCity":
				valueInt, err = strconv.Atoi(valueString)
				_, _ = concerns.CAddCity(planet, city, int32(valueInt), planetVectors, folder, node, false)
			case "UpdateName":
				_, _ = concerns.CUpdateName(planet, city, valueString, planetVectors, folder, node, false)
			case "UpdateNumber":
				valueInt, err = strconv.Atoi(valueString)
				_, _ = concerns.CUpdateNumber(planet, city, int32(valueInt), planetVectors, folder, node, false)
			case "DeleteCity":
				_, _ = concerns.CDeleteCity(planet, city, planetVectors, folder, node, false)
			}

			// Se obtiene el reloj de vector
			vectorStringArray := strings.Split(logArray[1], ", ")
			vector := [3]int32{}
			for index, str := range vectorStringArray {
				var vectorIndex int
				vectorIndex, err = strconv.Atoi(str)
				vector[index] = int32(vectorIndex)
			}

			// Se actualiza el reloj de vector actual
			planetVectors[planet] = maxEntries(concerns.CGetVector(planet, planetVectors), vector)

			// Se añade el planeta al arreglo de planetas que han sido cambiados
			touchedPlanets = append(touchedPlanets, planet)
		}
	}

	// Se propagan los cambios enviando una string por cada
	// planeta / vector / archivo cambiado
	files := []string{}

	for _, planet := range touchedPlanets {
		vectorString := ""

		for index, entry := range planetVectors[planet] {
			vectorString += string(entry)
			if index < 2 {
				vectorString += ", "
			}
		}

		input, err := ioutil.ReadFile(folder + "Registro_" + planet + ".txt")
		if err != nil {
			log.Fatal(err)
		}

		files = append(files, planet+"\n"+vectorString+"\n"+string(input))
	}

	// Pack request
	reqm := &fulcrumpb.MergeRequest{
		Files: files,
	}

	// Send request
	resm2, err := cf2.Merge(context.Background(), reqm)
	if err != nil {
		log.Fatalf("Error Call RPC: %v", err)
	}

	// Send request
	resm3, err := cf3.Merge(context.Background(), reqm)
	if err != nil {
		log.Fatalf("Error Call RPC: %v", err)
	}

	if resm2.Success && resm3.Success {
		fmt.Println("Merge propagated successfully.")
	} else {
		fmt.Println("Failed to propagate merge.")
	}
}

func maxEntries(vec1 [3]int32, vec2 [3]int32) [3]int32 {
	res := [3]int32{}

	for index, entry := range vec1 {
		if entry >= vec2[index] {
			res[index] = entry
		} else {
			res[index] = vec2[index]
		}
	}

	return res
}

func (*server) GetVector(ctx context.Context, req *fulcrumpb.GetVectorRequest) (*fulcrumpb.GetVectorResponse, error) {
	// Unpack request
	planet := req.GetPlanet()

	// Pack response
	vector := concerns.CGetVector(planet, planetVectors)

	// Send response
	res := &fulcrumpb.GetVectorResponse{
		Vector: vector[:],
	}
	return res, nil
}

func (*server) GetNumberRebelsFulcrum(ctx context.Context, req *fulcrumpb.GetNumberRebelsFulcrumRequest) (*fulcrumpb.GetNumberRebelsFulcrumResponse, error) {
	// Unpack request
	planet := req.GetPlanet()
	city := req.GetCity()

	// Pack response
	success, number := concerns.CGetRebels(planet, city, folder)

	vector := concerns.CGetVector(planet, planetVectors)

	// Send response
	res := &fulcrumpb.GetNumberRebelsFulcrumResponse{
		Success: success,
		Number:  int32(number),
		Vector:  vector[:],
	}
	return res, nil
}

func (*server) GetLogs(ctx context.Context, req *fulcrumpb.GetLogsRequest) (*fulcrumpb.GetLogsResponse, error) {

	// Pack response
	logs := concerns.CGetLogs(planetVectors, folder)

	// Send response
	res := &fulcrumpb.GetLogsResponse{
		Logs: logs,
	}
	return res, nil
}

func (*server) Merge(ctx context.Context, req *fulcrumpb.MergeRequest) (*fulcrumpb.MergeResponse, error) {
	// Unpack request
	files := req.GetFiles()

	// Pack response
	var success bool
	success, planetVectors = concerns.CMerge(files, planetVectors, folder)

	// Send response
	res := &fulcrumpb.MergeResponse{
		Success: success,
	}
	return res, nil
}

func (*server) AddCity(ctx context.Context, req *fulcrumpb.AddCityRequest) (*fulcrumpb.AddCityResponse, error) {
	// Unpack request
	planet := req.GetPlanet()
	city := req.GetCity()
	number := req.GetNumber()

	fmt.Println(planet, city, number)

	// Pack response
	var success bool
	success, planetVectors = concerns.CAddCity(planet, city, number, planetVectors, folder, node, true)
	vector := concerns.CGetVector(planet, planetVectors)

	// Send response
	res := &fulcrumpb.AddCityResponse{
		Success: success,
		Vector:  vector[:],
	}
	return res, nil
}

func (*server) UpdateName(ctx context.Context, req *fulcrumpb.UpdateNameRequest) (*fulcrumpb.UpdateNameResponse, error) {
	// Unpack request
	planet := req.GetPlanet()
	oldCity := req.GetOldCity()
	newCity := req.GetNewCity()

	// Pack response
	var success bool
	success, planetVectors = concerns.CUpdateName(planet, oldCity, newCity, planetVectors, folder, node, true)
	vector := concerns.CGetVector(planet, planetVectors)

	// Send response
	res := &fulcrumpb.UpdateNameResponse{
		Success: success,
		Vector:  vector[:],
	}
	return res, nil
}

func (*server) UpdateNumber(ctx context.Context, req *fulcrumpb.UpdateNumberRequest) (*fulcrumpb.UpdateNumberResponse, error) {
	// Unpack request
	planet := req.GetPlanet()
	city := req.GetCity()
	number := req.GetNumber()

	// Pack response
	var success bool
	success, planetVectors = concerns.CUpdateNumber(planet, city, number, planetVectors, folder, node, true)
	vector := concerns.CGetVector(planet, planetVectors)

	// Send response
	res := &fulcrumpb.UpdateNumberResponse{
		Success: success,
		Vector:  vector[:],
	}
	return res, nil
}

func (*server) DeleteCity(ctx context.Context, req *fulcrumpb.DeleteCityRequest) (*fulcrumpb.DeleteCityResponse, error) {
	// Unpack request
	planet := req.GetPlanet()
	city := req.GetCity()

	println(planet, city)

	// Pack response
	var success bool
	success, planetVectors = concerns.CDeleteCity(planet, city, planetVectors, folder, node, true)
	vector := concerns.CGetVector(planet, planetVectors)

	// Send response
	res := &fulcrumpb.DeleteCityResponse{
		Success: success,
		Vector:  vector[:],
	}
	return res, nil
}
