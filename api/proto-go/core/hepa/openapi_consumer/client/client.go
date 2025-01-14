// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: openapi_consumer.proto

package client

import (
	context "context"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/hepa/openapi_consumer/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// OpenapiConsumerService openapi_consumer.proto
	OpenapiConsumerService() pb.OpenapiConsumerServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		openapiConsumerService: pb.NewOpenapiConsumerServiceClient(cc),
	}
}

type serviceClients struct {
	openapiConsumerService pb.OpenapiConsumerServiceClient
}

func (c *serviceClients) OpenapiConsumerService() pb.OpenapiConsumerServiceClient {
	return c.openapiConsumerService
}

type openapiConsumerServiceWrapper struct {
	client pb.OpenapiConsumerServiceClient
	opts   []grpc1.CallOption
}

func (s *openapiConsumerServiceWrapper) GetConsumers(ctx context.Context, req *pb.GetConsumersRequest) (*pb.GetConsumersResponse, error) {
	return s.client.GetConsumers(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *openapiConsumerServiceWrapper) CreateConsumer(ctx context.Context, req *pb.CreateConsumerRequest) (*pb.CreateConsumerResponse, error) {
	return s.client.CreateConsumer(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *openapiConsumerServiceWrapper) UpdateConsumer(ctx context.Context, req *pb.UpdateConsumerRequest) (*pb.UpdateConsumerResponse, error) {
	return s.client.UpdateConsumer(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *openapiConsumerServiceWrapper) DeleteConsumer(ctx context.Context, req *pb.DeleteConsumerRequest) (*pb.DeleteConsumerResponse, error) {
	return s.client.DeleteConsumer(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *openapiConsumerServiceWrapper) GetConsumersName(ctx context.Context, req *pb.GetConsumersNameRequest) (*pb.GetConsumersNameResponse, error) {
	return s.client.GetConsumersName(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *openapiConsumerServiceWrapper) GetConsumerAcl(ctx context.Context, req *pb.GetConsumerAclRequest) (*pb.GetConsumerAclResponse, error) {
	return s.client.GetConsumerAcl(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *openapiConsumerServiceWrapper) UpdateConsumerAcl(ctx context.Context, req *pb.UpdateConsumerAclRequest) (*pb.UpdateConsumerAclResponse, error) {
	return s.client.UpdateConsumerAcl(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *openapiConsumerServiceWrapper) GetConsumerAuth(ctx context.Context, req *pb.GetConsumerAuthRequest) (*pb.GetConsumerAuthResponse, error) {
	return s.client.GetConsumerAuth(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *openapiConsumerServiceWrapper) UpdateConsumerAuth(ctx context.Context, req *pb.UpdateConsumerAuthRequest) (*pb.UpdateConsumerAuthResponse, error) {
	return s.client.UpdateConsumerAuth(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *openapiConsumerServiceWrapper) GetEndpointAcl(ctx context.Context, req *pb.GetEndpointAclRequest) (*pb.GetEndpointAclResponse, error) {
	return s.client.GetEndpointAcl(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *openapiConsumerServiceWrapper) UpdateEndpointAcl(ctx context.Context, req *pb.UpdateEndpointAclRequest) (*pb.UpdateEndpointAclResponse, error) {
	return s.client.UpdateEndpointAcl(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *openapiConsumerServiceWrapper) GetEndpointApiAcl(ctx context.Context, req *pb.GetEndpointApiAclRequest) (*pb.GetEndpointApiAclResponse, error) {
	return s.client.GetEndpointApiAcl(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *openapiConsumerServiceWrapper) UpdateEndpointApiAcl(ctx context.Context, req *pb.UpdateEndpointApiAclRequest) (*pb.UpdateEndpointApiAclResponse, error) {
	return s.client.UpdateEndpointApiAcl(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}
