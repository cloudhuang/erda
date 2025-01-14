// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: openapi_consumer.proto

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

// OpenapiConsumerServiceClient is the client API for OpenapiConsumerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OpenapiConsumerServiceClient interface {
	// +publish path: "/api/gateway/openapi/consumers"
	GetConsumers(ctx context.Context, in *GetConsumersRequest, opts ...grpc.CallOption) (*GetConsumersResponse, error)
	// +publish path: "/api/gateway/openapi/consumers"
	CreateConsumer(ctx context.Context, in *CreateConsumerRequest, opts ...grpc.CallOption) (*CreateConsumerResponse, error)
	// +publish path: "/api/gateway/openapi/consumers/{consumerId}"
	UpdateConsumer(ctx context.Context, in *UpdateConsumerRequest, opts ...grpc.CallOption) (*UpdateConsumerResponse, error)
	// +publish path: "/api/gateway/openapi/consumers/{consumerId}"
	DeleteConsumer(ctx context.Context, in *DeleteConsumerRequest, opts ...grpc.CallOption) (*DeleteConsumerResponse, error)
	// +publish path: "/api/gateway/openapi/consumers-name"
	GetConsumersName(ctx context.Context, in *GetConsumersNameRequest, opts ...grpc.CallOption) (*GetConsumersNameResponse, error)
	// +publish path: "/api/gateway/openapi/consumers/{consumerId}/packages"
	GetConsumerAcl(ctx context.Context, in *GetConsumerAclRequest, opts ...grpc.CallOption) (*GetConsumerAclResponse, error)
	// +publish path: "/api/gateway/openapi/consumers/{consumerId}/packages"
	UpdateConsumerAcl(ctx context.Context, in *UpdateConsumerAclRequest, opts ...grpc.CallOption) (*UpdateConsumerAclResponse, error)
	// +publish path: "/api/gateway/openapi/consumers/{consumerId}/credentials"
	GetConsumerAuth(ctx context.Context, in *GetConsumerAuthRequest, opts ...grpc.CallOption) (*GetConsumerAuthResponse, error)
	// +publish path: "/api/gateway/openapi/consumers/{consumerId}/credentials"
	UpdateConsumerAuth(ctx context.Context, in *UpdateConsumerAuthRequest, opts ...grpc.CallOption) (*UpdateConsumerAuthResponse, error)
	// +publish path: "/api/gateway/openapi/packages/{packageId}/consumers"
	GetEndpointAcl(ctx context.Context, in *GetEndpointAclRequest, opts ...grpc.CallOption) (*GetEndpointAclResponse, error)
	// +publish path: "/api/gateway/openapi/packages/{packageId}/consumers"
	UpdateEndpointAcl(ctx context.Context, in *UpdateEndpointAclRequest, opts ...grpc.CallOption) (*UpdateEndpointAclResponse, error)
	// +publish path: "/api/gateway/openapi/packages/{packageId}/apis/{apiId}/authz"
	GetEndpointApiAcl(ctx context.Context, in *GetEndpointApiAclRequest, opts ...grpc.CallOption) (*GetEndpointApiAclResponse, error)
	// +publish path: "/api/gateway/openapi/packages/{packageId}/apis/{apiId}/authz"
	UpdateEndpointApiAcl(ctx context.Context, in *UpdateEndpointApiAclRequest, opts ...grpc.CallOption) (*UpdateEndpointApiAclResponse, error)
}

type openapiConsumerServiceClient struct {
	cc grpc1.ClientConnInterface
}

func NewOpenapiConsumerServiceClient(cc grpc1.ClientConnInterface) OpenapiConsumerServiceClient {
	return &openapiConsumerServiceClient{cc}
}

func (c *openapiConsumerServiceClient) GetConsumers(ctx context.Context, in *GetConsumersRequest, opts ...grpc.CallOption) (*GetConsumersResponse, error) {
	out := new(GetConsumersResponse)
	err := c.cc.Invoke(ctx, "/erda.core.hepa.openapi_consumer.OpenapiConsumerService/GetConsumers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openapiConsumerServiceClient) CreateConsumer(ctx context.Context, in *CreateConsumerRequest, opts ...grpc.CallOption) (*CreateConsumerResponse, error) {
	out := new(CreateConsumerResponse)
	err := c.cc.Invoke(ctx, "/erda.core.hepa.openapi_consumer.OpenapiConsumerService/CreateConsumer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openapiConsumerServiceClient) UpdateConsumer(ctx context.Context, in *UpdateConsumerRequest, opts ...grpc.CallOption) (*UpdateConsumerResponse, error) {
	out := new(UpdateConsumerResponse)
	err := c.cc.Invoke(ctx, "/erda.core.hepa.openapi_consumer.OpenapiConsumerService/UpdateConsumer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openapiConsumerServiceClient) DeleteConsumer(ctx context.Context, in *DeleteConsumerRequest, opts ...grpc.CallOption) (*DeleteConsumerResponse, error) {
	out := new(DeleteConsumerResponse)
	err := c.cc.Invoke(ctx, "/erda.core.hepa.openapi_consumer.OpenapiConsumerService/DeleteConsumer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openapiConsumerServiceClient) GetConsumersName(ctx context.Context, in *GetConsumersNameRequest, opts ...grpc.CallOption) (*GetConsumersNameResponse, error) {
	out := new(GetConsumersNameResponse)
	err := c.cc.Invoke(ctx, "/erda.core.hepa.openapi_consumer.OpenapiConsumerService/GetConsumersName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openapiConsumerServiceClient) GetConsumerAcl(ctx context.Context, in *GetConsumerAclRequest, opts ...grpc.CallOption) (*GetConsumerAclResponse, error) {
	out := new(GetConsumerAclResponse)
	err := c.cc.Invoke(ctx, "/erda.core.hepa.openapi_consumer.OpenapiConsumerService/GetConsumerAcl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openapiConsumerServiceClient) UpdateConsumerAcl(ctx context.Context, in *UpdateConsumerAclRequest, opts ...grpc.CallOption) (*UpdateConsumerAclResponse, error) {
	out := new(UpdateConsumerAclResponse)
	err := c.cc.Invoke(ctx, "/erda.core.hepa.openapi_consumer.OpenapiConsumerService/UpdateConsumerAcl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openapiConsumerServiceClient) GetConsumerAuth(ctx context.Context, in *GetConsumerAuthRequest, opts ...grpc.CallOption) (*GetConsumerAuthResponse, error) {
	out := new(GetConsumerAuthResponse)
	err := c.cc.Invoke(ctx, "/erda.core.hepa.openapi_consumer.OpenapiConsumerService/GetConsumerAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openapiConsumerServiceClient) UpdateConsumerAuth(ctx context.Context, in *UpdateConsumerAuthRequest, opts ...grpc.CallOption) (*UpdateConsumerAuthResponse, error) {
	out := new(UpdateConsumerAuthResponse)
	err := c.cc.Invoke(ctx, "/erda.core.hepa.openapi_consumer.OpenapiConsumerService/UpdateConsumerAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openapiConsumerServiceClient) GetEndpointAcl(ctx context.Context, in *GetEndpointAclRequest, opts ...grpc.CallOption) (*GetEndpointAclResponse, error) {
	out := new(GetEndpointAclResponse)
	err := c.cc.Invoke(ctx, "/erda.core.hepa.openapi_consumer.OpenapiConsumerService/GetEndpointAcl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openapiConsumerServiceClient) UpdateEndpointAcl(ctx context.Context, in *UpdateEndpointAclRequest, opts ...grpc.CallOption) (*UpdateEndpointAclResponse, error) {
	out := new(UpdateEndpointAclResponse)
	err := c.cc.Invoke(ctx, "/erda.core.hepa.openapi_consumer.OpenapiConsumerService/UpdateEndpointAcl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openapiConsumerServiceClient) GetEndpointApiAcl(ctx context.Context, in *GetEndpointApiAclRequest, opts ...grpc.CallOption) (*GetEndpointApiAclResponse, error) {
	out := new(GetEndpointApiAclResponse)
	err := c.cc.Invoke(ctx, "/erda.core.hepa.openapi_consumer.OpenapiConsumerService/GetEndpointApiAcl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openapiConsumerServiceClient) UpdateEndpointApiAcl(ctx context.Context, in *UpdateEndpointApiAclRequest, opts ...grpc.CallOption) (*UpdateEndpointApiAclResponse, error) {
	out := new(UpdateEndpointApiAclResponse)
	err := c.cc.Invoke(ctx, "/erda.core.hepa.openapi_consumer.OpenapiConsumerService/UpdateEndpointApiAcl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OpenapiConsumerServiceServer is the server API for OpenapiConsumerService service.
// All implementations should embed UnimplementedOpenapiConsumerServiceServer
// for forward compatibility
type OpenapiConsumerServiceServer interface {
	// +publish path: "/api/gateway/openapi/consumers"
	GetConsumers(context.Context, *GetConsumersRequest) (*GetConsumersResponse, error)
	// +publish path: "/api/gateway/openapi/consumers"
	CreateConsumer(context.Context, *CreateConsumerRequest) (*CreateConsumerResponse, error)
	// +publish path: "/api/gateway/openapi/consumers/{consumerId}"
	UpdateConsumer(context.Context, *UpdateConsumerRequest) (*UpdateConsumerResponse, error)
	// +publish path: "/api/gateway/openapi/consumers/{consumerId}"
	DeleteConsumer(context.Context, *DeleteConsumerRequest) (*DeleteConsumerResponse, error)
	// +publish path: "/api/gateway/openapi/consumers-name"
	GetConsumersName(context.Context, *GetConsumersNameRequest) (*GetConsumersNameResponse, error)
	// +publish path: "/api/gateway/openapi/consumers/{consumerId}/packages"
	GetConsumerAcl(context.Context, *GetConsumerAclRequest) (*GetConsumerAclResponse, error)
	// +publish path: "/api/gateway/openapi/consumers/{consumerId}/packages"
	UpdateConsumerAcl(context.Context, *UpdateConsumerAclRequest) (*UpdateConsumerAclResponse, error)
	// +publish path: "/api/gateway/openapi/consumers/{consumerId}/credentials"
	GetConsumerAuth(context.Context, *GetConsumerAuthRequest) (*GetConsumerAuthResponse, error)
	// +publish path: "/api/gateway/openapi/consumers/{consumerId}/credentials"
	UpdateConsumerAuth(context.Context, *UpdateConsumerAuthRequest) (*UpdateConsumerAuthResponse, error)
	// +publish path: "/api/gateway/openapi/packages/{packageId}/consumers"
	GetEndpointAcl(context.Context, *GetEndpointAclRequest) (*GetEndpointAclResponse, error)
	// +publish path: "/api/gateway/openapi/packages/{packageId}/consumers"
	UpdateEndpointAcl(context.Context, *UpdateEndpointAclRequest) (*UpdateEndpointAclResponse, error)
	// +publish path: "/api/gateway/openapi/packages/{packageId}/apis/{apiId}/authz"
	GetEndpointApiAcl(context.Context, *GetEndpointApiAclRequest) (*GetEndpointApiAclResponse, error)
	// +publish path: "/api/gateway/openapi/packages/{packageId}/apis/{apiId}/authz"
	UpdateEndpointApiAcl(context.Context, *UpdateEndpointApiAclRequest) (*UpdateEndpointApiAclResponse, error)
}

// UnimplementedOpenapiConsumerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedOpenapiConsumerServiceServer struct {
}

func (*UnimplementedOpenapiConsumerServiceServer) GetConsumers(context.Context, *GetConsumersRequest) (*GetConsumersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConsumers not implemented")
}
func (*UnimplementedOpenapiConsumerServiceServer) CreateConsumer(context.Context, *CreateConsumerRequest) (*CreateConsumerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConsumer not implemented")
}
func (*UnimplementedOpenapiConsumerServiceServer) UpdateConsumer(context.Context, *UpdateConsumerRequest) (*UpdateConsumerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConsumer not implemented")
}
func (*UnimplementedOpenapiConsumerServiceServer) DeleteConsumer(context.Context, *DeleteConsumerRequest) (*DeleteConsumerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConsumer not implemented")
}
func (*UnimplementedOpenapiConsumerServiceServer) GetConsumersName(context.Context, *GetConsumersNameRequest) (*GetConsumersNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConsumersName not implemented")
}
func (*UnimplementedOpenapiConsumerServiceServer) GetConsumerAcl(context.Context, *GetConsumerAclRequest) (*GetConsumerAclResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConsumerAcl not implemented")
}
func (*UnimplementedOpenapiConsumerServiceServer) UpdateConsumerAcl(context.Context, *UpdateConsumerAclRequest) (*UpdateConsumerAclResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConsumerAcl not implemented")
}
func (*UnimplementedOpenapiConsumerServiceServer) GetConsumerAuth(context.Context, *GetConsumerAuthRequest) (*GetConsumerAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConsumerAuth not implemented")
}
func (*UnimplementedOpenapiConsumerServiceServer) UpdateConsumerAuth(context.Context, *UpdateConsumerAuthRequest) (*UpdateConsumerAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConsumerAuth not implemented")
}
func (*UnimplementedOpenapiConsumerServiceServer) GetEndpointAcl(context.Context, *GetEndpointAclRequest) (*GetEndpointAclResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEndpointAcl not implemented")
}
func (*UnimplementedOpenapiConsumerServiceServer) UpdateEndpointAcl(context.Context, *UpdateEndpointAclRequest) (*UpdateEndpointAclResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEndpointAcl not implemented")
}
func (*UnimplementedOpenapiConsumerServiceServer) GetEndpointApiAcl(context.Context, *GetEndpointApiAclRequest) (*GetEndpointApiAclResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEndpointApiAcl not implemented")
}
func (*UnimplementedOpenapiConsumerServiceServer) UpdateEndpointApiAcl(context.Context, *UpdateEndpointApiAclRequest) (*UpdateEndpointApiAclResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEndpointApiAcl not implemented")
}

func RegisterOpenapiConsumerServiceServer(s grpc1.ServiceRegistrar, srv OpenapiConsumerServiceServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_OpenapiConsumerService_serviceDesc(srv, opts...), srv)
}

var _OpenapiConsumerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.core.hepa.openapi_consumer.OpenapiConsumerService",
	HandlerType: (*OpenapiConsumerServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "openapi_consumer.proto",
}

func _get_OpenapiConsumerService_serviceDesc(srv OpenapiConsumerServiceServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_OpenapiConsumerService_GetConsumers_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetConsumers(ctx, req.(*GetConsumersRequest))
	}
	var _OpenapiConsumerService_GetConsumers_info transport.ServiceInfo
	if h.Interceptor != nil {
		_OpenapiConsumerService_GetConsumers_info = transport.NewServiceInfo("erda.core.hepa.openapi_consumer.OpenapiConsumerService", "GetConsumers", srv)
		_OpenapiConsumerService_GetConsumers_Handler = h.Interceptor(_OpenapiConsumerService_GetConsumers_Handler)
	}

	_OpenapiConsumerService_CreateConsumer_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.CreateConsumer(ctx, req.(*CreateConsumerRequest))
	}
	var _OpenapiConsumerService_CreateConsumer_info transport.ServiceInfo
	if h.Interceptor != nil {
		_OpenapiConsumerService_CreateConsumer_info = transport.NewServiceInfo("erda.core.hepa.openapi_consumer.OpenapiConsumerService", "CreateConsumer", srv)
		_OpenapiConsumerService_CreateConsumer_Handler = h.Interceptor(_OpenapiConsumerService_CreateConsumer_Handler)
	}

	_OpenapiConsumerService_UpdateConsumer_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.UpdateConsumer(ctx, req.(*UpdateConsumerRequest))
	}
	var _OpenapiConsumerService_UpdateConsumer_info transport.ServiceInfo
	if h.Interceptor != nil {
		_OpenapiConsumerService_UpdateConsumer_info = transport.NewServiceInfo("erda.core.hepa.openapi_consumer.OpenapiConsumerService", "UpdateConsumer", srv)
		_OpenapiConsumerService_UpdateConsumer_Handler = h.Interceptor(_OpenapiConsumerService_UpdateConsumer_Handler)
	}

	_OpenapiConsumerService_DeleteConsumer_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.DeleteConsumer(ctx, req.(*DeleteConsumerRequest))
	}
	var _OpenapiConsumerService_DeleteConsumer_info transport.ServiceInfo
	if h.Interceptor != nil {
		_OpenapiConsumerService_DeleteConsumer_info = transport.NewServiceInfo("erda.core.hepa.openapi_consumer.OpenapiConsumerService", "DeleteConsumer", srv)
		_OpenapiConsumerService_DeleteConsumer_Handler = h.Interceptor(_OpenapiConsumerService_DeleteConsumer_Handler)
	}

	_OpenapiConsumerService_GetConsumersName_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetConsumersName(ctx, req.(*GetConsumersNameRequest))
	}
	var _OpenapiConsumerService_GetConsumersName_info transport.ServiceInfo
	if h.Interceptor != nil {
		_OpenapiConsumerService_GetConsumersName_info = transport.NewServiceInfo("erda.core.hepa.openapi_consumer.OpenapiConsumerService", "GetConsumersName", srv)
		_OpenapiConsumerService_GetConsumersName_Handler = h.Interceptor(_OpenapiConsumerService_GetConsumersName_Handler)
	}

	_OpenapiConsumerService_GetConsumerAcl_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetConsumerAcl(ctx, req.(*GetConsumerAclRequest))
	}
	var _OpenapiConsumerService_GetConsumerAcl_info transport.ServiceInfo
	if h.Interceptor != nil {
		_OpenapiConsumerService_GetConsumerAcl_info = transport.NewServiceInfo("erda.core.hepa.openapi_consumer.OpenapiConsumerService", "GetConsumerAcl", srv)
		_OpenapiConsumerService_GetConsumerAcl_Handler = h.Interceptor(_OpenapiConsumerService_GetConsumerAcl_Handler)
	}

	_OpenapiConsumerService_UpdateConsumerAcl_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.UpdateConsumerAcl(ctx, req.(*UpdateConsumerAclRequest))
	}
	var _OpenapiConsumerService_UpdateConsumerAcl_info transport.ServiceInfo
	if h.Interceptor != nil {
		_OpenapiConsumerService_UpdateConsumerAcl_info = transport.NewServiceInfo("erda.core.hepa.openapi_consumer.OpenapiConsumerService", "UpdateConsumerAcl", srv)
		_OpenapiConsumerService_UpdateConsumerAcl_Handler = h.Interceptor(_OpenapiConsumerService_UpdateConsumerAcl_Handler)
	}

	_OpenapiConsumerService_GetConsumerAuth_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetConsumerAuth(ctx, req.(*GetConsumerAuthRequest))
	}
	var _OpenapiConsumerService_GetConsumerAuth_info transport.ServiceInfo
	if h.Interceptor != nil {
		_OpenapiConsumerService_GetConsumerAuth_info = transport.NewServiceInfo("erda.core.hepa.openapi_consumer.OpenapiConsumerService", "GetConsumerAuth", srv)
		_OpenapiConsumerService_GetConsumerAuth_Handler = h.Interceptor(_OpenapiConsumerService_GetConsumerAuth_Handler)
	}

	_OpenapiConsumerService_UpdateConsumerAuth_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.UpdateConsumerAuth(ctx, req.(*UpdateConsumerAuthRequest))
	}
	var _OpenapiConsumerService_UpdateConsumerAuth_info transport.ServiceInfo
	if h.Interceptor != nil {
		_OpenapiConsumerService_UpdateConsumerAuth_info = transport.NewServiceInfo("erda.core.hepa.openapi_consumer.OpenapiConsumerService", "UpdateConsumerAuth", srv)
		_OpenapiConsumerService_UpdateConsumerAuth_Handler = h.Interceptor(_OpenapiConsumerService_UpdateConsumerAuth_Handler)
	}

	_OpenapiConsumerService_GetEndpointAcl_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetEndpointAcl(ctx, req.(*GetEndpointAclRequest))
	}
	var _OpenapiConsumerService_GetEndpointAcl_info transport.ServiceInfo
	if h.Interceptor != nil {
		_OpenapiConsumerService_GetEndpointAcl_info = transport.NewServiceInfo("erda.core.hepa.openapi_consumer.OpenapiConsumerService", "GetEndpointAcl", srv)
		_OpenapiConsumerService_GetEndpointAcl_Handler = h.Interceptor(_OpenapiConsumerService_GetEndpointAcl_Handler)
	}

	_OpenapiConsumerService_UpdateEndpointAcl_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.UpdateEndpointAcl(ctx, req.(*UpdateEndpointAclRequest))
	}
	var _OpenapiConsumerService_UpdateEndpointAcl_info transport.ServiceInfo
	if h.Interceptor != nil {
		_OpenapiConsumerService_UpdateEndpointAcl_info = transport.NewServiceInfo("erda.core.hepa.openapi_consumer.OpenapiConsumerService", "UpdateEndpointAcl", srv)
		_OpenapiConsumerService_UpdateEndpointAcl_Handler = h.Interceptor(_OpenapiConsumerService_UpdateEndpointAcl_Handler)
	}

	_OpenapiConsumerService_GetEndpointApiAcl_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetEndpointApiAcl(ctx, req.(*GetEndpointApiAclRequest))
	}
	var _OpenapiConsumerService_GetEndpointApiAcl_info transport.ServiceInfo
	if h.Interceptor != nil {
		_OpenapiConsumerService_GetEndpointApiAcl_info = transport.NewServiceInfo("erda.core.hepa.openapi_consumer.OpenapiConsumerService", "GetEndpointApiAcl", srv)
		_OpenapiConsumerService_GetEndpointApiAcl_Handler = h.Interceptor(_OpenapiConsumerService_GetEndpointApiAcl_Handler)
	}

	_OpenapiConsumerService_UpdateEndpointApiAcl_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.UpdateEndpointApiAcl(ctx, req.(*UpdateEndpointApiAclRequest))
	}
	var _OpenapiConsumerService_UpdateEndpointApiAcl_info transport.ServiceInfo
	if h.Interceptor != nil {
		_OpenapiConsumerService_UpdateEndpointApiAcl_info = transport.NewServiceInfo("erda.core.hepa.openapi_consumer.OpenapiConsumerService", "UpdateEndpointApiAcl", srv)
		_OpenapiConsumerService_UpdateEndpointApiAcl_Handler = h.Interceptor(_OpenapiConsumerService_UpdateEndpointApiAcl_Handler)
	}

	var serviceDesc = _OpenapiConsumerService_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "GetConsumers",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetConsumersRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(OpenapiConsumerServiceServer).GetConsumers(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _OpenapiConsumerService_GetConsumers_info)
				}
				if interceptor == nil {
					return _OpenapiConsumerService_GetConsumers_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.hepa.openapi_consumer.OpenapiConsumerService/GetConsumers",
				}
				return interceptor(ctx, in, info, _OpenapiConsumerService_GetConsumers_Handler)
			},
		},
		{
			MethodName: "CreateConsumer",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(CreateConsumerRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(OpenapiConsumerServiceServer).CreateConsumer(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _OpenapiConsumerService_CreateConsumer_info)
				}
				if interceptor == nil {
					return _OpenapiConsumerService_CreateConsumer_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.hepa.openapi_consumer.OpenapiConsumerService/CreateConsumer",
				}
				return interceptor(ctx, in, info, _OpenapiConsumerService_CreateConsumer_Handler)
			},
		},
		{
			MethodName: "UpdateConsumer",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(UpdateConsumerRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(OpenapiConsumerServiceServer).UpdateConsumer(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _OpenapiConsumerService_UpdateConsumer_info)
				}
				if interceptor == nil {
					return _OpenapiConsumerService_UpdateConsumer_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.hepa.openapi_consumer.OpenapiConsumerService/UpdateConsumer",
				}
				return interceptor(ctx, in, info, _OpenapiConsumerService_UpdateConsumer_Handler)
			},
		},
		{
			MethodName: "DeleteConsumer",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(DeleteConsumerRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(OpenapiConsumerServiceServer).DeleteConsumer(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _OpenapiConsumerService_DeleteConsumer_info)
				}
				if interceptor == nil {
					return _OpenapiConsumerService_DeleteConsumer_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.hepa.openapi_consumer.OpenapiConsumerService/DeleteConsumer",
				}
				return interceptor(ctx, in, info, _OpenapiConsumerService_DeleteConsumer_Handler)
			},
		},
		{
			MethodName: "GetConsumersName",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetConsumersNameRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(OpenapiConsumerServiceServer).GetConsumersName(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _OpenapiConsumerService_GetConsumersName_info)
				}
				if interceptor == nil {
					return _OpenapiConsumerService_GetConsumersName_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.hepa.openapi_consumer.OpenapiConsumerService/GetConsumersName",
				}
				return interceptor(ctx, in, info, _OpenapiConsumerService_GetConsumersName_Handler)
			},
		},
		{
			MethodName: "GetConsumerAcl",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetConsumerAclRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(OpenapiConsumerServiceServer).GetConsumerAcl(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _OpenapiConsumerService_GetConsumerAcl_info)
				}
				if interceptor == nil {
					return _OpenapiConsumerService_GetConsumerAcl_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.hepa.openapi_consumer.OpenapiConsumerService/GetConsumerAcl",
				}
				return interceptor(ctx, in, info, _OpenapiConsumerService_GetConsumerAcl_Handler)
			},
		},
		{
			MethodName: "UpdateConsumerAcl",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(UpdateConsumerAclRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(OpenapiConsumerServiceServer).UpdateConsumerAcl(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _OpenapiConsumerService_UpdateConsumerAcl_info)
				}
				if interceptor == nil {
					return _OpenapiConsumerService_UpdateConsumerAcl_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.hepa.openapi_consumer.OpenapiConsumerService/UpdateConsumerAcl",
				}
				return interceptor(ctx, in, info, _OpenapiConsumerService_UpdateConsumerAcl_Handler)
			},
		},
		{
			MethodName: "GetConsumerAuth",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetConsumerAuthRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(OpenapiConsumerServiceServer).GetConsumerAuth(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _OpenapiConsumerService_GetConsumerAuth_info)
				}
				if interceptor == nil {
					return _OpenapiConsumerService_GetConsumerAuth_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.hepa.openapi_consumer.OpenapiConsumerService/GetConsumerAuth",
				}
				return interceptor(ctx, in, info, _OpenapiConsumerService_GetConsumerAuth_Handler)
			},
		},
		{
			MethodName: "UpdateConsumerAuth",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(UpdateConsumerAuthRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(OpenapiConsumerServiceServer).UpdateConsumerAuth(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _OpenapiConsumerService_UpdateConsumerAuth_info)
				}
				if interceptor == nil {
					return _OpenapiConsumerService_UpdateConsumerAuth_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.hepa.openapi_consumer.OpenapiConsumerService/UpdateConsumerAuth",
				}
				return interceptor(ctx, in, info, _OpenapiConsumerService_UpdateConsumerAuth_Handler)
			},
		},
		{
			MethodName: "GetEndpointAcl",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetEndpointAclRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(OpenapiConsumerServiceServer).GetEndpointAcl(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _OpenapiConsumerService_GetEndpointAcl_info)
				}
				if interceptor == nil {
					return _OpenapiConsumerService_GetEndpointAcl_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.hepa.openapi_consumer.OpenapiConsumerService/GetEndpointAcl",
				}
				return interceptor(ctx, in, info, _OpenapiConsumerService_GetEndpointAcl_Handler)
			},
		},
		{
			MethodName: "UpdateEndpointAcl",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(UpdateEndpointAclRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(OpenapiConsumerServiceServer).UpdateEndpointAcl(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _OpenapiConsumerService_UpdateEndpointAcl_info)
				}
				if interceptor == nil {
					return _OpenapiConsumerService_UpdateEndpointAcl_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.hepa.openapi_consumer.OpenapiConsumerService/UpdateEndpointAcl",
				}
				return interceptor(ctx, in, info, _OpenapiConsumerService_UpdateEndpointAcl_Handler)
			},
		},
		{
			MethodName: "GetEndpointApiAcl",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetEndpointApiAclRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(OpenapiConsumerServiceServer).GetEndpointApiAcl(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _OpenapiConsumerService_GetEndpointApiAcl_info)
				}
				if interceptor == nil {
					return _OpenapiConsumerService_GetEndpointApiAcl_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.hepa.openapi_consumer.OpenapiConsumerService/GetEndpointApiAcl",
				}
				return interceptor(ctx, in, info, _OpenapiConsumerService_GetEndpointApiAcl_Handler)
			},
		},
		{
			MethodName: "UpdateEndpointApiAcl",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(UpdateEndpointApiAclRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(OpenapiConsumerServiceServer).UpdateEndpointApiAcl(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _OpenapiConsumerService_UpdateEndpointApiAcl_info)
				}
				if interceptor == nil {
					return _OpenapiConsumerService_UpdateEndpointApiAcl_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.hepa.openapi_consumer.OpenapiConsumerService/UpdateEndpointApiAcl",
				}
				return interceptor(ctx, in, info, _OpenapiConsumerService_UpdateEndpointApiAcl_Handler)
			},
		},
	}
	return &serviceDesc
}
