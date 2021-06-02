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
	"github.com/go-openapi/swag"
)

// NewSearchQueueBucketParams creates a new SearchQueueBucketParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchQueueBucketParams() *SearchQueueBucketParams {
	return &SearchQueueBucketParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchQueueBucketParamsWithTimeout creates a new SearchQueueBucketParams object
// with the ability to set a timeout on a request.
func NewSearchQueueBucketParamsWithTimeout(timeout time.Duration) *SearchQueueBucketParams {
	return &SearchQueueBucketParams{
		timeout: timeout,
	}
}

// NewSearchQueueBucketParamsWithContext creates a new SearchQueueBucketParams object
// with the ability to set a context for a request.
func NewSearchQueueBucketParamsWithContext(ctx context.Context) *SearchQueueBucketParams {
	return &SearchQueueBucketParams{
		Context: ctx,
	}
}

// NewSearchQueueBucketParamsWithHTTPClient creates a new SearchQueueBucketParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchQueueBucketParamsWithHTTPClient(client *http.Client) *SearchQueueBucketParams {
	return &SearchQueueBucketParams{
		HTTPClient: client,
	}
}

/* SearchQueueBucketParams contains all the parameters to send to the API endpoint
   for the search queue bucket operation.

   Typically these are written to a http.Request.
*/
type SearchQueueBucketParams struct {

	// Fields.
	Fields []string

	// ID.
	ID []int64

	// Page.
	//
	// Format: int32
	Page *int32

	// Q.
	Q *string

	// QueueID.
	//
	// Format: int64
	QueueID string

	// Size.
	//
	// Format: int32
	Size *int32

	// Sort.
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search queue bucket params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchQueueBucketParams) WithDefaults() *SearchQueueBucketParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search queue bucket params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchQueueBucketParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search queue bucket params
func (o *SearchQueueBucketParams) WithTimeout(timeout time.Duration) *SearchQueueBucketParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search queue bucket params
func (o *SearchQueueBucketParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search queue bucket params
func (o *SearchQueueBucketParams) WithContext(ctx context.Context) *SearchQueueBucketParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search queue bucket params
func (o *SearchQueueBucketParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search queue bucket params
func (o *SearchQueueBucketParams) WithHTTPClient(client *http.Client) *SearchQueueBucketParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search queue bucket params
func (o *SearchQueueBucketParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the search queue bucket params
func (o *SearchQueueBucketParams) WithFields(fields []string) *SearchQueueBucketParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the search queue bucket params
func (o *SearchQueueBucketParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithID adds the id to the search queue bucket params
func (o *SearchQueueBucketParams) WithID(id []int64) *SearchQueueBucketParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the search queue bucket params
func (o *SearchQueueBucketParams) SetID(id []int64) {
	o.ID = id
}

// WithPage adds the page to the search queue bucket params
func (o *SearchQueueBucketParams) WithPage(page *int32) *SearchQueueBucketParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the search queue bucket params
func (o *SearchQueueBucketParams) SetPage(page *int32) {
	o.Page = page
}

// WithQ adds the q to the search queue bucket params
func (o *SearchQueueBucketParams) WithQ(q *string) *SearchQueueBucketParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the search queue bucket params
func (o *SearchQueueBucketParams) SetQ(q *string) {
	o.Q = q
}

// WithQueueID adds the queueID to the search queue bucket params
func (o *SearchQueueBucketParams) WithQueueID(queueID string) *SearchQueueBucketParams {
	o.SetQueueID(queueID)
	return o
}

// SetQueueID adds the queueId to the search queue bucket params
func (o *SearchQueueBucketParams) SetQueueID(queueID string) {
	o.QueueID = queueID
}

// WithSize adds the size to the search queue bucket params
func (o *SearchQueueBucketParams) WithSize(size *int32) *SearchQueueBucketParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the search queue bucket params
func (o *SearchQueueBucketParams) SetSize(size *int32) {
	o.Size = size
}

// WithSort adds the sort to the search queue bucket params
func (o *SearchQueueBucketParams) WithSort(sort *string) *SearchQueueBucketParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the search queue bucket params
func (o *SearchQueueBucketParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *SearchQueueBucketParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.ID != nil {

		// binding items for id
		joinedID := o.bindParamID(reg)

		// query array param id
		if err := r.SetQueryParam("id", joinedID...); err != nil {
			return err
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

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	// path param queue_id
	if err := r.SetPathParam("queue_id", o.QueueID); err != nil {
		return err
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

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSearchQueueBucket binds the parameter fields
func (o *SearchQueueBucketParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.Fields

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "multi"
	fieldsIS := swag.JoinByFormat(fieldsIC, "multi")

	return fieldsIS
}

// bindParamSearchQueueBucket binds the parameter id
func (o *SearchQueueBucketParams) bindParamID(formats strfmt.Registry) []string {
	iDIR := o.ID

	var iDIC []string
	for _, iDIIR := range iDIR { // explode []int64

		iDIIV := swag.FormatInt64(iDIIR) // int64 as string
		iDIC = append(iDIC, iDIIV)
	}

	// items.CollectionFormat: "multi"
	iDIS := swag.JoinByFormat(iDIC, "multi")

	return iDIS
}