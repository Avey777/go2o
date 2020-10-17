// Code generated by protoc-gen-go. DO NOT EDIT.
// source: advertisement_service.proto

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdvertisementServiceClient is the client API for AdvertisementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdvertisementServiceClient interface {
	GetAdGroupById(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*SAdGroup, error)
	DelAdGroup(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*Result, error)
	SaveAdGroup(ctx context.Context, in *SAdGroup, opts ...grpc.CallOption) (*Result, error)
	GetGroups(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AdGroupListResponse, error)
	GetPosition(ctx context.Context, in *AdPositionId, opts ...grpc.CallOption) (*SAdPosition, error)
	SaveAdPosition(ctx context.Context, in *SAdPosition, opts ...grpc.CallOption) (*Result, error)
	DelAdPosition(ctx context.Context, in *AdPositionId, opts ...grpc.CallOption) (*Result, error)
	// 设置广告位的默认广告
	SetDefaultAd(ctx context.Context, in *SetDefaultAdRequest, opts ...grpc.CallOption) (*Result, error)
	// 用户投放广告
	SetUserAd(ctx context.Context, in *SetUserAdRequest, opts ...grpc.CallOption) (*Result, error)
	// 获取广告
	GetAdvertisement(ctx context.Context, in *AdIdRequest, opts ...grpc.CallOption) (*SAd, error)
	// 获取广告及广告数据, 用于展示关高
	GetAdAndDataByKey(ctx context.Context, in *AdKeyRequest, opts ...grpc.CallOption) (*SAdDto, error)
	// 获取广告数据传输对象
	GetAdDto_(ctx context.Context, in *AdIdRequest, opts ...grpc.CallOption) (*SAdDto, error)
	// 保存广告,更新时不允许修改类型
	SaveAd(ctx context.Context, in *SaveAdRequest, opts ...grpc.CallOption) (*Result, error)
	// 删除广告
	DeleteAd(ctx context.Context, in *AdIdRequest, opts ...grpc.CallOption) (*Result, error)
	// 保存图片广告
	SaveHyperLinkAd(ctx context.Context, in *SaveLinkAdRequest, opts ...grpc.CallOption) (*Result, error)
	// 保存图片广告
	SaveImagOfAd(ctx context.Context, in *SaveImageAdRequest, opts ...grpc.CallOption) (*Result, error)
	// 获取广告图片
	GetValueAdImage(ctx context.Context, in *ImageAdIdRequest, opts ...grpc.CallOption) (*SImage, error)
	// 删除广告图片
	DelAdImage(ctx context.Context, in *ImageAdIdRequest, opts ...grpc.CallOption) (*Result, error)
}

type advertisementServiceClient struct {
	cc *grpc.ClientConn
}

func NewAdvertisementServiceClient(cc *grpc.ClientConn) AdvertisementServiceClient {
	return &advertisementServiceClient{cc}
}

func (c *advertisementServiceClient) GetAdGroupById(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*SAdGroup, error) {
	out := new(SAdGroup)
	err := c.cc.Invoke(ctx, "/AdvertisementService/GetAdGroupById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertisementServiceClient) DelAdGroup(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/AdvertisementService/DelAdGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertisementServiceClient) SaveAdGroup(ctx context.Context, in *SAdGroup, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/AdvertisementService/SaveAdGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertisementServiceClient) GetGroups(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AdGroupListResponse, error) {
	out := new(AdGroupListResponse)
	err := c.cc.Invoke(ctx, "/AdvertisementService/GetGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertisementServiceClient) GetPosition(ctx context.Context, in *AdPositionId, opts ...grpc.CallOption) (*SAdPosition, error) {
	out := new(SAdPosition)
	err := c.cc.Invoke(ctx, "/AdvertisementService/GetPosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertisementServiceClient) SaveAdPosition(ctx context.Context, in *SAdPosition, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/AdvertisementService/SaveAdPosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertisementServiceClient) DelAdPosition(ctx context.Context, in *AdPositionId, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/AdvertisementService/DelAdPosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertisementServiceClient) SetDefaultAd(ctx context.Context, in *SetDefaultAdRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/AdvertisementService/SetDefaultAd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertisementServiceClient) SetUserAd(ctx context.Context, in *SetUserAdRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/AdvertisementService/SetUserAd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertisementServiceClient) GetAdvertisement(ctx context.Context, in *AdIdRequest, opts ...grpc.CallOption) (*SAd, error) {
	out := new(SAd)
	err := c.cc.Invoke(ctx, "/AdvertisementService/GetAdvertisement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertisementServiceClient) GetAdAndDataByKey(ctx context.Context, in *AdKeyRequest, opts ...grpc.CallOption) (*SAdDto, error) {
	out := new(SAdDto)
	err := c.cc.Invoke(ctx, "/AdvertisementService/GetAdAndDataByKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertisementServiceClient) GetAdDto_(ctx context.Context, in *AdIdRequest, opts ...grpc.CallOption) (*SAdDto, error) {
	out := new(SAdDto)
	err := c.cc.Invoke(ctx, "/AdvertisementService/GetAdDto_", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertisementServiceClient) SaveAd(ctx context.Context, in *SaveAdRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/AdvertisementService/SaveAd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertisementServiceClient) DeleteAd(ctx context.Context, in *AdIdRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/AdvertisementService/DeleteAd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertisementServiceClient) SaveHyperLinkAd(ctx context.Context, in *SaveLinkAdRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/AdvertisementService/SaveHyperLinkAd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertisementServiceClient) SaveImagOfAd(ctx context.Context, in *SaveImageAdRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/AdvertisementService/SaveImagOfAd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertisementServiceClient) GetValueAdImage(ctx context.Context, in *ImageAdIdRequest, opts ...grpc.CallOption) (*SImage, error) {
	out := new(SImage)
	err := c.cc.Invoke(ctx, "/AdvertisementService/GetValueAdImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertisementServiceClient) DelAdImage(ctx context.Context, in *ImageAdIdRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/AdvertisementService/DelAdImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdvertisementServiceServer is the server API for AdvertisementService service.
type AdvertisementServiceServer interface {
	GetAdGroupById(context.Context, *Int64) (*SAdGroup, error)
	DelAdGroup(context.Context, *Int64) (*Result, error)
	SaveAdGroup(context.Context, *SAdGroup) (*Result, error)
	GetGroups(context.Context, *Empty) (*AdGroupListResponse, error)
	GetPosition(context.Context, *AdPositionId) (*SAdPosition, error)
	SaveAdPosition(context.Context, *SAdPosition) (*Result, error)
	DelAdPosition(context.Context, *AdPositionId) (*Result, error)
	// 设置广告位的默认广告
	SetDefaultAd(context.Context, *SetDefaultAdRequest) (*Result, error)
	// 用户投放广告
	SetUserAd(context.Context, *SetUserAdRequest) (*Result, error)
	// 获取广告
	GetAdvertisement(context.Context, *AdIdRequest) (*SAd, error)
	// 获取广告及广告数据, 用于展示关高
	GetAdAndDataByKey(context.Context, *AdKeyRequest) (*SAdDto, error)
	// 获取广告数据传输对象
	GetAdDto_(context.Context, *AdIdRequest) (*SAdDto, error)
	// 保存广告,更新时不允许修改类型
	SaveAd(context.Context, *SaveAdRequest) (*Result, error)
	// 删除广告
	DeleteAd(context.Context, *AdIdRequest) (*Result, error)
	// 保存图片广告
	SaveHyperLinkAd(context.Context, *SaveLinkAdRequest) (*Result, error)
	// 保存图片广告
	SaveImagOfAd(context.Context, *SaveImageAdRequest) (*Result, error)
	// 获取广告图片
	GetValueAdImage(context.Context, *ImageAdIdRequest) (*SImage, error)
	// 删除广告图片
	DelAdImage(context.Context, *ImageAdIdRequest) (*Result, error)
}

func RegisterAdvertisementServiceServer(s *grpc.Server, srv AdvertisementServiceServer) {
	s.RegisterService(&_AdvertisementService_serviceDesc, srv)
}

func _AdvertisementService_GetAdGroupById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int64)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertisementServiceServer).GetAdGroupById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AdvertisementService/GetAdGroupById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertisementServiceServer).GetAdGroupById(ctx, req.(*Int64))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdvertisementService_DelAdGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int64)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertisementServiceServer).DelAdGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AdvertisementService/DelAdGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertisementServiceServer).DelAdGroup(ctx, req.(*Int64))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdvertisementService_SaveAdGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SAdGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertisementServiceServer).SaveAdGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AdvertisementService/SaveAdGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertisementServiceServer).SaveAdGroup(ctx, req.(*SAdGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdvertisementService_GetGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertisementServiceServer).GetGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AdvertisementService/GetGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertisementServiceServer).GetGroups(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdvertisementService_GetPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdPositionId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertisementServiceServer).GetPosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AdvertisementService/GetPosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertisementServiceServer).GetPosition(ctx, req.(*AdPositionId))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdvertisementService_SaveAdPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SAdPosition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertisementServiceServer).SaveAdPosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AdvertisementService/SaveAdPosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertisementServiceServer).SaveAdPosition(ctx, req.(*SAdPosition))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdvertisementService_DelAdPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdPositionId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertisementServiceServer).DelAdPosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AdvertisementService/DelAdPosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertisementServiceServer).DelAdPosition(ctx, req.(*AdPositionId))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdvertisementService_SetDefaultAd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDefaultAdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertisementServiceServer).SetDefaultAd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AdvertisementService/SetDefaultAd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertisementServiceServer).SetDefaultAd(ctx, req.(*SetDefaultAdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdvertisementService_SetUserAd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserAdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertisementServiceServer).SetUserAd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AdvertisementService/SetUserAd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertisementServiceServer).SetUserAd(ctx, req.(*SetUserAdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdvertisementService_GetAdvertisement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertisementServiceServer).GetAdvertisement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AdvertisementService/GetAdvertisement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertisementServiceServer).GetAdvertisement(ctx, req.(*AdIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdvertisementService_GetAdAndDataByKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertisementServiceServer).GetAdAndDataByKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AdvertisementService/GetAdAndDataByKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertisementServiceServer).GetAdAndDataByKey(ctx, req.(*AdKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdvertisementService_GetAdDto__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertisementServiceServer).GetAdDto_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AdvertisementService/GetAdDto_",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertisementServiceServer).GetAdDto_(ctx, req.(*AdIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdvertisementService_SaveAd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveAdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertisementServiceServer).SaveAd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AdvertisementService/SaveAd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertisementServiceServer).SaveAd(ctx, req.(*SaveAdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdvertisementService_DeleteAd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertisementServiceServer).DeleteAd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AdvertisementService/DeleteAd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertisementServiceServer).DeleteAd(ctx, req.(*AdIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdvertisementService_SaveHyperLinkAd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveLinkAdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertisementServiceServer).SaveHyperLinkAd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AdvertisementService/SaveHyperLinkAd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertisementServiceServer).SaveHyperLinkAd(ctx, req.(*SaveLinkAdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdvertisementService_SaveImagOfAd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveImageAdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertisementServiceServer).SaveImagOfAd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AdvertisementService/SaveImagOfAd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertisementServiceServer).SaveImagOfAd(ctx, req.(*SaveImageAdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdvertisementService_GetValueAdImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageAdIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertisementServiceServer).GetValueAdImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AdvertisementService/GetValueAdImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertisementServiceServer).GetValueAdImage(ctx, req.(*ImageAdIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdvertisementService_DelAdImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageAdIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertisementServiceServer).DelAdImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AdvertisementService/DelAdImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertisementServiceServer).DelAdImage(ctx, req.(*ImageAdIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdvertisementService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AdvertisementService",
	HandlerType: (*AdvertisementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdGroupById",
			Handler:    _AdvertisementService_GetAdGroupById_Handler,
		},
		{
			MethodName: "DelAdGroup",
			Handler:    _AdvertisementService_DelAdGroup_Handler,
		},
		{
			MethodName: "SaveAdGroup",
			Handler:    _AdvertisementService_SaveAdGroup_Handler,
		},
		{
			MethodName: "GetGroups",
			Handler:    _AdvertisementService_GetGroups_Handler,
		},
		{
			MethodName: "GetPosition",
			Handler:    _AdvertisementService_GetPosition_Handler,
		},
		{
			MethodName: "SaveAdPosition",
			Handler:    _AdvertisementService_SaveAdPosition_Handler,
		},
		{
			MethodName: "DelAdPosition",
			Handler:    _AdvertisementService_DelAdPosition_Handler,
		},
		{
			MethodName: "SetDefaultAd",
			Handler:    _AdvertisementService_SetDefaultAd_Handler,
		},
		{
			MethodName: "SetUserAd",
			Handler:    _AdvertisementService_SetUserAd_Handler,
		},
		{
			MethodName: "GetAdvertisement",
			Handler:    _AdvertisementService_GetAdvertisement_Handler,
		},
		{
			MethodName: "GetAdAndDataByKey",
			Handler:    _AdvertisementService_GetAdAndDataByKey_Handler,
		},
		{
			MethodName: "GetAdDto_",
			Handler:    _AdvertisementService_GetAdDto__Handler,
		},
		{
			MethodName: "SaveAd",
			Handler:    _AdvertisementService_SaveAd_Handler,
		},
		{
			MethodName: "DeleteAd",
			Handler:    _AdvertisementService_DeleteAd_Handler,
		},
		{
			MethodName: "SaveHyperLinkAd",
			Handler:    _AdvertisementService_SaveHyperLinkAd_Handler,
		},
		{
			MethodName: "SaveImagOfAd",
			Handler:    _AdvertisementService_SaveImagOfAd_Handler,
		},
		{
			MethodName: "GetValueAdImage",
			Handler:    _AdvertisementService_GetValueAdImage_Handler,
		},
		{
			MethodName: "DelAdImage",
			Handler:    _AdvertisementService_DelAdImage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "advertisement_service.proto",
}

func init() {
	proto.RegisterFile("advertisement_service.proto", fileDescriptor_advertisement_service_885a5ba5bdf16483)
}

var fileDescriptor_advertisement_service_885a5ba5bdf16483 = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x4b, 0x6f, 0xd3, 0x40,
	0x10, 0xc7, 0x7d, 0x40, 0x2e, 0x99, 0xba, 0x2d, 0x19, 0x72, 0x4a, 0x0f, 0x48, 0x0e, 0x48, 0x94,
	0xc7, 0x06, 0x01, 0xe2, 0xc2, 0x69, 0x23, 0x23, 0x63, 0xb5, 0x12, 0x28, 0x16, 0x1c, 0xb8, 0x54,
	0x5b, 0x76, 0x1a, 0x59, 0x38, 0xb6, 0xf1, 0x8e, 0x23, 0xf9, 0x9b, 0xf0, 0x71, 0x91, 0xd7, 0x8f,
	0x3a, 0x55, 0xe0, 0xe4, 0xfd, 0xcf, 0xfe, 0xe6, 0x3f, 0x0f, 0xdb, 0x70, 0xae, 0xf4, 0x8e, 0x4a,
	0x4e, 0x0c, 0x6d, 0x29, 0xe3, 0x6b, 0x43, 0xe5, 0x2e, 0xf9, 0x49, 0xa2, 0x28, 0x73, 0xce, 0xe7,
	0xde, 0x26, 0xcd, 0x6f, 0x54, 0xda, 0xa9, 0xf3, 0x2d, 0x19, 0xa3, 0x36, 0xb4, 0xdc, 0x4b, 0x69,
	0x2f, 0xdf, 0xfe, 0x71, 0x61, 0x26, 0xc7, 0xf1, 0xb8, 0x75, 0xc2, 0x67, 0x70, 0x1a, 0x12, 0x4b,
	0x1d, 0x96, 0x79, 0x55, 0xac, 0xea, 0x48, 0xa3, 0x2b, 0xa2, 0x8c, 0x3f, 0xbc, 0x9f, 0x4f, 0x44,
	0xdc, 0x85, 0x7d, 0x07, 0x9f, 0x00, 0x04, 0x94, 0x76, 0x7a, 0x40, 0x8e, 0xc4, 0x9a, 0x4c, 0x95,
	0xb2, 0xef, 0xe0, 0x02, 0x8e, 0x63, 0xb5, 0xa3, 0x9e, 0xb8, 0x4b, 0x1e, 0x43, 0x2f, 0x61, 0x12,
	0x12, 0xdb, 0xb0, 0x41, 0x57, 0x7c, 0xda, 0x16, 0x5c, 0xcf, 0x67, 0xa2, 0x23, 0xaf, 0x12, 0xc3,
	0x6b, 0x32, 0x45, 0x9e, 0x19, 0xf2, 0x1d, 0x7c, 0x05, 0xc7, 0x21, 0xf1, 0xd7, 0xdc, 0x24, 0x9c,
	0xe4, 0x19, 0x9e, 0x08, 0xa9, 0x7b, 0x11, 0xe9, 0xb9, 0xd7, 0x14, 0xe8, 0xb5, 0xef, 0xe0, 0x05,
	0x9c, 0xb6, 0xf5, 0x87, 0x84, 0x3d, 0x62, 0xdc, 0xc5, 0x05, 0x9c, 0xd8, 0x59, 0xfe, 0x65, 0x3d,
	0x42, 0x97, 0xe0, 0xc5, 0xc4, 0x01, 0xdd, 0xaa, 0x2a, 0x65, 0xa9, 0x71, 0x26, 0xc6, 0x72, 0x4d,
	0xbf, 0x2b, 0x32, 0xbc, 0xef, 0x3d, 0x89, 0x89, 0xbf, 0x19, 0x2a, 0xa5, 0xc6, 0xa9, 0x18, 0xce,
	0x07, 0xd0, 0xe7, 0xf0, 0xc8, 0x6e, 0x7e, 0xf4, 0x52, 0xd0, 0x13, 0x52, 0x47, 0x03, 0xfc, 0xa0,
	0x99, 0xc0, 0x77, 0xf0, 0x35, 0x4c, 0x2d, 0x29, 0x33, 0x1d, 0x28, 0x56, 0xab, 0xfa, 0x92, 0x6a,
	0xdb, 0xf4, 0x25, 0xd5, 0x77, 0xc6, 0xb1, 0xd4, 0x01, 0xe7, 0xbe, 0x83, 0x4f, 0xed, 0x96, 0xad,
	0xba, 0xbe, 0xe7, 0x38, 0xa2, 0x16, 0xe0, 0xb6, 0x0b, 0xc3, 0x53, 0xd1, 0x1e, 0x0e, 0xf4, 0xb8,
	0x80, 0x87, 0x01, 0xa5, 0xc4, 0x0d, 0x76, 0xdf, 0x69, 0x80, 0xde, 0xc0, 0x59, 0x63, 0xf0, 0xb9,
	0x2e, 0xa8, 0xbc, 0x4a, 0xb2, 0x5f, 0x52, 0x23, 0x5a, 0xcb, 0x56, 0x1c, 0xc8, 0x10, 0xe0, 0x35,
	0xf7, 0xd1, 0x56, 0x6d, 0xbe, 0xdc, 0x4a, 0x8d, 0x8f, 0x45, 0x2f, 0x0f, 0xb7, 0xb1, 0x84, 0xb3,
	0x90, 0xf8, 0xbb, 0x4a, 0x2b, 0x92, 0xda, 0x62, 0x38, 0x15, 0x1d, 0xbe, 0x37, 0x9c, 0x8d, 0xf9,
	0x0e, 0xbe, 0xe8, 0x3e, 0xd7, 0xff, 0xb0, 0xbd, 0xf9, 0x6a, 0xf2, 0xe3, 0x48, 0x7c, 0xb4, 0x7f,
	0xc9, 0x8d, 0x6b, 0x1f, 0xef, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x84, 0x0b, 0x37, 0x76,
	0x03, 0x00, 0x00,
}
