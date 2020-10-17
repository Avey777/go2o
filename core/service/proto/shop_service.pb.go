// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shop_service.proto

package proto // import "."

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TurnShopRequest struct {
	ShopId               int64    `protobuf:"zigzag64,1,opt,name=shopId,proto3" json:"shopId,omitempty"`
	On                   bool     `protobuf:"varint,2,opt,name=on,proto3" json:"on,omitempty"`
	Reason               string   `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TurnShopRequest) Reset()         { *m = TurnShopRequest{} }
func (m *TurnShopRequest) String() string { return proto.CompactTextString(m) }
func (*TurnShopRequest) ProtoMessage()    {}
func (*TurnShopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_shop_service_88dc535edf9bd0f5, []int{0}
}
func (m *TurnShopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TurnShopRequest.Unmarshal(m, b)
}
func (m *TurnShopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TurnShopRequest.Marshal(b, m, deterministic)
}
func (dst *TurnShopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TurnShopRequest.Merge(dst, src)
}
func (m *TurnShopRequest) XXX_Size() int {
	return xxx_messageInfo_TurnShopRequest.Size(m)
}
func (m *TurnShopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TurnShopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TurnShopRequest proto.InternalMessageInfo

func (m *TurnShopRequest) GetShopId() int64 {
	if m != nil {
		return m.ShopId
	}
	return 0
}

func (m *TurnShopRequest) GetOn() bool {
	if m != nil {
		return m.On
	}
	return false
}

func (m *TurnShopRequest) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func init() {
	proto.RegisterType((*TurnShopRequest)(nil), "TurnShopRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ShopServiceClient is the client API for ShopService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShopServiceClient interface {
	// * 获取店铺,shopId
	GetShop(ctx context.Context, in *ShopId, opts ...grpc.CallOption) (*SShop, error)
	// rpc GetVendorShop_ (Int64) returns (SShop) {}
	// 检查商户是否开通店铺
	CheckMerchantShopState(ctx context.Context, in *MerchantId, opts ...grpc.CallOption) (*CheckShopResponse, error)
	// * 获取门店,storeId
	GetStore(ctx context.Context, in *StoreId, opts ...grpc.CallOption) (*SStore, error)
	// * 根据主机头获取店铺编号,host
	QueryShopByHost(ctx context.Context, in *String, opts ...grpc.CallOption) (*Int64, error)
	// 获取门店
	// rpc GetOfflineShop(1:sint32 shopId)returns(Shop)
	// 打开或关闭商店
	TurnShop(ctx context.Context, in *TurnShopRequest, opts ...grpc.CallOption) (*Result, error)
	// 保存门店
	SaveShop(ctx context.Context, in *SShop, opts ...grpc.CallOption) (*Result, error)
	// 保存门店
	SaveOfflineShop(ctx context.Context, in *SStore, opts ...grpc.CallOption) (*Result, error)
	// 删除商店
	DeleteStore(ctx context.Context, in *StoreId, opts ...grpc.CallOption) (*Result, error)
}

type shopServiceClient struct {
	cc *grpc.ClientConn
}

func NewShopServiceClient(cc *grpc.ClientConn) ShopServiceClient {
	return &shopServiceClient{cc}
}

func (c *shopServiceClient) GetShop(ctx context.Context, in *ShopId, opts ...grpc.CallOption) (*SShop, error) {
	out := new(SShop)
	err := c.cc.Invoke(ctx, "/ShopService/GetShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) CheckMerchantShopState(ctx context.Context, in *MerchantId, opts ...grpc.CallOption) (*CheckShopResponse, error) {
	out := new(CheckShopResponse)
	err := c.cc.Invoke(ctx, "/ShopService/CheckMerchantShopState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) GetStore(ctx context.Context, in *StoreId, opts ...grpc.CallOption) (*SStore, error) {
	out := new(SStore)
	err := c.cc.Invoke(ctx, "/ShopService/GetStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) QueryShopByHost(ctx context.Context, in *String, opts ...grpc.CallOption) (*Int64, error) {
	out := new(Int64)
	err := c.cc.Invoke(ctx, "/ShopService/QueryShopByHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) TurnShop(ctx context.Context, in *TurnShopRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ShopService/TurnShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) SaveShop(ctx context.Context, in *SShop, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ShopService/SaveShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) SaveOfflineShop(ctx context.Context, in *SStore, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ShopService/SaveOfflineShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) DeleteStore(ctx context.Context, in *StoreId, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ShopService/DeleteStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShopServiceServer is the server API for ShopService service.
type ShopServiceServer interface {
	// * 获取店铺,shopId
	GetShop(context.Context, *ShopId) (*SShop, error)
	// rpc GetVendorShop_ (Int64) returns (SShop) {}
	// 检查商户是否开通店铺
	CheckMerchantShopState(context.Context, *MerchantId) (*CheckShopResponse, error)
	// * 获取门店,storeId
	GetStore(context.Context, *StoreId) (*SStore, error)
	// * 根据主机头获取店铺编号,host
	QueryShopByHost(context.Context, *String) (*Int64, error)
	// 获取门店
	// rpc GetOfflineShop(1:sint32 shopId)returns(Shop)
	// 打开或关闭商店
	TurnShop(context.Context, *TurnShopRequest) (*Result, error)
	// 保存门店
	SaveShop(context.Context, *SShop) (*Result, error)
	// 保存门店
	SaveOfflineShop(context.Context, *SStore) (*Result, error)
	// 删除商店
	DeleteStore(context.Context, *StoreId) (*Result, error)
}

func RegisterShopServiceServer(s *grpc.Server, srv ShopServiceServer) {
	s.RegisterService(&_ShopService_serviceDesc, srv)
}

func _ShopService_GetShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShopId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).GetShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/GetShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).GetShop(ctx, req.(*ShopId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_CheckMerchantShopState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MerchantId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).CheckMerchantShopState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/CheckMerchantShopState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).CheckMerchantShopState(ctx, req.(*MerchantId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_GetStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).GetStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/GetStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).GetStore(ctx, req.(*StoreId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_QueryShopByHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).QueryShopByHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/QueryShopByHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).QueryShopByHost(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_TurnShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TurnShopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).TurnShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/TurnShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).TurnShop(ctx, req.(*TurnShopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_SaveShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SShop)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).SaveShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/SaveShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).SaveShop(ctx, req.(*SShop))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_SaveOfflineShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SStore)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).SaveOfflineShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/SaveOfflineShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).SaveOfflineShop(ctx, req.(*SStore))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_DeleteStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).DeleteStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/DeleteStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).DeleteStore(ctx, req.(*StoreId))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShopService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ShopService",
	HandlerType: (*ShopServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetShop",
			Handler:    _ShopService_GetShop_Handler,
		},
		{
			MethodName: "CheckMerchantShopState",
			Handler:    _ShopService_CheckMerchantShopState_Handler,
		},
		{
			MethodName: "GetStore",
			Handler:    _ShopService_GetStore_Handler,
		},
		{
			MethodName: "QueryShopByHost",
			Handler:    _ShopService_QueryShopByHost_Handler,
		},
		{
			MethodName: "TurnShop",
			Handler:    _ShopService_TurnShop_Handler,
		},
		{
			MethodName: "SaveShop",
			Handler:    _ShopService_SaveShop_Handler,
		},
		{
			MethodName: "SaveOfflineShop",
			Handler:    _ShopService_SaveOfflineShop_Handler,
		},
		{
			MethodName: "DeleteStore",
			Handler:    _ShopService_DeleteStore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shop_service.proto",
}

func init() { proto.RegisterFile("shop_service.proto", fileDescriptor_shop_service_88dc535edf9bd0f5) }

var fileDescriptor_shop_service_88dc535edf9bd0f5 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x5f, 0x4b, 0xf3, 0x30,
	0x14, 0xc6, 0xdb, 0xbe, 0xd0, 0x76, 0xd9, 0x8b, 0x93, 0x73, 0x31, 0x46, 0x51, 0x1c, 0x41, 0x71,
	0x57, 0x11, 0x54, 0xbc, 0xd9, 0xdd, 0x14, 0xb4, 0x17, 0x22, 0x6b, 0xbc, 0xf2, 0x46, 0xba, 0xed,
	0x6c, 0x1d, 0xd6, 0xa4, 0x26, 0xe9, 0x60, 0x5f, 0xcf, 0x4f, 0x26, 0x49, 0x5b, 0x18, 0x7a, 0x75,
	0xfe, 0xfd, 0xf2, 0xf0, 0x9c, 0x13, 0x02, 0xba, 0x90, 0xd5, 0xbb, 0x46, 0xb5, 0xdb, 0x2e, 0x91,
	0x55, 0x4a, 0x1a, 0x99, 0xfc, 0xdf, 0x94, 0x72, 0x91, 0x97, 0x6d, 0x05, 0x9f, 0xa8, 0x75, 0xbe,
	0xc1, 0x2b, 0x4b, 0x36, 0x3d, 0x3a, 0x27, 0x83, 0xd7, 0x5a, 0x09, 0x5e, 0xc8, 0x2a, 0xc3, 0xaf,
	0x1a, 0xb5, 0x81, 0x21, 0x09, 0x2d, 0x90, 0xae, 0x46, 0xfe, 0xd8, 0x9f, 0x40, 0xd6, 0x56, 0x70,
	0x44, 0x02, 0x29, 0x46, 0xc1, 0xd8, 0x9f, 0xc4, 0x59, 0x20, 0x85, 0xe5, 0x14, 0xe6, 0x5a, 0x8a,
	0xd1, 0xbf, 0xb1, 0x3f, 0xe9, 0x65, 0x6d, 0x75, 0xfd, 0x1d, 0x90, 0xbe, 0xd5, 0xe3, 0x8d, 0x15,
	0x38, 0x21, 0xd1, 0x23, 0x1a, 0xdb, 0x81, 0x88, 0x71, 0xa7, 0x95, 0x84, 0x8c, 0xdb, 0x8c, 0x7a,
	0x30, 0x25, 0xc3, 0xfb, 0x02, 0x97, 0x1f, 0xcf, 0xa8, 0x96, 0x45, 0x2e, 0x1c, 0xc7, 0x4d, 0x6e,
	0x10, 0xfa, 0xac, 0xeb, 0xa5, 0xab, 0x04, 0x98, 0xa3, 0x1a, 0x9f, 0xba, 0x92, 0x42, 0x23, 0xf5,
	0xe0, 0x8c, 0xc4, 0x56, 0xda, 0x48, 0x85, 0x10, 0x33, 0x17, 0xd3, 0x55, 0x12, 0x31, 0xee, 0x52,
	0xea, 0xc1, 0x39, 0x19, 0xcc, 0x6b, 0x54, 0x7b, 0xfb, 0x6e, 0xb6, 0x7f, 0x92, 0xda, 0x58, 0x0f,
	0x46, 0x6d, 0xc5, 0x26, 0x09, 0x59, 0x2a, 0xcc, 0xdd, 0x2d, 0xf5, 0xe0, 0x92, 0xc4, 0xdd, 0x11,
	0xe0, 0x98, 0xfd, 0xba, 0x47, 0x12, 0xb1, 0x0c, 0x75, 0x5d, 0x1a, 0xea, 0xc1, 0x29, 0x89, 0x79,
	0xbe, 0x43, 0x07, 0xb6, 0x2b, 0x1c, 0x8e, 0x2f, 0xc8, 0xc0, 0x8e, 0x5f, 0xd6, 0xeb, 0x72, 0x2b,
	0xb0, 0xdb, 0xd8, 0x79, 0x39, 0xc4, 0x28, 0xe9, 0x3f, 0x60, 0x89, 0x06, 0xff, 0x1a, 0xef, 0x98,
	0x59, 0xef, 0x2d, 0x62, 0x53, 0xf7, 0x45, 0x8b, 0xd0, 0x85, 0x9b, 0x9f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xa7, 0xa5, 0xc3, 0x2f, 0xe1, 0x01, 0x00, 0x00,
}
