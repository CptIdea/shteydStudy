package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	reference "reference/pkg/grpc"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial("127.0.0.1:8090", opts...)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	client := reference.NewGetterClient(conn)
	request := &reference.ReferenceName{Name: "power"}
	response, err := client.GetReferenceId(context.Background(), request)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	fmt.Println(response)
}
