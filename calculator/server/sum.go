package main

import (
	"context"
	"grpc-go-course/calculator/proto"
	"log"
)

func (s *Server) Sum(ctx context.Context, in *proto.SumRequest) (*proto.SumResponse, error) {
	log.Printf("Sum funcion was invoked with %v\n", in)

	return &proto.SumResponse{
		Result: in.GetNum1() + in.GetNum2(),
	}, nil
}
