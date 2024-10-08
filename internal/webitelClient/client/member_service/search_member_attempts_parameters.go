// Code generated by go-swagger; DO NOT EDIT.

package member_service

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

// NewSearchMemberAttemptsParams creates a new SearchMemberAttemptsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchMemberAttemptsParams() *SearchMemberAttemptsParams {
	return &SearchMemberAttemptsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchMemberAttemptsParamsWithTimeout creates a new SearchMemberAttemptsParams object
// with the ability to set a timeout on a request.
func NewSearchMemberAttemptsParamsWithTimeout(timeout time.Duration) *SearchMemberAttemptsParams {
	return &SearchMemberAttemptsParams{
		timeout: timeout,
	}
}

// NewSearchMemberAttemptsParamsWithContext creates a new SearchMemberAttemptsParams object
// with the ability to set a context for a request.
func NewSearchMemberAttemptsParamsWithContext(ctx context.Context) *SearchMemberAttemptsParams {
	return &SearchMemberAttemptsParams{
		Context: ctx,
	}
}

// NewSearchMemberAttemptsParamsWithHTTPClient creates a new SearchMemberAttemptsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchMemberAttemptsParamsWithHTTPClient(client *http.Client) *SearchMemberAttemptsParams {
	return &SearchMemberAttemptsParams{
		HTTPClient: client,
	}
}

/* SearchMemberAttemptsParams contains all the parameters to send to the API endpoint
   for the search member attempts operation.

   Typically these are written to a http.Request.
*/
type SearchMemberAttemptsParams struct {

	// DomainID.
	//
	// Format: int64
	DomainID *string

	// MemberID.
	//
	// Format: int64
	MemberID string

	// QueueID.
	//
	// Format: int64
	QueueID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search member attempts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchMemberAttemptsParams) WithDefaults() *SearchMemberAttemptsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search member attempts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchMemberAttemptsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search member attempts params
func (o *SearchMemberAttemptsParams) WithTimeout(timeout time.Duration) *SearchMemberAttemptsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search member attempts params
func (o *SearchMemberAttemptsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search member attempts params
func (o *SearchMemberAttemptsParams) WithContext(ctx context.Context) *SearchMemberAttemptsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search member attempts params
func (o *SearchMemberAttemptsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search member attempts params
func (o *SearchMemberAttemptsParams) WithHTTPClient(client *http.Client) *SearchMemberAttemptsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search member attempts params
func (o *SearchMemberAttemptsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the search member attempts params
func (o *SearchMemberAttemptsParams) WithDomainID(domainID *string) *SearchMemberAttemptsParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the search member attempts params
func (o *SearchMemberAttemptsParams) SetDomainID(domainID *string) {
	o.DomainID = domainID
}

// WithMemberID adds the memberID to the search member attempts params
func (o *SearchMemberAttemptsParams) WithMemberID(memberID string) *SearchMemberAttemptsParams {
	o.SetMemberID(memberID)
	return o
}

// SetMemberID adds the memberId to the search member attempts params
func (o *SearchMemberAttemptsParams) SetMemberID(memberID string) {
	o.MemberID = memberID
}

// WithQueueID adds the queueID to the search member attempts params
func (o *SearchMemberAttemptsParams) WithQueueID(queueID string) *SearchMemberAttemptsParams {
	o.SetQueueID(queueID)
	return o
}

// SetQueueID adds the queueId to the search member attempts params
func (o *SearchMemberAttemptsParams) SetQueueID(queueID string) {
	o.QueueID = queueID
}

// WriteToRequest writes these params to a swagger request
func (o *SearchMemberAttemptsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param member_id
	if err := r.SetPathParam("member_id", o.MemberID); err != nil {
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
