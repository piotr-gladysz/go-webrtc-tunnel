// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: tunnel.proto

package cliapi

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

type CreateTunnelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerId     string `protobuf:"bytes,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	LocalPort  uint32 `protobuf:"varint,2,opt,name=local_port,json=localPort,proto3" json:"local_port,omitempty"`
	RemoteHost string `protobuf:"bytes,3,opt,name=remote_host,json=remoteHost,proto3" json:"remote_host,omitempty"`
	RemotePort uint32 `protobuf:"varint,4,opt,name=remote_port,json=remotePort,proto3" json:"remote_port,omitempty"`
}

func (x *CreateTunnelRequest) Reset() {
	*x = CreateTunnelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tunnel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTunnelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTunnelRequest) ProtoMessage() {}

func (x *CreateTunnelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tunnel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTunnelRequest.ProtoReflect.Descriptor instead.
func (*CreateTunnelRequest) Descriptor() ([]byte, []int) {
	return file_tunnel_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTunnelRequest) GetPeerId() string {
	if x != nil {
		return x.PeerId
	}
	return ""
}

func (x *CreateTunnelRequest) GetLocalPort() uint32 {
	if x != nil {
		return x.LocalPort
	}
	return 0
}

func (x *CreateTunnelRequest) GetRemoteHost() string {
	if x != nil {
		return x.RemoteHost
	}
	return ""
}

func (x *CreateTunnelRequest) GetRemotePort() uint32 {
	if x != nil {
		return x.RemotePort
	}
	return 0
}

type TunnelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PeerId     string `protobuf:"bytes,2,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	LocalPort  uint32 `protobuf:"varint,3,opt,name=local_port,json=localPort,proto3" json:"local_port,omitempty"`
	RemotePort uint32 `protobuf:"varint,4,opt,name=remote_port,json=remotePort,proto3" json:"remote_port,omitempty"`
}

func (x *TunnelResponse) Reset() {
	*x = TunnelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tunnel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TunnelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TunnelResponse) ProtoMessage() {}

func (x *TunnelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tunnel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TunnelResponse.ProtoReflect.Descriptor instead.
func (*TunnelResponse) Descriptor() ([]byte, []int) {
	return file_tunnel_proto_rawDescGZIP(), []int{1}
}

func (x *TunnelResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TunnelResponse) GetPeerId() string {
	if x != nil {
		return x.PeerId
	}
	return ""
}

func (x *TunnelResponse) GetLocalPort() uint32 {
	if x != nil {
		return x.LocalPort
	}
	return 0
}

func (x *TunnelResponse) GetRemotePort() uint32 {
	if x != nil {
		return x.RemotePort
	}
	return 0
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tunnels []*TunnelResponse `protobuf:"bytes,1,rep,name=tunnels,proto3" json:"tunnels,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tunnel_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tunnel_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_tunnel_proto_rawDescGZIP(), []int{2}
}

func (x *ListResponse) GetTunnels() []*TunnelResponse {
	if x != nil {
		return x.Tunnels
	}
	return nil
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tunnel_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tunnel_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_tunnel_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_tunnel_proto protoreflect.FileDescriptor

var file_tunnel_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x63, 0x6c, 0x69, 0x61, 0x70, 0x69, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x75,
	0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70,
	0x65, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x50,
	0x6f, 0x72, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x68, 0x6f,
	0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x48, 0x6f, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x79, 0x0a, 0x0e, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x65, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x6f, 0x72, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74,
	0x22, 0x40, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x30, 0x0a, 0x07, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6c, 0x69, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x75, 0x6e, 0x6e, 0x65,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07, 0x74, 0x75, 0x6e, 0x6e, 0x65,
	0x6c, 0x73, 0x22, 0x1f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x32, 0xb6, 0x01, 0x0a, 0x06, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x3d,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x63, 0x6c, 0x69, 0x61, 0x70,
	0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x6c, 0x69, 0x61, 0x70, 0x69, 0x2e, 0x54,
	0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e,
	0x63, 0x6c, 0x69, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x15, 0x2e,
	0x63, 0x6c, 0x69, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x83, 0x01, 0x0a,
	0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6c, 0x69, 0x61, 0x70, 0x69, 0x42, 0x0b, 0x54, 0x75, 0x6e,
	0x6e, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x69, 0x6f, 0x74, 0x72, 0x2d, 0x67, 0x6c, 0x61,
	0x64, 0x79, 0x73, 0x7a, 0x2f, 0x67, 0x6f, 0x2d, 0x77, 0x65, 0x62, 0x72, 0x74, 0x63, 0x2d, 0x74,
	0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x63, 0x6c, 0x69, 0x61, 0x70, 0x69, 0xa2, 0x02, 0x03, 0x43,
	0x58, 0x58, 0xaa, 0x02, 0x06, 0x43, 0x6c, 0x69, 0x61, 0x70, 0x69, 0xca, 0x02, 0x06, 0x43, 0x6c,
	0x69, 0x61, 0x70, 0x69, 0xe2, 0x02, 0x12, 0x43, 0x6c, 0x69, 0x61, 0x70, 0x69, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x06, 0x43, 0x6c, 0x69, 0x61,
	0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tunnel_proto_rawDescOnce sync.Once
	file_tunnel_proto_rawDescData = file_tunnel_proto_rawDesc
)

func file_tunnel_proto_rawDescGZIP() []byte {
	file_tunnel_proto_rawDescOnce.Do(func() {
		file_tunnel_proto_rawDescData = protoimpl.X.CompressGZIP(file_tunnel_proto_rawDescData)
	})
	return file_tunnel_proto_rawDescData
}

var file_tunnel_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tunnel_proto_goTypes = []interface{}{
	(*CreateTunnelRequest)(nil), // 0: cliapi.CreateTunnelRequest
	(*TunnelResponse)(nil),      // 1: cliapi.TunnelResponse
	(*ListResponse)(nil),        // 2: cliapi.ListResponse
	(*DeleteRequest)(nil),       // 3: cliapi.DeleteRequest
	(*emptypb.Empty)(nil),       // 4: google.protobuf.Empty
}
var file_tunnel_proto_depIdxs = []int32{
	1, // 0: cliapi.ListResponse.tunnels:type_name -> cliapi.TunnelResponse
	0, // 1: cliapi.Tunnel.Create:input_type -> cliapi.CreateTunnelRequest
	4, // 2: cliapi.Tunnel.List:input_type -> google.protobuf.Empty
	3, // 3: cliapi.Tunnel.Delete:input_type -> cliapi.DeleteRequest
	1, // 4: cliapi.Tunnel.Create:output_type -> cliapi.TunnelResponse
	2, // 5: cliapi.Tunnel.List:output_type -> cliapi.ListResponse
	4, // 6: cliapi.Tunnel.Delete:output_type -> google.protobuf.Empty
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tunnel_proto_init() }
func file_tunnel_proto_init() {
	if File_tunnel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tunnel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTunnelRequest); i {
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
		file_tunnel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TunnelResponse); i {
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
		file_tunnel_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_tunnel_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
			RawDescriptor: file_tunnel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tunnel_proto_goTypes,
		DependencyIndexes: file_tunnel_proto_depIdxs,
		MessageInfos:      file_tunnel_proto_msgTypes,
	}.Build()
	File_tunnel_proto = out.File
	file_tunnel_proto_rawDesc = nil
	file_tunnel_proto_goTypes = nil
	file_tunnel_proto_depIdxs = nil
}
