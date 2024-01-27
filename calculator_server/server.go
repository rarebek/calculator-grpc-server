package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc-server/genprotos"
	"log"
	"net"
)

type CalculatorServer struct {
	pb.UnimplementedCalculatorServer
}

func (s *CalculatorServer) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	result := req.A + req.B
	return &pb.AddResponse{Result: result}, nil
}

func (s *CalculatorServer) Subtract(ctx context.Context, req *pb.SubtractRequest) (*pb.SubtractResponse, error) {
	result := req.A - req.B
	return &pb.SubtractResponse{Result: result}, nil
}

func (s *CalculatorServer) Multiply(ctx context.Context, req *pb.MultiplyRequest) (*pb.MultiplyResponse, error) {
	result := req.A * req.B
	return &pb.MultiplyResponse{Result: result}, nil
}

func (s *CalculatorServer) Divide(ctx context.Context, req *pb.DivideRequest) (*pb.DivideResponse, error) {
	result := req.A / req.B
	return &pb.DivideResponse{Result: result}, nil
}

func main() {
	connection, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln("Failed to listen: ", err)
	}

	server := grpc.NewServer()
	pb.RegisterCalculatorServer(server, &CalculatorServer{})
	server.Serve(connection)
}
