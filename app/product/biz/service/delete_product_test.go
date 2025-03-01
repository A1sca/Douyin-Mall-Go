package service

import (
	"context"
	"testing"

	product "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/product/kitex_gen"
)

func TestDeleteProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeleteProductService(ctx)
	// init req and assert value

	req := &product.DeleteProductReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
