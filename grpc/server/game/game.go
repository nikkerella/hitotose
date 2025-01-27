package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc/protobuf/game"

	svc "github.com/nikkerella/hitotose/gin/svc/game"
	"google.golang.org/grpc"
)

// GameServiceServer is the implementation of the gRPC service.
type GameServiceServer struct {
	pb.UnimplementedGameServiceServer
}

// QueryByStatus is the RPC method that queries games by status.
func (s *GameServiceServer) QueryByStatus(ctx context.Context, req *pb.QueryReq) (*pb.QueryResp, error) {
	// Fetch the data from your MongoDB service based on the status from the request
	status := req.GetStatus()

	// Assuming svc.ByStatus returns a list of games
	games := svc.ByStatus(status)

	// Map the results to the gRPC response
	var gameList []*game.Game
	for _, g := range games {
		gameList = append(gameList, &game.Game{
			Id:    g.ID,
			Title: g.Title,
			Genre: g.Genre,
		})
	}

	// Return the games in the response
	return &game.QueryResponse{Games: gameList}, nil
}

func main() {
	// Set up a listener for gRPC server
	listener, err := net.Listen("tcp", ":50051") // gRPC typically uses TCP
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a gRPC server
	grpcServer := grpc.NewServer()

	// Register the GameServiceServer to handle the RPC requests
	game.RegisterGameServiceServer(grpcServer, &GameServiceServer{})

	// Start the server
	fmt.Println("gRPC server is running on port :50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
