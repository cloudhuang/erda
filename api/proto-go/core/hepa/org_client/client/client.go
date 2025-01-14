// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: org_client.proto

package client

import (
	context "context"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/hepa/org_client/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// OrgClientService org_client.proto
	OrgClientService() pb.OrgClientServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		orgClientService: pb.NewOrgClientServiceClient(cc),
	}
}

type serviceClients struct {
	orgClientService pb.OrgClientServiceClient
}

func (c *serviceClients) OrgClientService() pb.OrgClientServiceClient {
	return c.orgClientService
}

type orgClientServiceWrapper struct {
	client pb.OrgClientServiceClient
	opts   []grpc1.CallOption
}

func (s *orgClientServiceWrapper) CreateClient(ctx context.Context, req *pb.CreateClientRequest) (*pb.CreateClientResponse, error) {
	return s.client.CreateClient(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *orgClientServiceWrapper) DeleteClient(ctx context.Context, req *pb.DeleteClientRequest) (*pb.DeleteClientResponse, error) {
	return s.client.DeleteClient(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *orgClientServiceWrapper) GetCredentials(ctx context.Context, req *pb.GetCredentialsRequest) (*pb.GetCredentialsResponse, error) {
	return s.client.GetCredentials(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *orgClientServiceWrapper) UpdateCredentials(ctx context.Context, req *pb.UpdateCredentialsRequest) (*pb.UpdateCredentialsResponse, error) {
	return s.client.UpdateCredentials(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *orgClientServiceWrapper) GrantEndpoint(ctx context.Context, req *pb.GrantEndpointRequest) (*pb.GrantEndpointResponse, error) {
	return s.client.GrantEndpoint(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *orgClientServiceWrapper) RevokeEndpoint(ctx context.Context, req *pb.RevokeEndpointRequest) (*pb.RevokeEndpointResponse, error) {
	return s.client.RevokeEndpoint(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *orgClientServiceWrapper) ChangeClientLimit(ctx context.Context, req *pb.ChangeClientLimitRequest) (*pb.ChangeClientLimitResponse, error) {
	return s.client.ChangeClientLimit(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}
