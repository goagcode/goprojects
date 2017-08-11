package main

import (
	"fmt"
	"log"
	"os"

	"github.com/miguellgt/goprojects/grpc/math/math_service"

	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	client := math_service.NewMathClient(conn)
	fmt.Println("Valid Request: 10/4")
	result, err := client.Divide(context.Background(), &math_service.Operands{
		Dividend: 10,
		Divisor:  4,
	})
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println(result.Quotient)

	fmt.Println("Valid Request: 10/0")
	_, err = client.Divide(context.Background(), &math_service.Operands{
		Dividend: 10,
		Divisor:  0,
	})
	if err != nil {
		log.Fatalf("%v", err)
	}
}
