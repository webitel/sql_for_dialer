// Code generated by go-swagger; DO NOT EDIT.

package media_file_service

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
)

// NewDeleteMediaFileParams creates a new DeleteMediaFileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteMediaFileParams() *DeleteMediaFileParams {
	return &DeleteMediaFileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMediaFileParamsWithTimeout creates a new DeleteMediaFileParams object
// with the ability to set a timeout on a request.
func NewDeleteMediaFileParamsWithTimeout(timeout time.Duration) *DeleteMediaFileParams {
	return &DeleteMediaFileParams{
		timeout: timeout,
	}
}

// NewDeleteMediaFileParamsWithContext creates a new DeleteMediaFileParams object
// with the ability to set a context for a request.
func NewDeleteMediaFileParamsWithContext(ctx context.Context) *DeleteMediaFileParams {
	return &DeleteMediaFileParams{
		Context: ctx,
	}
}

// NewDeleteMediaFileParamsWithHTTPClient creates a new DeleteMediaFileParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteMediaFileParamsWithHTTPClient(client *http.Client) *DeleteMediaFileParams {
	return &DeleteMediaFileParams{
		HTTPClient: client,
	}
}

/* DeleteMediaFileParams contains all the parameters to send to the API endpoint
   for the delete media file operation.

   Typically these are written to a http.Request.
*/
type DeleteMediaFileParams struct {

	// DomainID.
	//
	// Format: int64
	DomainID *string

	// ID.
	//
	// Format: int64
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete media file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMediaFileParams) WithDefaults() *DeleteMediaFileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete media file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMediaFileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete media file params
func (o *DeleteMediaFileParams) WithTimeout(timeout time.Duration) *DeleteMediaFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete media file params
func (o *DeleteMediaFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete media file params
func (o *DeleteMediaFileParams) WithContext(ctx context.Context) *DeleteMediaFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete media file params
func (o *DeleteMediaFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete media file params
func (o *DeleteMediaFileParams) WithHTTPClient(client *http.Client) *DeleteMediaFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete media file params
func (o *DeleteMediaFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the delete media file params
func (o *DeleteMediaFileParams) WithDomainID(domainID *string) *DeleteMediaFileParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the delete media file params
func (o *DeleteMediaFileParams) SetDomainID(domainID *string) {
	o.DomainID = domainID
}

// WithID adds the id to the delete media file params
func (o *DeleteMediaFileParams) WithID(id string) *DeleteMediaFileParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete media file params
func (o *DeleteMediaFileParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMediaFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DomainID != nil {

		// query param domain_id
		var qrDomainID string

		if o.DomainID != nil {
			qrDomainID = *o.DomainID
		}
		qDomainID := qrDomainID
		if qDomainID != "" {

			if err := r.SetQueryParam("domain_id", qDomainID); err != nil {
				return err
			}
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