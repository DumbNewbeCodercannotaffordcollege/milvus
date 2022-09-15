// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proxy.proto

package proxypb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	commonpb "github.com/milvus-io/milvus/api/commonpb"
	milvuspb "github.com/milvus-io/milvus/api/milvuspb"
	internalpb "github.com/milvus-io/milvus/internal/proto/internalpb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type InvalidateCollMetaCacheRequest struct {
	Base                 *commonpb.MsgBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	DbName               string            `protobuf:"bytes,2,opt,name=db_name,json=dbName,proto3" json:"db_name,omitempty"`
	CollectionName       string            `protobuf:"bytes,3,opt,name=collection_name,json=collectionName,proto3" json:"collection_name,omitempty"`
	CollectionID         int64             `protobuf:"varint,4,opt,name=collectionID,proto3" json:"collectionID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *InvalidateCollMetaCacheRequest) Reset()         { *m = InvalidateCollMetaCacheRequest{} }
func (m *InvalidateCollMetaCacheRequest) String() string { return proto.CompactTextString(m) }
func (*InvalidateCollMetaCacheRequest) ProtoMessage()    {}
func (*InvalidateCollMetaCacheRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_700b50b08ed8dbaf, []int{0}
}

func (m *InvalidateCollMetaCacheRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvalidateCollMetaCacheRequest.Unmarshal(m, b)
}
func (m *InvalidateCollMetaCacheRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvalidateCollMetaCacheRequest.Marshal(b, m, deterministic)
}
func (m *InvalidateCollMetaCacheRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvalidateCollMetaCacheRequest.Merge(m, src)
}
func (m *InvalidateCollMetaCacheRequest) XXX_Size() int {
	return xxx_messageInfo_InvalidateCollMetaCacheRequest.Size(m)
}
func (m *InvalidateCollMetaCacheRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InvalidateCollMetaCacheRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InvalidateCollMetaCacheRequest proto.InternalMessageInfo

func (m *InvalidateCollMetaCacheRequest) GetBase() *commonpb.MsgBase {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *InvalidateCollMetaCacheRequest) GetDbName() string {
	if m != nil {
		return m.DbName
	}
	return ""
}

func (m *InvalidateCollMetaCacheRequest) GetCollectionName() string {
	if m != nil {
		return m.CollectionName
	}
	return ""
}

func (m *InvalidateCollMetaCacheRequest) GetCollectionID() int64 {
	if m != nil {
		return m.CollectionID
	}
	return 0
}

type InvalidateCredCacheRequest struct {
	Base                 *commonpb.MsgBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Username             string            `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *InvalidateCredCacheRequest) Reset()         { *m = InvalidateCredCacheRequest{} }
func (m *InvalidateCredCacheRequest) String() string { return proto.CompactTextString(m) }
func (*InvalidateCredCacheRequest) ProtoMessage()    {}
func (*InvalidateCredCacheRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_700b50b08ed8dbaf, []int{1}
}

func (m *InvalidateCredCacheRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvalidateCredCacheRequest.Unmarshal(m, b)
}
func (m *InvalidateCredCacheRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvalidateCredCacheRequest.Marshal(b, m, deterministic)
}
func (m *InvalidateCredCacheRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvalidateCredCacheRequest.Merge(m, src)
}
func (m *InvalidateCredCacheRequest) XXX_Size() int {
	return xxx_messageInfo_InvalidateCredCacheRequest.Size(m)
}
func (m *InvalidateCredCacheRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InvalidateCredCacheRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InvalidateCredCacheRequest proto.InternalMessageInfo

func (m *InvalidateCredCacheRequest) GetBase() *commonpb.MsgBase {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *InvalidateCredCacheRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type UpdateCredCacheRequest struct {
	Base     *commonpb.MsgBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Username string            `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// password stored in cache
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateCredCacheRequest) Reset()         { *m = UpdateCredCacheRequest{} }
func (m *UpdateCredCacheRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateCredCacheRequest) ProtoMessage()    {}
func (*UpdateCredCacheRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_700b50b08ed8dbaf, []int{2}
}

func (m *UpdateCredCacheRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCredCacheRequest.Unmarshal(m, b)
}
func (m *UpdateCredCacheRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCredCacheRequest.Marshal(b, m, deterministic)
}
func (m *UpdateCredCacheRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCredCacheRequest.Merge(m, src)
}
func (m *UpdateCredCacheRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateCredCacheRequest.Size(m)
}
func (m *UpdateCredCacheRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCredCacheRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCredCacheRequest proto.InternalMessageInfo

func (m *UpdateCredCacheRequest) GetBase() *commonpb.MsgBase {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *UpdateCredCacheRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UpdateCredCacheRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RefreshPolicyInfoCacheRequest struct {
	Base                 *commonpb.MsgBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	OpType               int32             `protobuf:"varint,2,opt,name=opType,proto3" json:"opType,omitempty"`
	OpKey                string            `protobuf:"bytes,3,opt,name=opKey,proto3" json:"opKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RefreshPolicyInfoCacheRequest) Reset()         { *m = RefreshPolicyInfoCacheRequest{} }
func (m *RefreshPolicyInfoCacheRequest) String() string { return proto.CompactTextString(m) }
func (*RefreshPolicyInfoCacheRequest) ProtoMessage()    {}
func (*RefreshPolicyInfoCacheRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_700b50b08ed8dbaf, []int{3}
}

func (m *RefreshPolicyInfoCacheRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshPolicyInfoCacheRequest.Unmarshal(m, b)
}
func (m *RefreshPolicyInfoCacheRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshPolicyInfoCacheRequest.Marshal(b, m, deterministic)
}
func (m *RefreshPolicyInfoCacheRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshPolicyInfoCacheRequest.Merge(m, src)
}
func (m *RefreshPolicyInfoCacheRequest) XXX_Size() int {
	return xxx_messageInfo_RefreshPolicyInfoCacheRequest.Size(m)
}
func (m *RefreshPolicyInfoCacheRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshPolicyInfoCacheRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshPolicyInfoCacheRequest proto.InternalMessageInfo

func (m *RefreshPolicyInfoCacheRequest) GetBase() *commonpb.MsgBase {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *RefreshPolicyInfoCacheRequest) GetOpType() int32 {
	if m != nil {
		return m.OpType
	}
	return 0
}

func (m *RefreshPolicyInfoCacheRequest) GetOpKey() string {
	if m != nil {
		return m.OpKey
	}
	return ""
}

func init() {
	proto.RegisterType((*InvalidateCollMetaCacheRequest)(nil), "milvus.proto.proxy.InvalidateCollMetaCacheRequest")
	proto.RegisterType((*InvalidateCredCacheRequest)(nil), "milvus.proto.proxy.InvalidateCredCacheRequest")
	proto.RegisterType((*UpdateCredCacheRequest)(nil), "milvus.proto.proxy.UpdateCredCacheRequest")
	proto.RegisterType((*RefreshPolicyInfoCacheRequest)(nil), "milvus.proto.proxy.RefreshPolicyInfoCacheRequest")
}

func init() { proto.RegisterFile("proxy.proto", fileDescriptor_700b50b08ed8dbaf) }

var fileDescriptor_700b50b08ed8dbaf = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x5d, 0xd8, 0x5a, 0xe0, 0xae, 0x1a, 0x92, 0x35, 0x4a, 0x09, 0x0c, 0x55, 0x41, 0x82, 0x6a,
	0x12, 0xed, 0x28, 0xfc, 0x82, 0x75, 0x52, 0x55, 0xa1, 0xa1, 0x29, 0x83, 0x17, 0x5e, 0x90, 0x93,
	0xdc, 0xb5, 0xae, 0x1c, 0xdb, 0x8b, 0x9d, 0x41, 0x9f, 0x90, 0xf8, 0x55, 0xfc, 0x3b, 0x50, 0x3e,
	0x9a, 0x36, 0x5d, 0xba, 0x0a, 0x55, 0x7b, 0xcb, 0xb1, 0xcf, 0xd5, 0x39, 0x27, 0xbe, 0x07, 0xf6,
	0x55, 0x24, 0x7f, 0xce, 0xba, 0x2a, 0x92, 0x46, 0x12, 0x12, 0x32, 0x7e, 0x13, 0xeb, 0x0c, 0x75,
	0xd3, 0x1b, 0xbb, 0xe1, 0xcb, 0x30, 0x94, 0x22, 0x3b, 0xb3, 0x0f, 0x98, 0x30, 0x18, 0x09, 0xca,
	0x73, 0xdc, 0x58, 0x9e, 0x70, 0xfe, 0x58, 0xf0, 0x6a, 0x24, 0x6e, 0x28, 0x67, 0x01, 0x35, 0x38,
	0x90, 0x9c, 0x9f, 0xa3, 0xa1, 0x03, 0xea, 0x4f, 0xd0, 0xc5, 0xeb, 0x18, 0xb5, 0x21, 0x27, 0xb0,
	0xe7, 0x51, 0x8d, 0x2d, 0xab, 0x6d, 0x75, 0xf6, 0xfb, 0x2f, 0xbb, 0x25, 0xc5, 0x5c, 0xea, 0x5c,
	0x8f, 0x4f, 0xa9, 0x46, 0x37, 0x65, 0x92, 0x67, 0xf0, 0x30, 0xf0, 0xbe, 0x0b, 0x1a, 0x62, 0xeb,
	0x41, 0xdb, 0xea, 0x3c, 0x76, 0xeb, 0x81, 0xf7, 0x99, 0x86, 0x48, 0xde, 0xc2, 0x13, 0x5f, 0x72,
	0x8e, 0xbe, 0x61, 0x52, 0x64, 0x84, 0xdd, 0x94, 0x70, 0xb0, 0x38, 0x4e, 0x89, 0x0e, 0x34, 0x16,
	0x27, 0xa3, 0xb3, 0xd6, 0x5e, 0xdb, 0xea, 0xec, 0xba, 0xa5, 0x33, 0x67, 0x0a, 0xf6, 0x92, 0xf3,
	0x08, 0x83, 0x2d, 0x5d, 0xdb, 0xf0, 0x28, 0xd6, 0xc9, 0x9f, 0x2a, 0x6c, 0x17, 0xd8, 0xf9, 0x6d,
	0x41, 0xf3, 0xab, 0xba, 0x7f, 0xa1, 0xe4, 0x4e, 0x51, 0xad, 0x7f, 0xc8, 0x28, 0xc8, 0x7f, 0x4d,
	0x81, 0x9d, 0x5f, 0x70, 0xe4, 0xe2, 0x55, 0x84, 0x7a, 0x72, 0x21, 0x39, 0xf3, 0x67, 0x23, 0x71,
	0x25, 0xb7, 0xb4, 0xd2, 0x84, 0xba, 0x54, 0x5f, 0x66, 0x2a, 0x33, 0x52, 0x73, 0x73, 0x44, 0x0e,
	0xa1, 0x26, 0xd5, 0x27, 0x9c, 0xe5, 0x1e, 0x32, 0xd0, 0xff, 0x5b, 0x83, 0xda, 0x45, 0xb2, 0x62,
	0x44, 0x01, 0x19, 0xa2, 0x19, 0xc8, 0x50, 0x49, 0x81, 0xc2, 0x5c, 0x1a, 0x6a, 0x50, 0x93, 0x93,
	0xb2, 0x62, 0xb1, 0x78, 0xb7, 0xa9, 0xb9, 0x63, 0xfb, 0xcd, 0x9a, 0x89, 0x15, 0xba, 0xb3, 0x43,
	0xae, 0xe1, 0x70, 0x88, 0x29, 0x64, 0xda, 0x30, 0x5f, 0x0f, 0x26, 0x54, 0x08, 0xe4, 0xa4, 0xbf,
	0x5e, 0xf3, 0x16, 0x79, 0xae, 0xfa, 0xba, 0x3c, 0x93, 0x83, 0x4b, 0x13, 0x31, 0x31, 0x76, 0x51,
	0x2b, 0x29, 0x34, 0x3a, 0x3b, 0x24, 0x82, 0xa3, 0x72, 0x35, 0xb2, 0xd5, 0x2b, 0x0a, 0xb2, 0xaa,
	0x9d, 0xf5, 0xf2, 0xee, 0x36, 0xd9, 0x2f, 0x2a, 0x5f, 0x25, 0xb1, 0x1a, 0x27, 0x31, 0x29, 0x34,
	0x86, 0x68, 0xce, 0x82, 0x79, 0xbc, 0xe3, 0xf5, 0xf1, 0x0a, 0xd2, 0x7f, 0xc6, 0x9a, 0xc2, 0xf3,
	0x72, 0x6f, 0x50, 0x18, 0x46, 0x79, 0x16, 0xa9, 0xbb, 0x21, 0xd2, 0xca, 0xf6, 0x6f, 0x8a, 0xe3,
	0xc1, 0xd3, 0x45, 0x6d, 0x96, 0x75, 0x8e, 0xab, 0x74, 0xaa, 0x1b, 0xb6, 0x49, 0x63, 0x0a, 0xcd,
	0xea, 0x5a, 0x90, 0xf7, 0x55, 0x22, 0x77, 0x56, 0x68, 0x83, 0xd6, 0xe9, 0xc7, 0x6f, 0xfd, 0x31,
	0x33, 0x93, 0xd8, 0x4b, 0x6e, 0x7a, 0x19, 0xf5, 0x1d, 0x93, 0xf9, 0x57, 0x6f, 0xfe, 0x3c, 0xbd,
	0x74, 0xba, 0x97, 0x0a, 0x2a, 0xcf, 0xab, 0xa7, 0xf0, 0xc3, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x33, 0xad, 0xd0, 0x9b, 0xba, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProxyClient is the client API for Proxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProxyClient interface {
	GetComponentStates(ctx context.Context, in *internalpb.GetComponentStatesRequest, opts ...grpc.CallOption) (*internalpb.ComponentStates, error)
	GetStatisticsChannel(ctx context.Context, in *internalpb.GetStatisticsChannelRequest, opts ...grpc.CallOption) (*milvuspb.StringResponse, error)
	InvalidateCollectionMetaCache(ctx context.Context, in *InvalidateCollMetaCacheRequest, opts ...grpc.CallOption) (*commonpb.Status, error)
	GetDdChannel(ctx context.Context, in *internalpb.GetDdChannelRequest, opts ...grpc.CallOption) (*milvuspb.StringResponse, error)
	InvalidateCredentialCache(ctx context.Context, in *InvalidateCredCacheRequest, opts ...grpc.CallOption) (*commonpb.Status, error)
	UpdateCredentialCache(ctx context.Context, in *UpdateCredCacheRequest, opts ...grpc.CallOption) (*commonpb.Status, error)
	RefreshPolicyInfoCache(ctx context.Context, in *RefreshPolicyInfoCacheRequest, opts ...grpc.CallOption) (*commonpb.Status, error)
}

type proxyClient struct {
	cc *grpc.ClientConn
}

func NewProxyClient(cc *grpc.ClientConn) ProxyClient {
	return &proxyClient{cc}
}

func (c *proxyClient) GetComponentStates(ctx context.Context, in *internalpb.GetComponentStatesRequest, opts ...grpc.CallOption) (*internalpb.ComponentStates, error) {
	out := new(internalpb.ComponentStates)
	err := c.cc.Invoke(ctx, "/milvus.proto.proxy.Proxy/GetComponentStates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) GetStatisticsChannel(ctx context.Context, in *internalpb.GetStatisticsChannelRequest, opts ...grpc.CallOption) (*milvuspb.StringResponse, error) {
	out := new(milvuspb.StringResponse)
	err := c.cc.Invoke(ctx, "/milvus.proto.proxy.Proxy/GetStatisticsChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) InvalidateCollectionMetaCache(ctx context.Context, in *InvalidateCollMetaCacheRequest, opts ...grpc.CallOption) (*commonpb.Status, error) {
	out := new(commonpb.Status)
	err := c.cc.Invoke(ctx, "/milvus.proto.proxy.Proxy/InvalidateCollectionMetaCache", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) GetDdChannel(ctx context.Context, in *internalpb.GetDdChannelRequest, opts ...grpc.CallOption) (*milvuspb.StringResponse, error) {
	out := new(milvuspb.StringResponse)
	err := c.cc.Invoke(ctx, "/milvus.proto.proxy.Proxy/GetDdChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) InvalidateCredentialCache(ctx context.Context, in *InvalidateCredCacheRequest, opts ...grpc.CallOption) (*commonpb.Status, error) {
	out := new(commonpb.Status)
	err := c.cc.Invoke(ctx, "/milvus.proto.proxy.Proxy/InvalidateCredentialCache", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) UpdateCredentialCache(ctx context.Context, in *UpdateCredCacheRequest, opts ...grpc.CallOption) (*commonpb.Status, error) {
	out := new(commonpb.Status)
	err := c.cc.Invoke(ctx, "/milvus.proto.proxy.Proxy/UpdateCredentialCache", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) RefreshPolicyInfoCache(ctx context.Context, in *RefreshPolicyInfoCacheRequest, opts ...grpc.CallOption) (*commonpb.Status, error) {
	out := new(commonpb.Status)
	err := c.cc.Invoke(ctx, "/milvus.proto.proxy.Proxy/RefreshPolicyInfoCache", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProxyServer is the server API for Proxy service.
type ProxyServer interface {
	GetComponentStates(context.Context, *internalpb.GetComponentStatesRequest) (*internalpb.ComponentStates, error)
	GetStatisticsChannel(context.Context, *internalpb.GetStatisticsChannelRequest) (*milvuspb.StringResponse, error)
	InvalidateCollectionMetaCache(context.Context, *InvalidateCollMetaCacheRequest) (*commonpb.Status, error)
	GetDdChannel(context.Context, *internalpb.GetDdChannelRequest) (*milvuspb.StringResponse, error)
	InvalidateCredentialCache(context.Context, *InvalidateCredCacheRequest) (*commonpb.Status, error)
	UpdateCredentialCache(context.Context, *UpdateCredCacheRequest) (*commonpb.Status, error)
	RefreshPolicyInfoCache(context.Context, *RefreshPolicyInfoCacheRequest) (*commonpb.Status, error)
}

// UnimplementedProxyServer can be embedded to have forward compatible implementations.
type UnimplementedProxyServer struct {
}

func (*UnimplementedProxyServer) GetComponentStates(ctx context.Context, req *internalpb.GetComponentStatesRequest) (*internalpb.ComponentStates, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComponentStates not implemented")
}
func (*UnimplementedProxyServer) GetStatisticsChannel(ctx context.Context, req *internalpb.GetStatisticsChannelRequest) (*milvuspb.StringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatisticsChannel not implemented")
}
func (*UnimplementedProxyServer) InvalidateCollectionMetaCache(ctx context.Context, req *InvalidateCollMetaCacheRequest) (*commonpb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvalidateCollectionMetaCache not implemented")
}
func (*UnimplementedProxyServer) GetDdChannel(ctx context.Context, req *internalpb.GetDdChannelRequest) (*milvuspb.StringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDdChannel not implemented")
}
func (*UnimplementedProxyServer) InvalidateCredentialCache(ctx context.Context, req *InvalidateCredCacheRequest) (*commonpb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvalidateCredentialCache not implemented")
}
func (*UnimplementedProxyServer) UpdateCredentialCache(ctx context.Context, req *UpdateCredCacheRequest) (*commonpb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCredentialCache not implemented")
}
func (*UnimplementedProxyServer) RefreshPolicyInfoCache(ctx context.Context, req *RefreshPolicyInfoCacheRequest) (*commonpb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshPolicyInfoCache not implemented")
}

func RegisterProxyServer(s *grpc.Server, srv ProxyServer) {
	s.RegisterService(&_Proxy_serviceDesc, srv)
}

func _Proxy_GetComponentStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(internalpb.GetComponentStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).GetComponentStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.proxy.Proxy/GetComponentStates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).GetComponentStates(ctx, req.(*internalpb.GetComponentStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_GetStatisticsChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(internalpb.GetStatisticsChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).GetStatisticsChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.proxy.Proxy/GetStatisticsChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).GetStatisticsChannel(ctx, req.(*internalpb.GetStatisticsChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_InvalidateCollectionMetaCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvalidateCollMetaCacheRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).InvalidateCollectionMetaCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.proxy.Proxy/InvalidateCollectionMetaCache",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).InvalidateCollectionMetaCache(ctx, req.(*InvalidateCollMetaCacheRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_GetDdChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(internalpb.GetDdChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).GetDdChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.proxy.Proxy/GetDdChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).GetDdChannel(ctx, req.(*internalpb.GetDdChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_InvalidateCredentialCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvalidateCredCacheRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).InvalidateCredentialCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.proxy.Proxy/InvalidateCredentialCache",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).InvalidateCredentialCache(ctx, req.(*InvalidateCredCacheRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_UpdateCredentialCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCredCacheRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).UpdateCredentialCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.proxy.Proxy/UpdateCredentialCache",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).UpdateCredentialCache(ctx, req.(*UpdateCredCacheRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_RefreshPolicyInfoCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshPolicyInfoCacheRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).RefreshPolicyInfoCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.proxy.Proxy/RefreshPolicyInfoCache",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).RefreshPolicyInfoCache(ctx, req.(*RefreshPolicyInfoCacheRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Proxy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "milvus.proto.proxy.Proxy",
	HandlerType: (*ProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetComponentStates",
			Handler:    _Proxy_GetComponentStates_Handler,
		},
		{
			MethodName: "GetStatisticsChannel",
			Handler:    _Proxy_GetStatisticsChannel_Handler,
		},
		{
			MethodName: "InvalidateCollectionMetaCache",
			Handler:    _Proxy_InvalidateCollectionMetaCache_Handler,
		},
		{
			MethodName: "GetDdChannel",
			Handler:    _Proxy_GetDdChannel_Handler,
		},
		{
			MethodName: "InvalidateCredentialCache",
			Handler:    _Proxy_InvalidateCredentialCache_Handler,
		},
		{
			MethodName: "UpdateCredentialCache",
			Handler:    _Proxy_UpdateCredentialCache_Handler,
		},
		{
			MethodName: "RefreshPolicyInfoCache",
			Handler:    _Proxy_RefreshPolicyInfoCache_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proxy.proto",
}
