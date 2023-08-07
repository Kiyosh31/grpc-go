package main

import (
	"grpc-go-course/greet/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	c := proto.NewGreetServiceClient(conn)

	// doGreet(c)
	// doGreetManyTimes(c)
	// doLongGreet(c)
	doGreetEveryone(c)
}
