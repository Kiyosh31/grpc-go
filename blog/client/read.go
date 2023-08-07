package main

import (
	"context"
	"grpc-go-course/blog/proto"
	"log"
)

func readBlog(id string, c proto.BlogServiceClient) *proto.Blog {
	log.Printf("readBlog was invoked")

	req := &proto.BlogId{
		Id: id,
	}
	res, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("Error happened while reading: %v", err)
	}

	log.Printf("Log was read: %v", res)
	return res
}
