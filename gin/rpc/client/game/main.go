package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/nikkerella/hitotose/gin/rpc/protobuf/game"

	"google.golang.org/grpc"
)

func main() {
	// Connect to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGameServiceClient(conn)

	// Make a query request
	resp, err := client.QueryByStatus(context.Background(), &pb.QueryReq{
		Status: "Playing",
	})
	if err != nil {
		log.Fatalf("could not query: %v", err)
	}

	fmt.Println("Resp:", resp)
}
