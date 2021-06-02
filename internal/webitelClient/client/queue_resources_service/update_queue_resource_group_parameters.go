// Code generated by go-swagger; DO NOT EDIT.

package queue_resources_service

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

// NewUpdateQueueResourceGroupParams creates a new UpdateQueueResourceGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateQueueResourceGroupParams() *UpdateQueueResourceGroupParams {
	return &UpdateQueueResourceGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateQueueResourceGroupParamsWithTimeout creates a new UpdateQueueResourceGroupParams object
// with the ability to set a timeout on a request.
func NewUpdateQueueResourceGroupParamsWithTimeout(timeout time.Duration) *UpdateQueueResourceGroupParams {
	return &UpdateQueueResourceGroupParams{
		timeout: timeout,
	}
}

// NewUpdateQueueResourceGroupParamsWithContext creates a new UpdateQueueResourceGroupParams object
// with the ability to set a context for a request.
func NewUpdateQueueResourceGroupParamsWithContext(ctx context.Context) *UpdateQueueResourceGroupParams {
	return &UpdateQueueResourceGroupParams{
		Context: ctx,
	}
}

// NewUpdateQueueResourceGroupParamsWithHTTPClient creates a new UpdateQueueResourceGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateQueueResourceGroupParamsWithHTTPClient(client *http.Client) *UpdateQueueResourceGroupParams {
	return &UpdateQueueResourceGroupParams{
		HTTPClient: client,
	}
}

/* UpdateQueueResourceGroupParams contains all the parameters to send to the API endpoint
   for the update queue resource group operation.

   Typically these are written to a http.Request.
*/
type UpdateQueueResourceGroupParams struct {

	// Body.
	Body *models.EngineUpdateQueueResourceGroupRequest

	// ID.
	//
	// Format: int64
	ID string

	// QueueID.
	//
	// Format: int64
	QueueID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update queue resource group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateQueueResourceGroupParams) WithDefaults() *UpdateQueueResourceGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update queue resource group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateQueueResourceGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update queue resource group params
func (o *UpdateQueueResourceGroupParams) WithTimeout(timeout time.Duration) *UpdateQueueResourceGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update queue resource group params
func (o *UpdateQueueResourceGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update queue resource group params
func (o *UpdateQueueResourceGroupParams) WithContext(ctx context.Context) *UpdateQueueResourceGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update queue resource group params
func (o *UpdateQueueResourceGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update queue resource group params
func (o *UpdateQueueResourceGroupParams) WithHTTPClient(client *http.Client) *UpdateQueueResourceGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update queue resource group params
func (o *UpdateQueueResourceGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update queue resource group params
func (o *UpdateQueueResourceGroupParams) WithBody(body *models.EngineUpdateQueueResourceGroupRequest) *UpdateQueueResourceGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update queue resource group params
func (o *UpdateQueueResourceGroupParams) SetBody(body *models.EngineUpdateQueueResourceGroupRequest) {
	o.Body = body
}

// WithID adds the id to the update queue resource group params
func (o *UpdateQueueResourceGroupParams) WithID(id string) *UpdateQueueResourceGroupParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update queue resource group params
func (o *UpdateQueueResourceGroupParams) SetID(id string) {
	o.ID = id
}

// WithQueueID adds the queueID to the update queue resource group params
func (o *UpdateQueueResourceGroupParams) WithQueueID(queueID string) *UpdateQueueResourceGroupParams {
	o.SetQueueID(queueID)
	return o
}

// SetQueueID adds the queueId to the update queue resource group params
func (o *UpdateQueueResourceGroupParams) SetQueueID(queueID string) {
	o.QueueID = queueID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateQueueResourceGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param queue_id
	if err := r.SetPathParam("queue_id", o.QueueID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
