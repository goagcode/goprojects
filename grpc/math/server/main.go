package main

import (
	"fmt"

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
	fmt.Println("Running tcp server")
}
