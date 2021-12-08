package main

import (
	"context"
	"fmt"
	"go-starwars/api/fulcrumpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

var cf2 fulcrumpb.FulcrumServiceClient
var cf3 fulcrumpb.FulcrumServiceClient
var s *grpc.Server

func main() {
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

func (*server) GetVector(ctx context.Context, req *fulcrumpb.GetVectorRequest) (*fulcrumpb.GetVectorResponse, error) {
	// Unpack request
	planet := req.GetPlanet()

	fmt.Println(planet)

	// Pack response
	vector := []int32{0, 0, 0}

	// Send response
	res := &fulcrumpb.GetVectorResponse{
		Vector: vector,
	}
	return res, nil
}

func (*server) GetNumberRebelsFulcrum(ctx context.Context, req *fulcrumpb.GetNumberRebelsFulcrumRequest) (*fulcrumpb.GetNumberRebelsFulcrumResponse, error) {
	// Unpack request
	planet := req.GetPlanet()
	city := req.GetCity()

	fmt.Println(planet, city)

	// Pack response
	success := false
	number := -1

	// Send response
	res := &fulcrumpb.GetNumberRebelsFulcrumResponse{
		Success: success,
		Number:  int32(number),
	}
	return res, nil
}

func (*server) GetLogs(ctx context.Context, req *fulcrumpb.GetLogsRequest) (*fulcrumpb.GetLogsResponse, error) {
	// Unpack request
	value := req.GetValue()

	fmt.Println(value)

	// Pack response
	logs := []string{}

	// Send response
	res := &fulcrumpb.GetLogsResponse{
		Logs: logs,
	}
	return res, nil
}

func (*server) Merge(ctx context.Context, req *fulcrumpb.MergeRequest) (*fulcrumpb.MergeResponse, error) {
	// Unpack request
	files := req.GetFiles()

	fmt.Println(files)

	// Pack response
	success := false

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
	success := false
	vector := []int32{0, 0, 0}

	// Send response
	res := &fulcrumpb.AddCityResponse{
		Success: success,
		Vector:  vector,
	}
	return res, nil
}

func (*server) UpdateName(ctx context.Context, req *fulcrumpb.UpdateNameRequest) (*fulcrumpb.UpdateNameResponse, error) {
	// Unpack request
	planet := req.GetPlanet()
	oldCity := req.GetOldCity()
	newCity := req.GetNewCity()

	fmt.Println(planet, oldCity, newCity)

	// Pack response
	success := false
	vector := []int32{0, 0, 0}

	// Send response
	res := &fulcrumpb.UpdateNameResponse{
		Success: success,
		Vector:  vector,
	}
	return res, nil
}

func (*server) UpdateNumber(ctx context.Context, req *fulcrumpb.UpdateNumberRequest) (*fulcrumpb.UpdateNumberResponse, error) {
	// Unpack request
	planet := req.GetPlanet()
	city := req.GetCity()
	number := req.GetNumber()

	fmt.Println(planet, city, number)

	// Pack response
	success := false
	vector := []int32{0, 0, 0}

	// Send response
	res := &fulcrumpb.UpdateNumberResponse{
		Success: success,
		Vector:  vector,
	}
	return res, nil
}

func (*server) DeleteCity(ctx context.Context, req *fulcrumpb.DeleteCityRequest) (*fulcrumpb.DeleteCityResponse, error) {
	// Unpack request
	planet := req.GetPlanet()
	city := req.GetCity()

	fmt.Println(planet, city)

	// Pack response
	success := false
	vector := []int32{0, 0, 0}

	// Send response
	res := &fulcrumpb.DeleteCityResponse{
		Success: success,
		Vector:  vector,
	}
	return res, nil
}
