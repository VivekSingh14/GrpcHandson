package main

import (
	"log"

	pb "github.com/VivekSingh14/GrpcHandson/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v \n", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	//doGreet(c)

	//doGreetManyTimes(c)

	doGreetEveryone(c)

}
