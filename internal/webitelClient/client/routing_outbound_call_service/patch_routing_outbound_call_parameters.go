// Code generated by go-swagger; DO NOT EDIT.

package routing_outbound_call_service

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

	"github.com/webitel/sql_for_dialer/internal/webitelClient/models"
)

// NewPatchRoutingOutboundCallParams creates a new PatchRoutingOutboundCallParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchRoutingOutboundCallParams() *PatchRoutingOutboundCallParams {
	return &PatchRoutingOutboundCallParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchRoutingOutboundCallParamsWithTimeout creates a new PatchRoutingOutboundCallParams object
// with the ability to set a timeout on a request.
func NewPatchRoutingOutboundCallParamsWithTimeout(timeout time.Duration) *PatchRoutingOutboundCallParams {
	return &PatchRoutingOutboundCallParams{
		timeout: timeout,
	}
}

// NewPatchRoutingOutboundCallParamsWithContext creates a new PatchRoutingOutboundCallParams object
// with the ability to set a context for a request.
func NewPatchRoutingOutboundCallParamsWithContext(ctx context.Context) *PatchRoutingOutboundCallParams {
	return &PatchRoutingOutboundCallParams{
		Context: ctx,
	}
}

// NewPatchRoutingOutboundCallParamsWithHTTPClient creates a new PatchRoutingOutboundCallParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchRoutingOutboundCallParamsWithHTTPClient(client *http.Client) *PatchRoutingOutboundCallParams {
	return &PatchRoutingOutboundCallParams{
		HTTPClient: client,
	}
}

/* PatchRoutingOutboundCallParams contains all the parameters to send to the API endpoint
   for the patch routing outbound call operation.

   Typically these are written to a http.Request.
*/
type PatchRoutingOutboundCallParams struct {

	// Body.
	Body *models.EnginePatchRoutingOutboundCallRequest

	// ID.
	//
	// Format: int64
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch routing outbound call params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchRoutingOutboundCallParams) WithDefaults() *PatchRoutingOutboundCallParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch routing outbound call params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchRoutingOutboundCallParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch routing outbound call params
func (o *PatchRoutingOutboundCallParams) WithTimeout(timeout time.Duration) *PatchRoutingOutboundCallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch routing outbound call params
func (o *PatchRoutingOutboundCallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch routing outbound call params
func (o *PatchRoutingOutboundCallParams) WithContext(ctx context.Context) *PatchRoutingOutboundCallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch routing outbound call params
func (o *PatchRoutingOutboundCallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch routing outbound call params
func (o *PatchRoutingOutboundCallParams) WithHTTPClient(client *http.Client) *PatchRoutingOutboundCallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch routing outbound call params
func (o *PatchRoutingOutboundCallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch routing outbound call params
func (o *PatchRoutingOutboundCallParams) WithBody(body *models.EnginePatchRoutingOutboundCallRequest) *PatchRoutingOutboundCallParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch routing outbound call params
func (o *PatchRoutingOutboundCallParams) SetBody(body *models.EnginePatchRoutingOutboundCallRequest) {
	o.Body = body
}

// WithID adds the id to the patch routing outbound call params
func (o *PatchRoutingOutboundCallParams) WithID(id string) *PatchRoutingOutboundCallParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch routing outbound call params
func (o *PatchRoutingOutboundCallParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchRoutingOutboundCallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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
