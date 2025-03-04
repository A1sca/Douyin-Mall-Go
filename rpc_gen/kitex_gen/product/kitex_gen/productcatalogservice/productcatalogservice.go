// Code generated by Kitex v0.12.2. DO NOT EDIT.

package productcatalogservice

import (
	"context"
	"errors"
	kitex_gen "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/product/kitex_gen"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"ListProducts": kitex.NewMethodInfo(
		listProductsHandler,
		newListProductsArgs,
		newListProductsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetProduct": kitex.NewMethodInfo(
		getProductHandler,
		newGetProductArgs,
		newGetProductResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"SearchProducts": kitex.NewMethodInfo(
		searchProductsHandler,
		newSearchProductsArgs,
		newSearchProductsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"CreateProduct": kitex.NewMethodInfo(
		createProductHandler,
		newCreateProductArgs,
		newCreateProductResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"UpdateProduct": kitex.NewMethodInfo(
		updateProductHandler,
		newUpdateProductArgs,
		newUpdateProductResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"DeleteProduct": kitex.NewMethodInfo(
		deleteProductHandler,
		newDeleteProductArgs,
		newDeleteProductResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	productCatalogServiceServiceInfo                = NewServiceInfo()
	productCatalogServiceServiceInfoForClient       = NewServiceInfoForClient()
	productCatalogServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return productCatalogServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return productCatalogServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return productCatalogServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "ProductCatalogService"
	handlerType := (*kitex_gen.ProductCatalogService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.12.2",
		Extra:           extra,
	}
	return svcInfo
}

func listProductsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(kitex_gen.ListProductsReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(kitex_gen.ProductCatalogService).ListProducts(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ListProductsArgs:
		success, err := handler.(kitex_gen.ProductCatalogService).ListProducts(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ListProductsResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newListProductsArgs() interface{} {
	return &ListProductsArgs{}
}

func newListProductsResult() interface{} {
	return &ListProductsResult{}
}

type ListProductsArgs struct {
	Req *kitex_gen.ListProductsReq
}

func (p *ListProductsArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(kitex_gen.ListProductsReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ListProductsArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ListProductsArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ListProductsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ListProductsArgs) Unmarshal(in []byte) error {
	msg := new(kitex_gen.ListProductsReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ListProductsArgs_Req_DEFAULT *kitex_gen.ListProductsReq

func (p *ListProductsArgs) GetReq() *kitex_gen.ListProductsReq {
	if !p.IsSetReq() {
		return ListProductsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ListProductsArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ListProductsArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ListProductsResult struct {
	Success *kitex_gen.ListProductsResp
}

var ListProductsResult_Success_DEFAULT *kitex_gen.ListProductsResp

func (p *ListProductsResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(kitex_gen.ListProductsResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ListProductsResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ListProductsResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ListProductsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ListProductsResult) Unmarshal(in []byte) error {
	msg := new(kitex_gen.ListProductsResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ListProductsResult) GetSuccess() *kitex_gen.ListProductsResp {
	if !p.IsSetSuccess() {
		return ListProductsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ListProductsResult) SetSuccess(x interface{}) {
	p.Success = x.(*kitex_gen.ListProductsResp)
}

func (p *ListProductsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ListProductsResult) GetResult() interface{} {
	return p.Success
}

func getProductHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(kitex_gen.GetProductReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(kitex_gen.ProductCatalogService).GetProduct(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetProductArgs:
		success, err := handler.(kitex_gen.ProductCatalogService).GetProduct(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetProductResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetProductArgs() interface{} {
	return &GetProductArgs{}
}

func newGetProductResult() interface{} {
	return &GetProductResult{}
}

type GetProductArgs struct {
	Req *kitex_gen.GetProductReq
}

func (p *GetProductArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(kitex_gen.GetProductReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetProductArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetProductArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetProductArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetProductArgs) Unmarshal(in []byte) error {
	msg := new(kitex_gen.GetProductReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetProductArgs_Req_DEFAULT *kitex_gen.GetProductReq

func (p *GetProductArgs) GetReq() *kitex_gen.GetProductReq {
	if !p.IsSetReq() {
		return GetProductArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetProductArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetProductArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetProductResult struct {
	Success *kitex_gen.GetProductResp
}

var GetProductResult_Success_DEFAULT *kitex_gen.GetProductResp

func (p *GetProductResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(kitex_gen.GetProductResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetProductResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetProductResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetProductResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetProductResult) Unmarshal(in []byte) error {
	msg := new(kitex_gen.GetProductResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetProductResult) GetSuccess() *kitex_gen.GetProductResp {
	if !p.IsSetSuccess() {
		return GetProductResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetProductResult) SetSuccess(x interface{}) {
	p.Success = x.(*kitex_gen.GetProductResp)
}

func (p *GetProductResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetProductResult) GetResult() interface{} {
	return p.Success
}

func searchProductsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(kitex_gen.SearchProductsReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(kitex_gen.ProductCatalogService).SearchProducts(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *SearchProductsArgs:
		success, err := handler.(kitex_gen.ProductCatalogService).SearchProducts(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SearchProductsResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newSearchProductsArgs() interface{} {
	return &SearchProductsArgs{}
}

func newSearchProductsResult() interface{} {
	return &SearchProductsResult{}
}

type SearchProductsArgs struct {
	Req *kitex_gen.SearchProductsReq
}

func (p *SearchProductsArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(kitex_gen.SearchProductsReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *SearchProductsArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *SearchProductsArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *SearchProductsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *SearchProductsArgs) Unmarshal(in []byte) error {
	msg := new(kitex_gen.SearchProductsReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SearchProductsArgs_Req_DEFAULT *kitex_gen.SearchProductsReq

func (p *SearchProductsArgs) GetReq() *kitex_gen.SearchProductsReq {
	if !p.IsSetReq() {
		return SearchProductsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SearchProductsArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *SearchProductsArgs) GetFirstArgument() interface{} {
	return p.Req
}

type SearchProductsResult struct {
	Success *kitex_gen.SearchProductsResp
}

var SearchProductsResult_Success_DEFAULT *kitex_gen.SearchProductsResp

func (p *SearchProductsResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(kitex_gen.SearchProductsResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *SearchProductsResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *SearchProductsResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *SearchProductsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *SearchProductsResult) Unmarshal(in []byte) error {
	msg := new(kitex_gen.SearchProductsResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SearchProductsResult) GetSuccess() *kitex_gen.SearchProductsResp {
	if !p.IsSetSuccess() {
		return SearchProductsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SearchProductsResult) SetSuccess(x interface{}) {
	p.Success = x.(*kitex_gen.SearchProductsResp)
}

func (p *SearchProductsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SearchProductsResult) GetResult() interface{} {
	return p.Success
}

func createProductHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(kitex_gen.CreateProductReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(kitex_gen.ProductCatalogService).CreateProduct(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *CreateProductArgs:
		success, err := handler.(kitex_gen.ProductCatalogService).CreateProduct(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateProductResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newCreateProductArgs() interface{} {
	return &CreateProductArgs{}
}

func newCreateProductResult() interface{} {
	return &CreateProductResult{}
}

type CreateProductArgs struct {
	Req *kitex_gen.CreateProductReq
}

func (p *CreateProductArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(kitex_gen.CreateProductReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CreateProductArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CreateProductArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CreateProductArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *CreateProductArgs) Unmarshal(in []byte) error {
	msg := new(kitex_gen.CreateProductReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateProductArgs_Req_DEFAULT *kitex_gen.CreateProductReq

func (p *CreateProductArgs) GetReq() *kitex_gen.CreateProductReq {
	if !p.IsSetReq() {
		return CreateProductArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateProductArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CreateProductArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CreateProductResult struct {
	Success *kitex_gen.CreateProductResp
}

var CreateProductResult_Success_DEFAULT *kitex_gen.CreateProductResp

func (p *CreateProductResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(kitex_gen.CreateProductResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CreateProductResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CreateProductResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CreateProductResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *CreateProductResult) Unmarshal(in []byte) error {
	msg := new(kitex_gen.CreateProductResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateProductResult) GetSuccess() *kitex_gen.CreateProductResp {
	if !p.IsSetSuccess() {
		return CreateProductResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateProductResult) SetSuccess(x interface{}) {
	p.Success = x.(*kitex_gen.CreateProductResp)
}

func (p *CreateProductResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CreateProductResult) GetResult() interface{} {
	return p.Success
}

func updateProductHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(kitex_gen.UpdateProductReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(kitex_gen.ProductCatalogService).UpdateProduct(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *UpdateProductArgs:
		success, err := handler.(kitex_gen.ProductCatalogService).UpdateProduct(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdateProductResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newUpdateProductArgs() interface{} {
	return &UpdateProductArgs{}
}

func newUpdateProductResult() interface{} {
	return &UpdateProductResult{}
}

type UpdateProductArgs struct {
	Req *kitex_gen.UpdateProductReq
}

func (p *UpdateProductArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(kitex_gen.UpdateProductReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UpdateProductArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UpdateProductArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UpdateProductArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *UpdateProductArgs) Unmarshal(in []byte) error {
	msg := new(kitex_gen.UpdateProductReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdateProductArgs_Req_DEFAULT *kitex_gen.UpdateProductReq

func (p *UpdateProductArgs) GetReq() *kitex_gen.UpdateProductReq {
	if !p.IsSetReq() {
		return UpdateProductArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateProductArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UpdateProductArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UpdateProductResult struct {
	Success *kitex_gen.UpdateProductResp
}

var UpdateProductResult_Success_DEFAULT *kitex_gen.UpdateProductResp

func (p *UpdateProductResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(kitex_gen.UpdateProductResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UpdateProductResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UpdateProductResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UpdateProductResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *UpdateProductResult) Unmarshal(in []byte) error {
	msg := new(kitex_gen.UpdateProductResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateProductResult) GetSuccess() *kitex_gen.UpdateProductResp {
	if !p.IsSetSuccess() {
		return UpdateProductResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateProductResult) SetSuccess(x interface{}) {
	p.Success = x.(*kitex_gen.UpdateProductResp)
}

func (p *UpdateProductResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateProductResult) GetResult() interface{} {
	return p.Success
}

func deleteProductHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(kitex_gen.DeleteProductReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(kitex_gen.ProductCatalogService).DeleteProduct(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *DeleteProductArgs:
		success, err := handler.(kitex_gen.ProductCatalogService).DeleteProduct(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeleteProductResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newDeleteProductArgs() interface{} {
	return &DeleteProductArgs{}
}

func newDeleteProductResult() interface{} {
	return &DeleteProductResult{}
}

type DeleteProductArgs struct {
	Req *kitex_gen.DeleteProductReq
}

func (p *DeleteProductArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(kitex_gen.DeleteProductReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DeleteProductArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DeleteProductArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DeleteProductArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *DeleteProductArgs) Unmarshal(in []byte) error {
	msg := new(kitex_gen.DeleteProductReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeleteProductArgs_Req_DEFAULT *kitex_gen.DeleteProductReq

func (p *DeleteProductArgs) GetReq() *kitex_gen.DeleteProductReq {
	if !p.IsSetReq() {
		return DeleteProductArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteProductArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *DeleteProductArgs) GetFirstArgument() interface{} {
	return p.Req
}

type DeleteProductResult struct {
	Success *kitex_gen.DeleteProductResp
}

var DeleteProductResult_Success_DEFAULT *kitex_gen.DeleteProductResp

func (p *DeleteProductResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(kitex_gen.DeleteProductResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DeleteProductResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DeleteProductResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DeleteProductResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *DeleteProductResult) Unmarshal(in []byte) error {
	msg := new(kitex_gen.DeleteProductResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteProductResult) GetSuccess() *kitex_gen.DeleteProductResp {
	if !p.IsSetSuccess() {
		return DeleteProductResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteProductResult) SetSuccess(x interface{}) {
	p.Success = x.(*kitex_gen.DeleteProductResp)
}

func (p *DeleteProductResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteProductResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) ListProducts(ctx context.Context, Req *kitex_gen.ListProductsReq) (r *kitex_gen.ListProductsResp, err error) {
	var _args ListProductsArgs
	_args.Req = Req
	var _result ListProductsResult
	if err = p.c.Call(ctx, "ListProducts", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetProduct(ctx context.Context, Req *kitex_gen.GetProductReq) (r *kitex_gen.GetProductResp, err error) {
	var _args GetProductArgs
	_args.Req = Req
	var _result GetProductResult
	if err = p.c.Call(ctx, "GetProduct", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SearchProducts(ctx context.Context, Req *kitex_gen.SearchProductsReq) (r *kitex_gen.SearchProductsResp, err error) {
	var _args SearchProductsArgs
	_args.Req = Req
	var _result SearchProductsResult
	if err = p.c.Call(ctx, "SearchProducts", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateProduct(ctx context.Context, Req *kitex_gen.CreateProductReq) (r *kitex_gen.CreateProductResp, err error) {
	var _args CreateProductArgs
	_args.Req = Req
	var _result CreateProductResult
	if err = p.c.Call(ctx, "CreateProduct", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateProduct(ctx context.Context, Req *kitex_gen.UpdateProductReq) (r *kitex_gen.UpdateProductResp, err error) {
	var _args UpdateProductArgs
	_args.Req = Req
	var _result UpdateProductResult
	if err = p.c.Call(ctx, "UpdateProduct", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteProduct(ctx context.Context, Req *kitex_gen.DeleteProductReq) (r *kitex_gen.DeleteProductResp, err error) {
	var _args DeleteProductArgs
	_args.Req = Req
	var _result DeleteProductResult
	if err = p.c.Call(ctx, "DeleteProduct", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
