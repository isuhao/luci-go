// Code generated by protoc-gen-go.
// source: printer_test.proto
// DO NOT EDIT!

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	printer_test.proto

It has these top-level messages:
	M1
	M2
	M3
	NestedMessageParent
*/
package main

import prpccommon "github.com/luci/luci-go/common/prpc"
import prpc "github.com/luci/luci-go/server/prpc"

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

// Enum comment.
// next line.
type E1 int32

const (
	// V0
	// next line.
	E1_V0 E1 = 0
	// V1
	// next line.
	E1_V1 E1 = 1
)

var E1_name = map[int32]string{
	0: "V0",
	1: "V1",
}
var E1_value = map[string]int32{
	"V0": 0,
	"V1": 1,
}

func (x E1) String() string {
	return proto.EnumName(E1_name, int32(x))
}
func (E1) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type NestedMessageParent_NestedEnum int32

const (
	NestedMessageParent_V0 NestedMessageParent_NestedEnum = 0
	NestedMessageParent_V1 NestedMessageParent_NestedEnum = 1
)

var NestedMessageParent_NestedEnum_name = map[int32]string{
	0: "V0",
	1: "V1",
}
var NestedMessageParent_NestedEnum_value = map[string]int32{
	"V0": 0,
	"V1": 1,
}

func (x NestedMessageParent_NestedEnum) String() string {
	return proto.EnumName(NestedMessageParent_NestedEnum_name, int32(x))
}
func (NestedMessageParent_NestedEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0}
}

// M1
// next line.
type M1 struct {
	// f1
	// next line.
	F1 string `protobuf:"bytes,1,opt,name=f1" json:"f1,omitempty"`
}

func (m *M1) Reset()                    { *m = M1{} }
func (m *M1) String() string            { return proto.CompactTextString(m) }
func (*M1) ProtoMessage()               {}
func (*M1) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// M2
// next line.
type M2 struct {
	// f1
	// next line.
	F1 []*M1 `protobuf:"bytes,1,rep,name=f1" json:"f1,omitempty"`
	// f2
	// next line.
	F2 E1 `protobuf:"varint,2,opt,name=f2,enum=main.E1" json:"f2,omitempty"`
}

func (m *M2) Reset()                    { *m = M2{} }
func (m *M2) String() string            { return proto.CompactTextString(m) }
func (*M2) ProtoMessage()               {}
func (*M2) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *M2) GetF1() []*M1 {
	if m != nil {
		return m.F1
	}
	return nil
}

// M3
type M3 struct {
	// O1
	//
	// Types that are valid to be assigned to O1:
	//	*M3_F1
	//	*M3_F2
	O1 isM3_O1 `protobuf_oneof:"O1"`
	// O2
	//
	// Types that are valid to be assigned to O2:
	//	*M3_F3
	//	*M3_F4
	O2 isM3_O2 `protobuf_oneof:"O2"`
	// f5
	F5 string `protobuf:"bytes,5,opt,name=f5" json:"f5,omitempty"`
	// f6
	F6 int32 `protobuf:"varint,6,opt,name=f6" json:"f6,omitempty"`
}

func (m *M3) Reset()                    { *m = M3{} }
func (m *M3) String() string            { return proto.CompactTextString(m) }
func (*M3) ProtoMessage()               {}
func (*M3) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isM3_O1 interface {
	isM3_O1()
}
type isM3_O2 interface {
	isM3_O2()
}

type M3_F1 struct {
	F1 int32 `protobuf:"varint,1,opt,name=f1,oneof"`
}
type M3_F2 struct {
	F2 int32 `protobuf:"varint,2,opt,name=f2,oneof"`
}
type M3_F3 struct {
	F3 int32 `protobuf:"varint,3,opt,name=f3,oneof"`
}
type M3_F4 struct {
	F4 int32 `protobuf:"varint,4,opt,name=f4,oneof"`
}

func (*M3_F1) isM3_O1() {}
func (*M3_F2) isM3_O1() {}
func (*M3_F3) isM3_O2() {}
func (*M3_F4) isM3_O2() {}

func (m *M3) GetO1() isM3_O1 {
	if m != nil {
		return m.O1
	}
	return nil
}
func (m *M3) GetO2() isM3_O2 {
	if m != nil {
		return m.O2
	}
	return nil
}

func (m *M3) GetF1() int32 {
	if x, ok := m.GetO1().(*M3_F1); ok {
		return x.F1
	}
	return 0
}

func (m *M3) GetF2() int32 {
	if x, ok := m.GetO1().(*M3_F2); ok {
		return x.F2
	}
	return 0
}

func (m *M3) GetF3() int32 {
	if x, ok := m.GetO2().(*M3_F3); ok {
		return x.F3
	}
	return 0
}

func (m *M3) GetF4() int32 {
	if x, ok := m.GetO2().(*M3_F4); ok {
		return x.F4
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*M3) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _M3_OneofMarshaler, _M3_OneofUnmarshaler, _M3_OneofSizer, []interface{}{
		(*M3_F1)(nil),
		(*M3_F2)(nil),
		(*M3_F3)(nil),
		(*M3_F4)(nil),
	}
}

func _M3_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*M3)
	// O1
	switch x := m.O1.(type) {
	case *M3_F1:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.F1))
	case *M3_F2:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.F2))
	case nil:
	default:
		return fmt.Errorf("M3.O1 has unexpected type %T", x)
	}
	// O2
	switch x := m.O2.(type) {
	case *M3_F3:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.F3))
	case *M3_F4:
		b.EncodeVarint(4<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.F4))
	case nil:
	default:
		return fmt.Errorf("M3.O2 has unexpected type %T", x)
	}
	return nil
}

func _M3_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*M3)
	switch tag {
	case 1: // O1.f1
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.O1 = &M3_F1{int32(x)}
		return true, err
	case 2: // O1.f2
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.O1 = &M3_F2{int32(x)}
		return true, err
	case 3: // O2.f3
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.O2 = &M3_F3{int32(x)}
		return true, err
	case 4: // O2.f4
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.O2 = &M3_F4{int32(x)}
		return true, err
	default:
		return false, nil
	}
}

func _M3_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*M3)
	// O1
	switch x := m.O1.(type) {
	case *M3_F1:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.F1))
	case *M3_F2:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.F2))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	// O2
	switch x := m.O2.(type) {
	case *M3_F3:
		n += proto.SizeVarint(3<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.F3))
	case *M3_F4:
		n += proto.SizeVarint(4<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.F4))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type NestedMessageParent struct {
}

func (m *NestedMessageParent) Reset()                    { *m = NestedMessageParent{} }
func (m *NestedMessageParent) String() string            { return proto.CompactTextString(m) }
func (*NestedMessageParent) ProtoMessage()               {}
func (*NestedMessageParent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type NestedMessageParent_NestedMessage struct {
	F1 int32 `protobuf:"varint,1,opt,name=f1" json:"f1,omitempty"`
	F2 int32 `protobuf:"varint,2,opt,name=f2" json:"f2,omitempty"`
}

func (m *NestedMessageParent_NestedMessage) Reset()         { *m = NestedMessageParent_NestedMessage{} }
func (m *NestedMessageParent_NestedMessage) String() string { return proto.CompactTextString(m) }
func (*NestedMessageParent_NestedMessage) ProtoMessage()    {}
func (*NestedMessageParent_NestedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0}
}

func init() {
	proto.RegisterType((*M1)(nil), "main.M1")
	proto.RegisterType((*M2)(nil), "main.M2")
	proto.RegisterType((*M3)(nil), "main.M3")
	proto.RegisterType((*NestedMessageParent)(nil), "main.NestedMessageParent")
	proto.RegisterType((*NestedMessageParent_NestedMessage)(nil), "main.NestedMessageParent.NestedMessage")
	proto.RegisterEnum("main.E1", E1_name, E1_value)
	proto.RegisterEnum("main.NestedMessageParent_NestedEnum", NestedMessageParent_NestedEnum_name, NestedMessageParent_NestedEnum_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for S1 service

type S1Client interface {
	// R1
	R1(ctx context.Context, in *M1, opts ...grpc.CallOption) (*M2, error)
}
type s1PRPCClient struct {
	client *prpccommon.Client
}

func NewS1PRPCClient(client *prpccommon.Client) S1Client {
	return &s1PRPCClient{client}
}

func (c *s1PRPCClient) R1(ctx context.Context, in *M1, opts ...grpc.CallOption) (*M2, error) {
	out := new(M2)
	err := c.client.Call(ctx, "main.S1", "R1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type s1Client struct {
	cc *grpc.ClientConn
}

func NewS1Client(cc *grpc.ClientConn) S1Client {
	return &s1Client{cc}
}

func (c *s1Client) R1(ctx context.Context, in *M1, opts ...grpc.CallOption) (*M2, error) {
	out := new(M2)
	err := grpc.Invoke(ctx, "/main.S1/R1", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for S1 service

type S1Server interface {
	// R1
	R1(context.Context, *M1) (*M2, error)
}

func RegisterS1Server(s prpc.Registrar, srv S1Server) {
	s.RegisterService(&_S1_serviceDesc, srv)
}

func _S1_R1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(M1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S1Server).R1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.S1/R1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S1Server).R1(ctx, req.(*M1))
	}
	return interceptor(ctx, in, info, handler)
}

var _S1_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.S1",
	HandlerType: (*S1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "R1",
			Handler:    _S1_R1_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

// Client API for S2 service

type S2Client interface {
	// R1
	R1(ctx context.Context, in *M1, opts ...grpc.CallOption) (*M2, error)
	// R2
	R2(ctx context.Context, in *M1, opts ...grpc.CallOption) (*M2, error)
}
type s2PRPCClient struct {
	client *prpccommon.Client
}

func NewS2PRPCClient(client *prpccommon.Client) S2Client {
	return &s2PRPCClient{client}
}

func (c *s2PRPCClient) R1(ctx context.Context, in *M1, opts ...grpc.CallOption) (*M2, error) {
	out := new(M2)
	err := c.client.Call(ctx, "main.S2", "R1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s2PRPCClient) R2(ctx context.Context, in *M1, opts ...grpc.CallOption) (*M2, error) {
	out := new(M2)
	err := c.client.Call(ctx, "main.S2", "R2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type s2Client struct {
	cc *grpc.ClientConn
}

func NewS2Client(cc *grpc.ClientConn) S2Client {
	return &s2Client{cc}
}

func (c *s2Client) R1(ctx context.Context, in *M1, opts ...grpc.CallOption) (*M2, error) {
	out := new(M2)
	err := grpc.Invoke(ctx, "/main.S2/R1", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s2Client) R2(ctx context.Context, in *M1, opts ...grpc.CallOption) (*M2, error) {
	out := new(M2)
	err := grpc.Invoke(ctx, "/main.S2/R2", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for S2 service

type S2Server interface {
	// R1
	R1(context.Context, *M1) (*M2, error)
	// R2
	R2(context.Context, *M1) (*M2, error)
}

func RegisterS2Server(s prpc.Registrar, srv S2Server) {
	s.RegisterService(&_S2_serviceDesc, srv)
}

func _S2_R1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(M1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S2Server).R1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.S2/R1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S2Server).R1(ctx, req.(*M1))
	}
	return interceptor(ctx, in, info, handler)
}

func _S2_R2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(M1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S2Server).R2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.S2/R2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S2Server).R2(ctx, req.(*M1))
	}
	return interceptor(ctx, in, info, handler)
}

var _S2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.S2",
	HandlerType: (*S2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "R1",
			Handler:    _S2_R1_Handler,
		},
		{
			MethodName: "R2",
			Handler:    _S2_R2_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x93, 0x6d, 0x12, 0x74, 0xc4, 0x10, 0x56, 0x0f, 0x4b, 0xf0, 0x50, 0xf6, 0x54, 0x3c,
	0x44, 0x77, 0xd2, 0x16, 0x0f, 0x9e, 0x84, 0x80, 0x97, 0xa8, 0x54, 0xf0, 0x2a, 0x91, 0xae, 0xd2,
	0x43, 0xd3, 0x92, 0xac, 0xe0, 0xc7, 0x77, 0x32, 0x69, 0xe3, 0x1f, 0xc4, 0x4b, 0xb2, 0xfb, 0x7b,
	0xef, 0x85, 0x37, 0x13, 0x90, 0xdb, 0x66, 0x55, 0x3b, 0xdb, 0x3c, 0x3b, 0xdb, 0xba, 0x6c, 0xdb,
	0x6c, 0xdc, 0x46, 0x06, 0xeb, 0x6a, 0x55, 0xeb, 0x53, 0x10, 0xa5, 0x91, 0x31, 0x88, 0x57, 0xa3,
	0xfc, 0xb1, 0x3f, 0x39, 0x5c, 0xd0, 0x49, 0x5f, 0x11, 0x45, 0xa9, 0x76, 0x74, 0x34, 0x39, 0xc2,
	0x83, 0xac, 0xb3, 0x67, 0xa5, 0xe9, 0x74, 0x56, 0x50, 0x09, 0xf2, 0xc7, 0x7b, 0xa5, 0xe8, 0x14,
	0xd4, 0x1f, 0x94, 0xcc, 0x65, 0x32, 0x7c, 0x2f, 0xbc, 0xf5, 0x38, 0x91, 0x0c, 0x89, 0x9e, 0x20,
	0x93, 0x5c, 0x8d, 0x98, 0xf8, 0x44, 0xfa, 0xd4, 0x54, 0x05, 0x03, 0x99, 0x72, 0xaf, 0x99, 0x0a,
	0x77, 0xbd, 0x66, 0x7c, 0x9f, 0xab, 0xa8, 0x73, 0xd0, 0x7d, 0x7e, 0x13, 0x80, 0xb8, 0x37, 0xfc,
	0x44, 0xbd, 0x84, 0x93, 0x3b, 0x9a, 0xce, 0x2e, 0x4b, 0xdb, 0xb6, 0xd5, 0x9b, 0x7d, 0xa8, 0x1a,
	0x5b, 0xbb, 0xf4, 0x02, 0x8e, 0x7f, 0xe0, 0x6f, 0xb3, 0x86, 0xdc, 0x2c, 0xfe, 0x6a, 0xc6, 0x13,
	0x9c, 0x01, 0xf4, 0x81, 0xa2, 0x7e, 0x5f, 0xcb, 0x08, 0xc4, 0xd3, 0x65, 0xe2, 0xf1, 0xdb, 0x24,
	0xfe, 0x39, 0xed, 0xab, 0x30, 0xbf, 0x29, 0x8e, 0x41, 0x3c, 0x1a, 0x99, 0x82, 0x58, 0x18, 0x39,
	0x6c, 0x2a, 0xdd, 0x9f, 0x50, 0x7b, 0x78, 0x4d, 0x0e, 0xfc, 0xcf, 0xc1, 0x1a, 0xfe, 0xad, 0xbd,
	0x44, 0xfc, 0xcb, 0xf2, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb5, 0xd3, 0xd3, 0x89, 0xc8, 0x01,
	0x00, 0x00,
}
