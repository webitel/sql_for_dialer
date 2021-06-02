// Code generated by go-swagger; DO NOT EDIT.

package email_profile_service

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

// NewDeleteEmailProfileParams creates a new DeleteEmailProfileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteEmailProfileParams() *DeleteEmailProfileParams {
	return &DeleteEmailProfileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteEmailProfileParamsWithTimeout creates a new DeleteEmailProfileParams object
// with the ability to set a timeout on a request.
func NewDeleteEmailProfileParamsWithTimeout(timeout time.Duration) *DeleteEmailProfileParams {
	return &DeleteEmailProfileParams{
		timeout: timeout,
	}
}

// NewDeleteEmailProfileParamsWithContext creates a new DeleteEmailProfileParams object
// with the ability to set a context for a request.
func NewDeleteEmailProfileParamsWithContext(ctx context.Context) *DeleteEmailProfileParams {
	return &DeleteEmailProfileParams{
		Context: ctx,
	}
}

// NewDeleteEmailProfileParamsWithHTTPClient creates a new DeleteEmailProfileParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteEmailProfileParamsWithHTTPClient(client *http.Client) *DeleteEmailProfileParams {
	return &DeleteEmailProfileParams{
		HTTPClient: client,
	}
}

/* DeleteEmailProfileParams contains all the parameters to send to the API endpoint
   for the delete email profile operation.

   Typically these are written to a http.Request.
*/
type DeleteEmailProfileParams struct {

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

// WithDefaults hydrates default values in the delete email profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteEmailProfileParams) WithDefaults() *DeleteEmailProfileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete email profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteEmailProfileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete email profile params
func (o *DeleteEmailProfileParams) WithTimeout(timeout time.Duration) *DeleteEmailProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete email profile params
func (o *DeleteEmailProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete email profile params
func (o *DeleteEmailProfileParams) WithContext(ctx context.Context) *DeleteEmailProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete email profile params
func (o *DeleteEmailProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete email profile params
func (o *DeleteEmailProfileParams) WithHTTPClient(client *http.Client) *DeleteEmailProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete email profile params
func (o *DeleteEmailProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the delete email profile params
func (o *DeleteEmailProfileParams) WithDomainID(domainID *string) *DeleteEmailProfileParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the delete email profile params
func (o *DeleteEmailProfileParams) SetDomainID(domainID *string) {
	o.DomainID = domainID
}

// WithID adds the id to the delete email profile params
func (o *DeleteEmailProfileParams) WithID(id string) *DeleteEmailProfileParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete email profile params
func (o *DeleteEmailProfileParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteEmailProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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