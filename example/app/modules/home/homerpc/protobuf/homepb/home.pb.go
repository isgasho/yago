// Code generated by protoc-gen-go. DO NOT EDIT.
// source: homepb/home.proto

package homepb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b4a53407e735cc4, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloReply struct {
	Errno                int32    `protobuf:"varint,1,opt,name=errno,proto3" json:"errno,omitempty"`
	Errmsg               string   `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	Data                 string   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b4a53407e735cc4, []int{1}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetErrno() int32 {
	if m != nil {
		return m.Errno
	}
	return 0
}

func (m *HelloReply) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *HelloReply) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "homepb.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "homepb.HelloReply")
}

func init() { proto.RegisterFile("homepb/home.proto", fileDescriptor_1b4a53407e735cc4) }

var fileDescriptor_1b4a53407e735cc4 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcc, 0xc8, 0xcf, 0x4d,
	0x2d, 0x48, 0xd2, 0x07, 0x51, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x6c, 0x10, 0x21, 0x25,
	0x25, 0x2e, 0x1e, 0x8f, 0xd4, 0x9c, 0x9c, 0xfc, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21,
	0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b,
	0xc9, 0x8f, 0x8b, 0x0b, 0xaa, 0xa6, 0x20, 0xa7, 0x52, 0x48, 0x84, 0x8b, 0x35, 0xb5, 0xa8, 0x28,
	0x2f, 0x1f, 0xac, 0x84, 0x35, 0x08, 0xc2, 0x11, 0x12, 0xe3, 0x62, 0x4b, 0x2d, 0x2a, 0xca, 0x2d,
	0x4e, 0x97, 0x60, 0x02, 0xeb, 0x84, 0xf2, 0x40, 0xe6, 0xa5, 0x24, 0x96, 0x24, 0x4a, 0x30, 0x43,
	0xcc, 0x03, 0xb1, 0x8d, 0xac, 0xb9, 0x58, 0x3c, 0xf2, 0x73, 0x53, 0x85, 0x8c, 0xb9, 0x58, 0xc1,
	0xe6, 0x0a, 0x89, 0xe8, 0x41, 0x5c, 0xa3, 0x87, 0xec, 0x14, 0x29, 0x21, 0x34, 0xd1, 0x82, 0x9c,
	0x4a, 0x25, 0x86, 0x24, 0x36, 0xb0, 0xfb, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x58,
	0xcd, 0x98, 0xd4, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HomeClient is the client API for Home service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HomeClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type homeClient struct {
	cc *grpc.ClientConn
}

func NewHomeClient(cc *grpc.ClientConn) HomeClient {
	return &homeClient{cc}
}

func (c *homeClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/homepb.Home/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HomeServer is the server API for Home service.
type HomeServer interface {
	Hello(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterHomeServer(s *grpc.Server, srv HomeServer) {
	s.RegisterService(&_Home_serviceDesc, srv)
}

func _Home_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/homepb.Home/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Home_serviceDesc = grpc.ServiceDesc{
	ServiceName: "homepb.Home",
	HandlerType: (*HomeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Home_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "homepb/home.proto",
}