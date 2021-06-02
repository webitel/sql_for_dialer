// Code generated by go-swagger; DO NOT EDIT.

package backend_profile_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new backend profile service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for backend profile service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateBackendProfile(params *CreateBackendProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateBackendProfileOK, error)

	DeleteBackendProfile(params *DeleteBackendProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteBackendProfileOK, error)

	PatchBackendProfile(params *PatchBackendProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchBackendProfileOK, error)

	ReadBackendProfile(params *ReadBackendProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReadBackendProfileOK, error)

	SearchBackendProfile(params *SearchBackendProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchBackendProfileOK, error)

	UpdateBackendProfile(params *UpdateBackendProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateBackendProfileOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateBackendProfile creates backend profile
*/
func (a *Client) CreateBackendProfile(params *CreateBackendProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateBackendProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateBackendProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateBackendProfile",
		Method:             "POST",
		PathPattern:        "/storage/backend_profiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateBackendProfileReader{formats: a.formats},
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
	success, ok := result.(*CreateBackendProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateBackendProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteBackendProfile removes backend profile
*/
func (a *Client) DeleteBackendProfile(params *DeleteBackendProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteBackendProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBackendProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteBackendProfile",
		Method:             "DELETE",
		PathPattern:        "/storage/backend_profiles/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteBackendProfileReader{formats: a.formats},
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
	success, ok := result.(*DeleteBackendProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteBackendProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchBackendProfile patches backend profile
*/
func (a *Client) PatchBackendProfile(params *PatchBackendProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchBackendProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchBackendProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchBackendProfile",
		Method:             "PATCH",
		PathPattern:        "/storage/backend_profiles/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PatchBackendProfileReader{formats: a.formats},
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
	success, ok := result.(*PatchBackendProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchBackendProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReadBackendProfile backends profile item
*/
func (a *Client) ReadBackendProfile(params *ReadBackendProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReadBackendProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadBackendProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReadBackendProfile",
		Method:             "GET",
		PathPattern:        "/storage/backend_profiles/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ReadBackendProfileReader{formats: a.formats},
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
	success, ok := result.(*ReadBackendProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ReadBackendProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SearchBackendProfile lists of backend profile
*/
func (a *Client) SearchBackendProfile(params *SearchBackendProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchBackendProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchBackendProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchBackendProfile",
		Method:             "GET",
		PathPattern:        "/storage/backend_profiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SearchBackendProfileReader{formats: a.formats},
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
	success, ok := result.(*SearchBackendProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SearchBackendProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateBackendProfile updates backend profile
*/
func (a *Client) UpdateBackendProfile(params *UpdateBackendProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateBackendProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateBackendProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateBackendProfile",
		Method:             "PUT",
		PathPattern:        "/storage/backend_profiles/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateBackendProfileReader{formats: a.formats},
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
	success, ok := result.(*UpdateBackendProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateBackendProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}