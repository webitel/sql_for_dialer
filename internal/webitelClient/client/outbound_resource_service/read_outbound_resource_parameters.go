// Code generated by go-swagger; DO NOT EDIT.

package outbound_resource_service

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

// NewReadOutboundResourceParams creates a new ReadOutboundResourceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReadOutboundResourceParams() *ReadOutboundResourceParams {
	return &ReadOutboundResourceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReadOutboundResourceParamsWithTimeout creates a new ReadOutboundResourceParams object
// with the ability to set a timeout on a request.
func NewReadOutboundResourceParamsWithTimeout(timeout time.Duration) *ReadOutboundResourceParams {
	return &ReadOutboundResourceParams{
		timeout: timeout,
	}
}

// NewReadOutboundResourceParamsWithContext creates a new ReadOutboundResourceParams object
// with the ability to set a context for a request.
func NewReadOutboundResourceParamsWithContext(ctx context.Context) *ReadOutboundResourceParams {
	return &ReadOutboundResourceParams{
		Context: ctx,
	}
}

// NewReadOutboundResourceParamsWithHTTPClient creates a new ReadOutboundResourceParams object
// with the ability to set a custom HTTPClient for a request.
func NewReadOutboundResourceParamsWithHTTPClient(client *http.Client) *ReadOutboundResourceParams {
	return &ReadOutboundResourceParams{
		HTTPClient: client,
	}
}

/* ReadOutboundResourceParams contains all the parameters to send to the API endpoint
   for the read outbound resource operation.

   Typically these are written to a http.Request.
*/
type ReadOutboundResourceParams struct {

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

// WithDefaults hydrates default values in the read outbound resource params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadOutboundResourceParams) WithDefaults() *ReadOutboundResourceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the read outbound resource params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadOutboundResourceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the read outbound resource params
func (o *ReadOutboundResourceParams) WithTimeout(timeout time.Duration) *ReadOutboundResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read outbound resource params
func (o *ReadOutboundResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read outbound resource params
func (o *ReadOutboundResourceParams) WithContext(ctx context.Context) *ReadOutboundResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read outbound resource params
func (o *ReadOutboundResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read outbound resource params
func (o *ReadOutboundResourceParams) WithHTTPClient(client *http.Client) *ReadOutboundResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read outbound resource params
func (o *ReadOutboundResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the read outbound resource params
func (o *ReadOutboundResourceParams) WithDomainID(domainID *string) *ReadOutboundResourceParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the read outbound resource params
func (o *ReadOutboundResourceParams) SetDomainID(domainID *string) {
	o.DomainID = domainID
}

// WithID adds the id to the read outbound resource params
func (o *ReadOutboundResourceParams) WithID(id string) *ReadOutboundResourceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the read outbound resource params
func (o *ReadOutboundResourceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReadOutboundResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
