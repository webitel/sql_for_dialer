// Code generated by go-swagger; DO NOT EDIT.

package agent_service

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

// NewPatchAgentParams creates a new PatchAgentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchAgentParams() *PatchAgentParams {
	return &PatchAgentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAgentParamsWithTimeout creates a new PatchAgentParams object
// with the ability to set a timeout on a request.
func NewPatchAgentParamsWithTimeout(timeout time.Duration) *PatchAgentParams {
	return &PatchAgentParams{
		timeout: timeout,
	}
}

// NewPatchAgentParamsWithContext creates a new PatchAgentParams object
// with the ability to set a context for a request.
func NewPatchAgentParamsWithContext(ctx context.Context) *PatchAgentParams {
	return &PatchAgentParams{
		Context: ctx,
	}
}

// NewPatchAgentParamsWithHTTPClient creates a new PatchAgentParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchAgentParamsWithHTTPClient(client *http.Client) *PatchAgentParams {
	return &PatchAgentParams{
		HTTPClient: client,
	}
}

/* PatchAgentParams contains all the parameters to send to the API endpoint
   for the patch agent operation.

   Typically these are written to a http.Request.
*/
type PatchAgentParams struct {

	// Body.
	Body *models.EnginePatchAgentRequest

	// ID.
	//
	// Format: int64
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch agent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAgentParams) WithDefaults() *PatchAgentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch agent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAgentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch agent params
func (o *PatchAgentParams) WithTimeout(timeout time.Duration) *PatchAgentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch agent params
func (o *PatchAgentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch agent params
func (o *PatchAgentParams) WithContext(ctx context.Context) *PatchAgentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch agent params
func (o *PatchAgentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch agent params
func (o *PatchAgentParams) WithHTTPClient(client *http.Client) *PatchAgentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch agent params
func (o *PatchAgentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch agent params
func (o *PatchAgentParams) WithBody(body *models.EnginePatchAgentRequest) *PatchAgentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch agent params
func (o *PatchAgentParams) SetBody(body *models.EnginePatchAgentRequest) {
	o.Body = body
}

// WithID adds the id to the patch agent params
func (o *PatchAgentParams) WithID(id string) *PatchAgentParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch agent params
func (o *PatchAgentParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAgentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
