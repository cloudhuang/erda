// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: endpoint_api.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"
	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*ChangeEndpointRootResponse)(nil)
var _ json.Unmarshaler = (*ChangeEndpointRootResponse)(nil)
var _ json.Marshaler = (*ChangeEndpointRootRequest)(nil)
var _ json.Unmarshaler = (*ChangeEndpointRootRequest)(nil)
var _ json.Marshaler = (*DeleteEndpointApiRequest)(nil)
var _ json.Unmarshaler = (*DeleteEndpointApiRequest)(nil)
var _ json.Marshaler = (*DeleteEndpointApiResponse)(nil)
var _ json.Unmarshaler = (*DeleteEndpointApiResponse)(nil)
var _ json.Marshaler = (*UpdateEndpointApiResponse)(nil)
var _ json.Unmarshaler = (*UpdateEndpointApiResponse)(nil)
var _ json.Marshaler = (*UpdateEndpointApiRequest)(nil)
var _ json.Unmarshaler = (*UpdateEndpointApiRequest)(nil)
var _ json.Marshaler = (*CreateEndpointApiResponse)(nil)
var _ json.Unmarshaler = (*CreateEndpointApiResponse)(nil)
var _ json.Marshaler = (*EndpointApi)(nil)
var _ json.Unmarshaler = (*EndpointApi)(nil)
var _ json.Marshaler = (*CreateEndpointApiRequest)(nil)
var _ json.Unmarshaler = (*CreateEndpointApiRequest)(nil)
var _ json.Marshaler = (*GetEndpointApisRequest)(nil)
var _ json.Unmarshaler = (*GetEndpointApisRequest)(nil)
var _ json.Marshaler = (*GetEndpointApisResponse)(nil)
var _ json.Unmarshaler = (*GetEndpointApisResponse)(nil)
var _ json.Marshaler = (*DeleteEndpointRequest)(nil)
var _ json.Unmarshaler = (*DeleteEndpointRequest)(nil)
var _ json.Marshaler = (*DeleteEndpointResponse)(nil)
var _ json.Unmarshaler = (*DeleteEndpointResponse)(nil)
var _ json.Marshaler = (*UpdateEndpointResponse)(nil)
var _ json.Unmarshaler = (*UpdateEndpointResponse)(nil)
var _ json.Marshaler = (*UpdateEndpointRequest)(nil)
var _ json.Unmarshaler = (*UpdateEndpointRequest)(nil)
var _ json.Marshaler = (*CreateEndpointResponse)(nil)
var _ json.Unmarshaler = (*CreateEndpointResponse)(nil)
var _ json.Marshaler = (*CreateEndpointRequest)(nil)
var _ json.Unmarshaler = (*CreateEndpointRequest)(nil)
var _ json.Marshaler = (*GetEndpointsNameRequest)(nil)
var _ json.Unmarshaler = (*GetEndpointsNameRequest)(nil)
var _ json.Marshaler = (*Endpoint)(nil)
var _ json.Unmarshaler = (*Endpoint)(nil)
var _ json.Marshaler = (*GetEndpointsNameResponse)(nil)
var _ json.Unmarshaler = (*GetEndpointsNameResponse)(nil)
var _ json.Marshaler = (*GetEndpointsRequest)(nil)
var _ json.Unmarshaler = (*GetEndpointsRequest)(nil)
var _ json.Marshaler = (*GetEndpointsResponse)(nil)
var _ json.Unmarshaler = (*GetEndpointsResponse)(nil)
var _ json.Marshaler = (*GetEndpointRequest)(nil)
var _ json.Unmarshaler = (*GetEndpointRequest)(nil)
var _ json.Marshaler = (*GetEndpointResponse)(nil)
var _ json.Unmarshaler = (*GetEndpointResponse)(nil)

// ChangeEndpointRootResponse implement json.Marshaler.
func (m *ChangeEndpointRootResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ChangeEndpointRootResponse implement json.Marshaler.
func (m *ChangeEndpointRootResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ChangeEndpointRootRequest implement json.Marshaler.
func (m *ChangeEndpointRootRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ChangeEndpointRootRequest implement json.Marshaler.
func (m *ChangeEndpointRootRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteEndpointApiRequest implement json.Marshaler.
func (m *DeleteEndpointApiRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteEndpointApiRequest implement json.Marshaler.
func (m *DeleteEndpointApiRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteEndpointApiResponse implement json.Marshaler.
func (m *DeleteEndpointApiResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteEndpointApiResponse implement json.Marshaler.
func (m *DeleteEndpointApiResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateEndpointApiResponse implement json.Marshaler.
func (m *UpdateEndpointApiResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateEndpointApiResponse implement json.Marshaler.
func (m *UpdateEndpointApiResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateEndpointApiRequest implement json.Marshaler.
func (m *UpdateEndpointApiRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateEndpointApiRequest implement json.Marshaler.
func (m *UpdateEndpointApiRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateEndpointApiResponse implement json.Marshaler.
func (m *CreateEndpointApiResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateEndpointApiResponse implement json.Marshaler.
func (m *CreateEndpointApiResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// EndpointApi implement json.Marshaler.
func (m *EndpointApi) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// EndpointApi implement json.Marshaler.
func (m *EndpointApi) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateEndpointApiRequest implement json.Marshaler.
func (m *CreateEndpointApiRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateEndpointApiRequest implement json.Marshaler.
func (m *CreateEndpointApiRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetEndpointApisRequest implement json.Marshaler.
func (m *GetEndpointApisRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetEndpointApisRequest implement json.Marshaler.
func (m *GetEndpointApisRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetEndpointApisResponse implement json.Marshaler.
func (m *GetEndpointApisResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetEndpointApisResponse implement json.Marshaler.
func (m *GetEndpointApisResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteEndpointRequest implement json.Marshaler.
func (m *DeleteEndpointRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteEndpointRequest implement json.Marshaler.
func (m *DeleteEndpointRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteEndpointResponse implement json.Marshaler.
func (m *DeleteEndpointResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteEndpointResponse implement json.Marshaler.
func (m *DeleteEndpointResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateEndpointResponse implement json.Marshaler.
func (m *UpdateEndpointResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateEndpointResponse implement json.Marshaler.
func (m *UpdateEndpointResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateEndpointRequest implement json.Marshaler.
func (m *UpdateEndpointRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateEndpointRequest implement json.Marshaler.
func (m *UpdateEndpointRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateEndpointResponse implement json.Marshaler.
func (m *CreateEndpointResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateEndpointResponse implement json.Marshaler.
func (m *CreateEndpointResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateEndpointRequest implement json.Marshaler.
func (m *CreateEndpointRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateEndpointRequest implement json.Marshaler.
func (m *CreateEndpointRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetEndpointsNameRequest implement json.Marshaler.
func (m *GetEndpointsNameRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetEndpointsNameRequest implement json.Marshaler.
func (m *GetEndpointsNameRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Endpoint implement json.Marshaler.
func (m *Endpoint) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Endpoint implement json.Marshaler.
func (m *Endpoint) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetEndpointsNameResponse implement json.Marshaler.
func (m *GetEndpointsNameResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetEndpointsNameResponse implement json.Marshaler.
func (m *GetEndpointsNameResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetEndpointsRequest implement json.Marshaler.
func (m *GetEndpointsRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetEndpointsRequest implement json.Marshaler.
func (m *GetEndpointsRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetEndpointsResponse implement json.Marshaler.
func (m *GetEndpointsResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetEndpointsResponse implement json.Marshaler.
func (m *GetEndpointsResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetEndpointRequest implement json.Marshaler.
func (m *GetEndpointRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetEndpointRequest implement json.Marshaler.
func (m *GetEndpointRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetEndpointResponse implement json.Marshaler.
func (m *GetEndpointResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetEndpointResponse implement json.Marshaler.
func (m *GetEndpointResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
