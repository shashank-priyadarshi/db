package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "dbconn/db"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDatabaseServiceServer(s, pb.UnimplementedDatabaseServiceServer{})
	log.Println("Starting server on :8080")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
