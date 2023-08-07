package main

import (
	"grpc-go-course/calculator/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

var addr = "0.0.0.0:50051"

type Server struct {
	proto.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Listening on: %s\n", addr)

	svc := grpc.NewServer()
	proto.RegisterCalculatorServiceServer(svc, &Server{})

	if err = svc.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
