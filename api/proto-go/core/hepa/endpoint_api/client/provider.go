// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: endpoint_api.proto

package client

import (
	fmt "fmt"
	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/hepa/endpoint_api/pb"
	grpc1 "google.golang.org/grpc"
	reflect "reflect"
	strings "strings"
)

var dependencies = []string{
	"grpc-client@erda.core.hepa.endpoint_api",
	"grpc-client",
}

// +provider
type provider struct {
	client Client
}

func (p *provider) Init(ctx servicehub.Context) error {
	var conn grpc.ClientConnInterface
	for _, dep := range dependencies {
		c, ok := ctx.Service(dep).(grpc.ClientConnInterface)
		if ok {
			conn = c
			break
		}
	}
	if conn == nil {
		return fmt.Errorf("not found connector in (%s)", strings.Join(dependencies, ", "))
	}
	p.client = New(conn)
	return nil
}

var (
	clientsType                  = reflect.TypeOf((*Client)(nil)).Elem()
	endpointApiServiceClientType = reflect.TypeOf((*pb.EndpointApiServiceClient)(nil)).Elem()
	endpointApiServiceServerType = reflect.TypeOf((*pb.EndpointApiServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.core.hepa.endpoint_api-client":
		return p.client
	case "erda.core.hepa.endpoint_api.EndpointApiService":
		return &endpointApiServiceWrapper{client: p.client.EndpointApiService(), opts: opts}
	case "erda.core.hepa.endpoint_api.EndpointApiService.client":
		return p.client.EndpointApiService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case endpointApiServiceClientType:
		return p.client.EndpointApiService()
	case endpointApiServiceServerType:
		return &endpointApiServiceWrapper{client: p.client.EndpointApiService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.core.hepa.endpoint_api-client", &servicehub.Spec{
		Services: []string{
			"erda.core.hepa.endpoint_api.EndpointApiService",
			"erda.core.hepa.endpoint_api-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			endpointApiServiceClientType,
			// server types
			endpointApiServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
