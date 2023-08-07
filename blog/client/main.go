package main

import (
	"grpc-go-course/blog/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error listening: %v", err)
	}
	defer conn.Close()

	c := proto.NewBlogServiceClient(conn)

	// createBlog(c)
	// readBlog("64d15e4a0d083169f468158a", c)
	// updateBlog("64d15e4a0d083169f468158a", c)
	// listBlogs(c)
	deleteBlog("64d15e4a0d083169f468158a", c)
}
