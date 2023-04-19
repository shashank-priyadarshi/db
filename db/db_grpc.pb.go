// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: db.proto

package db

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

// DatabaseServiceClient is the client API for DatabaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DatabaseServiceClient interface {
	GetData(ctx context.Context, in *DatabaseRequest, opts ...grpc.CallOption) (*DatabaseResponse, error)
	InsertData(ctx context.Context, in *DatabaseRequest, opts ...grpc.CallOption) (*DatabaseResponse, error)
}

type databaseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDatabaseServiceClient(cc grpc.ClientConnInterface) DatabaseServiceClient {
	return &databaseServiceClient{cc}
}

func (c *databaseServiceClient) GetData(ctx context.Context, in *DatabaseRequest, opts ...grpc.CallOption) (*DatabaseResponse, error) {
	out := new(DatabaseResponse)
	err := c.cc.Invoke(ctx, "/db_service.DatabaseService/GetData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) InsertData(ctx context.Context, in *DatabaseRequest, opts ...grpc.CallOption) (*DatabaseResponse, error) {
	out := new(DatabaseResponse)
	err := c.cc.Invoke(ctx, "/db_service.DatabaseService/InsertData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatabaseServiceServer is the server API for DatabaseService service.
// All implementations must embed UnimplementedDatabaseServiceServer
// for forward compatibility
type DatabaseServiceServer interface {
	GetData(context.Context, *DatabaseRequest) (*DatabaseResponse, error)
	InsertData(context.Context, *DatabaseRequest) (*DatabaseResponse, error)
	mustEmbedUnimplementedDatabaseServiceServer()
}

// UnimplementedDatabaseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDatabaseServiceServer struct {
}

func (UnimplementedDatabaseServiceServer) GetData(context.Context, *DatabaseRequest) (*DatabaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetData not implemented")
}
func (UnimplementedDatabaseServiceServer) InsertData(context.Context, *DatabaseRequest) (*DatabaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertData not implemented")
}
func (UnimplementedDatabaseServiceServer) mustEmbedUnimplementedDatabaseServiceServer() {}

// UnsafeDatabaseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatabaseServiceServer will
// result in compilation errors.
type UnsafeDatabaseServiceServer interface {
	mustEmbedUnimplementedDatabaseServiceServer()
}

func RegisterDatabaseServiceServer(s grpc.ServiceRegistrar, srv DatabaseServiceServer) {
	s.RegisterService(&DatabaseService_ServiceDesc, srv)
}

func _DatabaseService_GetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).GetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db_service.DatabaseService/GetData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).GetData(ctx, req.(*DatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_InsertData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).InsertData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db_service.DatabaseService/InsertData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).InsertData(ctx, req.(*DatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DatabaseService_ServiceDesc is the grpc.ServiceDesc for DatabaseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DatabaseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "db_service.DatabaseService",
	HandlerType: (*DatabaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetData",
			Handler:    _DatabaseService_GetData_Handler,
		},
		{
			MethodName: "InsertData",
			Handler:    _DatabaseService_InsertData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "db.proto",
}
