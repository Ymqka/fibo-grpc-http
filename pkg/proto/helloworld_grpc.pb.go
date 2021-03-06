// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// FibonacciCalculatorClient is the client API for FibonacciCalculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FibonacciCalculatorClient interface {
	GetFiboSequence(ctx context.Context, in *FiboRangeRequest, opts ...grpc.CallOption) (*FiboRangeResponse, error)
}

type fibonacciCalculatorClient struct {
	cc grpc.ClientConnInterface
}

func NewFibonacciCalculatorClient(cc grpc.ClientConnInterface) FibonacciCalculatorClient {
	return &fibonacciCalculatorClient{cc}
}

func (c *fibonacciCalculatorClient) GetFiboSequence(ctx context.Context, in *FiboRangeRequest, opts ...grpc.CallOption) (*FiboRangeResponse, error) {
	out := new(FiboRangeResponse)
	err := c.cc.Invoke(ctx, "/fibogrpchttp.FibonacciCalculator/GetFiboSequence", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FibonacciCalculatorServer is the server API for FibonacciCalculator service.
// All implementations must embed UnimplementedFibonacciCalculatorServer
// for forward compatibility
type FibonacciCalculatorServer interface {
	GetFiboSequence(context.Context, *FiboRangeRequest) (*FiboRangeResponse, error)
	mustEmbedUnimplementedFibonacciCalculatorServer()
}

// UnimplementedFibonacciCalculatorServer must be embedded to have forward compatible implementations.
type UnimplementedFibonacciCalculatorServer struct {
}

func (UnimplementedFibonacciCalculatorServer) GetFiboSequence(context.Context, *FiboRangeRequest) (*FiboRangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFiboSequence not implemented")
}
func (UnimplementedFibonacciCalculatorServer) mustEmbedUnimplementedFibonacciCalculatorServer() {}

// UnsafeFibonacciCalculatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FibonacciCalculatorServer will
// result in compilation errors.
type UnsafeFibonacciCalculatorServer interface {
	mustEmbedUnimplementedFibonacciCalculatorServer()
}

func RegisterFibonacciCalculatorServer(s grpc.ServiceRegistrar, srv FibonacciCalculatorServer) {
	s.RegisterService(&_FibonacciCalculator_serviceDesc, srv)
}

func _FibonacciCalculator_GetFiboSequence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FiboRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FibonacciCalculatorServer).GetFiboSequence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fibogrpchttp.FibonacciCalculator/GetFiboSequence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FibonacciCalculatorServer).GetFiboSequence(ctx, req.(*FiboRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FibonacciCalculator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fibogrpchttp.FibonacciCalculator",
	HandlerType: (*FibonacciCalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFiboSequence",
			Handler:    _FibonacciCalculator_GetFiboSequence_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld.proto",
}
