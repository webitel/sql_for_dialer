// Code generated by go-swagger; DO NOT EDIT.

package outbound_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new outbound resource service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for outbound resource service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateOutboundResource(params *CreateOutboundResourceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateOutboundResourceOK, error)

	CreateOutboundResourceDisplay(params *CreateOutboundResourceDisplayParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateOutboundResourceDisplayOK, error)

	DeleteOutboundResource(params *DeleteOutboundResourceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOutboundResourceOK, error)

	DeleteOutboundResourceDisplay(params *DeleteOutboundResourceDisplayParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOutboundResourceDisplayOK, error)

	PatchOutboundResource(params *PatchOutboundResourceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchOutboundResourceOK, error)

	ReadOutboundResource(params *ReadOutboundResourceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReadOutboundResourceOK, error)

	ReadOutboundResourceDisplay(params *ReadOutboundResourceDisplayParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReadOutboundResourceDisplayOK, error)

	SearchOutboundResource(params *SearchOutboundResourceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchOutboundResourceOK, error)

	SearchOutboundResourceDisplay(params *SearchOutboundResourceDisplayParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchOutboundResourceDisplayOK, error)

	UpdateOutboundResource(params *UpdateOutboundResourceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateOutboundResourceOK, error)

	UpdateOutboundResourceDisplay(params *UpdateOutboundResourceDisplayParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateOutboundResourceDisplayOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateOutboundResource creates outbound resource
*/
func (a *Client) CreateOutboundResource(params *CreateOutboundResourceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateOutboundResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateOutboundResourceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateOutboundResource",
		Method:             "POST",
		PathPattern:        "/call_center/resources",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateOutboundResourceReader{formats: a.formats},
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
	success, ok := result.(*CreateOutboundResourceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateOutboundResource: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateOutboundResourceDisplay creates create outbound resource display
*/
func (a *Client) CreateOutboundResourceDisplay(params *CreateOutboundResourceDisplayParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateOutboundResourceDisplayOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateOutboundResourceDisplayParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateOutboundResourceDisplay",
		Method:             "POST",
		PathPattern:        "/call_center/resources/{resource_id}/display",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateOutboundResourceDisplayReader{formats: a.formats},
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
	success, ok := result.(*CreateOutboundResourceDisplayOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateOutboundResourceDisplay: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteOutboundResource removes outbound resource
*/
func (a *Client) DeleteOutboundResource(params *DeleteOutboundResourceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOutboundResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOutboundResourceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteOutboundResource",
		Method:             "DELETE",
		PathPattern:        "/call_center/resources/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteOutboundResourceReader{formats: a.formats},
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
	success, ok := result.(*DeleteOutboundResourceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteOutboundResource: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteOutboundResourceDisplay removes resource display
*/
func (a *Client) DeleteOutboundResourceDisplay(params *DeleteOutboundResourceDisplayParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOutboundResourceDisplayOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOutboundResourceDisplayParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteOutboundResourceDisplay",
		Method:             "DELETE",
		PathPattern:        "/call_center/resources/{resource_id}/display/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteOutboundResourceDisplayReader{formats: a.formats},
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
	success, ok := result.(*DeleteOutboundResourceDisplayOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteOutboundResourceDisplay: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchOutboundResource patches outbound resource
*/
func (a *Client) PatchOutboundResource(params *PatchOutboundResourceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchOutboundResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchOutboundResourceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchOutboundResource",
		Method:             "PATCH",
		PathPattern:        "/call_center/resources/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PatchOutboundResourceReader{formats: a.formats},
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
	success, ok := result.(*PatchOutboundResourceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchOutboundResource: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReadOutboundResource outbounds resource item
*/
func (a *Client) ReadOutboundResource(params *ReadOutboundResourceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReadOutboundResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadOutboundResourceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReadOutboundResource",
		Method:             "GET",
		PathPattern:        "/call_center/resources/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ReadOutboundResourceReader{formats: a.formats},
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
	success, ok := result.(*ReadOutboundResourceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ReadOutboundResource: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReadOutboundResourceDisplay resources display item
*/
func (a *Client) ReadOutboundResourceDisplay(params *ReadOutboundResourceDisplayParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReadOutboundResourceDisplayOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadOutboundResourceDisplayParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReadOutboundResourceDisplay",
		Method:             "GET",
		PathPattern:        "/call_center/resources/{resource_id}/display/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ReadOutboundResourceDisplayReader{formats: a.formats},
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
	success, ok := result.(*ReadOutboundResourceDisplayOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ReadOutboundResourceDisplay: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SearchOutboundResource lists of outbound resource
*/
func (a *Client) SearchOutboundResource(params *SearchOutboundResourceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchOutboundResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchOutboundResourceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchOutboundResource",
		Method:             "GET",
		PathPattern:        "/call_center/resources",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SearchOutboundResourceReader{formats: a.formats},
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
	success, ok := result.(*SearchOutboundResourceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SearchOutboundResource: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SearchOutboundResourceDisplay lists of resource display
*/
func (a *Client) SearchOutboundResourceDisplay(params *SearchOutboundResourceDisplayParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchOutboundResourceDisplayOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchOutboundResourceDisplayParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchOutboundResourceDisplay",
		Method:             "GET",
		PathPattern:        "/call_center/resources/{resource_id}/display",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SearchOutboundResourceDisplayReader{formats: a.formats},
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
	success, ok := result.(*SearchOutboundResourceDisplayOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SearchOutboundResourceDisplay: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateOutboundResource updates outbound resource
*/
func (a *Client) UpdateOutboundResource(params *UpdateOutboundResourceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateOutboundResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOutboundResourceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateOutboundResource",
		Method:             "PUT",
		PathPattern:        "/call_center/resources/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateOutboundResourceReader{formats: a.formats},
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
	success, ok := result.(*UpdateOutboundResourceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateOutboundResource: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateOutboundResourceDisplay updates resource display
*/
func (a *Client) UpdateOutboundResourceDisplay(params *UpdateOutboundResourceDisplayParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateOutboundResourceDisplayOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOutboundResourceDisplayParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateOutboundResourceDisplay",
		Method:             "PUT",
		PathPattern:        "/call_center/resources/{resource_id}/display/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateOutboundResourceDisplayReader{formats: a.formats},
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
	success, ok := result.(*UpdateOutboundResourceDisplayOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateOutboundResourceDisplay: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
