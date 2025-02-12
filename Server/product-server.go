package main

import (
	"context"
	"log"
	"net"
	"strings"
	"sync"

	pb "github.com/A1sca/Douyin-Mall-Go/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedProductCatalogServiceServer
	mu       sync.RWMutex
	Products map[uint32]*pb.Product
	nextID   uint32
}

func newServer() *server {
	return &server{
		Products: make(map[uint32]*pb.Product),
		nextID:   1,
	}
}

// CreateProduct 创建商品
func (s *server) CreateProduct(ctx context.Context, req *pb.CreateProductReq) (*pb.CreateProductResp, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	product := &pb.Product{
		Id:          s.nextID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Categories:  req.Categories,
	}

	s.Products[s.nextID] = product
	resp := &pb.CreateProductResp{Id: s.nextID}
	s.nextID++
	return resp, nil
}

// DeleteProduct 删除商品
func (s *server) DeleteProduct(ctx context.Context, req *pb.DeleteProductReq) (*pb.DeleteProductResp, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.Products[req.Id]; !exists {
		return nil, status.Errorf(codes.NotFound, "未找到产品")
	}

	delete(s.Products, req.Id)
	return &pb.DeleteProductResp{Success: true}, nil
}

// UpdateProduct 更新商品
func (s *server) UpdateProduct(ctx context.Context, req *pb.UpdateProductReq) (*pb.UpdateProductResp, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	product, exists := s.Products[req.Id]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "未找到产品")
	}

	// 更新字段
	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price
	product.Categories = req.Categories

	return &pb.UpdateProductResp{Success: true}, nil
}

// ListProducts 返回分页的产品列表
func (s *server) ListProducts(ctx context.Context, req *pb.ListProductsReq) (*pb.ListProductsResp, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	allProducts := make([]*pb.Product, 0, len(s.Products))
	for _, p := range s.Products {
		allProducts = append(allProducts, p)
	}

	// 分页实现
	start := (int(req.Page) - 1) * int(req.PageSize)
	end := start + int(req.PageSize)

	if start > len(allProducts) {
		start = len(allProducts)
	}
	if end > len(allProducts) {
		end = len(allProducts)
	}

	return &pb.ListProductsResp{
		Products: allProducts[start:end],
	}, nil
}

// GetProduct 根据 ID 返回单个产品
func (s *server) GetProduct(ctx context.Context, req *pb.GetProductReq) (*pb.GetProductResp, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	product, exists := s.Products[req.Id]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "未找到产品")
	}

	return &pb.GetProductResp{Product: product}, nil
}

// SearchProducts 根据查询字符串搜索产品
func (s *server) SearchProducts(ctx context.Context, req *pb.SearchProductsReq) (*pb.SearchProductsResp, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	results := make([]*pb.Product, 0)
	for _, p := range s.Products {
		if strings.Contains(p.Name, req.Query) || strings.Contains(p.Description, req.Query) {
			results = append(results, p)
		}
	}

	return &pb.SearchProductsResp{
		Results: results,
	}, nil
}

// containsString 辅助函数，检查字符串中是否包含子字符串
func containsString(str, substr string) bool {
	// 简单实现，可根据需求优化
	for i := 0; i <= len(str)-len(substr); i++ {
		if str[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func main() {

	// 启动 gRPC 服务器
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterProductCatalogServiceServer(s, newServer())
	log.Printf("服务器正在监听 %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}

}
