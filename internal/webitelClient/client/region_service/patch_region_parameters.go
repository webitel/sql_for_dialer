// Code generated by go-swagger; DO NOT EDIT.

package region_service

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

// NewPatchRegionParams creates a new PatchRegionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchRegionParams() *PatchRegionParams {
	return &PatchRegionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchRegionParamsWithTimeout creates a new PatchRegionParams object
// with the ability to set a timeout on a request.
func NewPatchRegionParamsWithTimeout(timeout time.Duration) *PatchRegionParams {
	return &PatchRegionParams{
		timeout: timeout,
	}
}

// NewPatchRegionParamsWithContext creates a new PatchRegionParams object
// with the ability to set a context for a request.
func NewPatchRegionParamsWithContext(ctx context.Context) *PatchRegionParams {
	return &PatchRegionParams{
		Context: ctx,
	}
}

// NewPatchRegionParamsWithHTTPClient creates a new PatchRegionParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchRegionParamsWithHTTPClient(client *http.Client) *PatchRegionParams {
	return &PatchRegionParams{
		HTTPClient: client,
	}
}

/* PatchRegionParams contains all the parameters to send to the API endpoint
   for the patch region operation.

   Typically these are written to a http.Request.
*/
type PatchRegionParams struct {

	// Body.
	Body *models.EnginePatchRegionRequest

	// ID.
	//
	// Format: int64
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch region params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchRegionParams) WithDefaults() *PatchRegionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch region params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchRegionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch region params
func (o *PatchRegionParams) WithTimeout(timeout time.Duration) *PatchRegionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch region params
func (o *PatchRegionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch region params
func (o *PatchRegionParams) WithContext(ctx context.Context) *PatchRegionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch region params
func (o *PatchRegionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch region params
func (o *PatchRegionParams) WithHTTPClient(client *http.Client) *PatchRegionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch region params
func (o *PatchRegionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch region params
func (o *PatchRegionParams) WithBody(body *models.EnginePatchRegionRequest) *PatchRegionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch region params
func (o *PatchRegionParams) SetBody(body *models.EnginePatchRegionRequest) {
	o.Body = body
}

// WithID adds the id to the patch region params
func (o *PatchRegionParams) WithID(id int64) *PatchRegionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch region params
func (o *PatchRegionParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchRegionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}