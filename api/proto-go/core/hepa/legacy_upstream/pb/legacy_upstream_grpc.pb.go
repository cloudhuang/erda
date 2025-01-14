// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: legacy_upstream.proto

package pb

import (
	context "context"
	transport "github.com/erda-project/erda-infra/pkg/transport"
	grpc1 "github.com/erda-project/erda-infra/pkg/transport/grpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion5

// UpstreamServiceClient is the client API for UpstreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UpstreamServiceClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	AsyncRegister(ctx context.Context, in *AsyncRegisterRequest, opts ...grpc.CallOption) (*AsyncRegisterResponse, error)
}

type upstreamServiceClient struct {
	cc grpc1.ClientConnInterface
}

func NewUpstreamServiceClient(cc grpc1.ClientConnInterface) UpstreamServiceClient {
	return &upstreamServiceClient{cc}
}

func (c *upstreamServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/erda.core.hepa.legacy_upstream.UpstreamService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *upstreamServiceClient) AsyncRegister(ctx context.Context, in *AsyncRegisterRequest, opts ...grpc.CallOption) (*AsyncRegisterResponse, error) {
	out := new(AsyncRegisterResponse)
	err := c.cc.Invoke(ctx, "/erda.core.hepa.legacy_upstream.UpstreamService/AsyncRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UpstreamServiceServer is the server API for UpstreamService service.
// All implementations should embed UnimplementedUpstreamServiceServer
// for forward compatibility
type UpstreamServiceServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	AsyncRegister(context.Context, *AsyncRegisterRequest) (*AsyncRegisterResponse, error)
}

// UnimplementedUpstreamServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUpstreamServiceServer struct {
}

func (*UnimplementedUpstreamServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedUpstreamServiceServer) AsyncRegister(context.Context, *AsyncRegisterRequest) (*AsyncRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AsyncRegister not implemented")
}

func RegisterUpstreamServiceServer(s grpc1.ServiceRegistrar, srv UpstreamServiceServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_UpstreamService_serviceDesc(srv, opts...), srv)
}

var _UpstreamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.core.hepa.legacy_upstream.UpstreamService",
	HandlerType: (*UpstreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "legacy_upstream.proto",
}

func _get_UpstreamService_serviceDesc(srv UpstreamServiceServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_UpstreamService_Register_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.Register(ctx, req.(*RegisterRequest))
	}
	var _UpstreamService_Register_info transport.ServiceInfo
	if h.Interceptor != nil {
		_UpstreamService_Register_info = transport.NewServiceInfo("erda.core.hepa.legacy_upstream.UpstreamService", "Register", srv)
		_UpstreamService_Register_Handler = h.Interceptor(_UpstreamService_Register_Handler)
	}

	_UpstreamService_AsyncRegister_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.AsyncRegister(ctx, req.(*AsyncRegisterRequest))
	}
	var _UpstreamService_AsyncRegister_info transport.ServiceInfo
	if h.Interceptor != nil {
		_UpstreamService_AsyncRegister_info = transport.NewServiceInfo("erda.core.hepa.legacy_upstream.UpstreamService", "AsyncRegister", srv)
		_UpstreamService_AsyncRegister_Handler = h.Interceptor(_UpstreamService_AsyncRegister_Handler)
	}

	var serviceDesc = _UpstreamService_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(RegisterRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(UpstreamServiceServer).Register(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _UpstreamService_Register_info)
				}
				if interceptor == nil {
					return _UpstreamService_Register_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.hepa.legacy_upstream.UpstreamService/Register",
				}
				return interceptor(ctx, in, info, _UpstreamService_Register_Handler)
			},
		},
		{
			MethodName: "AsyncRegister",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(AsyncRegisterRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(UpstreamServiceServer).AsyncRegister(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _UpstreamService_AsyncRegister_info)
				}
				if interceptor == nil {
					return _UpstreamService_AsyncRegister_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.hepa.legacy_upstream.UpstreamService/AsyncRegister",
				}
				return interceptor(ctx, in, info, _UpstreamService_AsyncRegister_Handler)
			},
		},
	}
	return &serviceDesc
}
