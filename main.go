package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"buf.build/go/protovalidate"
	"github.com/gilperopiola/go-demo-grpc/pbs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	validator, err := protovalidate.New()
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	service := &service{validator: validator}

	pbs.RegisterAuthServiceServer(grpcServer, service)
	pbs.RegisterUsersServiceServer(grpcServer, service)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	go func() {
		log.Println("Starting server on :50051")
		if err := grpcServer.Serve(listener); err != nil {
			panic(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	log.Println("Shutting down server...")
	grpcServer.GracefulStop()
}
