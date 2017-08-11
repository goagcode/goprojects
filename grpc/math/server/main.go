package main

import (
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	ms "github.com/miguellgt/goprojects/grpc/math/math_service"

	context "golang.org/x/net/context"
)

type mathServer struct{}

func (m *mathServer) Divide(ctx context.Context, in *ms.Operands) (*ms.Result, error) {
	// Handle divide by zero
	if in.Dividend == 0 {
		return nil, grpc.Errorf(codes.InvalidArgument, "Cannot divide by zero")
	}
	// cast, divide, return
	return &ms.Result{Quotient: float32(in.Dividend) / float32(in.Divisor)}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Println("Could not listen on 50051")
		os.Exit(-1)
	}
	// new server, register, and serve
	server := grpc.NewServer()
	ms.RegisterMathServer(server, mathServer{})
	server.Serve(listener)
}
