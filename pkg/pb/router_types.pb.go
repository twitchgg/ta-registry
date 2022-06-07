// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: pkg/proto/router_types.proto

package pb

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

type RouterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WgConfig      []*WireguardConfig `protobuf:"bytes,1,rep,name=wgConfig,proto3" json:"wgConfig,omitempty"`
	IptablesRules []*IPTablesRules   `protobuf:"bytes,2,rep,name=iptablesRules,proto3" json:"iptablesRules,omitempty"`
}

func (x *RouterConfig) Reset() {
	*x = RouterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_router_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouterConfig) ProtoMessage() {}

func (x *RouterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_router_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouterConfig.ProtoReflect.Descriptor instead.
func (*RouterConfig) Descriptor() ([]byte, []int) {
	return file_pkg_proto_router_types_proto_rawDescGZIP(), []int{0}
}

func (x *RouterConfig) GetWgConfig() []*WireguardConfig {
	if x != nil {
		return x.WgConfig
	}
	return nil
}

func (x *RouterConfig) GetIptablesRules() []*IPTablesRules {
	if x != nil {
		return x.IptablesRules
	}
	return nil
}

type WireguardConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	InterfaceDef *WireguardInterface `protobuf:"bytes,2,opt,name=interfaceDef,proto3" json:"interfaceDef,omitempty"`
	Peers        []*WireguardPeer    `protobuf:"bytes,3,rep,name=peers,proto3" json:"peers,omitempty"`
}

func (x *WireguardConfig) Reset() {
	*x = WireguardConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_router_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WireguardConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WireguardConfig) ProtoMessage() {}

func (x *WireguardConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_router_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WireguardConfig.ProtoReflect.Descriptor instead.
func (*WireguardConfig) Descriptor() ([]byte, []int) {
	return file_pkg_proto_router_types_proto_rawDescGZIP(), []int{1}
}

func (x *WireguardConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WireguardConfig) GetInterfaceDef() *WireguardInterface {
	if x != nil {
		return x.InterfaceDef
	}
	return nil
}

func (x *WireguardConfig) GetPeers() []*WireguardPeer {
	if x != nil {
		return x.Peers
	}
	return nil
}

type WireguardInterface struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Port    int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	PrivKey string `protobuf:"bytes,3,opt,name=privKey,proto3" json:"privKey,omitempty"`
}

func (x *WireguardInterface) Reset() {
	*x = WireguardInterface{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_router_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WireguardInterface) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WireguardInterface) ProtoMessage() {}

func (x *WireguardInterface) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_router_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WireguardInterface.ProtoReflect.Descriptor instead.
func (*WireguardInterface) Descriptor() ([]byte, []int) {
	return file_pkg_proto_router_types_proto_rawDescGZIP(), []int{2}
}

func (x *WireguardInterface) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *WireguardInterface) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *WireguardInterface) GetPrivKey() string {
	if x != nil {
		return x.PrivKey
	}
	return ""
}

type WireguardPeer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Desc      string   `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	PubKey    string   `protobuf:"bytes,3,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	PsKey     string   `protobuf:"bytes,4,opt,name=psKey,proto3" json:"psKey,omitempty"`
	Endpoint  string   `protobuf:"bytes,5,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	AllowIPs  []string `protobuf:"bytes,6,rep,name=allowIPs,proto3" json:"allowIPs,omitempty"`
	Keepalive int32    `protobuf:"varint,7,opt,name=keepalive,proto3" json:"keepalive,omitempty"`
}

func (x *WireguardPeer) Reset() {
	*x = WireguardPeer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_router_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WireguardPeer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WireguardPeer) ProtoMessage() {}

func (x *WireguardPeer) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_router_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WireguardPeer.ProtoReflect.Descriptor instead.
func (*WireguardPeer) Descriptor() ([]byte, []int) {
	return file_pkg_proto_router_types_proto_rawDescGZIP(), []int{3}
}

func (x *WireguardPeer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WireguardPeer) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *WireguardPeer) GetPubKey() string {
	if x != nil {
		return x.PubKey
	}
	return ""
}

func (x *WireguardPeer) GetPsKey() string {
	if x != nil {
		return x.PsKey
	}
	return ""
}

func (x *WireguardPeer) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *WireguardPeer) GetAllowIPs() []string {
	if x != nil {
		return x.AllowIPs
	}
	return nil
}

func (x *WireguardPeer) GetKeepalive() int32 {
	if x != nil {
		return x.Keepalive
	}
	return 0
}

type IPTablesRules struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MasqueraedInterfaces []string `protobuf:"bytes,1,rep,name=masqueraedInterfaces,proto3" json:"masqueraedInterfaces,omitempty"`
	WhiteList            bool     `protobuf:"varint,2,opt,name=whiteList,proto3" json:"whiteList,omitempty"`
}

func (x *IPTablesRules) Reset() {
	*x = IPTablesRules{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_router_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPTablesRules) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPTablesRules) ProtoMessage() {}

func (x *IPTablesRules) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_router_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPTablesRules.ProtoReflect.Descriptor instead.
func (*IPTablesRules) Descriptor() ([]byte, []int) {
	return file_pkg_proto_router_types_proto_rawDescGZIP(), []int{4}
}

func (x *IPTablesRules) GetMasqueraedInterfaces() []string {
	if x != nil {
		return x.MasqueraedInterfaces
	}
	return nil
}

func (x *IPTablesRules) GetWhiteList() bool {
	if x != nil {
		return x.WhiteList
	}
	return false
}

var File_pkg_proto_router_types_proto protoreflect.FileDescriptor

var file_pkg_proto_router_types_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x74, 0x61, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22,
	0x92, 0x01, 0x0a, 0x0c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x3c, 0x0a, 0x08, 0x77, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x61, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x57, 0x69, 0x72, 0x65, 0x67, 0x75, 0x61, 0x72, 0x64, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x08, 0x77, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x44,
	0x0a, 0x0d, 0x69, 0x70, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x61, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x72, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x50, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73,
	0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x0d, 0x69, 0x70, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52,
	0x75, 0x6c, 0x65, 0x73, 0x22, 0xa4, 0x01, 0x0a, 0x0f, 0x57, 0x69, 0x72, 0x65, 0x67, 0x75, 0x61,
	0x72, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x47, 0x0a, 0x0c,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x44, 0x65, 0x66, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x74, 0x61, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x57, 0x69, 0x72, 0x65, 0x67, 0x75, 0x61, 0x72, 0x64, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x44, 0x65, 0x66, 0x12, 0x34, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x61, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x57, 0x69, 0x72, 0x65, 0x67, 0x75, 0x61, 0x72, 0x64,
	0x50, 0x65, 0x65, 0x72, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x22, 0x5c, 0x0a, 0x12, 0x57,
	0x69, 0x72, 0x65, 0x67, 0x75, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x76, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x72, 0x69, 0x76, 0x4b, 0x65, 0x79, 0x22, 0xbb, 0x01, 0x0a, 0x0d, 0x57, 0x69,
	0x72, 0x65, 0x67, 0x75, 0x61, 0x72, 0x64, 0x50, 0x65, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x65, 0x73, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x73, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x73, 0x4b, 0x65,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x49, 0x50, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x08, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x49, 0x50, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6b, 0x65, 0x65,
	0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6b, 0x65,
	0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x22, 0x61, 0x0a, 0x0d, 0x49, 0x50, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x73, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x14, 0x6d, 0x61, 0x73, 0x71,
	0x75, 0x65, 0x72, 0x61, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x14, 0x6d, 0x61, 0x73, 0x71, 0x75, 0x65, 0x72, 0x61,
	0x65, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x77, 0x68, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x77, 0x68, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x30, 0x0a, 0x13, 0x63, 0x6e,
	0x2e, 0x61, 0x63, 0x2e, 0x6e, 0x74, 0x73, 0x63, 0x2e, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x42, 0x10, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_router_types_proto_rawDescOnce sync.Once
	file_pkg_proto_router_types_proto_rawDescData = file_pkg_proto_router_types_proto_rawDesc
)

func file_pkg_proto_router_types_proto_rawDescGZIP() []byte {
	file_pkg_proto_router_types_proto_rawDescOnce.Do(func() {
		file_pkg_proto_router_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_router_types_proto_rawDescData)
	})
	return file_pkg_proto_router_types_proto_rawDescData
}

var file_pkg_proto_router_types_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_proto_router_types_proto_goTypes = []interface{}{
	(*RouterConfig)(nil),       // 0: ta.router.types.RouterConfig
	(*WireguardConfig)(nil),    // 1: ta.router.types.WireguardConfig
	(*WireguardInterface)(nil), // 2: ta.router.types.WireguardInterface
	(*WireguardPeer)(nil),      // 3: ta.router.types.WireguardPeer
	(*IPTablesRules)(nil),      // 4: ta.router.types.IPTablesRules
}
var file_pkg_proto_router_types_proto_depIdxs = []int32{
	1, // 0: ta.router.types.RouterConfig.wgConfig:type_name -> ta.router.types.WireguardConfig
	4, // 1: ta.router.types.RouterConfig.iptablesRules:type_name -> ta.router.types.IPTablesRules
	2, // 2: ta.router.types.WireguardConfig.interfaceDef:type_name -> ta.router.types.WireguardInterface
	3, // 3: ta.router.types.WireguardConfig.peers:type_name -> ta.router.types.WireguardPeer
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_proto_router_types_proto_init() }
func file_pkg_proto_router_types_proto_init() {
	if File_pkg_proto_router_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_router_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouterConfig); i {
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
		file_pkg_proto_router_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WireguardConfig); i {
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
		file_pkg_proto_router_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WireguardInterface); i {
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
		file_pkg_proto_router_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WireguardPeer); i {
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
		file_pkg_proto_router_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPTablesRules); i {
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
			RawDescriptor: file_pkg_proto_router_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_router_types_proto_goTypes,
		DependencyIndexes: file_pkg_proto_router_types_proto_depIdxs,
		MessageInfos:      file_pkg_proto_router_types_proto_msgTypes,
	}.Build()
	File_pkg_proto_router_types_proto = out.File
	file_pkg_proto_router_types_proto_rawDesc = nil
	file_pkg_proto_router_types_proto_goTypes = nil
	file_pkg_proto_router_types_proto_depIdxs = nil
}
