package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"reference/internal/reference"
	config "reference/pkg/config/reference"
	reference2 "reference/pkg/grpc"
)

func main() {
	cfg, err := config.GetRefConfig()
	if err != nil {
		log.Fatal(err)
	}

	getterServer := reference.NewService()

	server := grpc.NewServer()

	reference2.RegisterGetterServer(server, getterServer)

	ln, err := net.Listen("tcp", ":"+cfg.GrpcPort)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	log.Println("serving on port", cfg.GrpcPort)
	if err := server.Serve(ln); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
