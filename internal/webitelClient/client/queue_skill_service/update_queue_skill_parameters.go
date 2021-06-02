// Code generated by go-swagger; DO NOT EDIT.

package queue_skill_service

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
	"github.com/go-openapi/swag"

	"github.com/webitel/sql_for_dialer/internal/webitelClient/models"
)

// NewUpdateQueueSkillParams creates a new UpdateQueueSkillParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateQueueSkillParams() *UpdateQueueSkillParams {
	return &UpdateQueueSkillParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateQueueSkillParamsWithTimeout creates a new UpdateQueueSkillParams object
// with the ability to set a timeout on a request.
func NewUpdateQueueSkillParamsWithTimeout(timeout time.Duration) *UpdateQueueSkillParams {
	return &UpdateQueueSkillParams{
		timeout: timeout,
	}
}

// NewUpdateQueueSkillParamsWithContext creates a new UpdateQueueSkillParams object
// with the ability to set a context for a request.
func NewUpdateQueueSkillParamsWithContext(ctx context.Context) *UpdateQueueSkillParams {
	return &UpdateQueueSkillParams{
		Context: ctx,
	}
}

// NewUpdateQueueSkillParamsWithHTTPClient creates a new UpdateQueueSkillParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateQueueSkillParamsWithHTTPClient(client *http.Client) *UpdateQueueSkillParams {
	return &UpdateQueueSkillParams{
		HTTPClient: client,
	}
}

/* UpdateQueueSkillParams contains all the parameters to send to the API endpoint
   for the update queue skill operation.

   Typically these are written to a http.Request.
*/
type UpdateQueueSkillParams struct {

	// Body.
	Body *models.EngineUpdateQueueSkillRequest

	// ID.
	//
	// Format: int64
	ID int64

	// QueueID.
	//
	// Format: int64
	QueueID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update queue skill params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateQueueSkillParams) WithDefaults() *UpdateQueueSkillParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update queue skill params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateQueueSkillParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update queue skill params
func (o *UpdateQueueSkillParams) WithTimeout(timeout time.Duration) *UpdateQueueSkillParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update queue skill params
func (o *UpdateQueueSkillParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update queue skill params
func (o *UpdateQueueSkillParams) WithContext(ctx context.Context) *UpdateQueueSkillParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update queue skill params
func (o *UpdateQueueSkillParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update queue skill params
func (o *UpdateQueueSkillParams) WithHTTPClient(client *http.Client) *UpdateQueueSkillParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update queue skill params
func (o *UpdateQueueSkillParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update queue skill params
func (o *UpdateQueueSkillParams) WithBody(body *models.EngineUpdateQueueSkillRequest) *UpdateQueueSkillParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update queue skill params
func (o *UpdateQueueSkillParams) SetBody(body *models.EngineUpdateQueueSkillRequest) {
	o.Body = body
}

// WithID adds the id to the update queue skill params
func (o *UpdateQueueSkillParams) WithID(id int64) *UpdateQueueSkillParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update queue skill params
func (o *UpdateQueueSkillParams) SetID(id int64) {
	o.ID = id
}

// WithQueueID adds the queueID to the update queue skill params
func (o *UpdateQueueSkillParams) WithQueueID(queueID int64) *UpdateQueueSkillParams {
	o.SetQueueID(queueID)
	return o
}

// SetQueueID adds the queueId to the update queue skill params
func (o *UpdateQueueSkillParams) SetQueueID(queueID int64) {
	o.QueueID = queueID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateQueueSkillParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// path param queue_id
	if err := r.SetPathParam("queue_id", swag.FormatInt64(o.QueueID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
