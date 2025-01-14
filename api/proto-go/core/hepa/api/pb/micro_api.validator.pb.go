// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: micro_api.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/known/structpb"
	_ "github.com/erda-project/erda-proto-go/core/hepa/pb"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *DeleteApiRequest) Validate() error {
	return nil
}
func (this *DeleteApiResponse) Validate() error {
	return nil
}
func (this *ApiRequest) Validate() error {
	return nil
}
func (this *UpdateApiRequest) Validate() error {
	if this.ApiRequest != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ApiRequest); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ApiRequest", err)
		}
	}
	return nil
}
func (this *Policy) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *UpdateApiResponse) Validate() error {
	if this.DisplayPathPrefix != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DisplayPathPrefix); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DisplayPathPrefix", err)
		}
	}
	if this.Method != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Method); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Method", err)
		}
	}
	for _, item := range this.Policies {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Policies", err)
			}
		}
	}
	return nil
}
func (this *GetApisRequest) Validate() error {
	return nil
}
func (this *GetApisResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateApiRequest) Validate() error {
	if this.ApiRequest != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ApiRequest); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ApiRequest", err)
		}
	}
	return nil
}
func (this *CreateApiResponse) Validate() error {
	return nil
}
