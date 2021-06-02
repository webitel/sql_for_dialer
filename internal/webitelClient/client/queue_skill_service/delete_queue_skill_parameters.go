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
)

// NewDeleteQueueSkillParams creates a new DeleteQueueSkillParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteQueueSkillParams() *DeleteQueueSkillParams {
	return &DeleteQueueSkillParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteQueueSkillParamsWithTimeout creates a new DeleteQueueSkillParams object
// with the ability to set a timeout on a request.
func NewDeleteQueueSkillParamsWithTimeout(timeout time.Duration) *DeleteQueueSkillParams {
	return &DeleteQueueSkillParams{
		timeout: timeout,
	}
}

// NewDeleteQueueSkillParamsWithContext creates a new DeleteQueueSkillParams object
// with the ability to set a context for a request.
func NewDeleteQueueSkillParamsWithContext(ctx context.Context) *DeleteQueueSkillParams {
	return &DeleteQueueSkillParams{
		Context: ctx,
	}
}

// NewDeleteQueueSkillParamsWithHTTPClient creates a new DeleteQueueSkillParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteQueueSkillParamsWithHTTPClient(client *http.Client) *DeleteQueueSkillParams {
	return &DeleteQueueSkillParams{
		HTTPClient: client,
	}
}

/* DeleteQueueSkillParams contains all the parameters to send to the API endpoint
   for the delete queue skill operation.

   Typically these are written to a http.Request.
*/
type DeleteQueueSkillParams struct {

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

// WithDefaults hydrates default values in the delete queue skill params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteQueueSkillParams) WithDefaults() *DeleteQueueSkillParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete queue skill params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteQueueSkillParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete queue skill params
func (o *DeleteQueueSkillParams) WithTimeout(timeout time.Duration) *DeleteQueueSkillParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete queue skill params
func (o *DeleteQueueSkillParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete queue skill params
func (o *DeleteQueueSkillParams) WithContext(ctx context.Context) *DeleteQueueSkillParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete queue skill params
func (o *DeleteQueueSkillParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete queue skill params
func (o *DeleteQueueSkillParams) WithHTTPClient(client *http.Client) *DeleteQueueSkillParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete queue skill params
func (o *DeleteQueueSkillParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete queue skill params
func (o *DeleteQueueSkillParams) WithID(id int64) *DeleteQueueSkillParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete queue skill params
func (o *DeleteQueueSkillParams) SetID(id int64) {
	o.ID = id
}

// WithQueueID adds the queueID to the delete queue skill params
func (o *DeleteQueueSkillParams) WithQueueID(queueID int64) *DeleteQueueSkillParams {
	o.SetQueueID(queueID)
	return o
}

// SetQueueID adds the queueId to the delete queue skill params
func (o *DeleteQueueSkillParams) SetQueueID(queueID int64) {
	o.QueueID = queueID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteQueueSkillParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
