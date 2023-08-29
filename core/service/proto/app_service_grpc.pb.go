// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: app_service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AppService_SaveProd_FullMethodName      = "/AppService/SaveProd"
	AppService_SaveVersion_FullMethodName   = "/AppService/SaveVersion"
	AppService_GetProd_FullMethodName       = "/AppService/GetProd"
	AppService_GetVersion_FullMethodName    = "/AppService/GetVersion"
	AppService_DeleteProd_FullMethodName    = "/AppService/DeleteProd"
	AppService_DeleteVersion_FullMethodName = "/AppService/DeleteVersion"
	AppService_CheckVersion_FullMethodName  = "/AppService/CheckVersion"
)

// AppServiceClient is the client API for AppService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppServiceClient interface {
	// 保存APP产品
	SaveProd(ctx context.Context, in *AppProdRequest, opts ...grpc.CallOption) (*Result, error)
	// 保存版本
	SaveVersion(ctx context.Context, in *AppVersionRequest, opts ...grpc.CallOption) (*Result, error)
	// 获取产品信息
	GetProd(ctx context.Context, in *AppId, opts ...grpc.CallOption) (*SAppProd, error)
	// 获取版本
	GetVersion(ctx context.Context, in *AppVersionId, opts ...grpc.CallOption) (*SAppVersion, error)
	// 删除产品
	DeleteProd(ctx context.Context, in *AppId, opts ...grpc.CallOption) (*Result, error)
	// 删除版本
	DeleteVersion(ctx context.Context, in *AppVersionId, opts ...grpc.CallOption) (*Result, error)
	// 检测版本更新
	CheckVersion(ctx context.Context, in *CheckVersionRequest, opts ...grpc.CallOption) (*CheckVersionResponse, error)
}

type appServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAppServiceClient(cc grpc.ClientConnInterface) AppServiceClient {
	return &appServiceClient{cc}
}

func (c *appServiceClient) SaveProd(ctx context.Context, in *AppProdRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, AppService_SaveProd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) SaveVersion(ctx context.Context, in *AppVersionRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, AppService_SaveVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) GetProd(ctx context.Context, in *AppId, opts ...grpc.CallOption) (*SAppProd, error) {
	out := new(SAppProd)
	err := c.cc.Invoke(ctx, AppService_GetProd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) GetVersion(ctx context.Context, in *AppVersionId, opts ...grpc.CallOption) (*SAppVersion, error) {
	out := new(SAppVersion)
	err := c.cc.Invoke(ctx, AppService_GetVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) DeleteProd(ctx context.Context, in *AppId, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, AppService_DeleteProd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) DeleteVersion(ctx context.Context, in *AppVersionId, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, AppService_DeleteVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) CheckVersion(ctx context.Context, in *CheckVersionRequest, opts ...grpc.CallOption) (*CheckVersionResponse, error) {
	out := new(CheckVersionResponse)
	err := c.cc.Invoke(ctx, AppService_CheckVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppServiceServer is the server API for AppService service.
// All implementations must embed UnimplementedAppServiceServer
// for forward compatibility
type AppServiceServer interface {
	// 保存APP产品
	SaveProd(context.Context, *AppProdRequest) (*Result, error)
	// 保存版本
	SaveVersion(context.Context, *AppVersionRequest) (*Result, error)
	// 获取产品信息
	GetProd(context.Context, *AppId) (*SAppProd, error)
	// 获取版本
	GetVersion(context.Context, *AppVersionId) (*SAppVersion, error)
	// 删除产品
	DeleteProd(context.Context, *AppId) (*Result, error)
	// 删除版本
	DeleteVersion(context.Context, *AppVersionId) (*Result, error)
	// 检测版本更新
	CheckVersion(context.Context, *CheckVersionRequest) (*CheckVersionResponse, error)
	mustEmbedUnimplementedAppServiceServer()
}

// UnimplementedAppServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAppServiceServer struct {
}

func (UnimplementedAppServiceServer) SaveProd(context.Context, *AppProdRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveProd not implemented")
}
func (UnimplementedAppServiceServer) SaveVersion(context.Context, *AppVersionRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveVersion not implemented")
}
func (UnimplementedAppServiceServer) GetProd(context.Context, *AppId) (*SAppProd, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProd not implemented")
}
func (UnimplementedAppServiceServer) GetVersion(context.Context, *AppVersionId) (*SAppVersion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}
func (UnimplementedAppServiceServer) DeleteProd(context.Context, *AppId) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProd not implemented")
}
func (UnimplementedAppServiceServer) DeleteVersion(context.Context, *AppVersionId) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVersion not implemented")
}
func (UnimplementedAppServiceServer) CheckVersion(context.Context, *CheckVersionRequest) (*CheckVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckVersion not implemented")
}
func (UnimplementedAppServiceServer) mustEmbedUnimplementedAppServiceServer() {}

// UnsafeAppServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppServiceServer will
// result in compilation errors.
type UnsafeAppServiceServer interface {
	mustEmbedUnimplementedAppServiceServer()
}

func RegisterAppServiceServer(s grpc.ServiceRegistrar, srv AppServiceServer) {
	s.RegisterService(&AppService_ServiceDesc, srv)
}

func _AppService_SaveProd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppProdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).SaveProd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppService_SaveProd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).SaveProd(ctx, req.(*AppProdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_SaveVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).SaveVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppService_SaveVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).SaveVersion(ctx, req.(*AppVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_GetProd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).GetProd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppService_GetProd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).GetProd(ctx, req.(*AppId))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppVersionId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppService_GetVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).GetVersion(ctx, req.(*AppVersionId))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_DeleteProd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).DeleteProd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppService_DeleteProd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).DeleteProd(ctx, req.(*AppId))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_DeleteVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppVersionId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).DeleteVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppService_DeleteVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).DeleteVersion(ctx, req.(*AppVersionId))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_CheckVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).CheckVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppService_CheckVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).CheckVersion(ctx, req.(*CheckVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AppService_ServiceDesc is the grpc.ServiceDesc for AppService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AppService",
	HandlerType: (*AppServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveProd",
			Handler:    _AppService_SaveProd_Handler,
		},
		{
			MethodName: "SaveVersion",
			Handler:    _AppService_SaveVersion_Handler,
		},
		{
			MethodName: "GetProd",
			Handler:    _AppService_GetProd_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _AppService_GetVersion_Handler,
		},
		{
			MethodName: "DeleteProd",
			Handler:    _AppService_DeleteProd_Handler,
		},
		{
			MethodName: "DeleteVersion",
			Handler:    _AppService_DeleteVersion_Handler,
		},
		{
			MethodName: "CheckVersion",
			Handler:    _AppService_CheckVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app_service.proto",
}
