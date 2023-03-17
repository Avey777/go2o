// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: item_service.proto

package proto

import (
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

type ItemDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId   int64 `protobuf:"zigzag64,1,opt,name=ItemId,proto3" json:"ItemId"`
	ItemType int32 `protobuf:"zigzag32,2,opt,name=ItemType,proto3" json:"ItemType"`
}

func (x *ItemDetailRequest) Reset() {
	*x = ItemDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemDetailRequest) ProtoMessage() {}

func (x *ItemDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_item_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemDetailRequest.ProtoReflect.Descriptor instead.
func (*ItemDetailRequest) Descriptor() ([]byte, []int) {
	return file_item_service_proto_rawDescGZIP(), []int{0}
}

func (x *ItemDetailRequest) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *ItemDetailRequest) GetItemType() int32 {
	if x != nil {
		return x.ItemType
	}
	return 0
}

type SaleLabelItemsRequest_ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopId  int64          `protobuf:"varint,1,opt,name=shopId,proto3" json:"shopId"`
	LabelId int32          `protobuf:"varint,2,opt,name=labelId,proto3" json:"labelId"`
	Params  *SPagingParams `protobuf:"bytes,3,opt,name=params,proto3" json:"params"`
}

func (x *SaleLabelItemsRequest_) Reset() {
	*x = SaleLabelItemsRequest_{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaleLabelItemsRequest_) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaleLabelItemsRequest_) ProtoMessage() {}

func (x *SaleLabelItemsRequest_) ProtoReflect() protoreflect.Message {
	mi := &file_item_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaleLabelItemsRequest_.ProtoReflect.Descriptor instead.
func (*SaleLabelItemsRequest_) Descriptor() ([]byte, []int) {
	return file_item_service_proto_rawDescGZIP(), []int{1}
}

func (x *SaleLabelItemsRequest_) GetShopId() int64 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

func (x *SaleLabelItemsRequest_) GetLabelId() int32 {
	if x != nil {
		return x.LabelId
	}
	return 0
}

func (x *SaleLabelItemsRequest_) GetParams() *SPagingParams {
	if x != nil {
		return x.Params
	}
	return nil
}

var File_item_service_proto protoreflect.FileDescriptor

var file_item_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x69, 0x74, 0x65, 0x6d,
	0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x11, 0x49, 0x74,
	0x65, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x12, 0x52,
	0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x74, 0x65, 0x6d, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x08, 0x49, 0x74, 0x65, 0x6d, 0x54,
	0x79, 0x70, 0x65, 0x22, 0x72, 0x0a, 0x16, 0x53, 0x61, 0x6c, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73,
	0x68, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x49, 0x64, 0x12,
	0x26, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x53, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52,
	0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x32, 0xa8, 0x09, 0x0a, 0x0b, 0x49, 0x74, 0x65, 0x6d,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x53, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x08, 0x53, 0x61, 0x76,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x79, 0x53, 0x6b, 0x75, 0x12, 0x11, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x42, 0x79, 0x53, 0x6b, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x11, 0x2e, 0x53, 0x55, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x56, 0x69, 0x65, 0x77, 0x49, 0x74,
	0x65, 0x6d, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x41,
	0x6e, 0x64, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x1a, 0x2e, 0x47, 0x65, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x41, 0x6e, 0x64, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x2d, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x12, 0x06, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x1a, 0x0f, 0x2e, 0x53, 0x54,
	0x72, 0x61, 0x64, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x22, 0x00, 0x12, 0x19,
	0x0a, 0x06, 0x47, 0x65, 0x74, 0x53, 0x6b, 0x75, 0x12, 0x06, 0x2e, 0x53, 0x6b, 0x75, 0x49, 0x64,
	0x1a, 0x05, 0x2e, 0x53, 0x53, 0x6b, 0x75, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x0a, 0x52, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x0b, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x13, 0x2e, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0f, 0x53, 0x61, 0x76, 0x65, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x0d, 0x53,
	0x69, 0x67, 0x6e, 0x41, 0x73, 0x49, 0x6c, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x12, 0x13, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x49, 0x6c, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0e,
	0x53, 0x65, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x13,
	0x2e, 0x53, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x32,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x12, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x22, 0x00, 0x12, 0x34, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x10,
	0x2e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x57,
	0x68, 0x6f, 0x6c, 0x65, 0x73, 0x61, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x41, 0x72, 0x72,
	0x61, 0x79, 0x12, 0x06, 0x2e, 0x53, 0x6b, 0x75, 0x49, 0x64, 0x1a, 0x18, 0x2e, 0x53, 0x57, 0x73,
	0x53, 0x6b, 0x75, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x12, 0x53, 0x61, 0x76, 0x65, 0x57, 0x68,
	0x6f, 0x6c, 0x65, 0x73, 0x61, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x15, 0x2e, 0x53,
	0x61, 0x76, 0x65, 0x53, 0x6b, 0x75, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x52,
	0x0a, 0x19, 0x47, 0x65, 0x74, 0x57, 0x68, 0x6f, 0x6c, 0x65, 0x73, 0x61, 0x6c, 0x65, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x15, 0x2e, 0x47, 0x65,
	0x74, 0x57, 0x73, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x53, 0x57, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3c, 0x0a, 0x15, 0x53, 0x61, 0x76, 0x65, 0x57, 0x68, 0x6f, 0x6c, 0x65, 0x73,
	0x61, 0x6c, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x53, 0x61,
	0x76, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00,
	0x12, 0x34, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x61, 0x6c, 0x65, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6c,
	0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x09, 0x2e, 0x49, 0x64, 0x4f, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x1a, 0x0b, 0x2e, 0x53, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x00,
	0x12, 0x27, 0x0a, 0x0d, 0x53, 0x61, 0x76, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x12, 0x0b, 0x2e, 0x53, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x1a, 0x07,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x24, 0x0a, 0x0f, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x06, 0x2e, 0x49,
	0x6e, 0x74, 0x36, 0x34, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12,
	0x51, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x47, 0x6f, 0x6f, 0x64, 0x73, 0x42, 0x79, 0x53, 0x61, 0x6c, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x5f, 0x12, 0x17, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x1a, 0x14, 0x2e, 0x50, 0x61, 0x67,
	0x69, 0x6e, 0x67, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x1f, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x67, 0x6f, 0x32, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_item_service_proto_rawDescOnce sync.Once
	file_item_service_proto_rawDescData = file_item_service_proto_rawDesc
)

func file_item_service_proto_rawDescGZIP() []byte {
	file_item_service_proto_rawDescOnce.Do(func() {
		file_item_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_item_service_proto_rawDescData)
	})
	return file_item_service_proto_rawDescData
}

var file_item_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_item_service_proto_goTypes = []interface{}{
	(*ItemDetailRequest)(nil),           // 0: ItemDetailRequest
	(*SaleLabelItemsRequest_)(nil),      // 1: SaleLabelItemsRequest_
	(*SPagingParams)(nil),               // 2: SPagingParams
	(*GetItemRequest)(nil),              // 3: GetItemRequest
	(*SaveItemRequest)(nil),             // 4: SaveItemRequest
	(*ItemBySkuRequest)(nil),            // 5: ItemBySkuRequest
	(*GetItemAndSnapshotRequest)(nil),   // 6: GetItemAndSnapshotRequest
	(*Int64)(nil),                       // 7: Int64
	(*SkuId)(nil),                       // 8: SkuId
	(*ItemReviewRequest)(nil),           // 9: ItemReviewRequest
	(*RecycleItemRequest)(nil),          // 10: RecycleItemRequest
	(*SaveLevelPriceRequest)(nil),       // 11: SaveLevelPriceRequest
	(*ItemIllegalRequest)(nil),          // 12: ItemIllegalRequest
	(*ShelveStateRequest)(nil),          // 13: ShelveStateRequest
	(*GetItemsRequest)(nil),             // 14: GetItemsRequest
	(*SaveSkuPricesRequest)(nil),        // 15: SaveSkuPricesRequest
	(*GetWsDiscountRequest)(nil),        // 16: GetWsDiscountRequest
	(*SaveItemDiscountRequest)(nil),     // 17: SaveItemDiscountRequest
	(*Empty)(nil),                       // 18: Empty
	(*IdOrName)(nil),                    // 19: IdOrName
	(*SItemLabel)(nil),                  // 20: SItemLabel
	(*SItemDataResponse)(nil),           // 21: SItemDataResponse
	(*SaveItemResponse)(nil),            // 22: SaveItemResponse
	(*SUnifiedViewItem)(nil),            // 23: SUnifiedViewItem
	(*ItemSnapshotResponse)(nil),        // 24: ItemSnapshotResponse
	(*STradeSnapshot)(nil),              // 25: STradeSnapshot
	(*SSku)(nil),                        // 26: SSku
	(*Result)(nil),                      // 27: Result
	(*String)(nil),                      // 28: String
	(*PagingGoodsResponse)(nil),         // 29: PagingGoodsResponse
	(*SWsSkuPriceListResponse)(nil),     // 30: SWsSkuPriceListResponse
	(*SWsItemDiscountListResponse)(nil), // 31: SWsItemDiscountListResponse
	(*ItemLabelListResponse)(nil),       // 32: ItemLabelListResponse
}
var file_item_service_proto_depIdxs = []int32{
	2,  // 0: SaleLabelItemsRequest_.params:type_name -> SPagingParams
	3,  // 1: ItemService.GetItem:input_type -> GetItemRequest
	4,  // 2: ItemService.SaveItem:input_type -> SaveItemRequest
	5,  // 3: ItemService.GetItemBySku:input_type -> ItemBySkuRequest
	6,  // 4: ItemService.GetItemAndSnapshot:input_type -> GetItemAndSnapshotRequest
	7,  // 5: ItemService.GetTradeSnapshot:input_type -> Int64
	8,  // 6: ItemService.GetSku:input_type -> SkuId
	9,  // 7: ItemService.ReviewItem:input_type -> ItemReviewRequest
	10, // 8: ItemService.RecycleItem:input_type -> RecycleItemRequest
	11, // 9: ItemService.SaveLevelPrices:input_type -> SaveLevelPriceRequest
	12, // 10: ItemService.SignAsIllegal:input_type -> ItemIllegalRequest
	13, // 11: ItemService.SetShelveState:input_type -> ShelveStateRequest
	0,  // 12: ItemService.GetItemDetailData:input_type -> ItemDetailRequest
	14, // 13: ItemService.GetItems:input_type -> GetItemsRequest
	8,  // 14: ItemService.GetWholesalePriceArray:input_type -> SkuId
	15, // 15: ItemService.SaveWholesalePrice:input_type -> SaveSkuPricesRequest
	16, // 16: ItemService.GetWholesaleDiscountArray:input_type -> GetWsDiscountRequest
	17, // 17: ItemService.SaveWholesaleDiscount:input_type -> SaveItemDiscountRequest
	18, // 18: ItemService.GetAllSaleLabels:input_type -> Empty
	19, // 19: ItemService.GetSaleLabel:input_type -> IdOrName
	20, // 20: ItemService.SaveSaleLabel:input_type -> SItemLabel
	7,  // 21: ItemService.DeleteSaleLabel:input_type -> Int64
	1,  // 22: ItemService.GetPagedValueGoodsBySaleLabel_:input_type -> SaleLabelItemsRequest_
	21, // 23: ItemService.GetItem:output_type -> SItemDataResponse
	22, // 24: ItemService.SaveItem:output_type -> SaveItemResponse
	23, // 25: ItemService.GetItemBySku:output_type -> SUnifiedViewItem
	24, // 26: ItemService.GetItemAndSnapshot:output_type -> ItemSnapshotResponse
	25, // 27: ItemService.GetTradeSnapshot:output_type -> STradeSnapshot
	26, // 28: ItemService.GetSku:output_type -> SSku
	27, // 29: ItemService.ReviewItem:output_type -> Result
	27, // 30: ItemService.RecycleItem:output_type -> Result
	27, // 31: ItemService.SaveLevelPrices:output_type -> Result
	27, // 32: ItemService.SignAsIllegal:output_type -> Result
	27, // 33: ItemService.SetShelveState:output_type -> Result
	28, // 34: ItemService.GetItemDetailData:output_type -> String
	29, // 35: ItemService.GetItems:output_type -> PagingGoodsResponse
	30, // 36: ItemService.GetWholesalePriceArray:output_type -> SWsSkuPriceListResponse
	27, // 37: ItemService.SaveWholesalePrice:output_type -> Result
	31, // 38: ItemService.GetWholesaleDiscountArray:output_type -> SWsItemDiscountListResponse
	27, // 39: ItemService.SaveWholesaleDiscount:output_type -> Result
	32, // 40: ItemService.GetAllSaleLabels:output_type -> ItemLabelListResponse
	20, // 41: ItemService.GetSaleLabel:output_type -> SItemLabel
	27, // 42: ItemService.SaveSaleLabel:output_type -> Result
	27, // 43: ItemService.DeleteSaleLabel:output_type -> Result
	29, // 44: ItemService.GetPagedValueGoodsBySaleLabel_:output_type -> PagingGoodsResponse
	23, // [23:45] is the sub-list for method output_type
	1,  // [1:23] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_item_service_proto_init() }
func file_item_service_proto_init() {
	if File_item_service_proto != nil {
		return
	}
	file_global_proto_init()
	file_message_item_dto_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_item_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemDetailRequest); i {
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
		file_item_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaleLabelItemsRequest_); i {
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
			RawDescriptor: file_item_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_item_service_proto_goTypes,
		DependencyIndexes: file_item_service_proto_depIdxs,
		MessageInfos:      file_item_service_proto_msgTypes,
	}.Build()
	File_item_service_proto = out.File
	file_item_service_proto_rawDesc = nil
	file_item_service_proto_goTypes = nil
	file_item_service_proto_depIdxs = nil
}
