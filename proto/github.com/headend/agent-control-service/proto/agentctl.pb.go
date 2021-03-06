// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: agentctl.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type AgentCTLResponseStatus int32

const (
	AgentCTLResponseStatus_FAIL    AgentCTLResponseStatus = 0 /// Success
	AgentCTLResponseStatus_SUCCESS AgentCTLResponseStatus = 1 /// Failed
)

// Enum value maps for AgentCTLResponseStatus.
var (
	AgentCTLResponseStatus_name = map[int32]string{
		0: "FAIL",
		1: "SUCCESS",
	}
	AgentCTLResponseStatus_value = map[string]int32{
		"FAIL":    0,
		"SUCCESS": 1,
	}
)

func (x AgentCTLResponseStatus) Enum() *AgentCTLResponseStatus {
	p := new(AgentCTLResponseStatus)
	*p = x
	return p
}

func (x AgentCTLResponseStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AgentCTLResponseStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_agentctl_proto_enumTypes[0].Descriptor()
}

func (AgentCTLResponseStatus) Type() protoreflect.EnumType {
	return &file_agentctl_proto_enumTypes[0]
}

func (x AgentCTLResponseStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AgentCTLResponseStatus.Descriptor instead.
func (AgentCTLResponseStatus) EnumDescriptor() ([]byte, []int) {
	return file_agentctl_proto_rawDescGZIP(), []int{0}
}

//*
// Represents the params to identify agent.
type AgentCTLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// IP of agent
	AgentId     int64        `protobuf:"varint,1,opt,name=agent_id,proto3" json:"agent_id,omitempty"`
	ProfileIp   int64        `protobuf:"varint,2,opt,name=profile_ip,proto3" json:"profile_ip,omitempty"`
	ControlId   int64        `protobuf:"varint,3,opt,name=control_id,proto3" json:"control_id,omitempty"`
	ControlType int64        `protobuf:"varint,4,opt,name=control_type,proto3" json:"control_type,omitempty"`
	RunThread   int64        `protobuf:"varint,5,opt,name=run_thread,proto3" json:"run_thread,omitempty"`
	TunnelData  []*anypb.Any `protobuf:"bytes,6,rep,name=tunnel_data,proto3" json:"tunnel_data,omitempty"`
}

func (x *AgentCTLRequest) Reset() {
	*x = AgentCTLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agentctl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentCTLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentCTLRequest) ProtoMessage() {}

func (x *AgentCTLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agentctl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentCTLRequest.ProtoReflect.Descriptor instead.
func (*AgentCTLRequest) Descriptor() ([]byte, []int) {
	return file_agentctl_proto_rawDescGZIP(), []int{0}
}

func (x *AgentCTLRequest) GetAgentId() int64 {
	if x != nil {
		return x.AgentId
	}
	return 0
}

func (x *AgentCTLRequest) GetProfileIp() int64 {
	if x != nil {
		return x.ProfileIp
	}
	return 0
}

func (x *AgentCTLRequest) GetControlId() int64 {
	if x != nil {
		return x.ControlId
	}
	return 0
}

func (x *AgentCTLRequest) GetControlType() int64 {
	if x != nil {
		return x.ControlType
	}
	return 0
}

func (x *AgentCTLRequest) GetRunThread() int64 {
	if x != nil {
		return x.RunThread
	}
	return 0
}

func (x *AgentCTLRequest) GetTunnelData() []*anypb.Any {
	if x != nil {
		return x.TunnelData
	}
	return nil
}

type AgentCTLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// Status
	Status AgentCTLResponseStatus `protobuf:"varint,1,opt,name=status,proto3,enum=proto.AgentCTLResponseStatus" json:"status,omitempty"`
	//*
	// Slice of agent object
	Agentctls []*AgentCTLRequest `protobuf:"bytes,2,rep,name=agentctls,json=data,proto3" json:"agentctls,omitempty"`
}

func (x *AgentCTLResponse) Reset() {
	*x = AgentCTLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agentctl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentCTLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentCTLResponse) ProtoMessage() {}

func (x *AgentCTLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agentctl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentCTLResponse.ProtoReflect.Descriptor instead.
func (*AgentCTLResponse) Descriptor() ([]byte, []int) {
	return file_agentctl_proto_rawDescGZIP(), []int{1}
}

func (x *AgentCTLResponse) GetStatus() AgentCTLResponseStatus {
	if x != nil {
		return x.Status
	}
	return AgentCTLResponseStatus_FAIL
}

func (x *AgentCTLResponse) GetAgentctls() []*AgentCTLRequest {
	if x != nil {
		return x.Agentctls
	}
	return nil
}

var File_agentctl_proto protoreflect.FileDescriptor

var file_agentctl_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x63, 0x74, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xe9, 0x01, 0x0a, 0x0f, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x54, 0x4c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x69, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f,
	0x69, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x75, 0x6e, 0x5f, 0x74, 0x68,
	0x72, 0x65, 0x61, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x72, 0x75, 0x6e, 0x5f,
	0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x12, 0x36, 0x0a, 0x0b, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e,
	0x79, 0x52, 0x0b, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x22, 0x7a,
	0x0a, 0x10, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x54, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x43, 0x54, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a, 0x09, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x63, 0x74, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x54, 0x4c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x2f, 0x0a, 0x16, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x43, 0x54, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x32, 0x54, 0x0a, 0x0f, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x43, 0x54, 0x4c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41,
	0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x16,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x54, 0x4c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x43, 0x54, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2d, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_agentctl_proto_rawDescOnce sync.Once
	file_agentctl_proto_rawDescData = file_agentctl_proto_rawDesc
)

func file_agentctl_proto_rawDescGZIP() []byte {
	file_agentctl_proto_rawDescOnce.Do(func() {
		file_agentctl_proto_rawDescData = protoimpl.X.CompressGZIP(file_agentctl_proto_rawDescData)
	})
	return file_agentctl_proto_rawDescData
}

var file_agentctl_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_agentctl_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_agentctl_proto_goTypes = []interface{}{
	(AgentCTLResponseStatus)(0), // 0: proto.AgentCTLResponseStatus
	(*AgentCTLRequest)(nil),     // 1: proto.AgentCTLRequest
	(*AgentCTLResponse)(nil),    // 2: proto.AgentCTLResponse
	(*anypb.Any)(nil),           // 3: google.protobuf.Any
}
var file_agentctl_proto_depIdxs = []int32{
	3, // 0: proto.AgentCTLRequest.tunnel_data:type_name -> google.protobuf.Any
	0, // 1: proto.AgentCTLResponse.status:type_name -> proto.AgentCTLResponseStatus
	1, // 2: proto.AgentCTLResponse.agentctls:type_name -> proto.AgentCTLRequest
	1, // 3: proto.AgentCTLService.ControlAgent:input_type -> proto.AgentCTLRequest
	2, // 4: proto.AgentCTLService.ControlAgent:output_type -> proto.AgentCTLResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_agentctl_proto_init() }
func file_agentctl_proto_init() {
	if File_agentctl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_agentctl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentCTLRequest); i {
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
		file_agentctl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentCTLResponse); i {
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
			RawDescriptor: file_agentctl_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_agentctl_proto_goTypes,
		DependencyIndexes: file_agentctl_proto_depIdxs,
		EnumInfos:         file_agentctl_proto_enumTypes,
		MessageInfos:      file_agentctl_proto_msgTypes,
	}.Build()
	File_agentctl_proto = out.File
	file_agentctl_proto_rawDesc = nil
	file_agentctl_proto_goTypes = nil
	file_agentctl_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AgentCTLServiceClient is the client API for AgentCTLService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentCTLServiceClient interface {
	//*
	// Used to control an Agent
	ControlAgent(ctx context.Context, in *AgentCTLRequest, opts ...grpc.CallOption) (*AgentCTLResponse, error)
}

type agentCTLServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentCTLServiceClient(cc grpc.ClientConnInterface) AgentCTLServiceClient {
	return &agentCTLServiceClient{cc}
}

func (c *agentCTLServiceClient) ControlAgent(ctx context.Context, in *AgentCTLRequest, opts ...grpc.CallOption) (*AgentCTLResponse, error) {
	out := new(AgentCTLResponse)
	err := c.cc.Invoke(ctx, "/proto.AgentCTLService/ControlAgent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentCTLServiceServer is the server API for AgentCTLService service.
type AgentCTLServiceServer interface {
	//*
	// Used to control an Agent
	ControlAgent(context.Context, *AgentCTLRequest) (*AgentCTLResponse, error)
}

// UnimplementedAgentCTLServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAgentCTLServiceServer struct {
}

func (*UnimplementedAgentCTLServiceServer) ControlAgent(context.Context, *AgentCTLRequest) (*AgentCTLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ControlAgent not implemented")
}

func RegisterAgentCTLServiceServer(s *grpc.Server, srv AgentCTLServiceServer) {
	s.RegisterService(&_AgentCTLService_serviceDesc, srv)
}

func _AgentCTLService_ControlAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentCTLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentCTLServiceServer).ControlAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AgentCTLService/ControlAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentCTLServiceServer).ControlAgent(ctx, req.(*AgentCTLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AgentCTLService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AgentCTLService",
	HandlerType: (*AgentCTLServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ControlAgent",
			Handler:    _AgentCTLService_ControlAgent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agentctl.proto",
}
