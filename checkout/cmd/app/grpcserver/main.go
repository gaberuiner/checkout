package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "your-package-path/your-proto-package" // Замените на путь к вашему пакету протофайлов

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedYourServiceServer
}

func (s *server) YourRPCMethod(ctx context.Context, req *pb.YourRequest) (*pb.YourResponse, error) {
	// Обработка вашего запроса и возвращение ответа
	// Вместо "YourRPCMethod", "YourRequest" и "YourResponse" используйте ваши сгенерированные имена
	// и методы из протофайла
	response := &pb.YourResponse{
		Message: "Hello, gRPC Client!",
	}
	return response, nil
}

func main() {
	port := 50051

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterYourServiceServer(s, &server{})

	log.Printf("gRPC server listening on port %d", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
