// Code generated by protoc-gen-go.
// source: helloworld_test.proto
// DO NOT EDIT!

/*
Package prpc is a generated protocol buffer package.

It is generated from these files:
	helloworld_test.proto

It has these top-level messages:
	HelloRequest
	HelloReply
	MultiplyRequest
	MultiplyResponse
*/
package prpc

import prpccommon "github.com/luci/luci-go/common/prpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The request message containing the user's name.
type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// The response message containing the greetings
type HelloReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type MultiplyRequest struct {
	X int32 `protobuf:"varint,1,opt,name=x" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y" json:"y,omitempty"`
}

func (m *MultiplyRequest) Reset()                    { *m = MultiplyRequest{} }
func (m *MultiplyRequest) String() string            { return proto.CompactTextString(m) }
func (*MultiplyRequest) ProtoMessage()               {}
func (*MultiplyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type MultiplyResponse struct {
	Z int32 `protobuf:"varint,1,opt,name=z" json:"z,omitempty"`
}

func (m *MultiplyResponse) Reset()                    { *m = MultiplyResponse{} }
func (m *MultiplyResponse) String() string            { return proto.CompactTextString(m) }
func (*MultiplyResponse) ProtoMessage()               {}
func (*MultiplyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*HelloRequest)(nil), "prpc.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "prpc.HelloReply")
	proto.RegisterType((*MultiplyRequest)(nil), "prpc.MultiplyRequest")
	proto.RegisterType((*MultiplyResponse)(nil), "prpc.MultiplyResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Greeter service

type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}
type greeterPRPCClient struct {
	client *prpccommon.Client
}

func NewGreeterPRPCClient(client *prpccommon.Client) GreeterClient {
	return &greeterPRPCClient{client}
}

func (c *greeterPRPCClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.client.Call(ctx, "prpc.Greeter", "SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := grpc.Invoke(ctx, "/prpc.Greeter/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterGreeterServer(s Registrar, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prpc.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "prpc.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

// Client API for Calc service

type CalcClient interface {
	Multiply(ctx context.Context, in *MultiplyRequest, opts ...grpc.CallOption) (*MultiplyResponse, error)
}
type calcPRPCClient struct {
	client *prpccommon.Client
}

func NewCalcPRPCClient(client *prpccommon.Client) CalcClient {
	return &calcPRPCClient{client}
}

func (c *calcPRPCClient) Multiply(ctx context.Context, in *MultiplyRequest, opts ...grpc.CallOption) (*MultiplyResponse, error) {
	out := new(MultiplyResponse)
	err := c.client.Call(ctx, "prpc.Calc", "Multiply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type calcClient struct {
	cc *grpc.ClientConn
}

func NewCalcClient(cc *grpc.ClientConn) CalcClient {
	return &calcClient{cc}
}

func (c *calcClient) Multiply(ctx context.Context, in *MultiplyRequest, opts ...grpc.CallOption) (*MultiplyResponse, error) {
	out := new(MultiplyResponse)
	err := grpc.Invoke(ctx, "/prpc.Calc/Multiply", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Calc service

type CalcServer interface {
	Multiply(context.Context, *MultiplyRequest) (*MultiplyResponse, error)
}

func RegisterCalcServer(s Registrar, srv CalcServer) {
	s.RegisterService(&_Calc_serviceDesc, srv)
}

func _Calc_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prpc.Calc/Multiply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServer).Multiply(ctx, req.(*MultiplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "prpc.Calc",
	HandlerType: (*CalcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Multiply",
			Handler:    _Calc_Multiply_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x37, 0x52, 0xdd, 0x75, 0x58, 0x70, 0x19, 0x58, 0x29, 0x7b, 0x5a, 0x72, 0x10, 0x2f,
	0xf6, 0x50, 0x8f, 0xe2, 0x69, 0x0f, 0x7a, 0xf1, 0x52, 0x1f, 0x40, 0x62, 0x1d, 0x54, 0x48, 0x9b,
	0x98, 0xa4, 0xd8, 0xf4, 0xe9, 0x6d, 0x52, 0x83, 0x6e, 0x6f, 0xf9, 0x67, 0x3e, 0xfe, 0x7c, 0x0c,
	0x6c, 0x3f, 0x48, 0x4a, 0xf5, 0xad, 0x8c, 0x7c, 0x7b, 0x71, 0x64, 0x5d, 0xa1, 0x8d, 0x72, 0x0a,
	0x33, 0x6d, 0x74, 0xcd, 0x39, 0xac, 0x1f, 0xc3, 0xba, 0xa2, 0xaf, 0x6e, 0xdc, 0x21, 0x42, 0xd6,
	0x8a, 0x86, 0x72, 0xb6, 0x67, 0xd7, 0xe7, 0x55, 0x7c, 0xf3, 0x2b, 0x80, 0x5f, 0x46, 0x4b, 0x8f,
	0x39, 0x2c, 0x1b, 0xb2, 0x56, 0xbc, 0x27, 0x28, 0x45, 0x7e, 0x03, 0x17, 0x4f, 0x9d, 0x74, 0x9f,
	0x23, 0x95, 0xea, 0xd6, 0xc0, 0xfa, 0x88, 0x9d, 0x56, 0xac, 0x0f, 0xc9, 0xe7, 0x27, 0x53, 0xf2,
	0x7c, 0x0f, 0x9b, 0x3f, 0xdc, 0x6a, 0xd5, 0x5a, 0x0a, 0xc4, 0x90, 0xf8, 0xa1, 0xbc, 0x87, 0xe5,
	0x83, 0x21, 0x72, 0x64, 0xb0, 0x84, 0xd5, 0xb3, 0xf0, 0x51, 0x03, 0xb1, 0x08, 0xea, 0xc5, 0x7f,
	0xef, 0xdd, 0xe6, 0x68, 0x36, 0x56, 0xf2, 0x45, 0x79, 0x80, 0xec, 0x20, 0x64, 0x8d, 0x77, 0xb0,
	0x4a, 0x1f, 0xe1, 0x76, 0xe2, 0x66, 0x9e, 0xbb, 0xcb, 0xf9, 0x78, 0xf2, 0xe1, 0x8b, 0xd7, 0xb3,
	0x78, 0xad, 0xdb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x49, 0xa8, 0x7b, 0x20, 0x46, 0x01, 0x00,
	0x00,
}
