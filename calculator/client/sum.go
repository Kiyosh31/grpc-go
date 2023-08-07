package main

import (
	"context"
	"grpc-go-course/calculator/proto"
	"log"
)

func doSum(c proto.CalculatorServiceClient) {
	log.Println("doSum was invoked")

	res, err := c.Sum(context.Background(), &proto.SumRequest{
		Num1: 10,
		Num2: 10,
	})
	if err != nil {
		log.Fatalf("Could ot sum: %v", err)
	}

	log.Printf("Sum: %s", res.GetResult())
}
