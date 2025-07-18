// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protocol

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Weaviate_Search_FullMethodName       = "/weaviate.v1.Weaviate/Search"
	Weaviate_BatchObjects_FullMethodName = "/weaviate.v1.Weaviate/BatchObjects"
	Weaviate_BatchDelete_FullMethodName  = "/weaviate.v1.Weaviate/BatchDelete"
	Weaviate_TenantsGet_FullMethodName   = "/weaviate.v1.Weaviate/TenantsGet"
	Weaviate_Aggregate_FullMethodName    = "/weaviate.v1.Weaviate/Aggregate"
)

// WeaviateClient is the client API for Weaviate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WeaviateClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchReply, error)
	BatchObjects(ctx context.Context, in *BatchObjectsRequest, opts ...grpc.CallOption) (*BatchObjectsReply, error)
	BatchDelete(ctx context.Context, in *BatchDeleteRequest, opts ...grpc.CallOption) (*BatchDeleteReply, error)
	TenantsGet(ctx context.Context, in *TenantsGetRequest, opts ...grpc.CallOption) (*TenantsGetReply, error)
	Aggregate(ctx context.Context, in *AggregateRequest, opts ...grpc.CallOption) (*AggregateReply, error)
}

type weaviateClient struct {
	cc grpc.ClientConnInterface
}

func NewWeaviateClient(cc grpc.ClientConnInterface) WeaviateClient {
	return &weaviateClient{cc}
}

func (c *weaviateClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchReply)
	err := c.cc.Invoke(ctx, Weaviate_Search_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weaviateClient) BatchObjects(ctx context.Context, in *BatchObjectsRequest, opts ...grpc.CallOption) (*BatchObjectsReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BatchObjectsReply)
	err := c.cc.Invoke(ctx, Weaviate_BatchObjects_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weaviateClient) BatchDelete(ctx context.Context, in *BatchDeleteRequest, opts ...grpc.CallOption) (*BatchDeleteReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BatchDeleteReply)
	err := c.cc.Invoke(ctx, Weaviate_BatchDelete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weaviateClient) TenantsGet(ctx context.Context, in *TenantsGetRequest, opts ...grpc.CallOption) (*TenantsGetReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TenantsGetReply)
	err := c.cc.Invoke(ctx, Weaviate_TenantsGet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weaviateClient) Aggregate(ctx context.Context, in *AggregateRequest, opts ...grpc.CallOption) (*AggregateReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AggregateReply)
	err := c.cc.Invoke(ctx, Weaviate_Aggregate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WeaviateServer is the server API for Weaviate service.
// All implementations must embed UnimplementedWeaviateServer
// for forward compatibility.
type WeaviateServer interface {
	Search(context.Context, *SearchRequest) (*SearchReply, error)
	BatchObjects(context.Context, *BatchObjectsRequest) (*BatchObjectsReply, error)
	BatchDelete(context.Context, *BatchDeleteRequest) (*BatchDeleteReply, error)
	TenantsGet(context.Context, *TenantsGetRequest) (*TenantsGetReply, error)
	Aggregate(context.Context, *AggregateRequest) (*AggregateReply, error)
	mustEmbedUnimplementedWeaviateServer()
}

// UnimplementedWeaviateServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWeaviateServer struct{}

func (UnimplementedWeaviateServer) Search(context.Context, *SearchRequest) (*SearchReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedWeaviateServer) BatchObjects(context.Context, *BatchObjectsRequest) (*BatchObjectsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchObjects not implemented")
}
func (UnimplementedWeaviateServer) BatchDelete(context.Context, *BatchDeleteRequest) (*BatchDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDelete not implemented")
}
func (UnimplementedWeaviateServer) TenantsGet(context.Context, *TenantsGetRequest) (*TenantsGetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TenantsGet not implemented")
}
func (UnimplementedWeaviateServer) Aggregate(context.Context, *AggregateRequest) (*AggregateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Aggregate not implemented")
}
func (UnimplementedWeaviateServer) mustEmbedUnimplementedWeaviateServer() {}
func (UnimplementedWeaviateServer) testEmbeddedByValue()                  {}

// UnsafeWeaviateServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WeaviateServer will
// result in compilation errors.
type UnsafeWeaviateServer interface {
	mustEmbedUnimplementedWeaviateServer()
}

func RegisterWeaviateServer(s grpc.ServiceRegistrar, srv WeaviateServer) {
	// If the following call pancis, it indicates UnimplementedWeaviateServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Weaviate_ServiceDesc, srv)
}

func _Weaviate_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeaviateServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Weaviate_Search_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeaviateServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Weaviate_BatchObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchObjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeaviateServer).BatchObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Weaviate_BatchObjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeaviateServer).BatchObjects(ctx, req.(*BatchObjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Weaviate_BatchDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeaviateServer).BatchDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Weaviate_BatchDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeaviateServer).BatchDelete(ctx, req.(*BatchDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Weaviate_TenantsGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TenantsGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeaviateServer).TenantsGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Weaviate_TenantsGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeaviateServer).TenantsGet(ctx, req.(*TenantsGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Weaviate_Aggregate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AggregateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeaviateServer).Aggregate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Weaviate_Aggregate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeaviateServer).Aggregate(ctx, req.(*AggregateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Weaviate_ServiceDesc is the grpc.ServiceDesc for Weaviate service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Weaviate_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "weaviate.v1.Weaviate",
	HandlerType: (*WeaviateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _Weaviate_Search_Handler,
		},
		{
			MethodName: "BatchObjects",
			Handler:    _Weaviate_BatchObjects_Handler,
		},
		{
			MethodName: "BatchDelete",
			Handler:    _Weaviate_BatchDelete_Handler,
		},
		{
			MethodName: "TenantsGet",
			Handler:    _Weaviate_TenantsGet_Handler,
		},
		{
			MethodName: "Aggregate",
			Handler:    _Weaviate_Aggregate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/weaviate.proto",
}
