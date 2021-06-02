// Code generated by go-swagger; DO NOT EDIT.

package queue_resources_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new queue resources service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for queue resources service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateQueueResourceGroup(params *CreateQueueResourceGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateQueueResourceGroupOK, error)

	DeleteQueueResourceGroup(params *DeleteQueueResourceGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteQueueResourceGroupOK, error)

	ReadQueueResourceGroup(params *ReadQueueResourceGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReadQueueResourceGroupOK, error)

	SearchQueueResourceGroup(params *SearchQueueResourceGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchQueueResourceGroupOK, error)

	UpdateQueueResourceGroup(params *UpdateQueueResourceGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateQueueResourceGroupOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateQueueResourceGroup creates queue resource group
*/
func (a *Client) CreateQueueResourceGroup(params *CreateQueueResourceGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateQueueResourceGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateQueueResourceGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateQueueResourceGroup",
		Method:             "POST",
		PathPattern:        "/call_center/queues/{queue_id}/resource_groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateQueueResourceGroupReader{formats: a.formats},
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
	success, ok := result.(*CreateQueueResourceGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateQueueResourceGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteQueueResourceGroup deletes queue resource group
*/
func (a *Client) DeleteQueueResourceGroup(params *DeleteQueueResourceGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteQueueResourceGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteQueueResourceGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteQueueResourceGroup",
		Method:             "DELETE",
		PathPattern:        "/call_center/queues/{queue_id}/resource_groups/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteQueueResourceGroupReader{formats: a.formats},
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
	success, ok := result.(*DeleteQueueResourceGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteQueueResourceGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReadQueueResourceGroup reads queue resource group
*/
func (a *Client) ReadQueueResourceGroup(params *ReadQueueResourceGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReadQueueResourceGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadQueueResourceGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReadQueueResourceGroup",
		Method:             "GET",
		PathPattern:        "/call_center/queues/{queue_id}/resource_groups/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ReadQueueResourceGroupReader{formats: a.formats},
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
	success, ok := result.(*ReadQueueResourceGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ReadQueueResourceGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SearchQueueResourceGroup searches queue resource group
*/
func (a *Client) SearchQueueResourceGroup(params *SearchQueueResourceGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchQueueResourceGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchQueueResourceGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchQueueResourceGroup",
		Method:             "GET",
		PathPattern:        "/call_center/queues/{queue_id}/resource_groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SearchQueueResourceGroupReader{formats: a.formats},
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
	success, ok := result.(*SearchQueueResourceGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SearchQueueResourceGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateQueueResourceGroup updates queue resource group
*/
func (a *Client) UpdateQueueResourceGroup(params *UpdateQueueResourceGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateQueueResourceGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateQueueResourceGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateQueueResourceGroup",
		Method:             "PUT",
		PathPattern:        "/call_center/queues/{queue_id}/resource_groups/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateQueueResourceGroupReader{formats: a.formats},
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
	success, ok := result.(*UpdateQueueResourceGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateQueueResourceGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
