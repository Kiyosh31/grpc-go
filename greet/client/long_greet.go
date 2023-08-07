package main

import (
	"context"
	"grpc-go-course/greet/proto"
	"log"
	"time"
)

func doLongGreet(c proto.GreetServiceClient) {
	log.Printf("doLongGreet was invoked")

	reqs := []*proto.GreetRequest{
		{FirstName: "David"},
		{FirstName: "German"},
		{FirstName: "Aleman"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet: %s", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %s", err)
	}

	log.Printf("LongGreet: %s\n", res)

}
