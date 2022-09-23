// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: protos/main.proto

package protos

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

// StudentClient is the client API for Student service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StudentClient interface {
	RegisterStudent(ctx context.Context, in *StudentRequest, opts ...grpc.CallOption) (*Response, error)
	GetStudent(ctx context.Context, in *GetStudentRequest, opts ...grpc.CallOption) (*StudentResponse, error)
}

type studentClient struct {
	cc grpc.ClientConnInterface
}

func NewStudentClient(cc grpc.ClientConnInterface) StudentClient {
	return &studentClient{cc}
}

func (c *studentClient) RegisterStudent(ctx context.Context, in *StudentRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Student/RegisterStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentClient) GetStudent(ctx context.Context, in *GetStudentRequest, opts ...grpc.CallOption) (*StudentResponse, error) {
	out := new(StudentResponse)
	err := c.cc.Invoke(ctx, "/Student/GetStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentServer is the server API for Student service.
// All implementations should embed UnimplementedStudentServer
// for forward compatibility
type StudentServer interface {
	RegisterStudent(context.Context, *StudentRequest) (*Response, error)
	GetStudent(context.Context, *GetStudentRequest) (*StudentResponse, error)
}

// UnimplementedStudentServer should be embedded to have forward compatible implementations.
type UnimplementedStudentServer struct {
}

func (UnimplementedStudentServer) RegisterStudent(context.Context, *StudentRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterStudent not implemented")
}
func (UnimplementedStudentServer) GetStudent(context.Context, *GetStudentRequest) (*StudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudent not implemented")
}

// UnsafeStudentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StudentServer will
// result in compilation errors.
type UnsafeStudentServer interface {
	mustEmbedUnimplementedStudentServer()
}

func RegisterStudentServer(s grpc.ServiceRegistrar, srv StudentServer) {
	s.RegisterService(&Student_ServiceDesc, srv)
}

func _Student_RegisterStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServer).RegisterStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Student/RegisterStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServer).RegisterStudent(ctx, req.(*StudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Student_GetStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServer).GetStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Student/GetStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServer).GetStudent(ctx, req.(*GetStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Student_ServiceDesc is the grpc.ServiceDesc for Student service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Student_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Student",
	HandlerType: (*StudentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterStudent",
			Handler:    _Student_RegisterStudent_Handler,
		},
		{
			MethodName: "GetStudent",
			Handler:    _Student_GetStudent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/main.proto",
}
