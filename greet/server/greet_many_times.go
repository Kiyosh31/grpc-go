package main

import (
	"fmt"
	"grpc-go-course/greet/proto"
	"log"
)

func (s *Server) GreetManytimes(in *proto.GreetRequest, stream proto.GreetService_GreetManytimesServer) error {
	log.Printf("GreetManyTimes function was invoked %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", in.GetFirstName(), i)

		stream.Send(&proto.GreetResponse{
			Result: res,
		})
	}

	return nil
}
