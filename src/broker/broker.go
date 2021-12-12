package main

import (
	"context"
	"fmt"
	"go-starwars/api/brokerpb"
	"go-starwars/api/fulcrumpb"
	"log"
	"math/rand"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct{}

var cf1 fulcrumpb.FulcrumServiceClient
var cf2 fulcrumpb.FulcrumServiceClient
var cf3 fulcrumpb.FulcrumServiceClient
var s *grpc.Server

func main() {
	// Connect to fulcrum1 server
	fmt.Println("Starting Client...")
	cc, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()
	cf1 = fulcrumpb.NewFulcrumServiceClient(cc)

	// Connect to fulcrum2 server
	fmt.Println("Starting Client...")
	cc, err = grpc.Dial("0.0.0.0:50052", grpc.WithInsecure())
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

	// Start server
	fmt.Println("Starting server...")
	l, err := net.Listen("tcp", "0.0.0.0:50050")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	s = grpc.NewServer()
	brokerpb.RegisterBrokerServiceServer(s, &server{})
	if err := s.Serve(l); err != nil {
		log.Fatalf("Failed to server %v", err)
	}
}

// Si el primer vector es menor o igual al segundo vector
// coordenada a coordenada, entonces retorna true,
// si no, returna false.
func vectorLeq(first []int32, second []int32) bool {
	n := len(first)

	for i := 0; i < n; i++ {
		if first[i] > second[i] {
			return false
		}
	}

	return true
}

func (*server) GetFulcrum(ctx context.Context, req *brokerpb.GetFulcrumRequest) (*brokerpb.GetFulcrumResponse, error) {
	// Unpack request
	planet := req.GetPlanet()
	vector := req.GetVector()

	fmt.Println(planet, vector)

	var fulcrumId int

	rand.Seed(time.Now().UnixNano())

	validFulcrum := []int{}

	success := true

	// Si el informante ha consultado el registro antes, entonces
	// busca los fulcrum con vectores más recientes.
	if len(vector) > 0 {
		// Pack request
		reqf := &fulcrumpb.GetVectorRequest{
			Planet: planet,
		}

		// Send request
		res, err := cf1.GetVector(context.Background(), reqf)
		if err != nil {
			log.Fatalf("Error Call RPC: %v", err)
		}

		if vectorLeq(vector, res.Vector) {
			validFulcrum = append(validFulcrum, 1)
		}

		// Send request
		res, err = cf2.GetVector(context.Background(), reqf)
		if err != nil {
			log.Fatalf("Error Call RPC: %v", err)
		}

		if vectorLeq(vector, res.Vector) {
			validFulcrum = append(validFulcrum, 2)
		}

		// Send request
		res, err = cf3.GetVector(context.Background(), reqf)
		if err != nil {
			log.Fatalf("Error Call RPC: %v", err)
		}

		if vectorLeq(vector, res.Vector) {
			validFulcrum = append(validFulcrum, 3)
		}

		if len(validFulcrum) <= 0 {
			success = false
		}
	}

	// Si encontró un fulcrum con un vector más reciente, entonces
	// elige uno de estos al azar.
	// Si no, (es decir, si el informante no había consultado el registro antes,
	// o si no encontró un fulcrum con un vector más reciente), entonces
	// elige un fulcrum al azar.
	if len(validFulcrum) > 0 {
		fulcrumId = validFulcrum[rand.Intn(len(validFulcrum))]
	} else {
		fulcrumId = rand.Intn(3) + 1
	}

	// Send response
	res := &brokerpb.GetFulcrumResponse{
		Success:   success,
		FulcrumId: int32(fulcrumId),
	}
	return res, nil
}

func (*server) GetNumberRebels(ctx context.Context, req *brokerpb.GetNumberRebelsRequest) (*brokerpb.GetNumberRebelsResponse, error) {
	// Unpack request
	planet := req.GetPlanet()
	city := req.GetCity()
	vector := req.GetVector()

	fmt.Println(planet, city, vector)

	var fulcrumId int
	var number int32

	rand.Seed(time.Now().UnixNano())

	validFulcrum := []int{}

	success := true

	// Si Leia ha consultado el registro antes, entonces
	// busca los fulcrum con vectores más recientes.
	if len(vector) > 0 {
		// Pack request
		reqf := &fulcrumpb.GetVectorRequest{
			Planet: planet,
		}

		// Send request
		res, err := cf1.GetVector(context.Background(), reqf)
		if err != nil {
			log.Fatalf("Error Call RPC: %v", err)
		}

		if vectorLeq(vector, res.Vector) {
			validFulcrum = append(validFulcrum, 1)
		}

		// Send request
		res, err = cf2.GetVector(context.Background(), reqf)
		if err != nil {
			log.Fatalf("Error Call RPC: %v", err)
		}

		if vectorLeq(vector, res.Vector) {
			validFulcrum = append(validFulcrum, 2)
		}

		// Send request
		res, err = cf3.GetVector(context.Background(), reqf)
		if err != nil {
			log.Fatalf("Error Call RPC: %v", err)
		}

		if vectorLeq(vector, res.Vector) {
			validFulcrum = append(validFulcrum, 3)
		}

		if len(validFulcrum) <= 0 {
			success = false
		}
	}

	// Si encontró un fulcrum con un vector más reciente, entonces
	// elige uno de estos al azar.
	// Si no, (es decir, si Leia no había consultado el registro antes,
	// o si no encontró un fulcrum con un vector más reciente), entonces
	// elige un fulcrum al azar.
	if len(validFulcrum) > 0 {
		fulcrumId = validFulcrum[rand.Intn(len(validFulcrum))]
	} else {
		fulcrumId = rand.Intn(3) + 1
	}

	// Consulta el número de rebeldes en el fulcrum elegido.
	switch fulcrumId {
	case 1:
		// Pack request
		reqf := &fulcrumpb.GetNumberRebelsFulcrumRequest{
			Planet: planet,
			City:   city,
		}

		// Send request
		res, err := cf1.GetNumberRebelsFulcrum(context.Background(), reqf)
		if err != nil {
			log.Fatalf("Error Call RPC: %v", err)
		}

		if res.Success {
			number = res.Number
			vector = res.Vector
		} else {
			success = false
			number = res.Number
			vector = res.Vector
		}
	case 2:
		// Pack request
		reqf := &fulcrumpb.GetNumberRebelsFulcrumRequest{
			Planet: planet,
			City:   city,
		}

		// Send request
		res, err := cf2.GetNumberRebelsFulcrum(context.Background(), reqf)
		if err != nil {
			log.Fatalf("Error Call RPC: %v", err)
		}

		if res.Success {
			number = res.Number
			vector = res.Vector
		} else {
			success = false
			number = res.Number
			vector = res.Vector
		}
	case 3:
		// Pack request
		reqf := &fulcrumpb.GetNumberRebelsFulcrumRequest{
			Planet: planet,
			City:   city,
		}

		// Send request
		res, err := cf3.GetNumberRebelsFulcrum(context.Background(), reqf)
		if err != nil {
			log.Fatalf("Error Call RPC: %v", err)
		}

		if res.Success {
			number = res.Number
			vector = res.Vector
		} else {
			success = false
			number = res.Number
			vector = res.Vector
		}
	default:
		success = false
		number = -1
		vector = []int32{}
	}

	// Send response
	res := &brokerpb.GetNumberRebelsResponse{
		Success: success,
		Number:  int32(number),
		Vector:  vector,
	}
	return res, nil
}
