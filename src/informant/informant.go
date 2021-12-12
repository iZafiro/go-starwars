package main

import (
	"context"
	"fmt"
	"go-starwars/api/brokerpb"
	"go-starwars/api/fulcrumpb"
	"log"
	"strconv"
	"strings"

	"google.golang.org/grpc"
)

var cf1 fulcrumpb.FulcrumServiceClient
var cf2 fulcrumpb.FulcrumServiceClient
var cf3 fulcrumpb.FulcrumServiceClient
var cb1 brokerpb.BrokerServiceClient

type registry struct {
	Planet  string
	City    string
	Rebels  int
	Vector  []int32
	Fulcrum int
}

var consistency map[string]*registry

func main() {

	// Connect to broker server
	fmt.Println("Starting Client...")
	cc, err := grpc.Dial("0.0.0.0:50050", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()
	cb1 = brokerpb.NewBrokerServiceClient(cc)

	// Connect to fulcrum1 server
	fmt.Println("Starting Client...")
	cc, err = grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
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

	consistency = make(map[string]*registry)

	//informant loop
	for {
		fmt.Println("Ingrese comando")

		var input string
		fmt.Scanln(&input)

		if input == "end" {
			break
		} else {
			//gets fulcrum Id
			command := strings.Split(string(input), ",")
			registryName := command[1] + " " + command[2]

			//checks if it's already registred
			if _, ok := consistency[registryName]; !ok {
				reg := &registry{command[1], command[2], 0, []int32{}, 1}
				consistency[registryName] = reg
			}
			succ, fId := getFul(command[1], consistency[registryName].Vector, cb1)
			
			if succ {
				//proceeds with command on the correct fulcrumm

				var cf fulcrumpb.FulcrumServiceClient
				switch fId {
				case 1:
					cf = cf1
				case 2:
					cf = cf2
				case 3:
					cf = cf3
				}
				var vec []int32
				num := 0
				switch command[0] {
				case "AddCity":
					if len(command) >= 4 {
						num, _ = strconv.Atoi(command[3])
					}
					fmt.Println("asd")
					succ, vec = addCity(command[1], command[2], int32(num), cf)

				case "UpdateName":
					succ, vec = updateName(command[1], command[2], command[3], cf)

				case "UpdateNumber":
					num, _ = strconv.Atoi(command[3])
					succ, vec = updateNumber(command[1], command[2], int32(num), cf)

				case "DeleteCity":
					succ, vec = deleteCity(command[1], command[2], cf)
				}

				if succ {
					consistency[registryName].Vector = vec
					fmt.Println("La operación fue exitosa")
				} else {
					fmt.Println("La operación no se pudo realizar")
				}

			} else {
				fmt.Println("La operación no se pudo realizar")
			}
		}
	}

}

func getFul(planet string, vec []int32, cb1 brokerpb.BrokerServiceClient) (bool, int32) {

	//pack request
	req := &brokerpb.GetFulcrumRequest{
		Planet: planet,
		Vector: vec,
	}

	//send request
	res, err := cb1.GetFulcrum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error Call RPC %v", err)
	}
	return res.Success, res.FulcrumId
}

func addCity(planet string, city string, num int32, cf fulcrumpb.FulcrumServiceClient) (bool, []int32) {

	//pack request
	req := &fulcrumpb.AddCityRequest{
		Planet: planet,
		City:   city,
		Number: num,
	}

	//send request
	res, err := cf.AddCity(context.Background(), req)
	if err != nil {
		log.Fatalf("Error Call RPC %v", err)
	}
	return res.Success, res.Vector
}

func updateName(planet string, city string, newName string, cf fulcrumpb.FulcrumServiceClient) (bool, []int32) {

	//pack request
	req := &fulcrumpb.UpdateNameRequest{
		Planet:  planet,
		OldCity: city,
		NewCity: newName,
	}

	//send request
	res, err := cf.UpdateName(context.Background(), req)
	if err != nil {
		log.Fatalf("Error Call RPC %v", err)
	}
	return res.Success, res.Vector
}

func updateNumber(planet string, city string, num int32, cf fulcrumpb.FulcrumServiceClient) (bool, []int32) {

	//pack request
	req := &fulcrumpb.UpdateNumberRequest{
		Planet: planet,
		City:   city,
		Number: num,
	}

	//send request
	res, err := cf.UpdateNumber(context.Background(), req)
	if err != nil {
		log.Fatalf("Error Call RPC %v", err)
	}
	return res.Success, res.Vector
}

func deleteCity(planet string, city string, cf fulcrumpb.FulcrumServiceClient) (bool, []int32) {

	//pack request
	req := &fulcrumpb.DeleteCityRequest{
		Planet: planet,
		City:   city,
	}

	//send request
	res, err := cf.DeleteCity(context.Background(), req)
	if err != nil {
		log.Fatalf("Error Call RPC %v", err)
	}
	return res.Success, res.Vector
}
