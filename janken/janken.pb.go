// Code generated by protoc-gen-go. DO NOT EDIT.
// source: janken.proto

/*
Package jankenserve is a generated protocol buffer package.

It is generated from these files:
	janken.proto

It has these top-level messages:
	Janken
	Result
*/
package jankenserve

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

type Janken struct {
	Janken int64 `protobuf:"varint,1,opt,name=janken" json:"janken,omitempty"`
}

func (m *Janken) Reset()                    { *m = Janken{} }
func (m *Janken) String() string            { return proto.CompactTextString(m) }
func (*Janken) ProtoMessage()               {}
func (*Janken) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Janken) GetJanken() int64 {
	if m != nil {
		return m.Janken
	}
	return 0
}

type Result struct {
	Result string `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
	Koma   string `protobuf:"bytes,2,opt,name=koma" json:"koma,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Result) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *Result) GetKoma() string {
	if m != nil {
		return m.Koma
	}
	return ""
}

func init() {
	proto.RegisterType((*Janken)(nil), "jankenserve.Janken")
	proto.RegisterType((*Result)(nil), "jankenserve.Result")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Tasks service

type TasksClient interface {
	Jankening(ctx context.Context, in *Janken, opts ...grpc.CallOption) (*Result, error)
}

type tasksClient struct {
	cc *grpc.ClientConn
}

func NewTasksClient(cc *grpc.ClientConn) TasksClient {
	return &tasksClient{cc}
}

func (c *tasksClient) Jankening(ctx context.Context, in *Janken, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/jankenserve.Tasks/Jankening", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Tasks service

type TasksServer interface {
	Jankening(context.Context, *Janken) (*Result, error)
}

func RegisterTasksServer(s *grpc.Server, srv TasksServer) {
	s.RegisterService(&_Tasks_serviceDesc, srv)
}

func _Tasks_Jankening_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Janken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).Jankening(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jankenserve.Tasks/Jankening",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).Jankening(ctx, req.(*Janken))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tasks_serviceDesc = grpc.ServiceDesc{
	ServiceName: "jankenserve.Tasks",
	HandlerType: (*TasksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Jankening",
			Handler:    _Tasks_Jankening_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "janken.proto",
}

func init() { proto.RegisterFile("janken.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x4a, 0xcc, 0xcb,
	0x4e, 0xcd, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x86, 0xf0, 0x8a, 0x53, 0x8b, 0xca,
	0x52, 0x95, 0x14, 0xb8, 0xd8, 0xbc, 0xc0, 0x5c, 0x21, 0x31, 0x2e, 0x36, 0x88, 0x84, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0x73, 0x10, 0x94, 0xa7, 0x64, 0xc2, 0xc5, 0x16, 0x94, 0x5a, 0x5c, 0x9a, 0x53,
	0x02, 0x52, 0x51, 0x04, 0x66, 0x81, 0x55, 0x70, 0x06, 0x41, 0x79, 0x42, 0x42, 0x5c, 0x2c, 0xd9,
	0xf9, 0xb9, 0x89, 0x12, 0x4c, 0x60, 0x51, 0x30, 0xdb, 0xc8, 0x81, 0x8b, 0x35, 0x24, 0xb1, 0x38,
	0xbb, 0x58, 0xc8, 0x9c, 0x8b, 0x13, 0x62, 0x41, 0x66, 0x5e, 0xba, 0x90, 0xb0, 0x1e, 0x92, 0xdd,
	0x7a, 0x10, 0x71, 0x29, 0x54, 0x41, 0x88, 0x5d, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0xd7, 0x1a, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xdf, 0x32, 0xad, 0xb8, 0xbd, 0x00, 0x00, 0x00,
}
