package main

import (
	"context"
	"log"

	pb "github.com/VivekSingh14/GrpcHandson/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet fucntion was invoked with %v \n ", in)
	return &pb.GreetResponse{
		Result: "Hello" + in.FirstName,
	}, nil
}
