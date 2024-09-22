package main

import (
	"log"
	"net"

	"github.com/STO-KubSU/raptor-product-service/internal/service"
	pb "github.com/STO-KubSU/raptor-proto/productpb"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	productService := service.NewProductService()

	pb.RegisterProductServiceServer(grpcServer, productService)

	log.Println("Product service is running on port 50053...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
