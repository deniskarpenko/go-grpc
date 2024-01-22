package main

import (
	"log"
	"net"

	pb "github.com/deniskarpenko/go-grpc/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:9001"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatal("Failed to listen on: %v\n", err)
	}

	log.Printf("Listenong on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
