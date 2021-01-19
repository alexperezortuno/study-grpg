package main

import (
	"../hellopb"
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
}
