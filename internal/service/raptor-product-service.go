package service

import (
	"context"
	"fmt"

	pb "github.com/STO-KubSU/raptor-proto/productpb"
)

func i32Ptr(i int32) *int32 {
	return &i
}

type ProductService struct {
	pb.UnimplementedProductServiceServer
	products map[int32]*pb.GetProductResponse
}

func NewProductService() *ProductService {
	// Пример данных о продуктах
	products := map[int32]*pb.GetProductResponse{
		1: {Product: &pb.Product{Id: i32Ptr(1), Name: "Laptop", Quantity: 10, Price: 999.99}},
		2: {Product: &pb.Product{Id: i32Ptr(2), Name: "Smartphone", Quantity: 5, Price: 499.99}},
	}
	return &ProductService{products: products}
}

func (s *ProductService) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	product, exists := s.products[req.Id]
	if !exists {
		return nil, fmt.Errorf("product not found")
	}
	return product, nil
}

/*
func (s *ProductService) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	var products []*pb.GetProductResponse
	for _, product := range s.products {
		products = append(products, product)
	}
	return &pb.ListProductsResponse{Products: products}, nil
}
*/ // TODO: реализовать
