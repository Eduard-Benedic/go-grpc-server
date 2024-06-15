package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/Eduard-Benedic/course-protofiles/protogen/go/hello"
	"github.com/Eduard-Benedic/go-grpc-server/internal/port"
	"google.golang.org/grpc"
)

type GrpcAdapter struct {
	helloService port.HelloServicePort
	grpcPort     int
	server       *grpc.Server
	hello.HelloServiceServer
}

func NewGrpcAdapter(helloservice port.HelloServicePort, grpcPort int) *GrpcAdapter {
	return &GrpcAdapter{
		helloService: helloService,
		grpcPort:     grpcPort,
	}
}

func (a *GrpcAdapter) Run() {
	var err error
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.grpcPort))
	if err != nil {
		log.Fatalf("Failed to list on port %d: %v", a.grpcPort, err)
	}

	log.Printf("Server listening on port %d \n", a.grpcPort)

	grpcServer := grpc.NewServer()

	a.server = grpcServer
	hello.RegisterHelloServiceServer(grpcServer, a)

	if err = grpc.Serve(listen); err != nil {
		log.Fatalf("Failed to serve gRPC on port %d : %v\n", a.grpcPort, err)
	}
}

func (a *GrpcAdapter) Stop() {
	a.server.Stop()
}
