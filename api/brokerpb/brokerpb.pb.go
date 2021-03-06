// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/brokerpb.proto

package brokerpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetFulcrumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Planet string  `protobuf:"bytes,1,opt,name=planet,proto3" json:"planet,omitempty"`
	Vector []int32 `protobuf:"varint,2,rep,packed,name=vector,proto3" json:"vector,omitempty"`
}

func (x *GetFulcrumRequest) Reset() {
	*x = GetFulcrumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_brokerpb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFulcrumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFulcrumRequest) ProtoMessage() {}

func (x *GetFulcrumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_brokerpb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFulcrumRequest.ProtoReflect.Descriptor instead.
func (*GetFulcrumRequest) Descriptor() ([]byte, []int) {
	return file_api_brokerpb_proto_rawDescGZIP(), []int{0}
}

func (x *GetFulcrumRequest) GetPlanet() string {
	if x != nil {
		return x.Planet
	}
	return ""
}

func (x *GetFulcrumRequest) GetVector() []int32 {
	if x != nil {
		return x.Vector
	}
	return nil
}

type GetFulcrumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success   bool  `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	FulcrumId int32 `protobuf:"varint,2,opt,name=fulcrumId,proto3" json:"fulcrumId,omitempty"`
}

func (x *GetFulcrumResponse) Reset() {
	*x = GetFulcrumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_brokerpb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFulcrumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFulcrumResponse) ProtoMessage() {}

func (x *GetFulcrumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_brokerpb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFulcrumResponse.ProtoReflect.Descriptor instead.
func (*GetFulcrumResponse) Descriptor() ([]byte, []int) {
	return file_api_brokerpb_proto_rawDescGZIP(), []int{1}
}

func (x *GetFulcrumResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetFulcrumResponse) GetFulcrumId() int32 {
	if x != nil {
		return x.FulcrumId
	}
	return 0
}

type GetNumberRebelsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Planet string  `protobuf:"bytes,1,opt,name=planet,proto3" json:"planet,omitempty"`
	City   string  `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Vector []int32 `protobuf:"varint,3,rep,packed,name=vector,proto3" json:"vector,omitempty"`
}

func (x *GetNumberRebelsRequest) Reset() {
	*x = GetNumberRebelsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_brokerpb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNumberRebelsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNumberRebelsRequest) ProtoMessage() {}

func (x *GetNumberRebelsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_brokerpb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNumberRebelsRequest.ProtoReflect.Descriptor instead.
func (*GetNumberRebelsRequest) Descriptor() ([]byte, []int) {
	return file_api_brokerpb_proto_rawDescGZIP(), []int{2}
}

func (x *GetNumberRebelsRequest) GetPlanet() string {
	if x != nil {
		return x.Planet
	}
	return ""
}

func (x *GetNumberRebelsRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *GetNumberRebelsRequest) GetVector() []int32 {
	if x != nil {
		return x.Vector
	}
	return nil
}

type GetNumberRebelsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool    `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Number  int32   `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	Vector  []int32 `protobuf:"varint,3,rep,packed,name=vector,proto3" json:"vector,omitempty"`
}

func (x *GetNumberRebelsResponse) Reset() {
	*x = GetNumberRebelsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_brokerpb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNumberRebelsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNumberRebelsResponse) ProtoMessage() {}

func (x *GetNumberRebelsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_brokerpb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNumberRebelsResponse.ProtoReflect.Descriptor instead.
func (*GetNumberRebelsResponse) Descriptor() ([]byte, []int) {
	return file_api_brokerpb_proto_rawDescGZIP(), []int{3}
}

func (x *GetNumberRebelsResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetNumberRebelsResponse) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *GetNumberRebelsResponse) GetVector() []int32 {
	if x != nil {
		return x.Vector
	}
	return nil
}

var File_api_brokerpb_proto protoreflect.FileDescriptor

var file_api_brokerpb_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x70, 0x62, 0x22, 0x43,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x76,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x06, 0x76, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x22, 0x4c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6c, 0x63, 0x72, 0x75,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x66, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x49,
	0x64, 0x22, 0x5c, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65,
	0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x61,
	0x6e, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x06, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x22,
	0x63, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x62, 0x65,
	0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x06, 0x76, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x32, 0xb4, 0x01, 0x0a, 0x0d, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6c,
	0x63, 0x72, 0x75, 0x6d, 0x12, 0x1b, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x46, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x46, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x58, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0x20, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x62, 0x65, 0x6c, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x62, 0x65, 0x6c,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x61,
	0x70, 0x69, 0x2f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_brokerpb_proto_rawDescOnce sync.Once
	file_api_brokerpb_proto_rawDescData = file_api_brokerpb_proto_rawDesc
)

func file_api_brokerpb_proto_rawDescGZIP() []byte {
	file_api_brokerpb_proto_rawDescOnce.Do(func() {
		file_api_brokerpb_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_brokerpb_proto_rawDescData)
	})
	return file_api_brokerpb_proto_rawDescData
}

var file_api_brokerpb_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_brokerpb_proto_goTypes = []interface{}{
	(*GetFulcrumRequest)(nil),       // 0: brokerpb.GetFulcrumRequest
	(*GetFulcrumResponse)(nil),      // 1: brokerpb.GetFulcrumResponse
	(*GetNumberRebelsRequest)(nil),  // 2: brokerpb.GetNumberRebelsRequest
	(*GetNumberRebelsResponse)(nil), // 3: brokerpb.GetNumberRebelsResponse
}
var file_api_brokerpb_proto_depIdxs = []int32{
	0, // 0: brokerpb.BrokerService.GetFulcrum:input_type -> brokerpb.GetFulcrumRequest
	2, // 1: brokerpb.BrokerService.GetNumberRebels:input_type -> brokerpb.GetNumberRebelsRequest
	1, // 2: brokerpb.BrokerService.GetFulcrum:output_type -> brokerpb.GetFulcrumResponse
	3, // 3: brokerpb.BrokerService.GetNumberRebels:output_type -> brokerpb.GetNumberRebelsResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_brokerpb_proto_init() }
func file_api_brokerpb_proto_init() {
	if File_api_brokerpb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_brokerpb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFulcrumRequest); i {
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
		file_api_brokerpb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFulcrumResponse); i {
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
		file_api_brokerpb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNumberRebelsRequest); i {
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
		file_api_brokerpb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNumberRebelsResponse); i {
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
			RawDescriptor: file_api_brokerpb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_brokerpb_proto_goTypes,
		DependencyIndexes: file_api_brokerpb_proto_depIdxs,
		MessageInfos:      file_api_brokerpb_proto_msgTypes,
	}.Build()
	File_api_brokerpb_proto = out.File
	file_api_brokerpb_proto_rawDesc = nil
	file_api_brokerpb_proto_goTypes = nil
	file_api_brokerpb_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BrokerServiceClient is the client API for BrokerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BrokerServiceClient interface {
	GetFulcrum(ctx context.Context, in *GetFulcrumRequest, opts ...grpc.CallOption) (*GetFulcrumResponse, error)
	GetNumberRebels(ctx context.Context, in *GetNumberRebelsRequest, opts ...grpc.CallOption) (*GetNumberRebelsResponse, error)
}

type brokerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBrokerServiceClient(cc grpc.ClientConnInterface) BrokerServiceClient {
	return &brokerServiceClient{cc}
}

func (c *brokerServiceClient) GetFulcrum(ctx context.Context, in *GetFulcrumRequest, opts ...grpc.CallOption) (*GetFulcrumResponse, error) {
	out := new(GetFulcrumResponse)
	err := c.cc.Invoke(ctx, "/brokerpb.BrokerService/GetFulcrum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerServiceClient) GetNumberRebels(ctx context.Context, in *GetNumberRebelsRequest, opts ...grpc.CallOption) (*GetNumberRebelsResponse, error) {
	out := new(GetNumberRebelsResponse)
	err := c.cc.Invoke(ctx, "/brokerpb.BrokerService/GetNumberRebels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BrokerServiceServer is the server API for BrokerService service.
type BrokerServiceServer interface {
	GetFulcrum(context.Context, *GetFulcrumRequest) (*GetFulcrumResponse, error)
	GetNumberRebels(context.Context, *GetNumberRebelsRequest) (*GetNumberRebelsResponse, error)
}

// UnimplementedBrokerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBrokerServiceServer struct {
}

func (*UnimplementedBrokerServiceServer) GetFulcrum(context.Context, *GetFulcrumRequest) (*GetFulcrumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFulcrum not implemented")
}
func (*UnimplementedBrokerServiceServer) GetNumberRebels(context.Context, *GetNumberRebelsRequest) (*GetNumberRebelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNumberRebels not implemented")
}

func RegisterBrokerServiceServer(s *grpc.Server, srv BrokerServiceServer) {
	s.RegisterService(&_BrokerService_serviceDesc, srv)
}

func _BrokerService_GetFulcrum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFulcrumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServiceServer).GetFulcrum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brokerpb.BrokerService/GetFulcrum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServiceServer).GetFulcrum(ctx, req.(*GetFulcrumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrokerService_GetNumberRebels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNumberRebelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServiceServer).GetNumberRebels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brokerpb.BrokerService/GetNumberRebels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServiceServer).GetNumberRebels(ctx, req.(*GetNumberRebelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BrokerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "brokerpb.BrokerService",
	HandlerType: (*BrokerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFulcrum",
			Handler:    _BrokerService_GetFulcrum_Handler,
		},
		{
			MethodName: "GetNumberRebels",
			Handler:    _BrokerService_GetNumberRebels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/brokerpb.proto",
}
