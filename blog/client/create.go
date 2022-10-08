package main

import (
	"context"
	"log"

	pb "github.com/VivekSingh14/GrpcHandson/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("---createBlog was invoked.------")
	blog := &pb.Blog{
		AuthorId: "Vivek",
		Content:  "Temporary content",
		Title:    "My name is Vivek",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error: %v \n", err)
	}

	log.Printf("Blog has been created: %v \n", res.Id)
	return res.Id

}
