package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/miguellgt/goprojects/grpc/customer"
)

const (
	address = "localhost:50051"
)

// createCustomer calls the RPC method CreateCustomer of CustomerServer
func createCustomer(client pb.CustomerClient, customer *pb.CustomerRequest) {
	resp, err := client.CreateCustomer(context.Background(), customer)
	if err != nil {
		log.Fatalf("Could not create Customer: %v", err)
	}
	if resp.Success {
		log.Println("A new Customer has been added with id: %d", resp.Id)
	}
}

func main() {
	fmt.Println("Hello")
}
