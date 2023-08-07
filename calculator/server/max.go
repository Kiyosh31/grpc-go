package main

import (
	"grpc-go-course/calculator/proto"
	"log"
)

func (s *Server) Max(stream proto.CalculatorService_MaxServer) error {
	log.Printf("Max was invoked")
}
