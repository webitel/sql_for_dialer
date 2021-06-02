// Code generated by go-swagger; DO NOT EDIT.

package media_file_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new media file service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for media file service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteMediaFile(params *DeleteMediaFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteMediaFileOK, error)

	ReadMediaFile(params *ReadMediaFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReadMediaFileOK, error)

	SearchMediaFile(params *SearchMediaFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchMediaFileOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteMediaFile removes media file
*/
func (a *Client) DeleteMediaFile(params *DeleteMediaFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteMediaFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteMediaFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteMediaFile",
		Method:             "DELETE",
		PathPattern:        "/storage/media/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteMediaFileReader{formats: a.formats},
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
	success, ok := result.(*DeleteMediaFileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteMediaFile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReadMediaFile media file item
*/
func (a *Client) ReadMediaFile(params *ReadMediaFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReadMediaFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadMediaFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReadMediaFile",
		Method:             "GET",
		PathPattern:        "/storage/media/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ReadMediaFileReader{formats: a.formats},
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
	success, ok := result.(*ReadMediaFileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ReadMediaFile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SearchMediaFile searches media file
*/
func (a *Client) SearchMediaFile(params *SearchMediaFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchMediaFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchMediaFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchMediaFile",
		Method:             "GET",
		PathPattern:        "/storage/media",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SearchMediaFileReader{formats: a.formats},
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
	success, ok := result.(*SearchMediaFileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SearchMediaFile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}