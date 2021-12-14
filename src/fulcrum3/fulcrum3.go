package main

import (
	"context"
	"fmt"
	"go-starwars/api/fulcrumpb"
	"go-starwars/src/concerns"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

var planetVectors map[string][3]int32
var folder string
var node int32

var s *grpc.Server

func main() {
	planetVectors = make(map[string][3]int32)
	folder = "src/fulcrum3/out/"
	node = int32(3)
	concerns.CRemoveContents("src/fulcrum3/out")
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

	// Start server
	fmt.Println("Starting server...")
	l, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	s = grpc.NewServer()
	fulcrumpb.RegisterFulcrumServiceServer(s, &server{})
	if err := s.Serve(l); err != nil {
		log.Fatalf("Failed to server %v", err)
	}
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
