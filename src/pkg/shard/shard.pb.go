// Code generated by protoc-gen-go.
// source: pkg/shard/shard.proto
// DO NOT EDIT!

/*
Package shard is a generated protocol buffer package.

It is generated from these files:
	pkg/shard/shard.proto

It has these top-level messages:
	ServerState
	FrontendState
	ServerRole
	Addresses
	StartRegister
	FinishRegister
	Version
	StartAssignRoles
	FinishAssignRoles
	FailedToAssignRoles
	SetServerState
	SetFrontendState
	AddServerRole
	RemoveServerRole
	SetServerRole
	DeleteServerRole
	SetAddresses
	GetAddress
	GetShardToAddress
*/
package shard

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type ServerState struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Version int64  `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
}

func (m *ServerState) Reset()                    { *m = ServerState{} }
func (m *ServerState) String() string            { return proto.CompactTextString(m) }
func (*ServerState) ProtoMessage()               {}
func (*ServerState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type FrontendState struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Version int64  `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
}

func (m *FrontendState) Reset()                    { *m = FrontendState{} }
func (m *FrontendState) String() string            { return proto.CompactTextString(m) }
func (*FrontendState) ProtoMessage()               {}
func (*FrontendState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ServerRole struct {
	Address string          `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Version int64           `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
	Shards  map[uint64]bool `protobuf:"bytes,3,rep,name=shards" json:"shards,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *ServerRole) Reset()                    { *m = ServerRole{} }
func (m *ServerRole) String() string            { return proto.CompactTextString(m) }
func (*ServerRole) ProtoMessage()               {}
func (*ServerRole) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ServerRole) GetShards() map[uint64]bool {
	if m != nil {
		return m.Shards
	}
	return nil
}

type Addresses struct {
	Version   int64             `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Addresses map[uint64]string `protobuf:"bytes,2,rep,name=addresses" json:"addresses,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Addresses) Reset()                    { *m = Addresses{} }
func (m *Addresses) String() string            { return proto.CompactTextString(m) }
func (*Addresses) ProtoMessage()               {}
func (*Addresses) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Addresses) GetAddresses() map[uint64]string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type StartRegister struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
}

func (m *StartRegister) Reset()                    { *m = StartRegister{} }
func (m *StartRegister) String() string            { return proto.CompactTextString(m) }
func (*StartRegister) ProtoMessage()               {}
func (*StartRegister) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type FinishRegister struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *FinishRegister) Reset()                    { *m = FinishRegister{} }
func (m *FinishRegister) String() string            { return proto.CompactTextString(m) }
func (*FinishRegister) ProtoMessage()               {}
func (*FinishRegister) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type Version struct {
	Result int64  `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *Version) Reset()                    { *m = Version{} }
func (m *Version) String() string            { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()               {}
func (*Version) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type StartAssignRoles struct {
}

func (m *StartAssignRoles) Reset()                    { *m = StartAssignRoles{} }
func (m *StartAssignRoles) String() string            { return proto.CompactTextString(m) }
func (*StartAssignRoles) ProtoMessage()               {}
func (*StartAssignRoles) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type FinishAssignRoles struct {
	Error string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *FinishAssignRoles) Reset()                    { *m = FinishAssignRoles{} }
func (m *FinishAssignRoles) String() string            { return proto.CompactTextString(m) }
func (*FinishAssignRoles) ProtoMessage()               {}
func (*FinishAssignRoles) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type FailedToAssignRoles struct {
	ServerStates map[string]*ServerState `protobuf:"bytes,1,rep,name=server_states" json:"server_states,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	NumShards    uint64                  `protobuf:"varint,2,opt,name=num_shards" json:"num_shards,omitempty"`
	NumReplicas  uint64                  `protobuf:"varint,3,opt,name=num_replicas" json:"num_replicas,omitempty"`
}

func (m *FailedToAssignRoles) Reset()                    { *m = FailedToAssignRoles{} }
func (m *FailedToAssignRoles) String() string            { return proto.CompactTextString(m) }
func (*FailedToAssignRoles) ProtoMessage()               {}
func (*FailedToAssignRoles) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *FailedToAssignRoles) GetServerStates() map[string]*ServerState {
	if m != nil {
		return m.ServerStates
	}
	return nil
}

type SetServerState struct {
	ServerState *ServerState `protobuf:"bytes,1,opt,name=serverState" json:"serverState,omitempty"`
}

func (m *SetServerState) Reset()                    { *m = SetServerState{} }
func (m *SetServerState) String() string            { return proto.CompactTextString(m) }
func (*SetServerState) ProtoMessage()               {}
func (*SetServerState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *SetServerState) GetServerState() *ServerState {
	if m != nil {
		return m.ServerState
	}
	return nil
}

type SetFrontendState struct {
	FrontendState *FrontendState `protobuf:"bytes,1,opt,name=frontendState" json:"frontendState,omitempty"`
}

func (m *SetFrontendState) Reset()                    { *m = SetFrontendState{} }
func (m *SetFrontendState) String() string            { return proto.CompactTextString(m) }
func (*SetFrontendState) ProtoMessage()               {}
func (*SetFrontendState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *SetFrontendState) GetFrontendState() *FrontendState {
	if m != nil {
		return m.FrontendState
	}
	return nil
}

type AddServerRole struct {
	ServerRole *ServerRole `protobuf:"bytes,1,opt,name=serverRole" json:"serverRole,omitempty"`
	Error      string      `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *AddServerRole) Reset()                    { *m = AddServerRole{} }
func (m *AddServerRole) String() string            { return proto.CompactTextString(m) }
func (*AddServerRole) ProtoMessage()               {}
func (*AddServerRole) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *AddServerRole) GetServerRole() *ServerRole {
	if m != nil {
		return m.ServerRole
	}
	return nil
}

type RemoveServerRole struct {
	ServerRole *ServerRole `protobuf:"bytes,1,opt,name=serverRole" json:"serverRole,omitempty"`
	Error      string      `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *RemoveServerRole) Reset()                    { *m = RemoveServerRole{} }
func (m *RemoveServerRole) String() string            { return proto.CompactTextString(m) }
func (*RemoveServerRole) ProtoMessage()               {}
func (*RemoveServerRole) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *RemoveServerRole) GetServerRole() *ServerRole {
	if m != nil {
		return m.ServerRole
	}
	return nil
}

type SetServerRole struct {
	ServerRole *ServerRole `protobuf:"bytes,2,opt,name=serverRole" json:"serverRole,omitempty"`
}

func (m *SetServerRole) Reset()                    { *m = SetServerRole{} }
func (m *SetServerRole) String() string            { return proto.CompactTextString(m) }
func (*SetServerRole) ProtoMessage()               {}
func (*SetServerRole) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *SetServerRole) GetServerRole() *ServerRole {
	if m != nil {
		return m.ServerRole
	}
	return nil
}

type DeleteServerRole struct {
	ServerRole *ServerRole `protobuf:"bytes,2,opt,name=serverRole" json:"serverRole,omitempty"`
}

func (m *DeleteServerRole) Reset()                    { *m = DeleteServerRole{} }
func (m *DeleteServerRole) String() string            { return proto.CompactTextString(m) }
func (*DeleteServerRole) ProtoMessage()               {}
func (*DeleteServerRole) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *DeleteServerRole) GetServerRole() *ServerRole {
	if m != nil {
		return m.ServerRole
	}
	return nil
}

type SetAddresses struct {
	Addresses *Addresses `protobuf:"bytes,1,opt,name=addresses" json:"addresses,omitempty"`
}

func (m *SetAddresses) Reset()                    { *m = SetAddresses{} }
func (m *SetAddresses) String() string            { return proto.CompactTextString(m) }
func (*SetAddresses) ProtoMessage()               {}
func (*SetAddresses) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *SetAddresses) GetAddresses() *Addresses {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type GetAddress struct {
	Shard   uint64 `protobuf:"varint,1,opt,name=shard" json:"shard,omitempty"`
	Version int64  `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
	Result  string `protobuf:"bytes,3,opt,name=result" json:"result,omitempty"`
	Ok      bool   `protobuf:"varint,4,opt,name=ok" json:"ok,omitempty"`
	Error   string `protobuf:"bytes,5,opt,name=error" json:"error,omitempty"`
}

func (m *GetAddress) Reset()                    { *m = GetAddress{} }
func (m *GetAddress) String() string            { return proto.CompactTextString(m) }
func (*GetAddress) ProtoMessage()               {}
func (*GetAddress) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

type GetShardToAddress struct {
	Version int64             `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Result  map[uint64]string `protobuf:"bytes,2,rep,name=result" json:"result,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Error   string            `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *GetShardToAddress) Reset()                    { *m = GetShardToAddress{} }
func (m *GetShardToAddress) String() string            { return proto.CompactTextString(m) }
func (*GetShardToAddress) ProtoMessage()               {}
func (*GetShardToAddress) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *GetShardToAddress) GetResult() map[uint64]string {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*ServerState)(nil), "shard.ServerState")
	proto.RegisterType((*FrontendState)(nil), "shard.FrontendState")
	proto.RegisterType((*ServerRole)(nil), "shard.ServerRole")
	proto.RegisterType((*Addresses)(nil), "shard.Addresses")
	proto.RegisterType((*StartRegister)(nil), "shard.StartRegister")
	proto.RegisterType((*FinishRegister)(nil), "shard.FinishRegister")
	proto.RegisterType((*Version)(nil), "shard.Version")
	proto.RegisterType((*StartAssignRoles)(nil), "shard.StartAssignRoles")
	proto.RegisterType((*FinishAssignRoles)(nil), "shard.FinishAssignRoles")
	proto.RegisterType((*FailedToAssignRoles)(nil), "shard.FailedToAssignRoles")
	proto.RegisterType((*SetServerState)(nil), "shard.SetServerState")
	proto.RegisterType((*SetFrontendState)(nil), "shard.SetFrontendState")
	proto.RegisterType((*AddServerRole)(nil), "shard.AddServerRole")
	proto.RegisterType((*RemoveServerRole)(nil), "shard.RemoveServerRole")
	proto.RegisterType((*SetServerRole)(nil), "shard.SetServerRole")
	proto.RegisterType((*DeleteServerRole)(nil), "shard.DeleteServerRole")
	proto.RegisterType((*SetAddresses)(nil), "shard.SetAddresses")
	proto.RegisterType((*GetAddress)(nil), "shard.GetAddress")
	proto.RegisterType((*GetShardToAddress)(nil), "shard.GetShardToAddress")
}

var fileDescriptor0 = []byte{
	// 570 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x54, 0xd1, 0x6f, 0x12, 0x4f,
	0x10, 0xce, 0x41, 0xa1, 0x3f, 0x86, 0x1e, 0x3f, 0x58, 0x6b, 0x42, 0x9a, 0x18, 0x71, 0xd5, 0x48,
	0xa2, 0xa5, 0x4a, 0x8d, 0x51, 0x5f, 0x4c, 0xa3, 0xa5, 0x3e, 0x83, 0xd1, 0xc7, 0xe6, 0x94, 0x91,
	0x5e, 0xb8, 0xde, 0x91, 0xdd, 0x85, 0xa4, 0xaf, 0x3e, 0xfb, 0x07, 0xf8, 0x4f, 0xf9, 0x3f, 0xb9,
	0x3b, 0xbb, 0x70, 0x0b, 0x47, 0xb5, 0xc6, 0x17, 0xc2, 0xcd, 0xce, 0xf7, 0xcd, 0xb7, 0xdf, 0xce,
	0x0c, 0xdc, 0x9e, 0x4d, 0x27, 0x47, 0xf2, 0x22, 0x12, 0x63, 0xfb, 0xdb, 0x9b, 0x89, 0x4c, 0x65,
	0xac, 0x42, 0x1f, 0xfc, 0x08, 0xea, 0x23, 0x14, 0x0b, 0x14, 0x23, 0x15, 0x29, 0x64, 0xff, 0xc3,
	0x6e, 0x34, 0x1e, 0x0b, 0x94, 0xb2, 0x1d, 0x74, 0x82, 0x6e, 0xcd, 0x04, 0xf4, 0xa1, 0x8c, 0xb3,
	0xb4, 0x5d, 0xd2, 0x81, 0x32, 0x7f, 0x06, 0xe1, 0x40, 0x64, 0xa9, 0xc2, 0x74, 0x7c, 0x53, 0xc8,
	0xf7, 0x00, 0xc0, 0x16, 0x19, 0x66, 0xc9, 0x0d, 0x00, 0xec, 0x10, 0xaa, 0xa4, 0x4e, 0xb6, 0xcb,
	0x9d, 0x72, 0xb7, 0xde, 0xbf, 0xd3, 0xb3, 0xca, 0x73, 0x92, 0xde, 0x88, 0xce, 0x4f, 0x53, 0x25,
	0xae, 0x0e, 0x0e, 0xf5, 0x1d, 0xf2, 0x4f, 0x56, 0x87, 0xf2, 0x14, 0xaf, 0x88, 0x7b, 0x87, 0x85,
	0x50, 0x59, 0x44, 0xc9, 0x1c, 0x89, 0xf9, 0xbf, 0xd7, 0xa5, 0x97, 0x01, 0xff, 0x16, 0x40, 0xed,
	0xc4, 0x0a, 0x40, 0xe9, 0x17, 0x0f, 0xa8, 0x78, 0x1f, 0x6a, 0xd1, 0xf2, 0x54, 0xa3, 0x4c, 0xfd,
	0xbb, 0xae, 0xfe, 0x0a, 0x95, 0xff, 0xb3, 0x0a, 0x9e, 0x42, 0x63, 0x3d, 0xf2, 0x1b, 0x11, 0x35,
	0x12, 0xd1, 0x81, 0x50, 0xdb, 0x27, 0xd4, 0x10, 0x27, 0xb1, 0x54, 0x28, 0x0a, 0xae, 0x70, 0xcd,
	0x39, 0x88, 0xd3, 0x58, 0x5e, 0x5c, 0x9b, 0x62, 0x78, 0x51, 0x88, 0x4c, 0x58, 0x5e, 0xde, 0x85,
	0xdd, 0x8f, 0xf6, 0x2a, 0xac, 0x01, 0x55, 0x9d, 0x37, 0x4f, 0x94, 0xbb, 0xd4, 0x46, 0x26, 0x83,
	0x26, 0x55, 0x3f, 0x91, 0x32, 0x9e, 0xa4, 0xc6, 0x50, 0xc9, 0x39, 0xb4, 0x6c, 0x3d, 0x2f, 0x98,
	0xe3, 0xac, 0xa6, 0x9f, 0x01, 0xdc, 0x1a, 0x44, 0x71, 0x82, 0xe3, 0x0f, 0x99, 0x9f, 0xf6, 0x16,
	0x42, 0x49, 0x6f, 0x73, 0x2e, 0x4d, 0x4f, 0x18, 0x7d, 0xc6, 0xb7, 0x27, 0xce, 0xb7, 0x2d, 0x90,
	0x9e, 0xd7, 0x75, 0xce, 0x32, 0x06, 0x90, 0xce, 0x2f, 0xcf, 0xdd, 0xcb, 0x97, 0xc8, 0xb9, 0x7d,
	0xd8, 0x33, 0x31, 0x81, 0xb3, 0x24, 0xfe, 0x12, 0x99, 0x7e, 0xd0, 0xd1, 0x83, 0x33, 0x68, 0x15,
	0xe1, 0x9e, 0xe3, 0x35, 0x76, 0xcf, 0x77, 0xbc, 0xde, 0x67, 0x6b, 0x0d, 0x44, 0x28, 0x7a, 0x85,
	0x57, 0xd0, 0x18, 0xa1, 0xf2, 0x07, 0xe0, 0x11, 0xd4, 0x65, 0xfe, 0x49, 0x6c, 0x5b, 0xe1, 0xfc,
	0x8d, 0xb6, 0x10, 0xd5, 0xfa, 0x28, 0x3c, 0x86, 0xf0, 0xab, 0x1f, 0x70, 0xf0, 0xfd, 0xa5, 0x0d,
	0xfe, 0x19, 0x3f, 0x85, 0x50, 0xf7, 0x8c, 0x37, 0x17, 0x0f, 0x01, 0xe4, 0xea, 0xcb, 0x41, 0x5b,
	0x85, 0xce, 0xdf, 0x7c, 0xca, 0xf7, 0xd0, 0x1c, 0xe2, 0x65, 0xb6, 0xc0, 0x7f, 0x66, 0x7a, 0xa1,
	0x5b, 0x72, 0x69, 0xc6, 0x16, 0x9a, 0xd2, 0x35, 0x34, 0xda, 0xc4, 0xe6, 0x3b, 0x4c, 0x50, 0xe1,
	0xdf, 0x43, 0x8f, 0x61, 0x4f, 0x97, 0xcc, 0x87, 0xf1, 0xbe, 0x3f, 0x7b, 0x56, 0x77, 0x73, 0x73,
	0xf6, 0xf8, 0x27, 0x80, 0xb3, 0x15, 0xc8, 0x5c, 0x82, 0x12, 0xdc, 0xa8, 0x15, 0x76, 0x49, 0x3e,
	0x09, 0x65, 0xea, 0x0c, 0x80, 0x52, 0x36, 0x6d, 0xef, 0x98, 0x6d, 0x90, 0x1b, 0x50, 0x21, 0x03,
	0x7e, 0x04, 0xd0, 0xd2, 0xcc, 0xb4, 0x4b, 0x74, 0xb3, 0xba, 0x02, 0x85, 0x05, 0xf1, 0x7c, 0xc5,
	0x68, 0xb7, 0xc3, 0x03, 0xa7, 0xb0, 0x00, 0xed, 0x0d, 0x29, 0xcd, 0xb6, 0xe7, 0xaa, 0x16, 0xc9,
	0x30, 0x3b, 0xcb, 0x3f, 0xfd, 0xc3, 0xba, 0xf8, 0x5c, 0xa5, 0xa5, 0x7d, 0xfc, 0x2b, 0x00, 0x00,
	0xff, 0xff, 0xf1, 0x6f, 0x86, 0xcd, 0xcd, 0x05, 0x00, 0x00,
}
