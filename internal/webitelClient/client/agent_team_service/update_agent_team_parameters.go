// Code generated by go-swagger; DO NOT EDIT.

package agent_team_service

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

// NewUpdateAgentTeamParams creates a new UpdateAgentTeamParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAgentTeamParams() *UpdateAgentTeamParams {
	return &UpdateAgentTeamParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAgentTeamParamsWithTimeout creates a new UpdateAgentTeamParams object
// with the ability to set a timeout on a request.
func NewUpdateAgentTeamParamsWithTimeout(timeout time.Duration) *UpdateAgentTeamParams {
	return &UpdateAgentTeamParams{
		timeout: timeout,
	}
}

// NewUpdateAgentTeamParamsWithContext creates a new UpdateAgentTeamParams object
// with the ability to set a context for a request.
func NewUpdateAgentTeamParamsWithContext(ctx context.Context) *UpdateAgentTeamParams {
	return &UpdateAgentTeamParams{
		Context: ctx,
	}
}

// NewUpdateAgentTeamParamsWithHTTPClient creates a new UpdateAgentTeamParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAgentTeamParamsWithHTTPClient(client *http.Client) *UpdateAgentTeamParams {
	return &UpdateAgentTeamParams{
		HTTPClient: client,
	}
}

/* UpdateAgentTeamParams contains all the parameters to send to the API endpoint
   for the update agent team operation.

   Typically these are written to a http.Request.
*/
type UpdateAgentTeamParams struct {

	// Body.
	Body *models.EngineUpdateAgentTeamRequest

	// ID.
	//
	// Format: int64
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update agent team params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAgentTeamParams) WithDefaults() *UpdateAgentTeamParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update agent team params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAgentTeamParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update agent team params
func (o *UpdateAgentTeamParams) WithTimeout(timeout time.Duration) *UpdateAgentTeamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update agent team params
func (o *UpdateAgentTeamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update agent team params
func (o *UpdateAgentTeamParams) WithContext(ctx context.Context) *UpdateAgentTeamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update agent team params
func (o *UpdateAgentTeamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update agent team params
func (o *UpdateAgentTeamParams) WithHTTPClient(client *http.Client) *UpdateAgentTeamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update agent team params
func (o *UpdateAgentTeamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update agent team params
func (o *UpdateAgentTeamParams) WithBody(body *models.EngineUpdateAgentTeamRequest) *UpdateAgentTeamParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update agent team params
func (o *UpdateAgentTeamParams) SetBody(body *models.EngineUpdateAgentTeamRequest) {
	o.Body = body
}

// WithID adds the id to the update agent team params
func (o *UpdateAgentTeamParams) WithID(id string) *UpdateAgentTeamParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update agent team params
func (o *UpdateAgentTeamParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAgentTeamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
