package main

import (
	"context"
	"log"

	pb "github.com/deniskarpenko/go-grpc/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Denys",
	})

	if err != nil {
		log.Fatal("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
