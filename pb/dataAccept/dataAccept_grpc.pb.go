// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: dataAccept.proto

package dataAccept

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DataAccept_GetDataAccept_FullMethodName = "/dataAccept.DataAccept/GetDataAccept"
)

// DataAcceptClient is the client API for DataAccept service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataAcceptClient interface {
	GetDataAccept(ctx context.Context, in *GetDateAcceptRequest, opts ...grpc.CallOption) (*GetDateAcceptResponse, error)
}

type dataAcceptClient struct {
	cc grpc.ClientConnInterface
}

func NewDataAcceptClient(cc grpc.ClientConnInterface) DataAcceptClient {
	return &dataAcceptClient{cc}
}

func (c *dataAcceptClient) GetDataAccept(ctx context.Context, in *GetDateAcceptRequest, opts ...grpc.CallOption) (*GetDateAcceptResponse, error) {
	out := new(GetDateAcceptResponse)
	err := c.cc.Invoke(ctx, DataAccept_GetDataAccept_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataAcceptServer is the server API for DataAccept service.
// All implementations must embed UnimplementedDataAcceptServer
// for forward compatibility
type DataAcceptServer interface {
	GetDataAccept(context.Context, *GetDateAcceptRequest) (*GetDateAcceptResponse, error)
	mustEmbedUnimplementedDataAcceptServer()
}

// UnimplementedDataAcceptServer must be embedded to have forward compatible implementations.
type UnimplementedDataAcceptServer struct {
}

func (UnimplementedDataAcceptServer) GetDataAccept(context.Context, *GetDateAcceptRequest) (*GetDateAcceptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDataAccept not implemented")
}
func (UnimplementedDataAcceptServer) mustEmbedUnimplementedDataAcceptServer() {}

// UnsafeDataAcceptServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataAcceptServer will
// result in compilation errors.
type UnsafeDataAcceptServer interface {
	mustEmbedUnimplementedDataAcceptServer()
}

func RegisterDataAcceptServer(s grpc.ServiceRegistrar, srv DataAcceptServer) {
	s.RegisterService(&DataAccept_ServiceDesc, srv)
}

func _DataAccept_GetDataAccept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDateAcceptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataAcceptServer).GetDataAccept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataAccept_GetDataAccept_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataAcceptServer).GetDataAccept(ctx, req.(*GetDateAcceptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataAccept_ServiceDesc is the grpc.ServiceDesc for DataAccept service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataAccept_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dataAccept.DataAccept",
	HandlerType: (*DataAcceptServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDataAccept",
			Handler:    _DataAccept_GetDataAccept_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dataAccept.proto",
}
