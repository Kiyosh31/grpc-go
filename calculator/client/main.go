package main

import (
	"grpc-go-course/calculator/proto"
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

	c := proto.NewCalculatorServiceClient(conn)

	// doSum(c)
	// doPrimeManyTimes(c)
	doAverage(c)
}
