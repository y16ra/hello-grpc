package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/y16ra/hello-grpc/hello"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	hello.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *hello.SayHelloRequest) (*hello.SayHelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &hello.SayHelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	// Specify the port number
	listenPort, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	hello.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", listenPort.Addr())
	if err := s.Serve(listenPort); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
