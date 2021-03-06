package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os/exec"

	"google.golang.org/grpc"

	"github.com/Sirupsen/logrus"
	pb "github.com/miguellgt/goprojects/grpc/say-grpc/api"
	"golang.org/x/net/context"
)

func main() {
	port := flag.Int("p", 8080, "port to listen to")
	flag.Parse()

	logrus.Infof("listening to port: %d", *port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logrus.Fatalf("could not listen to port %d: %v", *port, err)
	}

	s := grpc.NewServer()
	pb.RegisterTextToSpeechServer(s, server{})
	err = s.Serve(lis)
	if err != nil {
		logrus.Fatalf("could not serve: %v", err)
	}
}

type server struct{}

func (server) Say(ctx context.Context, in *pb.Text) (*pb.Speech, error) {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		return nil, fmt.Errorf("could not create tmp file: %v", err)
	}
	if err := f.Close(); err != nil {
		return nil, fmt.Errorf("could not close tmp file %s: %v", f.Name(), err)
	}

	cmd := exec.Command("flite", "-t", in.Text, "-o", f.Name())
	if data, err := cmd.CombinedOutput(); err != nil {
		return nil, fmt.Errorf("flite failed: %s", data)
	}

	data, err := ioutil.ReadFile(f.Name())
	if err != nil {
		fmt.Errorf("could not read tmp file: %v", err)
	}

	return &pb.Speech{Audio: data}, nil
}
