// Code generated by go-swagger; DO NOT EDIT.

package bucket_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new bucket service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for bucket service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateBucket(params *CreateBucketParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateBucketOK, error)

	DeleteBucket(params *DeleteBucketParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteBucketOK, error)

	ReadBucket(params *ReadBucketParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReadBucketOK, error)

	SearchBucket(params *SearchBucketParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchBucketOK, error)

	UpdateBucket(params *UpdateBucketParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateBucketOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateBucket creates bucket
*/
func (a *Client) CreateBucket(params *CreateBucketParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateBucketOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateBucketParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateBucket",
		Method:             "POST",
		PathPattern:        "/call_center/buckets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateBucketReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateBucketOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateBucket: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteBucket removes bucket
*/
func (a *Client) DeleteBucket(params *DeleteBucketParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteBucketOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBucketParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteBucket",
		Method:             "DELETE",
		PathPattern:        "/call_center/buckets/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteBucketReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteBucketOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteBucket: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReadBucket buckets item
*/
func (a *Client) ReadBucket(params *ReadBucketParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReadBucketOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadBucketParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReadBucket",
		Method:             "GET",
		PathPattern:        "/call_center/buckets/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ReadBucketReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadBucketOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ReadBucket: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SearchBucket lists of bucket
*/
func (a *Client) SearchBucket(params *SearchBucketParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchBucketOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchBucketParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchBucket",
		Method:             "GET",
		PathPattern:        "/call_center/buckets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SearchBucketReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SearchBucketOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SearchBucket: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateBucket updates bucket
*/
func (a *Client) UpdateBucket(params *UpdateBucketParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateBucketOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateBucketParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateBucket",
		Method:             "PUT",
		PathPattern:        "/call_center/buckets/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateBucketReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateBucketOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateBucket: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
