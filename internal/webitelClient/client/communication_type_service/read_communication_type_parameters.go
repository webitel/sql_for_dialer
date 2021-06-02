// Code generated by go-swagger; DO NOT EDIT.

package communication_type_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewReadCommunicationTypeParams creates a new ReadCommunicationTypeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReadCommunicationTypeParams() *ReadCommunicationTypeParams {
	return &ReadCommunicationTypeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReadCommunicationTypeParamsWithTimeout creates a new ReadCommunicationTypeParams object
// with the ability to set a timeout on a request.
func NewReadCommunicationTypeParamsWithTimeout(timeout time.Duration) *ReadCommunicationTypeParams {
	return &ReadCommunicationTypeParams{
		timeout: timeout,
	}
}

// NewReadCommunicationTypeParamsWithContext creates a new ReadCommunicationTypeParams object
// with the ability to set a context for a request.
func NewReadCommunicationTypeParamsWithContext(ctx context.Context) *ReadCommunicationTypeParams {
	return &ReadCommunicationTypeParams{
		Context: ctx,
	}
}

// NewReadCommunicationTypeParamsWithHTTPClient creates a new ReadCommunicationTypeParams object
// with the ability to set a custom HTTPClient for a request.
func NewReadCommunicationTypeParamsWithHTTPClient(client *http.Client) *ReadCommunicationTypeParams {
	return &ReadCommunicationTypeParams{
		HTTPClient: client,
	}
}

/* ReadCommunicationTypeParams contains all the parameters to send to the API endpoint
   for the read communication type operation.

   Typically these are written to a http.Request.
*/
type ReadCommunicationTypeParams struct {

	// DomainID.
	//
	// Format: int64
	DomainID *string

	// ID.
	//
	// Format: int64
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the read communication type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadCommunicationTypeParams) WithDefaults() *ReadCommunicationTypeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the read communication type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadCommunicationTypeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the read communication type params
func (o *ReadCommunicationTypeParams) WithTimeout(timeout time.Duration) *ReadCommunicationTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read communication type params
func (o *ReadCommunicationTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read communication type params
func (o *ReadCommunicationTypeParams) WithContext(ctx context.Context) *ReadCommunicationTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read communication type params
func (o *ReadCommunicationTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read communication type params
func (o *ReadCommunicationTypeParams) WithHTTPClient(client *http.Client) *ReadCommunicationTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read communication type params
func (o *ReadCommunicationTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the read communication type params
func (o *ReadCommunicationTypeParams) WithDomainID(domainID *string) *ReadCommunicationTypeParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the read communication type params
func (o *ReadCommunicationTypeParams) SetDomainID(domainID *string) {
	o.DomainID = domainID
}

// WithID adds the id to the read communication type params
func (o *ReadCommunicationTypeParams) WithID(id string) *ReadCommunicationTypeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the read communication type params
func (o *ReadCommunicationTypeParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReadCommunicationTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DomainID != nil {

		// query param domain_id
		var qrDomainID string

		if o.DomainID != nil {
			qrDomainID = *o.DomainID
		}
		qDomainID := qrDomainID
		if qDomainID != "" {

			if err := r.SetQueryParam("domain_id", qDomainID); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}