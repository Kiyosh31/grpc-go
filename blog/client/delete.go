package main

import (
	"context"
	"grpc-go-course/blog/proto"
	"log"
)

func deleteBlog(id string, c proto.BlogServiceClient) {
	log.Printf("deleteBlog was invoked")

	_, err := c.DeleteBlog(context.Background(), &proto.BlogId{Id: id})
	if err != nil {
		log.Fatalf("Error while deleting: %v", err)
	}

	log.Println("Blog was deleted!")
}
