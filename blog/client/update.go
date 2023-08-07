package main

import (
	"context"
	"grpc-go-course/blog/proto"
	"log"
)

func updateBlog(id string, c proto.BlogServiceClient) {
	log.Printf("updateBlog was invoked")

	newBlog := &proto.Blog{
		Id:       id,
		AuthorId: "Now changed",
		Title:    "New title",
		Content:  "New content",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("Error happening while updating: %v", err)
	}

	log.Println("Blog updated!")
}
