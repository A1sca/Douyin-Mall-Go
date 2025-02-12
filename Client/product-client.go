package product_client

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/A1sca/Douyin-Mall-Go/grpc"
)

type ProductClient struct {
	conn   *grpc.ClientConn
	client pb.ProductCatalogServiceClient
}

// 初始化服务
func NewProductClient(serverAddr string) (*ProductClient, error) {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := pb.NewProductCatalogServiceClient(conn)
	return &ProductClient{
		conn:   conn,
		client: client,
	}, nil
}

func (c *ProductClient) Close() {
	c.conn.Close()
}

// CreateProduct 实现创建商品方法
func (c *ProductClient) CreateProduct(name, description string, price float32, categories []string) (uint32, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := c.client.CreateProduct(ctx, &pb.CreateProductReq{
		Name:        name,
		Description: description,
		Price:       price,
		Categories:  categories,
	})
	if err != nil {
		return 0, err
	}
	return resp.Id, nil
}

// DeleteProduct 实现删除商品方法
func (c *ProductClient) DeleteProduct(id uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := c.client.DeleteProduct(ctx, &pb.DeleteProductReq{Id: id})
	return err
}

// UpdateProduct 实现更新商品方法
func (c *ProductClient) UpdateProduct(id uint32, name, description string, price float32, categories []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := c.client.UpdateProduct(ctx, &pb.UpdateProductReq{
		Id:          id,
		Name:        name,
		Description: description,
		Price:       price,
		Categories:  categories,
	})
	return err
}

// ListProducts 实现获取产品列表方法
func (c *ProductClient) ListProducts(page int32, pageSize int64) ([]*pb.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := c.client.ListProducts(ctx, &pb.ListProductsReq{
		Page:     page,
		PageSize: pageSize,
	})
	if err != nil {
		return nil, err
	}
	return resp.Products, nil
}

// GetProduct 实现获取单个产品方法
func (c *ProductClient) GetProduct(id uint32) (*pb.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := c.client.GetProduct(ctx, &pb.GetProductReq{Id: id})
	if err != nil {
		if status, ok := status.FromError(err); ok && status.Code() == codes.NotFound {
			return nil, fmt.Errorf("product %d not found", id)
		}
		return nil, err
	}
	return resp.Product, nil
}

// SearchProducts 实现搜索产品方法
func (c *ProductClient) SearchProducts(query string) ([]*pb.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := c.client.SearchProducts(ctx, &pb.SearchProductsReq{Query: query})
	if err != nil {
		return nil, err
	}
	return resp.Results, nil
}

func main() {
	// 注册服务
	client, err := NewProductClient("localhost:50051")
	if err != nil {
		log.Fatalf("客户端创建失败: %v", err)
	}
	defer client.Close()

	// 测试用例
	testProduct := &struct {
		Name        string
		Description string
		Price       float32
		Categories  []string
	}{
		Name:        "刷子",
		Description: "用来清洁的用品",
		Price:       9.9,
		Categories:  []string{"工具", "清洁工具"},
	}

	// 创建商品
	id, err := client.CreateProduct(
		testProduct.Name,
		testProduct.Description,
		testProduct.Price,
		testProduct.Categories,
	)
	if err != nil {
		log.Fatalf("创建商品失败: %v", err)
	}
	log.Printf("创建的商品 ID: %d", id)

	// 获取商品详情
	product, err := client.GetProduct(id)
	if err != nil {
		log.Fatalf("获取商品失败: %v", err)
	}
	log.Printf("商品详情:\n%+v", product)

	// 更新商品
	newPrice := float32(5.9)
	err = client.UpdateProduct(
		id,
		product.Name,
		product.Description,
		newPrice,
		product.Categories,
	)
	if err != nil {
		log.Fatalf("更新商品失败: %v", err)
	}
	log.Println("刷新成功")

	// 搜索商品
	results, err := client.SearchProducts("刷子")
	if err != nil {
		log.Fatalf("搜索失败: %v", err)
	}
	log.Printf("搜索结果 (%d 项):", len(results))
	for _, p := range results {
		log.Printf("- %s ($%.2f)", p.Name, p.Price)
	}

	// 列出商品
	products, err := client.ListProducts(1, 10)
	if err != nil {
		log.Fatalf("陈列商品无效: %v", err)
	}
	log.Printf("第一页商品 (%d 项):", len(products))
	for _, p := range products {
		log.Printf("- %s ($%.2f)", p.Name, p.Price)
	}

	// 删除商品
	if err := client.DeleteProduct(id); err != nil {
		log.Fatalf("删除该商品失败: %v", err)
	}
	log.Println("该商品删除成功")
}
