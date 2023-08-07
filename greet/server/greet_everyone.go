package main

import (
	"grpc-go-course/greet/proto"
	"io"
	"log"
)

func (s *Server) GreetEveryone(stream proto.GreetService_GreetEveryoneServer) error {
	log.Printf("GreetEveryone was invoked")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %s", err)
		}

		res := "Hello " + req.GetFirstName() + "!"
		err = stream.Send(&proto.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Error while sendind data to client: %s", err)
		}
	}
}
