package main

import (
	"database/sql"
	"fmt"
	"learn-grpc/modules/books"
	proto_books "learn-grpc/modules/books/proto-books"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	proto_books.RegisterBookServer(grpcServer, books.NewBookServer(
		&sql.DB{},
	))

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:8090"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}

	// this function is only return if error happened
	log.Println("starting server on localhost:8090...")
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to start Server: %v", err)
	}
}
