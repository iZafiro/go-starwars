package main

import (
	"context"
	"fmt"
	"go-starwars/api/brokerpb"
	"log"
	"strings"

	"google.golang.org/grpc"
)

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

	// Connect to fulcrum3 server
	fmt.Println("Starting Client...")
	cc, err := grpc.Dial("0.0.0.0:50050", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()
	cb1 = brokerpb.NewBrokerServiceClient(cc)

	consistency = make(map[string]*registry)

	//leia loop
	for {
		fmt.Println("Ingrese comando")

		var input string
		fmt.Scanln(&input)

		if input == "end" {
			break
		} else {
			command := strings.Split(string(input), ",")
			if command[0] == "GetNumberRebels" {
				registryName := command[1] + " " + command[2]
				//checks if its already registered
				if _, ok := consistency[registryName]; !ok {
					reg := &registry{command[1], command[2], 0, []int32{}, 1}
					consistency[registryName] = reg
				}

				succ, num, vec := getNum(command[1], command[2], consistency[command[2]].Vector, cb1) /////////////////////
				if succ {
					fmt.Println("En la ciudad " + command[2] + "del planeta " + command[1] + "hay " + string(num))
					consistency[registryName].Vector = vec
				} else {
					fmt.Println("La operación no se pudo realizar")
				}

			} else {
				fmt.Println("Ingrese un commando válido")
			}
		}
	}
}

func getNum(planet string, city string, vec []int32, cb1 brokerpb.BrokerServiceClient) (bool, int32, []int32) {

	//pack request
	req := &brokerpb.GetNumberRebelsRequest{
		Planet: planet,
		City:   city,
		Vector: vec,
	}

	//send request
	res, err := cb1.GetNumberRebels(context.Background(), req)
	if err != nil {
		log.Fatalf("Error Call RPC %v", err)
	}
	return res.Success, res.Number, res.Vector
}
