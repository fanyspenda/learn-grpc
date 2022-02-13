package main

import (
	"flag"
	"fmt"
	"learn-grpc/modules/books"
	proto_books "learn-grpc/modules/books/proto-books"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:8090"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var group sync.WaitGroup

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	proto_books.RegisterBookServer(grpcServer, books.NewBookServer())

	go func() {
		group.Add(1)
		err = grpcServer.Serve(lis)
		group.Done()
	}()

	group.Wait()

	if err != nil {
		log.Fatalf("failed to start Server: %v", err)
	}

	fmt.Println("success starting gRPC server")

	<-make(chan bool, 1)
}
