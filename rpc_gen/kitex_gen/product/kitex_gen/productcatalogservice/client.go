// Code generated by Kitex v0.12.2. DO NOT EDIT.

package productcatalogservice

import (
	"context"
	kitex_gen "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/product/kitex_gen"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	ListProducts(ctx context.Context, Req *kitex_gen.ListProductsReq, callOptions ...callopt.Option) (r *kitex_gen.ListProductsResp, err error)
	GetProduct(ctx context.Context, Req *kitex_gen.GetProductReq, callOptions ...callopt.Option) (r *kitex_gen.GetProductResp, err error)
	SearchProducts(ctx context.Context, Req *kitex_gen.SearchProductsReq, callOptions ...callopt.Option) (r *kitex_gen.SearchProductsResp, err error)
	CreateProduct(ctx context.Context, Req *kitex_gen.CreateProductReq, callOptions ...callopt.Option) (r *kitex_gen.CreateProductResp, err error)
	UpdateProduct(ctx context.Context, Req *kitex_gen.UpdateProductReq, callOptions ...callopt.Option) (r *kitex_gen.UpdateProductResp, err error)
	DeleteProduct(ctx context.Context, Req *kitex_gen.DeleteProductReq, callOptions ...callopt.Option) (r *kitex_gen.DeleteProductResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kProductCatalogServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kProductCatalogServiceClient struct {
	*kClient
}

func (p *kProductCatalogServiceClient) ListProducts(ctx context.Context, Req *kitex_gen.ListProductsReq, callOptions ...callopt.Option) (r *kitex_gen.ListProductsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListProducts(ctx, Req)
}

func (p *kProductCatalogServiceClient) GetProduct(ctx context.Context, Req *kitex_gen.GetProductReq, callOptions ...callopt.Option) (r *kitex_gen.GetProductResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetProduct(ctx, Req)
}

func (p *kProductCatalogServiceClient) SearchProducts(ctx context.Context, Req *kitex_gen.SearchProductsReq, callOptions ...callopt.Option) (r *kitex_gen.SearchProductsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SearchProducts(ctx, Req)
}

func (p *kProductCatalogServiceClient) CreateProduct(ctx context.Context, Req *kitex_gen.CreateProductReq, callOptions ...callopt.Option) (r *kitex_gen.CreateProductResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateProduct(ctx, Req)
}

func (p *kProductCatalogServiceClient) UpdateProduct(ctx context.Context, Req *kitex_gen.UpdateProductReq, callOptions ...callopt.Option) (r *kitex_gen.UpdateProductResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateProduct(ctx, Req)
}

func (p *kProductCatalogServiceClient) DeleteProduct(ctx context.Context, Req *kitex_gen.DeleteProductReq, callOptions ...callopt.Option) (r *kitex_gen.DeleteProductResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteProduct(ctx, Req)
}
