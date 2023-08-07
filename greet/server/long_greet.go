package main

import (
	"fmt"
	"grpc-go-course/greet/proto"
	"io"
	"log"
)

func (s *Server) LongGreet(stream proto.GreetService_LongGreetServer) error {
	log.Printf("LongGreet function was invoked")

	res := ""

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.GreetResponse{
				Result: res,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}

		log.Printf("Receiving %v\n", req)

		res += fmt.Sprintf("Hello %s!\n", req.GetFirstName())
	}
}
