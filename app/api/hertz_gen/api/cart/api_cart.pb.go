// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.29.0
// source: api/api_cart.proto

package cart

import (
	_ "github.com/A1sca/Douyin-Mall-Go/app/api/hertz_gen/api"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CartItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId uint32 `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty" form:"product_id" query:"product_id"`
	Quantity  int32  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty" form:"quantity" query:"quantity"`
}

func (x *CartItem) Reset() {
	*x = CartItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_cart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItem) ProtoMessage() {}

func (x *CartItem) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_cart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItem.ProtoReflect.Descriptor instead.
func (*CartItem) Descriptor() ([]byte, []int) {
	return file_api_api_cart_proto_rawDescGZIP(), []int{0}
}

func (x *CartItem) GetProductId() uint32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *CartItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type AddItemReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" form:"user_id"`
	Item   *CartItem `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty" form:"order_id"`
}

func (x *AddItemReq) Reset() {
	*x = AddItemReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_cart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddItemReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddItemReq) ProtoMessage() {}

func (x *AddItemReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_cart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddItemReq.ProtoReflect.Descriptor instead.
func (*AddItemReq) Descriptor() ([]byte, []int) {
	return file_api_api_cart_proto_rawDescGZIP(), []int{1}
}

func (x *AddItemReq) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AddItemReq) GetItem() *CartItem {
	if x != nil {
		return x.Item
	}
	return nil
}

type AddItemResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddItemResp) Reset() {
	*x = AddItemResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_cart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddItemResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddItemResp) ProtoMessage() {}

func (x *AddItemResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_cart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddItemResp.ProtoReflect.Descriptor instead.
func (*AddItemResp) Descriptor() ([]byte, []int) {
	return file_api_api_cart_proto_rawDescGZIP(), []int{2}
}

type EmptyCartReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" path:"user_id"`
}

func (x *EmptyCartReq) Reset() {
	*x = EmptyCartReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_cart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyCartReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyCartReq) ProtoMessage() {}

func (x *EmptyCartReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_cart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyCartReq.ProtoReflect.Descriptor instead.
func (*EmptyCartReq) Descriptor() ([]byte, []int) {
	return file_api_api_cart_proto_rawDescGZIP(), []int{3}
}

func (x *EmptyCartReq) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetCartReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" path:"user_id"`
}

func (x *GetCartReq) Reset() {
	*x = GetCartReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_cart_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartReq) ProtoMessage() {}

func (x *GetCartReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_cart_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartReq.ProtoReflect.Descriptor instead.
func (*GetCartReq) Descriptor() ([]byte, []int) {
	return file_api_api_cart_proto_rawDescGZIP(), []int{4}
}

func (x *GetCartReq) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetCartResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cart *Cart `protobuf:"bytes,1,opt,name=cart,proto3" json:"cart,omitempty" form:"cart" query:"cart"`
}

func (x *GetCartResp) Reset() {
	*x = GetCartResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_cart_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartResp) ProtoMessage() {}

func (x *GetCartResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_cart_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartResp.ProtoReflect.Descriptor instead.
func (*GetCartResp) Descriptor() ([]byte, []int) {
	return file_api_api_cart_proto_rawDescGZIP(), []int{5}
}

func (x *GetCartResp) GetCart() *Cart {
	if x != nil {
		return x.Cart
	}
	return nil
}

type Cart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32      `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" form:"user_id" query:"user_id"`
	Items  []*CartItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty" form:"items" query:"items"`
}

func (x *Cart) Reset() {
	*x = Cart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_cart_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cart) ProtoMessage() {}

func (x *Cart) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_cart_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cart.ProtoReflect.Descriptor instead.
func (*Cart) Descriptor() ([]byte, []int) {
	return file_api_api_cart_proto_rawDescGZIP(), []int{6}
}

func (x *Cart) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Cart) GetItems() []*CartItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type EmptyCartResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyCartResp) Reset() {
	*x = EmptyCartResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_cart_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyCartResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyCartResp) ProtoMessage() {}

func (x *EmptyCartResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_cart_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyCartResp.ProtoReflect.Descriptor instead.
func (*EmptyCartResp) Descriptor() ([]byte, []int) {
	return file_api_api_cart_proto_rawDescGZIP(), []int{7}
}

var File_api_api_cart_proto protoreflect.FileDescriptor

var file_api_api_cart_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x1a, 0x0d,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a,
	0x08, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x22, 0x68, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x71, 0x12, 0x24, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x42, 0x0b, 0xe2, 0xbb, 0x18, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72,
	0x74, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x0c, 0xe2, 0xbb, 0x18, 0x08,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x0d,
	0x0a, 0x0b, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x22, 0x34, 0x0a,
	0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x12, 0x24, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0b,
	0xd2, 0xbb, 0x18, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x24, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x42, 0x0b, 0xd2, 0xbb, 0x18, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x31, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x22, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x04, 0x63, 0x61, 0x72, 0x74, 0x22, 0x49, 0x0a, 0x04, 0x43, 0x61,
	0x72, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x43, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x32, 0x85, 0x02, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x41, 0x64, 0x64,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61,
	0x72, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x22, 0x10,
	0xd2, 0xc1, 0x18, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x61, 0x64, 0x64,
	0x12, 0x51, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x12, 0x14, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x19, 0xca, 0xc1, 0x18, 0x15, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x3a, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x12, 0x59, 0x0a, 0x09, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x43, 0x61, 0x72, 0x74,
	0x12, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x61, 0x72, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x1b, 0xd2, 0xc1, 0x18, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2f, 0x3a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x3c,
	0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x31, 0x73,
	0x63, 0x61, 0x2f, 0x44, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2d, 0x4d, 0x61, 0x6c, 0x6c, 0x2d, 0x47,
	0x6f, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x5f,
	0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_api_cart_proto_rawDescOnce sync.Once
	file_api_api_cart_proto_rawDescData = file_api_api_cart_proto_rawDesc
)

func file_api_api_cart_proto_rawDescGZIP() []byte {
	file_api_api_cart_proto_rawDescOnce.Do(func() {
		file_api_api_cart_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_api_cart_proto_rawDescData)
	})
	return file_api_api_cart_proto_rawDescData
}

var file_api_api_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_api_cart_proto_goTypes = []interface{}{
	(*CartItem)(nil),      // 0: api.cart.CartItem
	(*AddItemReq)(nil),    // 1: api.cart.AddItemReq
	(*AddItemResp)(nil),   // 2: api.cart.AddItemResp
	(*EmptyCartReq)(nil),  // 3: api.cart.EmptyCartReq
	(*GetCartReq)(nil),    // 4: api.cart.GetCartReq
	(*GetCartResp)(nil),   // 5: api.cart.GetCartResp
	(*Cart)(nil),          // 6: api.cart.Cart
	(*EmptyCartResp)(nil), // 7: api.cart.EmptyCartResp
}
var file_api_api_cart_proto_depIdxs = []int32{
	0, // 0: api.cart.AddItemReq.item:type_name -> api.cart.CartItem
	6, // 1: api.cart.GetCartResp.cart:type_name -> api.cart.Cart
	0, // 2: api.cart.Cart.items:type_name -> api.cart.CartItem
	1, // 3: api.cart.CartService.AddItem:input_type -> api.cart.AddItemReq
	4, // 4: api.cart.CartService.GetCart:input_type -> api.cart.GetCartReq
	3, // 5: api.cart.CartService.EmptyCart:input_type -> api.cart.EmptyCartReq
	2, // 6: api.cart.CartService.AddItem:output_type -> api.cart.AddItemResp
	5, // 7: api.cart.CartService.GetCart:output_type -> api.cart.GetCartResp
	7, // 8: api.cart.CartService.EmptyCart:output_type -> api.cart.EmptyCartResp
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_api_cart_proto_init() }
func file_api_api_cart_proto_init() {
	if File_api_api_cart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_api_cart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_api_cart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddItemReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_api_cart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddItemResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_api_cart_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyCartReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_api_cart_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_api_cart_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_api_cart_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cart); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_api_cart_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyCartResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_api_cart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_api_cart_proto_goTypes,
		DependencyIndexes: file_api_api_cart_proto_depIdxs,
		MessageInfos:      file_api_api_cart_proto_msgTypes,
	}.Build()
	File_api_api_cart_proto = out.File
	file_api_api_cart_proto_rawDesc = nil
	file_api_api_cart_proto_goTypes = nil
	file_api_api_cart_proto_depIdxs = nil
}
