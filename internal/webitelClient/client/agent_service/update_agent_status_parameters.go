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

// NewUpdateAgentStatusParams creates a new UpdateAgentStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAgentStatusParams() *UpdateAgentStatusParams {
	return &UpdateAgentStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAgentStatusParamsWithTimeout creates a new UpdateAgentStatusParams object
// with the ability to set a timeout on a request.
func NewUpdateAgentStatusParamsWithTimeout(timeout time.Duration) *UpdateAgentStatusParams {
	return &UpdateAgentStatusParams{
		timeout: timeout,
	}
}

// NewUpdateAgentStatusParamsWithContext creates a new UpdateAgentStatusParams object
// with the ability to set a context for a request.
func NewUpdateAgentStatusParamsWithContext(ctx context.Context) *UpdateAgentStatusParams {
	return &UpdateAgentStatusParams{
		Context: ctx,
	}
}

// NewUpdateAgentStatusParamsWithHTTPClient creates a new UpdateAgentStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAgentStatusParamsWithHTTPClient(client *http.Client) *UpdateAgentStatusParams {
	return &UpdateAgentStatusParams{
		HTTPClient: client,
	}
}

/* UpdateAgentStatusParams contains all the parameters to send to the API endpoint
   for the update agent status operation.

   Typically these are written to a http.Request.
*/
type UpdateAgentStatusParams struct {

	// Body.
	Body *models.EngineAgentStatusRequest

	// ID.
	//
	// Format: int64
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update agent status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAgentStatusParams) WithDefaults() *UpdateAgentStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update agent status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAgentStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update agent status params
func (o *UpdateAgentStatusParams) WithTimeout(timeout time.Duration) *UpdateAgentStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update agent status params
func (o *UpdateAgentStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update agent status params
func (o *UpdateAgentStatusParams) WithContext(ctx context.Context) *UpdateAgentStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update agent status params
func (o *UpdateAgentStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update agent status params
func (o *UpdateAgentStatusParams) WithHTTPClient(client *http.Client) *UpdateAgentStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update agent status params
func (o *UpdateAgentStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update agent status params
func (o *UpdateAgentStatusParams) WithBody(body *models.EngineAgentStatusRequest) *UpdateAgentStatusParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update agent status params
func (o *UpdateAgentStatusParams) SetBody(body *models.EngineAgentStatusRequest) {
	o.Body = body
}

// WithID adds the id to the update agent status params
func (o *UpdateAgentStatusParams) WithID(id string) *UpdateAgentStatusParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update agent status params
func (o *UpdateAgentStatusParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAgentStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
