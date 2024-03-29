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

// NewSearchAttemptsParams creates a new SearchAttemptsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchAttemptsParams() *SearchAttemptsParams {
	return &SearchAttemptsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchAttemptsParamsWithTimeout creates a new SearchAttemptsParams object
// with the ability to set a timeout on a request.
func NewSearchAttemptsParamsWithTimeout(timeout time.Duration) *SearchAttemptsParams {
	return &SearchAttemptsParams{
		timeout: timeout,
	}
}

// NewSearchAttemptsParamsWithContext creates a new SearchAttemptsParams object
// with the ability to set a context for a request.
func NewSearchAttemptsParamsWithContext(ctx context.Context) *SearchAttemptsParams {
	return &SearchAttemptsParams{
		Context: ctx,
	}
}

// NewSearchAttemptsParamsWithHTTPClient creates a new SearchAttemptsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchAttemptsParamsWithHTTPClient(client *http.Client) *SearchAttemptsParams {
	return &SearchAttemptsParams{
		HTTPClient: client,
	}
}

/* SearchAttemptsParams contains all the parameters to send to the API endpoint
   for the search attempts operation.

   Typically these are written to a http.Request.
*/
type SearchAttemptsParams struct {

	// AgentID.
	AgentID []string

	// BucketID.
	BucketID []string

	// DomainID.
	//
	// Format: int64
	DomainID *string

	// Fields.
	Fields []string

	// ID.
	ID []string

	// JoinedAtFrom.
	//
	// Format: int64
	JoinedAtFrom *string

	// JoinedAtTo.
	//
	// Format: int64
	JoinedAtTo *string

	// MemberID.
	MemberID []string

	// Page.
	//
	// Format: int32
	Page *int32

	// QueueID.
	QueueID []string

	// Result.
	Result *string

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

// WithDefaults hydrates default values in the search attempts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchAttemptsParams) WithDefaults() *SearchAttemptsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search attempts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchAttemptsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search attempts params
func (o *SearchAttemptsParams) WithTimeout(timeout time.Duration) *SearchAttemptsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search attempts params
func (o *SearchAttemptsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search attempts params
func (o *SearchAttemptsParams) WithContext(ctx context.Context) *SearchAttemptsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search attempts params
func (o *SearchAttemptsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search attempts params
func (o *SearchAttemptsParams) WithHTTPClient(client *http.Client) *SearchAttemptsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search attempts params
func (o *SearchAttemptsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentID adds the agentID to the search attempts params
func (o *SearchAttemptsParams) WithAgentID(agentID []string) *SearchAttemptsParams {
	o.SetAgentID(agentID)
	return o
}

// SetAgentID adds the agentId to the search attempts params
func (o *SearchAttemptsParams) SetAgentID(agentID []string) {
	o.AgentID = agentID
}

// WithBucketID adds the bucketID to the search attempts params
func (o *SearchAttemptsParams) WithBucketID(bucketID []string) *SearchAttemptsParams {
	o.SetBucketID(bucketID)
	return o
}

// SetBucketID adds the bucketId to the search attempts params
func (o *SearchAttemptsParams) SetBucketID(bucketID []string) {
	o.BucketID = bucketID
}

// WithDomainID adds the domainID to the search attempts params
func (o *SearchAttemptsParams) WithDomainID(domainID *string) *SearchAttemptsParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the search attempts params
func (o *SearchAttemptsParams) SetDomainID(domainID *string) {
	o.DomainID = domainID
}

// WithFields adds the fields to the search attempts params
func (o *SearchAttemptsParams) WithFields(fields []string) *SearchAttemptsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the search attempts params
func (o *SearchAttemptsParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithID adds the id to the search attempts params
func (o *SearchAttemptsParams) WithID(id []string) *SearchAttemptsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the search attempts params
func (o *SearchAttemptsParams) SetID(id []string) {
	o.ID = id
}

// WithJoinedAtFrom adds the joinedAtFrom to the search attempts params
func (o *SearchAttemptsParams) WithJoinedAtFrom(joinedAtFrom *string) *SearchAttemptsParams {
	o.SetJoinedAtFrom(joinedAtFrom)
	return o
}

// SetJoinedAtFrom adds the joinedAtFrom to the search attempts params
func (o *SearchAttemptsParams) SetJoinedAtFrom(joinedAtFrom *string) {
	o.JoinedAtFrom = joinedAtFrom
}

// WithJoinedAtTo adds the joinedAtTo to the search attempts params
func (o *SearchAttemptsParams) WithJoinedAtTo(joinedAtTo *string) *SearchAttemptsParams {
	o.SetJoinedAtTo(joinedAtTo)
	return o
}

// SetJoinedAtTo adds the joinedAtTo to the search attempts params
func (o *SearchAttemptsParams) SetJoinedAtTo(joinedAtTo *string) {
	o.JoinedAtTo = joinedAtTo
}

// WithMemberID adds the memberID to the search attempts params
func (o *SearchAttemptsParams) WithMemberID(memberID []string) *SearchAttemptsParams {
	o.SetMemberID(memberID)
	return o
}

// SetMemberID adds the memberId to the search attempts params
func (o *SearchAttemptsParams) SetMemberID(memberID []string) {
	o.MemberID = memberID
}

// WithPage adds the page to the search attempts params
func (o *SearchAttemptsParams) WithPage(page *int32) *SearchAttemptsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the search attempts params
func (o *SearchAttemptsParams) SetPage(page *int32) {
	o.Page = page
}

// WithQueueID adds the queueID to the search attempts params
func (o *SearchAttemptsParams) WithQueueID(queueID []string) *SearchAttemptsParams {
	o.SetQueueID(queueID)
	return o
}

// SetQueueID adds the queueId to the search attempts params
func (o *SearchAttemptsParams) SetQueueID(queueID []string) {
	o.QueueID = queueID
}

// WithResult adds the result to the search attempts params
func (o *SearchAttemptsParams) WithResult(result *string) *SearchAttemptsParams {
	o.SetResult(result)
	return o
}

// SetResult adds the result to the search attempts params
func (o *SearchAttemptsParams) SetResult(result *string) {
	o.Result = result
}

// WithSize adds the size to the search attempts params
func (o *SearchAttemptsParams) WithSize(size *int32) *SearchAttemptsParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the search attempts params
func (o *SearchAttemptsParams) SetSize(size *int32) {
	o.Size = size
}

// WithSort adds the sort to the search attempts params
func (o *SearchAttemptsParams) WithSort(sort *string) *SearchAttemptsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the search attempts params
func (o *SearchAttemptsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *SearchAttemptsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AgentID != nil {

		// binding items for agent_id
		joinedAgentID := o.bindParamAgentID(reg)

		// query array param agent_id
		if err := r.SetQueryParam("agent_id", joinedAgentID...); err != nil {
			return err
		}
	}

	if o.BucketID != nil {

		// binding items for bucket_id
		joinedBucketID := o.bindParamBucketID(reg)

		// query array param bucket_id
		if err := r.SetQueryParam("bucket_id", joinedBucketID...); err != nil {
			return err
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

	if o.JoinedAtFrom != nil {

		// query param joined_at.from
		var qrJoinedAtFrom string

		if o.JoinedAtFrom != nil {
			qrJoinedAtFrom = *o.JoinedAtFrom
		}
		qJoinedAtFrom := qrJoinedAtFrom
		if qJoinedAtFrom != "" {

			if err := r.SetQueryParam("joined_at.from", qJoinedAtFrom); err != nil {
				return err
			}
		}
	}

	if o.JoinedAtTo != nil {

		// query param joined_at.to
		var qrJoinedAtTo string

		if o.JoinedAtTo != nil {
			qrJoinedAtTo = *o.JoinedAtTo
		}
		qJoinedAtTo := qrJoinedAtTo
		if qJoinedAtTo != "" {

			if err := r.SetQueryParam("joined_at.to", qJoinedAtTo); err != nil {
				return err
			}
		}
	}

	if o.MemberID != nil {

		// binding items for member_id
		joinedMemberID := o.bindParamMemberID(reg)

		// query array param member_id
		if err := r.SetQueryParam("member_id", joinedMemberID...); err != nil {
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

	if o.QueueID != nil {

		// binding items for queue_id
		joinedQueueID := o.bindParamQueueID(reg)

		// query array param queue_id
		if err := r.SetQueryParam("queue_id", joinedQueueID...); err != nil {
			return err
		}
	}

	if o.Result != nil {

		// query param result
		var qrResult string

		if o.Result != nil {
			qrResult = *o.Result
		}
		qResult := qrResult
		if qResult != "" {

			if err := r.SetQueryParam("result", qResult); err != nil {
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

// bindParamSearchAttempts binds the parameter agent_id
func (o *SearchAttemptsParams) bindParamAgentID(formats strfmt.Registry) []string {
	agentIDIR := o.AgentID

	var agentIDIC []string
	for _, agentIDIIR := range agentIDIR { // explode []string

		agentIDIIV := agentIDIIR // string as string
		agentIDIC = append(agentIDIC, agentIDIIV)
	}

	// items.CollectionFormat: "multi"
	agentIDIS := swag.JoinByFormat(agentIDIC, "multi")

	return agentIDIS
}

// bindParamSearchAttempts binds the parameter bucket_id
func (o *SearchAttemptsParams) bindParamBucketID(formats strfmt.Registry) []string {
	bucketIDIR := o.BucketID

	var bucketIDIC []string
	for _, bucketIDIIR := range bucketIDIR { // explode []string

		bucketIDIIV := bucketIDIIR // string as string
		bucketIDIC = append(bucketIDIC, bucketIDIIV)
	}

	// items.CollectionFormat: "multi"
	bucketIDIS := swag.JoinByFormat(bucketIDIC, "multi")

	return bucketIDIS
}

// bindParamSearchAttempts binds the parameter fields
func (o *SearchAttemptsParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamSearchAttempts binds the parameter id
func (o *SearchAttemptsParams) bindParamID(formats strfmt.Registry) []string {
	iDIR := o.ID

	var iDIC []string
	for _, iDIIR := range iDIR { // explode []string

		iDIIV := iDIIR // string as string
		iDIC = append(iDIC, iDIIV)
	}

	// items.CollectionFormat: "multi"
	iDIS := swag.JoinByFormat(iDIC, "multi")

	return iDIS
}

// bindParamSearchAttempts binds the parameter member_id
func (o *SearchAttemptsParams) bindParamMemberID(formats strfmt.Registry) []string {
	memberIDIR := o.MemberID

	var memberIDIC []string
	for _, memberIDIIR := range memberIDIR { // explode []string

		memberIDIIV := memberIDIIR // string as string
		memberIDIC = append(memberIDIC, memberIDIIV)
	}

	// items.CollectionFormat: "multi"
	memberIDIS := swag.JoinByFormat(memberIDIC, "multi")

	return memberIDIS
}

// bindParamSearchAttempts binds the parameter queue_id
func (o *SearchAttemptsParams) bindParamQueueID(formats strfmt.Registry) []string {
	queueIDIR := o.QueueID

	var queueIDIC []string
	for _, queueIDIIR := range queueIDIR { // explode []string

		queueIDIIV := queueIDIIR // string as string
		queueIDIC = append(queueIDIC, queueIDIIV)
	}

	// items.CollectionFormat: "multi"
	queueIDIS := swag.JoinByFormat(queueIDIC, "multi")

	return queueIDIS
}
