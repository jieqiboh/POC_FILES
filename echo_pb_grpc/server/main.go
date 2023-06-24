package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"server/model"

	"google.golang.org/grpc"
)

type server struct {
	model.UnimplementedEchoServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) Echo(ctx context.Context, in *model.Request) (*model.Response, error) {
	log.Printf("Received: %v", in.GetMessage())
	return &model.Response{Message: "Hello " + in.GetMessage()}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8888))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	model.RegisterEchoServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
