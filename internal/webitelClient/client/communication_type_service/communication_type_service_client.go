// Code generated by go-swagger; DO NOT EDIT.

package communication_type_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new communication type service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for communication type service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateCommunicationType(params *CreateCommunicationTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateCommunicationTypeOK, error)

	DeleteCommunicationType(params *DeleteCommunicationTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteCommunicationTypeOK, error)

	ReadCommunicationType(params *ReadCommunicationTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReadCommunicationTypeOK, error)

	SearchCommunicationType(params *SearchCommunicationTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchCommunicationTypeOK, error)

	UpdateCommunicationType(params *UpdateCommunicationTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateCommunicationTypeOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateCommunicationType creates communication type
*/
func (a *Client) CreateCommunicationType(params *CreateCommunicationTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateCommunicationTypeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateCommunicationTypeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateCommunicationType",
		Method:             "POST",
		PathPattern:        "/call_center/communication_type",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateCommunicationTypeReader{formats: a.formats},
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
	success, ok := result.(*CreateCommunicationTypeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateCommunicationType: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteCommunicationType removes communication type
*/
func (a *Client) DeleteCommunicationType(params *DeleteCommunicationTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteCommunicationTypeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCommunicationTypeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteCommunicationType",
		Method:             "DELETE",
		PathPattern:        "/call_center/communication_type/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteCommunicationTypeReader{formats: a.formats},
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
	success, ok := result.(*DeleteCommunicationTypeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteCommunicationType: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReadCommunicationType communications type item
*/
func (a *Client) ReadCommunicationType(params *ReadCommunicationTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReadCommunicationTypeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadCommunicationTypeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReadCommunicationType",
		Method:             "GET",
		PathPattern:        "/call_center/communication_type/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ReadCommunicationTypeReader{formats: a.formats},
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
	success, ok := result.(*ReadCommunicationTypeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ReadCommunicationType: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SearchCommunicationType lists of communication type
*/
func (a *Client) SearchCommunicationType(params *SearchCommunicationTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchCommunicationTypeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchCommunicationTypeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchCommunicationType",
		Method:             "GET",
		PathPattern:        "/call_center/communication_type",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SearchCommunicationTypeReader{formats: a.formats},
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
	success, ok := result.(*SearchCommunicationTypeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SearchCommunicationType: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateCommunicationType updates communication type
*/
func (a *Client) UpdateCommunicationType(params *UpdateCommunicationTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateCommunicationTypeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCommunicationTypeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateCommunicationType",
		Method:             "PUT",
		PathPattern:        "/call_center/communication_type/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateCommunicationTypeReader{formats: a.formats},
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
	success, ok := result.(*UpdateCommunicationTypeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateCommunicationType: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}