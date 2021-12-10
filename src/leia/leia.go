package main

import (
	"context"
	"fmt"
	"strings"
	"strconv"
	"go-starwars/api/fulcrumpb"
	"go-starwars/api/brokerpb"
	"go-starwars/src/concerns"
	"log"
	"net"

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

var consistency map[string]registry

func main() {

	// Connect to fulcrum3 server
	fmt.Println("Starting Client...")
	cc, err := grpc.Dial("0.0.0.0:50050", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()
	cb1 = brokerpb.NewBrokerServiceClient(cc)

	//leia loop
	for {
		fmt.Println("Ingrese comando")

		var input string;
		fmt.Scanln(&input)

		if input == "end"{
			break
		} else{
			command := strings.Split(input, " ")
			if command[0] == "GetNumberRebels" {
				resistryName := command[1] + " " + command[2]
				consistency[registryName].Planet = command[1]
				succ, num, vec, fulc, := getNumberRebels(command[1], command[2], consistency[command[2].Vector], br1) /////////////////////
				if succ {
					consistency[registryName].City = command[2]
					consistency[registryName].Rebels = num
					consistency[registryName].Fulcrum = fulc
					fmt.Println("En la ciudad %s del planeta %s hay %d\n\n", command[2], command[1], num)
				} else {
					fmt.Println("La operación no se pudo realizar\n\n")
				}

			} else {
				fmt.Println("Ingrese un commando válido")
			}
		}
	}
}

func getNumberRebels(string planet, string city, []int32 vector, br1 brokerpb.BrokerServiceClient) int32 {
	
	//pack request
	req := &brokerpb.GetNumberRebelsRequest{
		Planet: planet,
		City: city,
		Vector: vector 
	}

	//send request
	res, err := br1.GetNumberRebels(context.Background(), req)
	if err != nil{
		log.Fatalf("Error Call RPC %v", err)
	}
	return res.Number
}

