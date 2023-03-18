// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: order_service.proto

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

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	// 提交订单
	SubmitOrder(ctx context.Context, in *SubmitOrderRequest, opts ...grpc.CallOption) (*OrderSubmitResponse, error)
	// 预生成订单
	PrepareOrder(ctx context.Context, in *PrepareOrderRequest, opts ...grpc.CallOption) (*PrepareOrderResponse, error)
	// 获取订单信息
	GetParentOrder(ctx context.Context, in *OrderNoV2, opts ...grpc.CallOption) (*SParentOrder, error)
	// 获取子订单,orderId
	GetOrder(ctx context.Context, in *OrderNoV2, opts ...grpc.CallOption) (*SSingleOrder, error)
	// 交易单现金支付,orderId
	TradeOrderCashPay(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*Result, error)
	// 上传交易单发票
	TradeOrderUpdateTicket(ctx context.Context, in *TradeOrderTicketRequest, opts ...grpc.CallOption) (*Result, error)
	// 预生成订单，使用优惠券
	PrepareOrderWithCoupon_(ctx context.Context, in *PrepareOrderRequest, opts ...grpc.CallOption) (*StringMap, error)
	// 取消订单
	CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*Result, error)
	// 确定订单
	ConfirmOrder(ctx context.Context, in *OrderNo, opts ...grpc.CallOption) (*Result, error)
	// 更改收货地址
	ChangeShipmentAddress(ctx context.Context, in *ChangeOrderAddressRequest, opts ...grpc.CallOption) (*Result, error)
	// 备货完成
	PickUp(ctx context.Context, in *OrderNo, opts ...grpc.CallOption) (*Result, error)
	// 订单发货,并记录配送服务商编号及单号
	Ship(ctx context.Context, in *OrderShipmentRequest, opts ...grpc.CallOption) (*Result, error)
	// 买家收货
	BuyerReceived(ctx context.Context, in *OrderNo, opts ...grpc.CallOption) (*Result, error)
	// 删除订单
	Forbid(ctx context.Context, in *OrderNo, opts ...grpc.CallOption) (*Result, error)
	// 获取订单日志
	LogBytes(ctx context.Context, in *OrderNo, opts ...grpc.CallOption) (*String, error)
	// * 获取订单返利列表
	QueryRebateListList(ctx context.Context, in *QueryRebateListRequest, opts ...grpc.CallOption) (*QueryRebateListResponse, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) SubmitOrder(ctx context.Context, in *SubmitOrderRequest, opts ...grpc.CallOption) (*OrderSubmitResponse, error) {
	out := new(OrderSubmitResponse)
	err := c.cc.Invoke(ctx, "/OrderService/SubmitOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) PrepareOrder(ctx context.Context, in *PrepareOrderRequest, opts ...grpc.CallOption) (*PrepareOrderResponse, error) {
	out := new(PrepareOrderResponse)
	err := c.cc.Invoke(ctx, "/OrderService/PrepareOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetParentOrder(ctx context.Context, in *OrderNoV2, opts ...grpc.CallOption) (*SParentOrder, error) {
	out := new(SParentOrder)
	err := c.cc.Invoke(ctx, "/OrderService/GetParentOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrder(ctx context.Context, in *OrderNoV2, opts ...grpc.CallOption) (*SSingleOrder, error) {
	out := new(SSingleOrder)
	err := c.cc.Invoke(ctx, "/OrderService/GetOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) TradeOrderCashPay(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/OrderService/TradeOrderCashPay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) TradeOrderUpdateTicket(ctx context.Context, in *TradeOrderTicketRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/OrderService/TradeOrderUpdateTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) PrepareOrderWithCoupon_(ctx context.Context, in *PrepareOrderRequest, opts ...grpc.CallOption) (*StringMap, error) {
	out := new(StringMap)
	err := c.cc.Invoke(ctx, "/OrderService/PrepareOrderWithCoupon_", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/OrderService/CancelOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) ConfirmOrder(ctx context.Context, in *OrderNo, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/OrderService/ConfirmOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) ChangeShipmentAddress(ctx context.Context, in *ChangeOrderAddressRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/OrderService/ChangeShipmentAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) PickUp(ctx context.Context, in *OrderNo, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/OrderService/PickUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) Ship(ctx context.Context, in *OrderShipmentRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/OrderService/Ship", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) BuyerReceived(ctx context.Context, in *OrderNo, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/OrderService/BuyerReceived", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) Forbid(ctx context.Context, in *OrderNo, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/OrderService/forbid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) LogBytes(ctx context.Context, in *OrderNo, opts ...grpc.CallOption) (*String, error) {
	out := new(String)
	err := c.cc.Invoke(ctx, "/OrderService/LogBytes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) QueryRebateListList(ctx context.Context, in *QueryRebateListRequest, opts ...grpc.CallOption) (*QueryRebateListResponse, error) {
	out := new(QueryRebateListResponse)
	err := c.cc.Invoke(ctx, "/OrderService/QueryRebateListList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations must embed UnimplementedOrderServiceServer
// for forward compatibility
type OrderServiceServer interface {
	// 提交订单
	SubmitOrder(context.Context, *SubmitOrderRequest) (*OrderSubmitResponse, error)
	// 预生成订单
	PrepareOrder(context.Context, *PrepareOrderRequest) (*PrepareOrderResponse, error)
	// 获取订单信息
	GetParentOrder(context.Context, *OrderNoV2) (*SParentOrder, error)
	// 获取子订单,orderId
	GetOrder(context.Context, *OrderNoV2) (*SSingleOrder, error)
	// 交易单现金支付,orderId
	TradeOrderCashPay(context.Context, *Int64) (*Result, error)
	// 上传交易单发票
	TradeOrderUpdateTicket(context.Context, *TradeOrderTicketRequest) (*Result, error)
	// 预生成订单，使用优惠券
	PrepareOrderWithCoupon_(context.Context, *PrepareOrderRequest) (*StringMap, error)
	// 取消订单
	CancelOrder(context.Context, *CancelOrderRequest) (*Result, error)
	// 确定订单
	ConfirmOrder(context.Context, *OrderNo) (*Result, error)
	// 更改收货地址
	ChangeShipmentAddress(context.Context, *ChangeOrderAddressRequest) (*Result, error)
	// 备货完成
	PickUp(context.Context, *OrderNo) (*Result, error)
	// 订单发货,并记录配送服务商编号及单号
	Ship(context.Context, *OrderShipmentRequest) (*Result, error)
	// 买家收货
	BuyerReceived(context.Context, *OrderNo) (*Result, error)
	// 删除订单
	Forbid(context.Context, *OrderNo) (*Result, error)
	// 获取订单日志
	LogBytes(context.Context, *OrderNo) (*String, error)
	// * 获取订单返利列表
	QueryRebateListList(context.Context, *QueryRebateListRequest) (*QueryRebateListResponse, error)
	mustEmbedUnimplementedOrderServiceServer()
}

// UnimplementedOrderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (UnimplementedOrderServiceServer) SubmitOrder(context.Context, *SubmitOrderRequest) (*OrderSubmitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitOrder not implemented")
}
func (UnimplementedOrderServiceServer) PrepareOrder(context.Context, *PrepareOrderRequest) (*PrepareOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrepareOrder not implemented")
}
func (UnimplementedOrderServiceServer) GetParentOrder(context.Context, *OrderNoV2) (*SParentOrder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParentOrder not implemented")
}
func (UnimplementedOrderServiceServer) GetOrder(context.Context, *OrderNoV2) (*SSingleOrder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedOrderServiceServer) TradeOrderCashPay(context.Context, *Int64) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TradeOrderCashPay not implemented")
}
func (UnimplementedOrderServiceServer) TradeOrderUpdateTicket(context.Context, *TradeOrderTicketRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TradeOrderUpdateTicket not implemented")
}
func (UnimplementedOrderServiceServer) PrepareOrderWithCoupon_(context.Context, *PrepareOrderRequest) (*StringMap, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrepareOrderWithCoupon_ not implemented")
}
func (UnimplementedOrderServiceServer) CancelOrder(context.Context, *CancelOrderRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}
func (UnimplementedOrderServiceServer) ConfirmOrder(context.Context, *OrderNo) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmOrder not implemented")
}
func (UnimplementedOrderServiceServer) ChangeShipmentAddress(context.Context, *ChangeOrderAddressRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeShipmentAddress not implemented")
}
func (UnimplementedOrderServiceServer) PickUp(context.Context, *OrderNo) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PickUp not implemented")
}
func (UnimplementedOrderServiceServer) Ship(context.Context, *OrderShipmentRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ship not implemented")
}
func (UnimplementedOrderServiceServer) BuyerReceived(context.Context, *OrderNo) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuyerReceived not implemented")
}
func (UnimplementedOrderServiceServer) Forbid(context.Context, *OrderNo) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Forbid not implemented")
}
func (UnimplementedOrderServiceServer) LogBytes(context.Context, *OrderNo) (*String, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogBytes not implemented")
}
func (UnimplementedOrderServiceServer) QueryRebateListList(context.Context, *QueryRebateListRequest) (*QueryRebateListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryRebateListList not implemented")
}
func (UnimplementedOrderServiceServer) mustEmbedUnimplementedOrderServiceServer() {}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_SubmitOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).SubmitOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/SubmitOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).SubmitOrder(ctx, req.(*SubmitOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_PrepareOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).PrepareOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/PrepareOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).PrepareOrder(ctx, req.(*PrepareOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetParentOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderNoV2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetParentOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/GetParentOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetParentOrder(ctx, req.(*OrderNoV2))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderNoV2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/GetOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrder(ctx, req.(*OrderNoV2))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_TradeOrderCashPay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int64)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).TradeOrderCashPay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/TradeOrderCashPay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).TradeOrderCashPay(ctx, req.(*Int64))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_TradeOrderUpdateTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TradeOrderTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).TradeOrderUpdateTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/TradeOrderUpdateTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).TradeOrderUpdateTicket(ctx, req.(*TradeOrderTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_PrepareOrderWithCoupon__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).PrepareOrderWithCoupon_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/PrepareOrderWithCoupon_",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).PrepareOrderWithCoupon_(ctx, req.(*PrepareOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/CancelOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CancelOrder(ctx, req.(*CancelOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_ConfirmOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderNo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).ConfirmOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/ConfirmOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).ConfirmOrder(ctx, req.(*OrderNo))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_ChangeShipmentAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeOrderAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).ChangeShipmentAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/ChangeShipmentAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).ChangeShipmentAddress(ctx, req.(*ChangeOrderAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_PickUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderNo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).PickUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/PickUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).PickUp(ctx, req.(*OrderNo))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_Ship_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderShipmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).Ship(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/Ship",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).Ship(ctx, req.(*OrderShipmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_BuyerReceived_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderNo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).BuyerReceived(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/BuyerReceived",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).BuyerReceived(ctx, req.(*OrderNo))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_Forbid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderNo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).Forbid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/forbid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).Forbid(ctx, req.(*OrderNo))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_LogBytes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderNo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).LogBytes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/LogBytes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).LogBytes(ctx, req.(*OrderNo))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_QueryRebateListList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRebateListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).QueryRebateListList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/QueryRebateListList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).QueryRebateListList(ctx, req.(*QueryRebateListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitOrder",
			Handler:    _OrderService_SubmitOrder_Handler,
		},
		{
			MethodName: "PrepareOrder",
			Handler:    _OrderService_PrepareOrder_Handler,
		},
		{
			MethodName: "GetParentOrder",
			Handler:    _OrderService_GetParentOrder_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _OrderService_GetOrder_Handler,
		},
		{
			MethodName: "TradeOrderCashPay",
			Handler:    _OrderService_TradeOrderCashPay_Handler,
		},
		{
			MethodName: "TradeOrderUpdateTicket",
			Handler:    _OrderService_TradeOrderUpdateTicket_Handler,
		},
		{
			MethodName: "PrepareOrderWithCoupon_",
			Handler:    _OrderService_PrepareOrderWithCoupon__Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _OrderService_CancelOrder_Handler,
		},
		{
			MethodName: "ConfirmOrder",
			Handler:    _OrderService_ConfirmOrder_Handler,
		},
		{
			MethodName: "ChangeShipmentAddress",
			Handler:    _OrderService_ChangeShipmentAddress_Handler,
		},
		{
			MethodName: "PickUp",
			Handler:    _OrderService_PickUp_Handler,
		},
		{
			MethodName: "Ship",
			Handler:    _OrderService_Ship_Handler,
		},
		{
			MethodName: "BuyerReceived",
			Handler:    _OrderService_BuyerReceived_Handler,
		},
		{
			MethodName: "forbid",
			Handler:    _OrderService_Forbid_Handler,
		},
		{
			MethodName: "LogBytes",
			Handler:    _OrderService_LogBytes_Handler,
		},
		{
			MethodName: "QueryRebateListList",
			Handler:    _OrderService_QueryRebateListList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order_service.proto",
}
