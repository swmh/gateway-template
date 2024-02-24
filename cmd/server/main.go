package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "grpclearn/internal/genproto"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedNoteV1Server
}

func (s server) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	fmt.Printf("%v\n", request)
	return nil, nil
}

func (s server) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	fmt.Printf("%v\n", request)
	return nil, nil
}

func (s server) List(ctx context.Context, request *pb.ListRequest) (*pb.ListResponse, error) {
	fmt.Printf("%v\n", request)
	return nil, nil
}

func (s server) Update(ctx context.Context, request *pb.UpdateRequest) (*emptypb.Empty, error) {
	fmt.Printf("%v\n", request)
	return nil, nil
}

func (s server) Delete(ctx context.Context, request *pb.DeleteRequest) (*emptypb.Empty, error) {
	fmt.Printf("%v\n", request)
	return nil, nil
}

func main() {
	s := &server{}
	grpcServer := grpc.NewServer()
	pb.RegisterNoteV1Server(grpcServer, s)

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Server started")
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("Server stopped: %v", err)
	}
}
