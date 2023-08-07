package main

import (
	"grpc-go-course/greet/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	proto.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	log.Printf("Listening on: %s\n", addr)

	svc := grpc.NewServer()
	proto.RegisterGreetServiceServer(svc, &Server{})

	if err = svc.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
