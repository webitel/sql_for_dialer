// Code generated by go-swagger; DO NOT EDIT.

package queue_bucket_service

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

// NewCreateQueueBucketParams creates a new CreateQueueBucketParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateQueueBucketParams() *CreateQueueBucketParams {
	return &CreateQueueBucketParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateQueueBucketParamsWithTimeout creates a new CreateQueueBucketParams object
// with the ability to set a timeout on a request.
func NewCreateQueueBucketParamsWithTimeout(timeout time.Duration) *CreateQueueBucketParams {
	return &CreateQueueBucketParams{
		timeout: timeout,
	}
}

// NewCreateQueueBucketParamsWithContext creates a new CreateQueueBucketParams object
// with the ability to set a context for a request.
func NewCreateQueueBucketParamsWithContext(ctx context.Context) *CreateQueueBucketParams {
	return &CreateQueueBucketParams{
		Context: ctx,
	}
}

// NewCreateQueueBucketParamsWithHTTPClient creates a new CreateQueueBucketParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateQueueBucketParamsWithHTTPClient(client *http.Client) *CreateQueueBucketParams {
	return &CreateQueueBucketParams{
		HTTPClient: client,
	}
}

/* CreateQueueBucketParams contains all the parameters to send to the API endpoint
   for the create queue bucket operation.

   Typically these are written to a http.Request.
*/
type CreateQueueBucketParams struct {

	// Body.
	Body *models.EngineCreateQueueBucketRequest

	// QueueID.
	//
	// Format: int64
	QueueID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create queue bucket params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateQueueBucketParams) WithDefaults() *CreateQueueBucketParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create queue bucket params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateQueueBucketParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create queue bucket params
func (o *CreateQueueBucketParams) WithTimeout(timeout time.Duration) *CreateQueueBucketParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create queue bucket params
func (o *CreateQueueBucketParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create queue bucket params
func (o *CreateQueueBucketParams) WithContext(ctx context.Context) *CreateQueueBucketParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create queue bucket params
func (o *CreateQueueBucketParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create queue bucket params
func (o *CreateQueueBucketParams) WithHTTPClient(client *http.Client) *CreateQueueBucketParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create queue bucket params
func (o *CreateQueueBucketParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create queue bucket params
func (o *CreateQueueBucketParams) WithBody(body *models.EngineCreateQueueBucketRequest) *CreateQueueBucketParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create queue bucket params
func (o *CreateQueueBucketParams) SetBody(body *models.EngineCreateQueueBucketRequest) {
	o.Body = body
}

// WithQueueID adds the queueID to the create queue bucket params
func (o *CreateQueueBucketParams) WithQueueID(queueID string) *CreateQueueBucketParams {
	o.SetQueueID(queueID)
	return o
}

// SetQueueID adds the queueId to the create queue bucket params
func (o *CreateQueueBucketParams) SetQueueID(queueID string) {
	o.QueueID = queueID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateQueueBucketParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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