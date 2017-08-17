package main

import (
	"flag"
	"log"

	pb "github.com/miguellgt/goprojects/grpc/google/google_svc"

	"google.golang.org/grpc"
)

func main() {
	server := flag.String("s", ":8080", "address to grpc server")
	query := flag.String("q", "gRPC framework", "query to search int the google search")
	mode := flag.String("m", "search", "mode to use the service")

	flag.Parse()

	// Connect to the server
	conn, err := grpc.Dial(*server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to the server %s: %v", *server, err)
	}
	defer conn.Close()

	client := pb.NewGoogleClient(conn)

	// Run the RPC.
	switch *mode {
	case "search":
		search(client, *query)
	case "watch":
		watch(client, *query)
	default:
		log.Fatalf("Unknown mode: %q", *mode)
	}
}
