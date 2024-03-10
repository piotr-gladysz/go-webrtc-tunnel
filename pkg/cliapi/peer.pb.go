// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: peer.proto

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

type SetPeerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Ports   []uint32 `protobuf:"varint,2,rep,packed,name=ports,proto3" json:"ports,omitempty"`
	Connect bool     `protobuf:"varint,3,opt,name=connect,proto3" json:"connect,omitempty"`
}

func (x *SetPeerRequest) Reset() {
	*x = SetPeerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPeerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPeerRequest) ProtoMessage() {}

func (x *SetPeerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_peer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPeerRequest.ProtoReflect.Descriptor instead.
func (*SetPeerRequest) Descriptor() ([]byte, []int) {
	return file_peer_proto_rawDescGZIP(), []int{0}
}

func (x *SetPeerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SetPeerRequest) GetPorts() []uint32 {
	if x != nil {
		return x.Ports
	}
	return nil
}

func (x *SetPeerRequest) GetConnect() bool {
	if x != nil {
		return x.Connect
	}
	return false
}

type RemovePeerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RemovePeerRequest) Reset() {
	*x = RemovePeerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovePeerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePeerRequest) ProtoMessage() {}

func (x *RemovePeerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_peer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePeerRequest.ProtoReflect.Descriptor instead.
func (*RemovePeerRequest) Descriptor() ([]byte, []int) {
	return file_peer_proto_rawDescGZIP(), []int{1}
}

func (x *RemovePeerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PeerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Ports       []string `protobuf:"bytes,2,rep,name=ports,proto3" json:"ports,omitempty"`
	LocalPorts  []string `protobuf:"bytes,3,rep,name=local_ports,json=localPorts,proto3" json:"local_ports,omitempty"`
	RemotePorts []string `protobuf:"bytes,4,rep,name=remote_ports,json=remotePorts,proto3" json:"remote_ports,omitempty"`
	Connected   bool     `protobuf:"varint,5,opt,name=connected,proto3" json:"connected,omitempty"`
}

func (x *PeerResponse) Reset() {
	*x = PeerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerResponse) ProtoMessage() {}

func (x *PeerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_peer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerResponse.ProtoReflect.Descriptor instead.
func (*PeerResponse) Descriptor() ([]byte, []int) {
	return file_peer_proto_rawDescGZIP(), []int{2}
}

func (x *PeerResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PeerResponse) GetPorts() []string {
	if x != nil {
		return x.Ports
	}
	return nil
}

func (x *PeerResponse) GetLocalPorts() []string {
	if x != nil {
		return x.LocalPorts
	}
	return nil
}

func (x *PeerResponse) GetRemotePorts() []string {
	if x != nil {
		return x.RemotePorts
	}
	return nil
}

func (x *PeerResponse) GetConnected() bool {
	if x != nil {
		return x.Connected
	}
	return false
}

type PeerList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peers []*PeerResponse `protobuf:"bytes,1,rep,name=peers,proto3" json:"peers,omitempty"`
}

func (x *PeerList) Reset() {
	*x = PeerList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerList) ProtoMessage() {}

func (x *PeerList) ProtoReflect() protoreflect.Message {
	mi := &file_peer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerList.ProtoReflect.Descriptor instead.
func (*PeerList) Descriptor() ([]byte, []int) {
	return file_peer_proto_rawDescGZIP(), []int{3}
}

func (x *PeerList) GetPeers() []*PeerResponse {
	if x != nil {
		return x.Peers
	}
	return nil
}

var File_peer_proto protoreflect.FileDescriptor

var file_peer_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6c,
	0x69, 0x61, 0x70, 0x69, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x50, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x22, 0x23, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x65, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x96, 0x01, 0x0a, 0x0c, 0x50, 0x65, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x72,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x6f, 0x72, 0x74, 0x73,
	0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x6f,
	0x72, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x22, 0x36, 0x0a, 0x08, 0x50, 0x65, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2a, 0x0a,
	0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63,
	0x6c, 0x69, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x32, 0xb6, 0x01, 0x0a, 0x04, 0x50, 0x65,
	0x65, 0x72, 0x12, 0x34, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x50, 0x65, 0x65, 0x72, 0x73, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x63, 0x6c, 0x69, 0x61, 0x70, 0x69, 0x2e,
	0x50, 0x65, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x50,
	0x65, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x63, 0x6c, 0x69, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x74,
	0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x6c,
	0x69, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x65, 0x65, 0x72, 0x12,
	0x19, 0x2e, 0x63, 0x6c, 0x69, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50,
	0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x42, 0x81, 0x01, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6c, 0x69, 0x61, 0x70,
	0x69, 0x42, 0x09, 0x50, 0x65, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x69, 0x6f, 0x74, 0x72,
	0x2d, 0x67, 0x6c, 0x61, 0x64, 0x79, 0x73, 0x7a, 0x2f, 0x67, 0x6f, 0x2d, 0x77, 0x65, 0x62, 0x72,
	0x74, 0x63, 0x2d, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x63, 0x6c, 0x69, 0x61, 0x70, 0x69,
	0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x43, 0x6c, 0x69, 0x61, 0x70, 0x69, 0xca,
	0x02, 0x06, 0x43, 0x6c, 0x69, 0x61, 0x70, 0x69, 0xe2, 0x02, 0x12, 0x43, 0x6c, 0x69, 0x61, 0x70,
	0x69, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x06,
	0x43, 0x6c, 0x69, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_peer_proto_rawDescOnce sync.Once
	file_peer_proto_rawDescData = file_peer_proto_rawDesc
)

func file_peer_proto_rawDescGZIP() []byte {
	file_peer_proto_rawDescOnce.Do(func() {
		file_peer_proto_rawDescData = protoimpl.X.CompressGZIP(file_peer_proto_rawDescData)
	})
	return file_peer_proto_rawDescData
}

var file_peer_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_peer_proto_goTypes = []interface{}{
	(*SetPeerRequest)(nil),    // 0: cliapi.SetPeerRequest
	(*RemovePeerRequest)(nil), // 1: cliapi.RemovePeerRequest
	(*PeerResponse)(nil),      // 2: cliapi.PeerResponse
	(*PeerList)(nil),          // 3: cliapi.PeerList
	(*emptypb.Empty)(nil),     // 4: google.protobuf.Empty
}
var file_peer_proto_depIdxs = []int32{
	2, // 0: cliapi.PeerList.peers:type_name -> cliapi.PeerResponse
	4, // 1: cliapi.Peer.GetPeers:input_type -> google.protobuf.Empty
	0, // 2: cliapi.Peer.SetPeer:input_type -> cliapi.SetPeerRequest
	1, // 3: cliapi.Peer.RemovePeer:input_type -> cliapi.RemovePeerRequest
	3, // 4: cliapi.Peer.GetPeers:output_type -> cliapi.PeerList
	2, // 5: cliapi.Peer.SetPeer:output_type -> cliapi.PeerResponse
	4, // 6: cliapi.Peer.RemovePeer:output_type -> google.protobuf.Empty
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_peer_proto_init() }
func file_peer_proto_init() {
	if File_peer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_peer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPeerRequest); i {
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
		file_peer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovePeerRequest); i {
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
		file_peer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerResponse); i {
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
		file_peer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerList); i {
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
			RawDescriptor: file_peer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_peer_proto_goTypes,
		DependencyIndexes: file_peer_proto_depIdxs,
		MessageInfos:      file_peer_proto_msgTypes,
	}.Build()
	File_peer_proto = out.File
	file_peer_proto_rawDesc = nil
	file_peer_proto_goTypes = nil
	file_peer_proto_depIdxs = nil
}
