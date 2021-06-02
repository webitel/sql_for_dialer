// Code generated by go-swagger; DO NOT EDIT.

package routing_variable_service

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

// NewUpdateRoutingVariableParams creates a new UpdateRoutingVariableParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateRoutingVariableParams() *UpdateRoutingVariableParams {
	return &UpdateRoutingVariableParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRoutingVariableParamsWithTimeout creates a new UpdateRoutingVariableParams object
// with the ability to set a timeout on a request.
func NewUpdateRoutingVariableParamsWithTimeout(timeout time.Duration) *UpdateRoutingVariableParams {
	return &UpdateRoutingVariableParams{
		timeout: timeout,
	}
}

// NewUpdateRoutingVariableParamsWithContext creates a new UpdateRoutingVariableParams object
// with the ability to set a context for a request.
func NewUpdateRoutingVariableParamsWithContext(ctx context.Context) *UpdateRoutingVariableParams {
	return &UpdateRoutingVariableParams{
		Context: ctx,
	}
}

// NewUpdateRoutingVariableParamsWithHTTPClient creates a new UpdateRoutingVariableParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateRoutingVariableParamsWithHTTPClient(client *http.Client) *UpdateRoutingVariableParams {
	return &UpdateRoutingVariableParams{
		HTTPClient: client,
	}
}

/* UpdateRoutingVariableParams contains all the parameters to send to the API endpoint
   for the update routing variable operation.

   Typically these are written to a http.Request.
*/
type UpdateRoutingVariableParams struct {

	// Body.
	Body *models.EngineUpdateRoutingVariableRequest

	// ID.
	//
	// Format: int64
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update routing variable params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRoutingVariableParams) WithDefaults() *UpdateRoutingVariableParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update routing variable params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRoutingVariableParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update routing variable params
func (o *UpdateRoutingVariableParams) WithTimeout(timeout time.Duration) *UpdateRoutingVariableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update routing variable params
func (o *UpdateRoutingVariableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update routing variable params
func (o *UpdateRoutingVariableParams) WithContext(ctx context.Context) *UpdateRoutingVariableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update routing variable params
func (o *UpdateRoutingVariableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update routing variable params
func (o *UpdateRoutingVariableParams) WithHTTPClient(client *http.Client) *UpdateRoutingVariableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update routing variable params
func (o *UpdateRoutingVariableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update routing variable params
func (o *UpdateRoutingVariableParams) WithBody(body *models.EngineUpdateRoutingVariableRequest) *UpdateRoutingVariableParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update routing variable params
func (o *UpdateRoutingVariableParams) SetBody(body *models.EngineUpdateRoutingVariableRequest) {
	o.Body = body
}

// WithID adds the id to the update routing variable params
func (o *UpdateRoutingVariableParams) WithID(id string) *UpdateRoutingVariableParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update routing variable params
func (o *UpdateRoutingVariableParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRoutingVariableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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