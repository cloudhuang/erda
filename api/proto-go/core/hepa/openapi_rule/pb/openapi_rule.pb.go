// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: openapi_rule.proto

package pb

import (
	pb "github.com/erda-project/erda-proto-go/core/hepa/pb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DeleteLimitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data bool `protobuf:"varint,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DeleteLimitResponse) Reset() {
	*x = DeleteLimitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openapi_rule_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLimitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLimitResponse) ProtoMessage() {}

func (x *DeleteLimitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_openapi_rule_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLimitResponse.ProtoReflect.Descriptor instead.
func (*DeleteLimitResponse) Descriptor() ([]byte, []int) {
	return file_openapi_rule_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteLimitResponse) GetData() bool {
	if x != nil {
		return x.Data
	}
	return false
}

type DeleteLimitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RuleId string `protobuf:"bytes,1,opt,name=ruleId,proto3" json:"ruleId,omitempty"`
}

func (x *DeleteLimitRequest) Reset() {
	*x = DeleteLimitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openapi_rule_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLimitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLimitRequest) ProtoMessage() {}

func (x *DeleteLimitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_openapi_rule_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLimitRequest.ProtoReflect.Descriptor instead.
func (*DeleteLimitRequest) Descriptor() ([]byte, []int) {
	return file_openapi_rule_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteLimitRequest) GetRuleId() string {
	if x != nil {
		return x.RuleId
	}
	return ""
}

type LimitRuleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerId   string     `protobuf:"bytes,1,opt,name=consumerId,proto3" json:"consumerId,omitempty"`
	PackageId    string     `protobuf:"bytes,2,opt,name=packageId,proto3" json:"packageId,omitempty"`
	Method       string     `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	ApiPath      string     `protobuf:"bytes,4,opt,name=apiPath,proto3" json:"apiPath,omitempty"`
	Limit        *LimitType `protobuf:"bytes,5,opt,name=limit,proto3" json:"limit,omitempty"`
	Id           string     `protobuf:"bytes,6,opt,name=id,proto3" json:"id,omitempty"`
	CreateAt     string     `protobuf:"bytes,7,opt,name=createAt,proto3" json:"createAt,omitempty"`
	ConsumerName string     `protobuf:"bytes,8,opt,name=consumerName,proto3" json:"consumerName,omitempty"`
	PackageName  string     `protobuf:"bytes,9,opt,name=packageName,proto3" json:"packageName,omitempty"`
}

func (x *LimitRuleInfo) Reset() {
	*x = LimitRuleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openapi_rule_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitRuleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitRuleInfo) ProtoMessage() {}

func (x *LimitRuleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_openapi_rule_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitRuleInfo.ProtoReflect.Descriptor instead.
func (*LimitRuleInfo) Descriptor() ([]byte, []int) {
	return file_openapi_rule_proto_rawDescGZIP(), []int{2}
}

func (x *LimitRuleInfo) GetConsumerId() string {
	if x != nil {
		return x.ConsumerId
	}
	return ""
}

func (x *LimitRuleInfo) GetPackageId() string {
	if x != nil {
		return x.PackageId
	}
	return ""
}

func (x *LimitRuleInfo) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *LimitRuleInfo) GetApiPath() string {
	if x != nil {
		return x.ApiPath
	}
	return ""
}

func (x *LimitRuleInfo) GetLimit() *LimitType {
	if x != nil {
		return x.Limit
	}
	return nil
}

func (x *LimitRuleInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LimitRuleInfo) GetCreateAt() string {
	if x != nil {
		return x.CreateAt
	}
	return ""
}

func (x *LimitRuleInfo) GetConsumerName() string {
	if x != nil {
		return x.ConsumerName
	}
	return ""
}

func (x *LimitRuleInfo) GetPackageName() string {
	if x != nil {
		return x.PackageName
	}
	return ""
}

type UpdateLimitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *LimitRuleInfo `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UpdateLimitResponse) Reset() {
	*x = UpdateLimitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openapi_rule_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateLimitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLimitResponse) ProtoMessage() {}

func (x *UpdateLimitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_openapi_rule_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLimitResponse.ProtoReflect.Descriptor instead.
func (*UpdateLimitResponse) Descriptor() ([]byte, []int) {
	return file_openapi_rule_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateLimitResponse) GetData() *LimitRuleInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateLimitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RuleId       string        `protobuf:"bytes,1,opt,name=ruleId,proto3" json:"ruleId,omitempty"`
	LimitRequest *LimitRequest `protobuf:"bytes,2,opt,name=limitRequest,proto3" json:"limitRequest,omitempty"`
}

func (x *UpdateLimitRequest) Reset() {
	*x = UpdateLimitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openapi_rule_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateLimitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLimitRequest) ProtoMessage() {}

func (x *UpdateLimitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_openapi_rule_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLimitRequest.ProtoReflect.Descriptor instead.
func (*UpdateLimitRequest) Descriptor() ([]byte, []int) {
	return file_openapi_rule_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateLimitRequest) GetRuleId() string {
	if x != nil {
		return x.RuleId
	}
	return ""
}

func (x *UpdateLimitRequest) GetLimitRequest() *LimitRequest {
	if x != nil {
		return x.LimitRequest
	}
	return nil
}

type CreateLimitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data bool `protobuf:"varint,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateLimitResponse) Reset() {
	*x = CreateLimitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openapi_rule_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLimitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLimitResponse) ProtoMessage() {}

func (x *CreateLimitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_openapi_rule_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLimitResponse.ProtoReflect.Descriptor instead.
func (*CreateLimitResponse) Descriptor() ([]byte, []int) {
	return file_openapi_rule_proto_rawDescGZIP(), []int{5}
}

func (x *CreateLimitResponse) GetData() bool {
	if x != nil {
		return x.Data
	}
	return false
}

type LimitType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Qpd int32 `protobuf:"varint,1,opt,name=qpd,proto3" json:"qpd,omitempty"`
	Qph int32 `protobuf:"varint,2,opt,name=qph,proto3" json:"qph,omitempty"`
	Qpm int32 `protobuf:"varint,3,opt,name=qpm,proto3" json:"qpm,omitempty"`
	Qps int32 `protobuf:"varint,4,opt,name=qps,proto3" json:"qps,omitempty"`
}

func (x *LimitType) Reset() {
	*x = LimitType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openapi_rule_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitType) ProtoMessage() {}

func (x *LimitType) ProtoReflect() protoreflect.Message {
	mi := &file_openapi_rule_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitType.ProtoReflect.Descriptor instead.
func (*LimitType) Descriptor() ([]byte, []int) {
	return file_openapi_rule_proto_rawDescGZIP(), []int{6}
}

func (x *LimitType) GetQpd() int32 {
	if x != nil {
		return x.Qpd
	}
	return 0
}

func (x *LimitType) GetQph() int32 {
	if x != nil {
		return x.Qph
	}
	return 0
}

func (x *LimitType) GetQpm() int32 {
	if x != nil {
		return x.Qpm
	}
	return 0
}

func (x *LimitType) GetQps() int32 {
	if x != nil {
		return x.Qps
	}
	return 0
}

type LimitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerId string     `protobuf:"bytes,1,opt,name=consumerId,proto3" json:"consumerId,omitempty"`
	PackageId  string     `protobuf:"bytes,2,opt,name=packageId,proto3" json:"packageId,omitempty"`
	Method     string     `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	ApiPath    string     `protobuf:"bytes,4,opt,name=apiPath,proto3" json:"apiPath,omitempty"`
	Limit      *LimitType `protobuf:"bytes,5,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *LimitRequest) Reset() {
	*x = LimitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openapi_rule_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitRequest) ProtoMessage() {}

func (x *LimitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_openapi_rule_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitRequest.ProtoReflect.Descriptor instead.
func (*LimitRequest) Descriptor() ([]byte, []int) {
	return file_openapi_rule_proto_rawDescGZIP(), []int{7}
}

func (x *LimitRequest) GetConsumerId() string {
	if x != nil {
		return x.ConsumerId
	}
	return ""
}

func (x *LimitRequest) GetPackageId() string {
	if x != nil {
		return x.PackageId
	}
	return ""
}

func (x *LimitRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *LimitRequest) GetApiPath() string {
	if x != nil {
		return x.ApiPath
	}
	return ""
}

func (x *LimitRequest) GetLimit() *LimitType {
	if x != nil {
		return x.Limit
	}
	return nil
}

type CreateLimitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId    string        `protobuf:"bytes,1,opt,name=projectId,proto3" json:"projectId,omitempty"`
	Env          string        `protobuf:"bytes,2,opt,name=env,proto3" json:"env,omitempty"`
	LimitRequest *LimitRequest `protobuf:"bytes,3,opt,name=limitRequest,proto3" json:"limitRequest,omitempty"`
}

func (x *CreateLimitRequest) Reset() {
	*x = CreateLimitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openapi_rule_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLimitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLimitRequest) ProtoMessage() {}

func (x *CreateLimitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_openapi_rule_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLimitRequest.ProtoReflect.Descriptor instead.
func (*CreateLimitRequest) Descriptor() ([]byte, []int) {
	return file_openapi_rule_proto_rawDescGZIP(), []int{8}
}

func (x *CreateLimitRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *CreateLimitRequest) GetEnv() string {
	if x != nil {
		return x.Env
	}
	return ""
}

func (x *CreateLimitRequest) GetLimitRequest() *LimitRequest {
	if x != nil {
		return x.LimitRequest
	}
	return nil
}

type GetLimitsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerId string `protobuf:"bytes,1,opt,name=consumerId,proto3" json:"consumerId,omitempty"`
	PackageId  string `protobuf:"bytes,2,opt,name=packageId,proto3" json:"packageId,omitempty"`
	PageNo     int64  `protobuf:"varint,3,opt,name=pageNo,proto3" json:"pageNo,omitempty"`
	PageSize   int64  `protobuf:"varint,4,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *GetLimitsRequest) Reset() {
	*x = GetLimitsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openapi_rule_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLimitsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLimitsRequest) ProtoMessage() {}

func (x *GetLimitsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_openapi_rule_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLimitsRequest.ProtoReflect.Descriptor instead.
func (*GetLimitsRequest) Descriptor() ([]byte, []int) {
	return file_openapi_rule_proto_rawDescGZIP(), []int{9}
}

func (x *GetLimitsRequest) GetConsumerId() string {
	if x != nil {
		return x.ConsumerId
	}
	return ""
}

func (x *GetLimitsRequest) GetPackageId() string {
	if x != nil {
		return x.PackageId
	}
	return ""
}

func (x *GetLimitsRequest) GetPageNo() int64 {
	if x != nil {
		return x.PageNo
	}
	return 0
}

func (x *GetLimitsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type GetLimitsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *pb.NewPageResult `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetLimitsResponse) Reset() {
	*x = GetLimitsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openapi_rule_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLimitsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLimitsResponse) ProtoMessage() {}

func (x *GetLimitsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_openapi_rule_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLimitsResponse.ProtoReflect.Descriptor instead.
func (*GetLimitsResponse) Descriptor() ([]byte, []int) {
	return file_openapi_rule_proto_rawDescGZIP(), []int{10}
}

func (x *GetLimitsResponse) GetData() *pb.NewPageResult {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_openapi_rule_proto protoreflect.FileDescriptor

var file_openapi_rule_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x68, 0x65, 0x70, 0x61, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x75, 0x6c,
	0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x68, 0x65, 0x70, 0x61, 0x2f, 0x68, 0x65, 0x70, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2c,
	0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x22, 0xaf, 0x02, 0x0a,
	0x0d, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x69, 0x50, 0x61, 0x74, 0x68, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x69, 0x50, 0x61, 0x74, 0x68, 0x12, 0x3c,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x68, 0x65, 0x70, 0x61, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x55,
	0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x68, 0x65, 0x70, 0x61, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x75, 0x6c,
	0x65, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x7b, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x75, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x75, 0x6c,
	0x65, 0x49, 0x64, 0x12, 0x4d, 0x0a, 0x0c, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65, 0x72, 0x64, 0x61,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x68, 0x65, 0x70, 0x61, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x0c, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x29, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x53, 0x0a,
	0x09, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x70,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x71, 0x70, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x71, 0x70, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x71, 0x70, 0x68, 0x12, 0x10,
	0x0a, 0x03, 0x71, 0x70, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x71, 0x70, 0x6d,
	0x12, 0x10, 0x0a, 0x03, 0x71, 0x70, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x71,
	0x70, 0x73, 0x22, 0xbc, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x69,
	0x50, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x69, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x3c, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x68,
	0x65, 0x70, 0x61, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x75, 0x6c, 0x65,
	0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x22, 0x93, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x4d, 0x0a, 0x0c, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x68, 0x65, 0x70, 0x61, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0c, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x84, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61,
	0x67, 0x65, 0x4e, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x61, 0x67, 0x65,
	0x4e, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x46,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x68, 0x65,
	0x70, 0x61, 0x2e, 0x4e, 0x65, 0x77, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x8e, 0x06, 0x0a, 0x12, 0x4f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xe1, 0x01,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x2d, 0x2e, 0x65, 0x72,
	0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x68, 0x65, 0x70, 0x61, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x65, 0x72, 0x64,
	0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x68, 0x65, 0x70, 0x61, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x75, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x6f, 0x12, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x3f,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x3d, 0x7b, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x7d, 0x26, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x3d, 0x7b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x7d, 0x26, 0x70, 0x61,
	0x67, 0x65, 0x4e, 0x6f, 0x3d, 0x7b, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x7d, 0x26, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x3d, 0x7b, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x7d, 0x12, 0xc3, 0x01, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x2f, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x68, 0x65,
	0x70, 0x61, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x30, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x68,
	0x65, 0x70, 0x61, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x75, 0x6c, 0x65,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x51, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4b, 0x22, 0x3b, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x3f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x3d, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x7d, 0x26,
	0x65, 0x6e, 0x76, 0x3d, 0x7b, 0x65, 0x6e, 0x76, 0x7d, 0x3a, 0x0c, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0xac, 0x01, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x2f, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x68, 0x65, 0x70, 0x61, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x5f, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x68, 0x65, 0x70, 0x61, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3a, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x34, 0x32, 0x24, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2f,
	0x7b, 0x72, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x7d, 0x3a, 0x0c, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x9e, 0x01, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x2f, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x68, 0x65, 0x70, 0x61, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x5f,
	0x72, 0x75, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x68, 0x65, 0x70, 0x61, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x5f, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x26, 0x2a, 0x24, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2f, 0x7b,
	0x72, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x7d, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x72, 0x64, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2f, 0x65, 0x72, 0x64, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x68, 0x65, 0x70, 0x61, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_openapi_rule_proto_rawDescOnce sync.Once
	file_openapi_rule_proto_rawDescData = file_openapi_rule_proto_rawDesc
)

func file_openapi_rule_proto_rawDescGZIP() []byte {
	file_openapi_rule_proto_rawDescOnce.Do(func() {
		file_openapi_rule_proto_rawDescData = protoimpl.X.CompressGZIP(file_openapi_rule_proto_rawDescData)
	})
	return file_openapi_rule_proto_rawDescData
}

var file_openapi_rule_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_openapi_rule_proto_goTypes = []interface{}{
	(*DeleteLimitResponse)(nil), // 0: erda.core.hepa.openapi_rule.DeleteLimitResponse
	(*DeleteLimitRequest)(nil),  // 1: erda.core.hepa.openapi_rule.DeleteLimitRequest
	(*LimitRuleInfo)(nil),       // 2: erda.core.hepa.openapi_rule.LimitRuleInfo
	(*UpdateLimitResponse)(nil), // 3: erda.core.hepa.openapi_rule.UpdateLimitResponse
	(*UpdateLimitRequest)(nil),  // 4: erda.core.hepa.openapi_rule.UpdateLimitRequest
	(*CreateLimitResponse)(nil), // 5: erda.core.hepa.openapi_rule.CreateLimitResponse
	(*LimitType)(nil),           // 6: erda.core.hepa.openapi_rule.LimitType
	(*LimitRequest)(nil),        // 7: erda.core.hepa.openapi_rule.LimitRequest
	(*CreateLimitRequest)(nil),  // 8: erda.core.hepa.openapi_rule.CreateLimitRequest
	(*GetLimitsRequest)(nil),    // 9: erda.core.hepa.openapi_rule.GetLimitsRequest
	(*GetLimitsResponse)(nil),   // 10: erda.core.hepa.openapi_rule.GetLimitsResponse
	(*pb.NewPageResult)(nil),    // 11: erda.core.hepa.NewPageResult
}
var file_openapi_rule_proto_depIdxs = []int32{
	6,  // 0: erda.core.hepa.openapi_rule.LimitRuleInfo.limit:type_name -> erda.core.hepa.openapi_rule.LimitType
	2,  // 1: erda.core.hepa.openapi_rule.UpdateLimitResponse.data:type_name -> erda.core.hepa.openapi_rule.LimitRuleInfo
	7,  // 2: erda.core.hepa.openapi_rule.UpdateLimitRequest.limitRequest:type_name -> erda.core.hepa.openapi_rule.LimitRequest
	6,  // 3: erda.core.hepa.openapi_rule.LimitRequest.limit:type_name -> erda.core.hepa.openapi_rule.LimitType
	7,  // 4: erda.core.hepa.openapi_rule.CreateLimitRequest.limitRequest:type_name -> erda.core.hepa.openapi_rule.LimitRequest
	11, // 5: erda.core.hepa.openapi_rule.GetLimitsResponse.data:type_name -> erda.core.hepa.NewPageResult
	9,  // 6: erda.core.hepa.openapi_rule.OpenapiRuleService.GetLimits:input_type -> erda.core.hepa.openapi_rule.GetLimitsRequest
	8,  // 7: erda.core.hepa.openapi_rule.OpenapiRuleService.CreateLimit:input_type -> erda.core.hepa.openapi_rule.CreateLimitRequest
	4,  // 8: erda.core.hepa.openapi_rule.OpenapiRuleService.UpdateLimit:input_type -> erda.core.hepa.openapi_rule.UpdateLimitRequest
	1,  // 9: erda.core.hepa.openapi_rule.OpenapiRuleService.DeleteLimit:input_type -> erda.core.hepa.openapi_rule.DeleteLimitRequest
	10, // 10: erda.core.hepa.openapi_rule.OpenapiRuleService.GetLimits:output_type -> erda.core.hepa.openapi_rule.GetLimitsResponse
	5,  // 11: erda.core.hepa.openapi_rule.OpenapiRuleService.CreateLimit:output_type -> erda.core.hepa.openapi_rule.CreateLimitResponse
	3,  // 12: erda.core.hepa.openapi_rule.OpenapiRuleService.UpdateLimit:output_type -> erda.core.hepa.openapi_rule.UpdateLimitResponse
	0,  // 13: erda.core.hepa.openapi_rule.OpenapiRuleService.DeleteLimit:output_type -> erda.core.hepa.openapi_rule.DeleteLimitResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_openapi_rule_proto_init() }
func file_openapi_rule_proto_init() {
	if File_openapi_rule_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_openapi_rule_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLimitResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_openapi_rule_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLimitRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_openapi_rule_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitRuleInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_openapi_rule_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateLimitResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_openapi_rule_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateLimitRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_openapi_rule_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLimitResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_openapi_rule_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitType); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_openapi_rule_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_openapi_rule_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLimitRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_openapi_rule_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLimitsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_openapi_rule_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLimitsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_openapi_rule_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_openapi_rule_proto_goTypes,
		DependencyIndexes: file_openapi_rule_proto_depIdxs,
		MessageInfos:      file_openapi_rule_proto_msgTypes,
	}.Build()
	File_openapi_rule_proto = out.File
	file_openapi_rule_proto_rawDesc = nil
	file_openapi_rule_proto_goTypes = nil
	file_openapi_rule_proto_depIdxs = nil
}
