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
	products map[int32]*pb.Product
}

func NewProductService() *ProductService {
	// Пример данных о продуктах
	products := map[int32]*pb.Product{
		0: {Id: i32Ptr(0), Name: "Phone", Quantity: 10, Price: 999.99},
		1: {Id: i32Ptr(1), Name: "Laptop", Quantity: 10, Price: 999.99},
		2: {Id: i32Ptr(2), Name: "Smartphone", Quantity: 5, Price: 499.99},
	}
	return &ProductService{products: products}
}

func (s *ProductService) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	product, exists := s.products[req.Id]
	if !exists {
		return nil, fmt.Errorf("product not found")
	}
	return &pb.GetProductResponse{Product: product}, nil
}

func (s *ProductService) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	product := req.Product
	product.Id = i32Ptr(int32(len(s.products)))
	s.products[*product.Id] = product
	return &pb.CreateProductResponse{Id: *product.Id}, nil

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
