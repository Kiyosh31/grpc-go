package main

import (
	"context"
	"grpc-go-course/blog/proto"
	"log"
)

func createBlog(c proto.BlogServiceClient) string {
	log.Printf("createBlog was invoked")

	blog := proto.Blog{
		AuthorId: "David",
		Title:    "My blog",
		Content:  "Content of the blog",
	}

	res, err := c.CreateBlog(context.Background(), &blog)
	if err != nil {
		log.Fatalf("Unexpected error: %s", err)
	}

	log.Printf("Blog has being created: %s\n", res)
	return res.GetId()
}
