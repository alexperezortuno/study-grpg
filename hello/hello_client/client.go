package main

import (
	"../hellopb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Go client is running")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect %v", err)
	}

	defer cc.Close()

	c := hellopb.NewHelloServiceClient(cc)

	fmt.Println("Starting unary RPC Hello")

	req := &hellopb.HelloRequest{
		Hello: &hellopb.Hello{
			FirstName: "Mark",
			Prefix:    "Joven",
		},
	}

	res, err := c.Hello(context.Background(), req)

	if err != nil {
		log.Fatalf("Error, calling Hello RPC: \n%v", err)
	}

	log.Printf("Responce Hello: %v", res.CustomHello)
}
