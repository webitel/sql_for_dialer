// Code generated by go-swagger; DO NOT EDIT.

package list_service

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

// NewUpdateListCommunicationParams creates a new UpdateListCommunicationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateListCommunicationParams() *UpdateListCommunicationParams {
	return &UpdateListCommunicationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateListCommunicationParamsWithTimeout creates a new UpdateListCommunicationParams object
// with the ability to set a timeout on a request.
func NewUpdateListCommunicationParamsWithTimeout(timeout time.Duration) *UpdateListCommunicationParams {
	return &UpdateListCommunicationParams{
		timeout: timeout,
	}
}

// NewUpdateListCommunicationParamsWithContext creates a new UpdateListCommunicationParams object
// with the ability to set a context for a request.
func NewUpdateListCommunicationParamsWithContext(ctx context.Context) *UpdateListCommunicationParams {
	return &UpdateListCommunicationParams{
		Context: ctx,
	}
}

// NewUpdateListCommunicationParamsWithHTTPClient creates a new UpdateListCommunicationParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateListCommunicationParamsWithHTTPClient(client *http.Client) *UpdateListCommunicationParams {
	return &UpdateListCommunicationParams{
		HTTPClient: client,
	}
}

/* UpdateListCommunicationParams contains all the parameters to send to the API endpoint
   for the update list communication operation.

   Typically these are written to a http.Request.
*/
type UpdateListCommunicationParams struct {

	// Body.
	Body *models.EngineUpdateListCommunicationRequest

	// ID.
	//
	// Format: int64
	ID string

	// ListID.
	//
	// Format: int64
	ListID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update list communication params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateListCommunicationParams) WithDefaults() *UpdateListCommunicationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update list communication params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateListCommunicationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update list communication params
func (o *UpdateListCommunicationParams) WithTimeout(timeout time.Duration) *UpdateListCommunicationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update list communication params
func (o *UpdateListCommunicationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update list communication params
func (o *UpdateListCommunicationParams) WithContext(ctx context.Context) *UpdateListCommunicationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update list communication params
func (o *UpdateListCommunicationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update list communication params
func (o *UpdateListCommunicationParams) WithHTTPClient(client *http.Client) *UpdateListCommunicationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update list communication params
func (o *UpdateListCommunicationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update list communication params
func (o *UpdateListCommunicationParams) WithBody(body *models.EngineUpdateListCommunicationRequest) *UpdateListCommunicationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update list communication params
func (o *UpdateListCommunicationParams) SetBody(body *models.EngineUpdateListCommunicationRequest) {
	o.Body = body
}

// WithID adds the id to the update list communication params
func (o *UpdateListCommunicationParams) WithID(id string) *UpdateListCommunicationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update list communication params
func (o *UpdateListCommunicationParams) SetID(id string) {
	o.ID = id
}

// WithListID adds the listID to the update list communication params
func (o *UpdateListCommunicationParams) WithListID(listID string) *UpdateListCommunicationParams {
	o.SetListID(listID)
	return o
}

// SetListID adds the listId to the update list communication params
func (o *UpdateListCommunicationParams) SetListID(listID string) {
	o.ListID = listID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateListCommunicationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param list_id
	if err := r.SetPathParam("list_id", o.ListID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}