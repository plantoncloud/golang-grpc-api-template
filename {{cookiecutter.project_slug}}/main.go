package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	healthcontroller "{{cookiecutter.go_module_name}}/controller/health"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"net"
)

const (
	Port = 8080
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	registerServices(grpcServer)
	log.Printf("grpc server listening on %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func registerServices(grpcServer *grpc.Server) {
	grpc_health_v1.RegisterHealthServer(grpcServer, &healthcontroller.Server{})
}
