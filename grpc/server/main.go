package main

import (
	"fmt"

	pb "github.com/miguellgt/goprojects/grpc/customer"
	"golang.org/x/net/context"
)

const (
	port = ":50051"
)

// server is used to implement customer.CustomerRequest
type server struct {
	savedCustomers []*pb.CustomerRequest
}

func (s *server) CreateCustomer(ctx context.Context, in *pb.CustomerRequest) (*pb.CustomerResponse, error) {
	s.savedCustomers = append(s.savedCustomers, in)
	return &pb.CustomerResponse{Id: in.Id, Success: true}, nil
}

func main() {
	fmt.Println("Hello grpc")
}
