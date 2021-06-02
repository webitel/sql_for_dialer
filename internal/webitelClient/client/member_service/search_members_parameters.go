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
	"github.com/go-openapi/swag"
)

// NewSearchMembersParams creates a new SearchMembersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchMembersParams() *SearchMembersParams {
	return &SearchMembersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchMembersParamsWithTimeout creates a new SearchMembersParams object
// with the ability to set a timeout on a request.
func NewSearchMembersParamsWithTimeout(timeout time.Duration) *SearchMembersParams {
	return &SearchMembersParams{
		timeout: timeout,
	}
}

// NewSearchMembersParamsWithContext creates a new SearchMembersParams object
// with the ability to set a context for a request.
func NewSearchMembersParamsWithContext(ctx context.Context) *SearchMembersParams {
	return &SearchMembersParams{
		Context: ctx,
	}
}

// NewSearchMembersParamsWithHTTPClient creates a new SearchMembersParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchMembersParamsWithHTTPClient(client *http.Client) *SearchMembersParams {
	return &SearchMembersParams{
		HTTPClient: client,
	}
}

/* SearchMembersParams contains all the parameters to send to the API endpoint
   for the search members operation.

   Typically these are written to a http.Request.
*/
type SearchMembersParams struct {

	// BucketID.
	//
	// Format: int32
	BucketID *int32

	// Destination.
	Destination *string

	// DomainID.
	//
	// Format: int64
	DomainID *string

	// ID.
	//
	// Format: int64
	ID *string

	// Page.
	//
	// Format: int32
	Page *int32

	// QueueID.
	//
	// Format: int64
	QueueID *string

	// Size.
	//
	// Format: int32
	Size *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search members params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchMembersParams) WithDefaults() *SearchMembersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search members params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchMembersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search members params
func (o *SearchMembersParams) WithTimeout(timeout time.Duration) *SearchMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search members params
func (o *SearchMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search members params
func (o *SearchMembersParams) WithContext(ctx context.Context) *SearchMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search members params
func (o *SearchMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search members params
func (o *SearchMembersParams) WithHTTPClient(client *http.Client) *SearchMembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search members params
func (o *SearchMembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBucketID adds the bucketID to the search members params
func (o *SearchMembersParams) WithBucketID(bucketID *int32) *SearchMembersParams {
	o.SetBucketID(bucketID)
	return o
}

// SetBucketID adds the bucketId to the search members params
func (o *SearchMembersParams) SetBucketID(bucketID *int32) {
	o.BucketID = bucketID
}

// WithDestination adds the destination to the search members params
func (o *SearchMembersParams) WithDestination(destination *string) *SearchMembersParams {
	o.SetDestination(destination)
	return o
}

// SetDestination adds the destination to the search members params
func (o *SearchMembersParams) SetDestination(destination *string) {
	o.Destination = destination
}

// WithDomainID adds the domainID to the search members params
func (o *SearchMembersParams) WithDomainID(domainID *string) *SearchMembersParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the search members params
func (o *SearchMembersParams) SetDomainID(domainID *string) {
	o.DomainID = domainID
}

// WithID adds the id to the search members params
func (o *SearchMembersParams) WithID(id *string) *SearchMembersParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the search members params
func (o *SearchMembersParams) SetID(id *string) {
	o.ID = id
}

// WithPage adds the page to the search members params
func (o *SearchMembersParams) WithPage(page *int32) *SearchMembersParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the search members params
func (o *SearchMembersParams) SetPage(page *int32) {
	o.Page = page
}

// WithQueueID adds the queueID to the search members params
func (o *SearchMembersParams) WithQueueID(queueID *string) *SearchMembersParams {
	o.SetQueueID(queueID)
	return o
}

// SetQueueID adds the queueId to the search members params
func (o *SearchMembersParams) SetQueueID(queueID *string) {
	o.QueueID = queueID
}

// WithSize adds the size to the search members params
func (o *SearchMembersParams) WithSize(size *int32) *SearchMembersParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the search members params
func (o *SearchMembersParams) SetSize(size *int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *SearchMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BucketID != nil {

		// query param bucket_id
		var qrBucketID int32

		if o.BucketID != nil {
			qrBucketID = *o.BucketID
		}
		qBucketID := swag.FormatInt32(qrBucketID)
		if qBucketID != "" {

			if err := r.SetQueryParam("bucket_id", qBucketID); err != nil {
				return err
			}
		}
	}

	if o.Destination != nil {

		// query param destination
		var qrDestination string

		if o.Destination != nil {
			qrDestination = *o.Destination
		}
		qDestination := qrDestination
		if qDestination != "" {

			if err := r.SetQueryParam("destination", qDestination); err != nil {
				return err
			}
		}
	}

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

	if o.ID != nil {

		// query param id
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.Page != nil {

		// query param page
		var qrPage int32

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.QueueID != nil {

		// query param queue_id
		var qrQueueID string

		if o.QueueID != nil {
			qrQueueID = *o.QueueID
		}
		qQueueID := qrQueueID
		if qQueueID != "" {

			if err := r.SetQueryParam("queue_id", qQueueID); err != nil {
				return err
			}
		}
	}

	if o.Size != nil {

		// query param size
		var qrSize int32

		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt32(qrSize)
		if qSize != "" {

			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
