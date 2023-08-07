package main

import (
	"context"
	"grpc-go-course/greet/proto"
	"io"
	"log"
)

func doGreetManyTimes(c proto.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	req := proto.GreetRequest{
		FirstName: "David",
	}

	stream, err := c.GreetManytimes(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %s", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v", err)
		}

		log.Printf("Result: %s\n", msg.GetResult())
	}
}
