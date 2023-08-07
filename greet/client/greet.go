package main

import (
	"context"
	"grpc-go-course/greet/proto"
	"log"
)

func doGreet(c proto.GreetServiceClient) {
	log.Println("doGreet was invoked")

	res, err := c.Greet(context.Background(), &proto.GreetRequest{
		FirstName: "David",
	})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	log.Printf("Greeting: %s", res.GetResult())
}
