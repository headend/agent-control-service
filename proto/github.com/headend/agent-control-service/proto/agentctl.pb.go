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
	AgentIp     string   `protobuf:"bytes,1,opt,name=agent_ip,proto3" json:"agent_ip,omitempty"`
	ProfileIp   int64    `protobuf:"varint,2,opt,name=profile_ip,proto3" json:"profile_ip,omitempty"`
	ControlId   int64    `protobuf:"varint,3,opt,name=control_id,proto3" json:"control_id,omitempty"`
	ControlType int64    `protobuf:"varint,4,opt,name=control_type,proto3" json:"control_type,omitempty"`
	TunnelData  []string `protobuf:"bytes,5,rep,name=tunnel_data,proto3" json:"tunnel_data,omitempty"`
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

func (x *AgentCTLRequest) GetAgentIp() string {
	if x != nil {
		return x.AgentIp
	}
	return ""
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

func (x *AgentCTLRequest) GetTunnelData() []string {
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
	AgentResponseStatus bool `protobuf:"varint,1,opt,name=AgentResponseStatus,json=status,proto3" json:"AgentResponseStatus,omitempty"`
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

func (x *AgentCTLResponse) GetAgentResponseStatus() bool {
	if x != nil {
		return x.AgentResponseStatus
	}
	return false
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
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x01, 0x0a, 0x0f, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x43, 0x54, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74,
	0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0b, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x22, 0x68, 0x0a,
	0x10, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x54, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x23, 0x0a, 0x13, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a, 0x09, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x63,
	0x74, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x54, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x2f, 0x0a, 0x16, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x43, 0x54, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53,
	0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x32, 0xb6, 0x03, 0x0a, 0x0f, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x43, 0x54, 0x4c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0b,
	0x53, 0x54, 0x41, 0x52, 0x54, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x54, 0x4c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x43, 0x54, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f,
	0x0a, 0x0a, 0x53, 0x54, 0x4f, 0x50, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x54, 0x4c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x43, 0x54, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x41, 0x0a, 0x0c, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12,
	0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x54, 0x4c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x54, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x48, 0x0a, 0x13, 0x52, 0x45, 0x46, 0x45, 0x53, 0x48, 0x4d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x54, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43,
	0x54, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x13,
	0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x43, 0x54, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x54, 0x4c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x14, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c,
	0x45, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x16,
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
}
var file_agentctl_proto_depIdxs = []int32{
	1, // 0: proto.AgentCTLResponse.agentctls:type_name -> proto.AgentCTLRequest
	1, // 1: proto.AgentCTLService.STARTWorker:input_type -> proto.AgentCTLRequest
	1, // 2: proto.AgentCTLService.STOPWorker:input_type -> proto.AgentCTLRequest
	1, // 3: proto.AgentCTLService.UPDATEWorker:input_type -> proto.AgentCTLRequest
	1, // 4: proto.AgentCTLService.REFESHMasterConnect:input_type -> proto.AgentCTLRequest
	1, // 5: proto.AgentCTLService.ENABLEMasterConnect:input_type -> proto.AgentCTLRequest
	1, // 6: proto.AgentCTLService.DISABLEMasterConnect:input_type -> proto.AgentCTLRequest
	2, // 7: proto.AgentCTLService.STARTWorker:output_type -> proto.AgentCTLResponse
	2, // 8: proto.AgentCTLService.STOPWorker:output_type -> proto.AgentCTLResponse
	2, // 9: proto.AgentCTLService.UPDATEWorker:output_type -> proto.AgentCTLResponse
	2, // 10: proto.AgentCTLService.REFESHMasterConnect:output_type -> proto.AgentCTLResponse
	2, // 11: proto.AgentCTLService.ENABLEMasterConnect:output_type -> proto.AgentCTLResponse
	2, // 12: proto.AgentCTLService.DISABLEMasterConnect:output_type -> proto.AgentCTLResponse
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
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
	// Used to start an Agent
	STARTWorker(ctx context.Context, in *AgentCTLRequest, opts ...grpc.CallOption) (*AgentCTLResponse, error)
	//*
	// Used to stop an Agent
	STOPWorker(ctx context.Context, in *AgentCTLRequest, opts ...grpc.CallOption) (*AgentCTLResponse, error)
	//*
	// Used to stop an Agent
	UPDATEWorker(ctx context.Context, in *AgentCTLRequest, opts ...grpc.CallOption) (*AgentCTLResponse, error)
	//*
	// Used to restart an Agent
	REFESHMasterConnect(ctx context.Context, in *AgentCTLRequest, opts ...grpc.CallOption) (*AgentCTLResponse, error)
	//*
	// Used to enable connect to an Agent,allow agent connect to agent's backend
	ENABLEMasterConnect(ctx context.Context, in *AgentCTLRequest, opts ...grpc.CallOption) (*AgentCTLResponse, error)
	//*
	// Used to disable connect to an Agent, reject agent connect to agent's backend
	DISABLEMasterConnect(ctx context.Context, in *AgentCTLRequest, opts ...grpc.CallOption) (*AgentCTLResponse, error)
}

type agentCTLServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentCTLServiceClient(cc grpc.ClientConnInterface) AgentCTLServiceClient {
	return &agentCTLServiceClient{cc}
}

func (c *agentCTLServiceClient) STARTWorker(ctx context.Context, in *AgentCTLRequest, opts ...grpc.CallOption) (*AgentCTLResponse, error) {
	out := new(AgentCTLResponse)
	err := c.cc.Invoke(ctx, "/proto.AgentCTLService/STARTWorker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentCTLServiceClient) STOPWorker(ctx context.Context, in *AgentCTLRequest, opts ...grpc.CallOption) (*AgentCTLResponse, error) {
	out := new(AgentCTLResponse)
	err := c.cc.Invoke(ctx, "/proto.AgentCTLService/STOPWorker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentCTLServiceClient) UPDATEWorker(ctx context.Context, in *AgentCTLRequest, opts ...grpc.CallOption) (*AgentCTLResponse, error) {
	out := new(AgentCTLResponse)
	err := c.cc.Invoke(ctx, "/proto.AgentCTLService/UPDATEWorker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentCTLServiceClient) REFESHMasterConnect(ctx context.Context, in *AgentCTLRequest, opts ...grpc.CallOption) (*AgentCTLResponse, error) {
	out := new(AgentCTLResponse)
	err := c.cc.Invoke(ctx, "/proto.AgentCTLService/REFESHMasterConnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentCTLServiceClient) ENABLEMasterConnect(ctx context.Context, in *AgentCTLRequest, opts ...grpc.CallOption) (*AgentCTLResponse, error) {
	out := new(AgentCTLResponse)
	err := c.cc.Invoke(ctx, "/proto.AgentCTLService/ENABLEMasterConnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentCTLServiceClient) DISABLEMasterConnect(ctx context.Context, in *AgentCTLRequest, opts ...grpc.CallOption) (*AgentCTLResponse, error) {
	out := new(AgentCTLResponse)
	err := c.cc.Invoke(ctx, "/proto.AgentCTLService/DISABLEMasterConnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentCTLServiceServer is the server API for AgentCTLService service.
type AgentCTLServiceServer interface {
	//*
	// Used to start an Agent
	STARTWorker(context.Context, *AgentCTLRequest) (*AgentCTLResponse, error)
	//*
	// Used to stop an Agent
	STOPWorker(context.Context, *AgentCTLRequest) (*AgentCTLResponse, error)
	//*
	// Used to stop an Agent
	UPDATEWorker(context.Context, *AgentCTLRequest) (*AgentCTLResponse, error)
	//*
	// Used to restart an Agent
	REFESHMasterConnect(context.Context, *AgentCTLRequest) (*AgentCTLResponse, error)
	//*
	// Used to enable connect to an Agent,allow agent connect to agent's backend
	ENABLEMasterConnect(context.Context, *AgentCTLRequest) (*AgentCTLResponse, error)
	//*
	// Used to disable connect to an Agent, reject agent connect to agent's backend
	DISABLEMasterConnect(context.Context, *AgentCTLRequest) (*AgentCTLResponse, error)
}

// UnimplementedAgentCTLServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAgentCTLServiceServer struct {
}

func (*UnimplementedAgentCTLServiceServer) STARTWorker(context.Context, *AgentCTLRequest) (*AgentCTLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method STARTWorker not implemented")
}
func (*UnimplementedAgentCTLServiceServer) STOPWorker(context.Context, *AgentCTLRequest) (*AgentCTLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method STOPWorker not implemented")
}
func (*UnimplementedAgentCTLServiceServer) UPDATEWorker(context.Context, *AgentCTLRequest) (*AgentCTLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UPDATEWorker not implemented")
}
func (*UnimplementedAgentCTLServiceServer) REFESHMasterConnect(context.Context, *AgentCTLRequest) (*AgentCTLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method REFESHMasterConnect not implemented")
}
func (*UnimplementedAgentCTLServiceServer) ENABLEMasterConnect(context.Context, *AgentCTLRequest) (*AgentCTLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ENABLEMasterConnect not implemented")
}
func (*UnimplementedAgentCTLServiceServer) DISABLEMasterConnect(context.Context, *AgentCTLRequest) (*AgentCTLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DISABLEMasterConnect not implemented")
}

func RegisterAgentCTLServiceServer(s *grpc.Server, srv AgentCTLServiceServer) {
	s.RegisterService(&_AgentCTLService_serviceDesc, srv)
}

func _AgentCTLService_STARTWorker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentCTLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentCTLServiceServer).STARTWorker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AgentCTLService/STARTWorker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentCTLServiceServer).STARTWorker(ctx, req.(*AgentCTLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentCTLService_STOPWorker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentCTLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentCTLServiceServer).STOPWorker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AgentCTLService/STOPWorker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentCTLServiceServer).STOPWorker(ctx, req.(*AgentCTLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentCTLService_UPDATEWorker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentCTLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentCTLServiceServer).UPDATEWorker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AgentCTLService/UPDATEWorker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentCTLServiceServer).UPDATEWorker(ctx, req.(*AgentCTLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentCTLService_REFESHMasterConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentCTLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentCTLServiceServer).REFESHMasterConnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AgentCTLService/REFESHMasterConnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentCTLServiceServer).REFESHMasterConnect(ctx, req.(*AgentCTLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentCTLService_ENABLEMasterConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentCTLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentCTLServiceServer).ENABLEMasterConnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AgentCTLService/ENABLEMasterConnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentCTLServiceServer).ENABLEMasterConnect(ctx, req.(*AgentCTLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentCTLService_DISABLEMasterConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentCTLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentCTLServiceServer).DISABLEMasterConnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AgentCTLService/DISABLEMasterConnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentCTLServiceServer).DISABLEMasterConnect(ctx, req.(*AgentCTLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AgentCTLService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AgentCTLService",
	HandlerType: (*AgentCTLServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "STARTWorker",
			Handler:    _AgentCTLService_STARTWorker_Handler,
		},
		{
			MethodName: "STOPWorker",
			Handler:    _AgentCTLService_STOPWorker_Handler,
		},
		{
			MethodName: "UPDATEWorker",
			Handler:    _AgentCTLService_UPDATEWorker_Handler,
		},
		{
			MethodName: "REFESHMasterConnect",
			Handler:    _AgentCTLService_REFESHMasterConnect_Handler,
		},
		{
			MethodName: "ENABLEMasterConnect",
			Handler:    _AgentCTLService_ENABLEMasterConnect_Handler,
		},
		{
			MethodName: "DISABLEMasterConnect",
			Handler:    _AgentCTLService_DISABLEMasterConnect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agentctl.proto",
}
