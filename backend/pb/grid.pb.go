// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: grid.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GridNodeTypeEnum int32

const (
	GridNodeTypeEnum_POWER_STATION GridNodeTypeEnum = 0
	GridNodeTypeEnum_SOLAR_PLANT   GridNodeTypeEnum = 1
	GridNodeTypeEnum_SUBSTATION    GridNodeTypeEnum = 2
	GridNodeTypeEnum_TRANSMISSION  GridNodeTypeEnum = 3
	GridNodeTypeEnum_CONSUMER      GridNodeTypeEnum = 4
)

// Enum value maps for GridNodeTypeEnum.
var (
	GridNodeTypeEnum_name = map[int32]string{
		0: "POWER_STATION",
		1: "SOLAR_PLANT",
		2: "SUBSTATION",
		3: "TRANSMISSION",
		4: "CONSUMER",
	}
	GridNodeTypeEnum_value = map[string]int32{
		"POWER_STATION": 0,
		"SOLAR_PLANT":   1,
		"SUBSTATION":    2,
		"TRANSMISSION":  3,
		"CONSUMER":      4,
	}
)

func (x GridNodeTypeEnum) Enum() *GridNodeTypeEnum {
	p := new(GridNodeTypeEnum)
	*p = x
	return p
}

func (x GridNodeTypeEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GridNodeTypeEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_grid_proto_enumTypes[0].Descriptor()
}

func (GridNodeTypeEnum) Type() protoreflect.EnumType {
	return &file_grid_proto_enumTypes[0]
}

func (x GridNodeTypeEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GridNodeTypeEnum.Descriptor instead.
func (GridNodeTypeEnum) EnumDescriptor() ([]byte, []int) {
	return file_grid_proto_rawDescGZIP(), []int{0}
}

type GridNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type  GridNodeTypeEnum `protobuf:"varint,3,opt,name=type,proto3,enum=grid.GridNodeTypeEnum" json:"type,omitempty"`
	Power float32          `protobuf:"fixed32,4,opt,name=power,proto3" json:"power,omitempty"`
}

func (x *GridNode) Reset() {
	*x = GridNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grid_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GridNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GridNode) ProtoMessage() {}

func (x *GridNode) ProtoReflect() protoreflect.Message {
	mi := &file_grid_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GridNode.ProtoReflect.Descriptor instead.
func (*GridNode) Descriptor() ([]byte, []int) {
	return file_grid_proto_rawDescGZIP(), []int{0}
}

func (x *GridNode) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GridNode) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GridNode) GetType() GridNodeTypeEnum {
	if x != nil {
		return x.Type
	}
	return GridNodeTypeEnum_POWER_STATION
}

func (x *GridNode) GetPower() float32 {
	if x != nil {
		return x.Power
	}
	return 0
}

type GetNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetNodeRequest) Reset() {
	*x = GetNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grid_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodeRequest) ProtoMessage() {}

func (x *GetNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grid_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodeRequest.ProtoReflect.Descriptor instead.
func (*GetNodeRequest) Descriptor() ([]byte, []int) {
	return file_grid_proto_rawDescGZIP(), []int{1}
}

func (x *GetNodeRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node *GridNode `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
}

func (x *UpdateNodeRequest) Reset() {
	*x = UpdateNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grid_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNodeRequest) ProtoMessage() {}

func (x *UpdateNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grid_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNodeRequest.ProtoReflect.Descriptor instead.
func (*UpdateNodeRequest) Descriptor() ([]byte, []int) {
	return file_grid_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateNodeRequest) GetNode() *GridNode {
	if x != nil {
		return x.Node
	}
	return nil
}

type GetNodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node *GridNode `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
}

func (x *GetNodeResponse) Reset() {
	*x = GetNodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grid_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodeResponse) ProtoMessage() {}

func (x *GetNodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grid_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodeResponse.ProtoReflect.Descriptor instead.
func (*GetNodeResponse) Descriptor() ([]byte, []int) {
	return file_grid_proto_rawDescGZIP(), []int{3}
}

func (x *GetNodeResponse) GetNode() *GridNode {
	if x != nil {
		return x.Node
	}
	return nil
}

type UpdateNodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateNodeResponse) Reset() {
	*x = UpdateNodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grid_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNodeResponse) ProtoMessage() {}

func (x *UpdateNodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grid_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNodeResponse.ProtoReflect.Descriptor instead.
func (*UpdateNodeResponse) Descriptor() ([]byte, []int) {
	return file_grid_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateNodeResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type NodeLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source int32 `protobuf:"varint,1,opt,name=source,proto3" json:"source,omitempty"`
	Target int32 `protobuf:"varint,2,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *NodeLink) Reset() {
	*x = NodeLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grid_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeLink) ProtoMessage() {}

func (x *NodeLink) ProtoReflect() protoreflect.Message {
	mi := &file_grid_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeLink.ProtoReflect.Descriptor instead.
func (*NodeLink) Descriptor() ([]byte, []int) {
	return file_grid_proto_rawDescGZIP(), []int{5}
}

func (x *NodeLink) GetSource() int32 {
	if x != nil {
		return x.Source
	}
	return 0
}

func (x *NodeLink) GetTarget() int32 {
	if x != nil {
		return x.Target
	}
	return 0
}

type GetGridResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*GridNode `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	Links []*NodeLink `protobuf:"bytes,2,rep,name=links,proto3" json:"links,omitempty"`
}

func (x *GetGridResponse) Reset() {
	*x = GetGridResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grid_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGridResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGridResponse) ProtoMessage() {}

func (x *GetGridResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grid_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGridResponse.ProtoReflect.Descriptor instead.
func (*GetGridResponse) Descriptor() ([]byte, []int) {
	return file_grid_proto_rawDescGZIP(), []int{6}
}

func (x *GetGridResponse) GetNodes() []*GridNode {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *GetGridResponse) GetLinks() []*NodeLink {
	if x != nil {
		return x.Links
	}
	return nil
}

var File_grid_proto protoreflect.FileDescriptor

var file_grid_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x72, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72,
	0x69, 0x64, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x70, 0x0a, 0x08, 0x47, 0x72, 0x69, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x2a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e,
	0x67, 0x72, 0x69, 0x64, 0x2e, 0x47, 0x72, 0x69, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x6f, 0x77, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x6f, 0x77, 0x65,
	0x72, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x37, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x67, 0x72, 0x69, 0x64, 0x2e, 0x47, 0x72,
	0x69, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x22, 0x35, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x22, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x67, 0x72, 0x69, 0x64, 0x2e, 0x47, 0x72, 0x69, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e,
	0x6f, 0x64, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3a, 0x0a, 0x08, 0x4e, 0x6f, 0x64,
	0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x5d, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x47, 0x72, 0x69, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x67, 0x72, 0x69, 0x64, 0x2e, 0x47,
	0x72, 0x69, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x24,
	0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x67, 0x72, 0x69, 0x64, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x05, 0x6c,
	0x69, 0x6e, 0x6b, 0x73, 0x2a, 0x66, 0x0a, 0x10, 0x47, 0x72, 0x69, 0x64, 0x4e, 0x6f, 0x64, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x4f, 0x57, 0x45,
	0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x53,
	0x4f, 0x4c, 0x41, 0x52, 0x5f, 0x50, 0x4c, 0x41, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a,
	0x53, 0x55, 0x42, 0x53, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c,
	0x54, 0x52, 0x41, 0x4e, 0x53, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x12, 0x0c,
	0x0a, 0x08, 0x43, 0x4f, 0x4e, 0x53, 0x55, 0x4d, 0x45, 0x52, 0x10, 0x04, 0x32, 0xc0, 0x01, 0x0a,
	0x0b, 0x47, 0x72, 0x69, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x47, 0x72, 0x69, 0x64, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x15, 0x2e, 0x67, 0x72, 0x69, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x69, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64,
	0x65, 0x12, 0x14, 0x2e, 0x67, 0x72, 0x69, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67, 0x72, 0x69, 0x64, 0x2e, 0x47,
	0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f,
	0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x2e, 0x67,
	0x72, 0x69, 0x64, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x72, 0x69, 0x64, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x14, 0x5a, 0x12, 0x65, 0x67, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grid_proto_rawDescOnce sync.Once
	file_grid_proto_rawDescData = file_grid_proto_rawDesc
)

func file_grid_proto_rawDescGZIP() []byte {
	file_grid_proto_rawDescOnce.Do(func() {
		file_grid_proto_rawDescData = protoimpl.X.CompressGZIP(file_grid_proto_rawDescData)
	})
	return file_grid_proto_rawDescData
}

var file_grid_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_grid_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_grid_proto_goTypes = []interface{}{
	(GridNodeTypeEnum)(0),      // 0: grid.GridNodeTypeEnum
	(*GridNode)(nil),           // 1: grid.GridNode
	(*GetNodeRequest)(nil),     // 2: grid.GetNodeRequest
	(*UpdateNodeRequest)(nil),  // 3: grid.UpdateNodeRequest
	(*GetNodeResponse)(nil),    // 4: grid.GetNodeResponse
	(*UpdateNodeResponse)(nil), // 5: grid.UpdateNodeResponse
	(*NodeLink)(nil),           // 6: grid.NodeLink
	(*GetGridResponse)(nil),    // 7: grid.GetGridResponse
	(*emptypb.Empty)(nil),      // 8: google.protobuf.Empty
}
var file_grid_proto_depIdxs = []int32{
	0, // 0: grid.GridNode.type:type_name -> grid.GridNodeTypeEnum
	1, // 1: grid.UpdateNodeRequest.node:type_name -> grid.GridNode
	1, // 2: grid.GetNodeResponse.node:type_name -> grid.GridNode
	1, // 3: grid.GetGridResponse.nodes:type_name -> grid.GridNode
	6, // 4: grid.GetGridResponse.links:type_name -> grid.NodeLink
	8, // 5: grid.GridService.GetGrid:input_type -> google.protobuf.Empty
	2, // 6: grid.GridService.GetNode:input_type -> grid.GetNodeRequest
	3, // 7: grid.GridService.UpdateNode:input_type -> grid.UpdateNodeRequest
	7, // 8: grid.GridService.GetGrid:output_type -> grid.GetGridResponse
	4, // 9: grid.GridService.GetNode:output_type -> grid.GetNodeResponse
	5, // 10: grid.GridService.UpdateNode:output_type -> grid.UpdateNodeResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_grid_proto_init() }
func file_grid_proto_init() {
	if File_grid_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grid_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GridNode); i {
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
		file_grid_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodeRequest); i {
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
		file_grid_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNodeRequest); i {
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
		file_grid_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodeResponse); i {
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
		file_grid_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNodeResponse); i {
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
		file_grid_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeLink); i {
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
		file_grid_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGridResponse); i {
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
			RawDescriptor: file_grid_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grid_proto_goTypes,
		DependencyIndexes: file_grid_proto_depIdxs,
		EnumInfos:         file_grid_proto_enumTypes,
		MessageInfos:      file_grid_proto_msgTypes,
	}.Build()
	File_grid_proto = out.File
	file_grid_proto_rawDesc = nil
	file_grid_proto_goTypes = nil
	file_grid_proto_depIdxs = nil
}
