package main

import (
	"context"
	"grpc-go-course/greet/proto"
	"io"
	"log"
	"time"
)

func doGreetEveryone(c proto.GreetServiceClient) {
	log.Printf("doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while casting stream: %s", err)
	}

	reqs := []*proto.GreetRequest{
		{FirstName: "David"},
		{FirstName: "Marie"},
		{FirstName: "Alexander"},
		{FirstName: "Popo"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v \n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
				break
			}

			log.Printf("Received: %v \n", res.GetResult())
		}
		close(waitc)
	}()

	<-waitc
}
