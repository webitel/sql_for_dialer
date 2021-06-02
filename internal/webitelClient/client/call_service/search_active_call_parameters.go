// Code generated by go-swagger; DO NOT EDIT.

package call_service

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

// NewSearchActiveCallParams creates a new SearchActiveCallParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchActiveCallParams() *SearchActiveCallParams {
	return &SearchActiveCallParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchActiveCallParamsWithTimeout creates a new SearchActiveCallParams object
// with the ability to set a timeout on a request.
func NewSearchActiveCallParamsWithTimeout(timeout time.Duration) *SearchActiveCallParams {
	return &SearchActiveCallParams{
		timeout: timeout,
	}
}

// NewSearchActiveCallParamsWithContext creates a new SearchActiveCallParams object
// with the ability to set a context for a request.
func NewSearchActiveCallParamsWithContext(ctx context.Context) *SearchActiveCallParams {
	return &SearchActiveCallParams{
		Context: ctx,
	}
}

// NewSearchActiveCallParamsWithHTTPClient creates a new SearchActiveCallParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchActiveCallParamsWithHTTPClient(client *http.Client) *SearchActiveCallParams {
	return &SearchActiveCallParams{
		HTTPClient: client,
	}
}

/* SearchActiveCallParams contains all the parameters to send to the API endpoint
   for the search active call operation.

   Typically these are written to a http.Request.
*/
type SearchActiveCallParams struct {

	// AgentID.
	AgentID []string

	// AnsweredAtFrom.
	//
	// Format: int64
	AnsweredAtFrom *string

	// AnsweredAtTo.
	//
	// Format: int64
	AnsweredAtTo *string

	// Cause.
	Cause []string

	// CreatedAtFrom.
	//
	// Format: int64
	CreatedAtFrom *string

	// CreatedAtTo.
	//
	// Format: int64
	CreatedAtTo *string

	// Direction.
	Direction []string

	// DomainID.
	//
	// Format: int64
	DomainID *string

	// DurationFrom.
	//
	// Format: int64
	DurationFrom *string

	// DurationTo.
	//
	// Format: int64
	DurationTo *string

	// ExistsFile.
	//
	// Format: boolean
	ExistsFile *bool

	// Fields.
	Fields []string

	// GatewayID.
	GatewayID []string

	// MemberID.
	MemberID []string

	// Missed.
	//
	// Format: boolean
	Missed *bool

	// Number.
	Number *string

	// Page.
	//
	// Format: int32
	Page *int32

	// ParentID.
	ParentID *string

	// Q.
	Q *string

	// QueueID.
	QueueID []string

	// Size.
	//
	// Format: int32
	Size *int32

	// SkipParent.
	//
	// Format: boolean
	SkipParent *bool

	// Sort.
	Sort *string

	// StoredAtFrom.
	//
	// Format: int64
	StoredAtFrom *string

	// StoredAtTo.
	//
	// Format: int64
	StoredAtTo *string

	// SupervisorID.
	SupervisorID []string

	// TeamID.
	TeamID []string

	// UserID.
	UserID []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search active call params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchActiveCallParams) WithDefaults() *SearchActiveCallParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search active call params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchActiveCallParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search active call params
func (o *SearchActiveCallParams) WithTimeout(timeout time.Duration) *SearchActiveCallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search active call params
func (o *SearchActiveCallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search active call params
func (o *SearchActiveCallParams) WithContext(ctx context.Context) *SearchActiveCallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search active call params
func (o *SearchActiveCallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search active call params
func (o *SearchActiveCallParams) WithHTTPClient(client *http.Client) *SearchActiveCallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search active call params
func (o *SearchActiveCallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentID adds the agentID to the search active call params
func (o *SearchActiveCallParams) WithAgentID(agentID []string) *SearchActiveCallParams {
	o.SetAgentID(agentID)
	return o
}

// SetAgentID adds the agentId to the search active call params
func (o *SearchActiveCallParams) SetAgentID(agentID []string) {
	o.AgentID = agentID
}

// WithAnsweredAtFrom adds the answeredAtFrom to the search active call params
func (o *SearchActiveCallParams) WithAnsweredAtFrom(answeredAtFrom *string) *SearchActiveCallParams {
	o.SetAnsweredAtFrom(answeredAtFrom)
	return o
}

// SetAnsweredAtFrom adds the answeredAtFrom to the search active call params
func (o *SearchActiveCallParams) SetAnsweredAtFrom(answeredAtFrom *string) {
	o.AnsweredAtFrom = answeredAtFrom
}

// WithAnsweredAtTo adds the answeredAtTo to the search active call params
func (o *SearchActiveCallParams) WithAnsweredAtTo(answeredAtTo *string) *SearchActiveCallParams {
	o.SetAnsweredAtTo(answeredAtTo)
	return o
}

// SetAnsweredAtTo adds the answeredAtTo to the search active call params
func (o *SearchActiveCallParams) SetAnsweredAtTo(answeredAtTo *string) {
	o.AnsweredAtTo = answeredAtTo
}

// WithCause adds the cause to the search active call params
func (o *SearchActiveCallParams) WithCause(cause []string) *SearchActiveCallParams {
	o.SetCause(cause)
	return o
}

// SetCause adds the cause to the search active call params
func (o *SearchActiveCallParams) SetCause(cause []string) {
	o.Cause = cause
}

// WithCreatedAtFrom adds the createdAtFrom to the search active call params
func (o *SearchActiveCallParams) WithCreatedAtFrom(createdAtFrom *string) *SearchActiveCallParams {
	o.SetCreatedAtFrom(createdAtFrom)
	return o
}

// SetCreatedAtFrom adds the createdAtFrom to the search active call params
func (o *SearchActiveCallParams) SetCreatedAtFrom(createdAtFrom *string) {
	o.CreatedAtFrom = createdAtFrom
}

// WithCreatedAtTo adds the createdAtTo to the search active call params
func (o *SearchActiveCallParams) WithCreatedAtTo(createdAtTo *string) *SearchActiveCallParams {
	o.SetCreatedAtTo(createdAtTo)
	return o
}

// SetCreatedAtTo adds the createdAtTo to the search active call params
func (o *SearchActiveCallParams) SetCreatedAtTo(createdAtTo *string) {
	o.CreatedAtTo = createdAtTo
}

// WithDirection adds the direction to the search active call params
func (o *SearchActiveCallParams) WithDirection(direction []string) *SearchActiveCallParams {
	o.SetDirection(direction)
	return o
}

// SetDirection adds the direction to the search active call params
func (o *SearchActiveCallParams) SetDirection(direction []string) {
	o.Direction = direction
}

// WithDomainID adds the domainID to the search active call params
func (o *SearchActiveCallParams) WithDomainID(domainID *string) *SearchActiveCallParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the search active call params
func (o *SearchActiveCallParams) SetDomainID(domainID *string) {
	o.DomainID = domainID
}

// WithDurationFrom adds the durationFrom to the search active call params
func (o *SearchActiveCallParams) WithDurationFrom(durationFrom *string) *SearchActiveCallParams {
	o.SetDurationFrom(durationFrom)
	return o
}

// SetDurationFrom adds the durationFrom to the search active call params
func (o *SearchActiveCallParams) SetDurationFrom(durationFrom *string) {
	o.DurationFrom = durationFrom
}

// WithDurationTo adds the durationTo to the search active call params
func (o *SearchActiveCallParams) WithDurationTo(durationTo *string) *SearchActiveCallParams {
	o.SetDurationTo(durationTo)
	return o
}

// SetDurationTo adds the durationTo to the search active call params
func (o *SearchActiveCallParams) SetDurationTo(durationTo *string) {
	o.DurationTo = durationTo
}

// WithExistsFile adds the existsFile to the search active call params
func (o *SearchActiveCallParams) WithExistsFile(existsFile *bool) *SearchActiveCallParams {
	o.SetExistsFile(existsFile)
	return o
}

// SetExistsFile adds the existsFile to the search active call params
func (o *SearchActiveCallParams) SetExistsFile(existsFile *bool) {
	o.ExistsFile = existsFile
}

// WithFields adds the fields to the search active call params
func (o *SearchActiveCallParams) WithFields(fields []string) *SearchActiveCallParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the search active call params
func (o *SearchActiveCallParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithGatewayID adds the gatewayID to the search active call params
func (o *SearchActiveCallParams) WithGatewayID(gatewayID []string) *SearchActiveCallParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the search active call params
func (o *SearchActiveCallParams) SetGatewayID(gatewayID []string) {
	o.GatewayID = gatewayID
}

// WithMemberID adds the memberID to the search active call params
func (o *SearchActiveCallParams) WithMemberID(memberID []string) *SearchActiveCallParams {
	o.SetMemberID(memberID)
	return o
}

// SetMemberID adds the memberId to the search active call params
func (o *SearchActiveCallParams) SetMemberID(memberID []string) {
	o.MemberID = memberID
}

// WithMissed adds the missed to the search active call params
func (o *SearchActiveCallParams) WithMissed(missed *bool) *SearchActiveCallParams {
	o.SetMissed(missed)
	return o
}

// SetMissed adds the missed to the search active call params
func (o *SearchActiveCallParams) SetMissed(missed *bool) {
	o.Missed = missed
}

// WithNumber adds the number to the search active call params
func (o *SearchActiveCallParams) WithNumber(number *string) *SearchActiveCallParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the search active call params
func (o *SearchActiveCallParams) SetNumber(number *string) {
	o.Number = number
}

// WithPage adds the page to the search active call params
func (o *SearchActiveCallParams) WithPage(page *int32) *SearchActiveCallParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the search active call params
func (o *SearchActiveCallParams) SetPage(page *int32) {
	o.Page = page
}

// WithParentID adds the parentID to the search active call params
func (o *SearchActiveCallParams) WithParentID(parentID *string) *SearchActiveCallParams {
	o.SetParentID(parentID)
	return o
}

// SetParentID adds the parentId to the search active call params
func (o *SearchActiveCallParams) SetParentID(parentID *string) {
	o.ParentID = parentID
}

// WithQ adds the q to the search active call params
func (o *SearchActiveCallParams) WithQ(q *string) *SearchActiveCallParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the search active call params
func (o *SearchActiveCallParams) SetQ(q *string) {
	o.Q = q
}

// WithQueueID adds the queueID to the search active call params
func (o *SearchActiveCallParams) WithQueueID(queueID []string) *SearchActiveCallParams {
	o.SetQueueID(queueID)
	return o
}

// SetQueueID adds the queueId to the search active call params
func (o *SearchActiveCallParams) SetQueueID(queueID []string) {
	o.QueueID = queueID
}

// WithSize adds the size to the search active call params
func (o *SearchActiveCallParams) WithSize(size *int32) *SearchActiveCallParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the search active call params
func (o *SearchActiveCallParams) SetSize(size *int32) {
	o.Size = size
}

// WithSkipParent adds the skipParent to the search active call params
func (o *SearchActiveCallParams) WithSkipParent(skipParent *bool) *SearchActiveCallParams {
	o.SetSkipParent(skipParent)
	return o
}

// SetSkipParent adds the skipParent to the search active call params
func (o *SearchActiveCallParams) SetSkipParent(skipParent *bool) {
	o.SkipParent = skipParent
}

// WithSort adds the sort to the search active call params
func (o *SearchActiveCallParams) WithSort(sort *string) *SearchActiveCallParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the search active call params
func (o *SearchActiveCallParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithStoredAtFrom adds the storedAtFrom to the search active call params
func (o *SearchActiveCallParams) WithStoredAtFrom(storedAtFrom *string) *SearchActiveCallParams {
	o.SetStoredAtFrom(storedAtFrom)
	return o
}

// SetStoredAtFrom adds the storedAtFrom to the search active call params
func (o *SearchActiveCallParams) SetStoredAtFrom(storedAtFrom *string) {
	o.StoredAtFrom = storedAtFrom
}

// WithStoredAtTo adds the storedAtTo to the search active call params
func (o *SearchActiveCallParams) WithStoredAtTo(storedAtTo *string) *SearchActiveCallParams {
	o.SetStoredAtTo(storedAtTo)
	return o
}

// SetStoredAtTo adds the storedAtTo to the search active call params
func (o *SearchActiveCallParams) SetStoredAtTo(storedAtTo *string) {
	o.StoredAtTo = storedAtTo
}

// WithSupervisorID adds the supervisorID to the search active call params
func (o *SearchActiveCallParams) WithSupervisorID(supervisorID []string) *SearchActiveCallParams {
	o.SetSupervisorID(supervisorID)
	return o
}

// SetSupervisorID adds the supervisorId to the search active call params
func (o *SearchActiveCallParams) SetSupervisorID(supervisorID []string) {
	o.SupervisorID = supervisorID
}

// WithTeamID adds the teamID to the search active call params
func (o *SearchActiveCallParams) WithTeamID(teamID []string) *SearchActiveCallParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the search active call params
func (o *SearchActiveCallParams) SetTeamID(teamID []string) {
	o.TeamID = teamID
}

// WithUserID adds the userID to the search active call params
func (o *SearchActiveCallParams) WithUserID(userID []string) *SearchActiveCallParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the search active call params
func (o *SearchActiveCallParams) SetUserID(userID []string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *SearchActiveCallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.AnsweredAtFrom != nil {

		// query param answered_at.from
		var qrAnsweredAtFrom string

		if o.AnsweredAtFrom != nil {
			qrAnsweredAtFrom = *o.AnsweredAtFrom
		}
		qAnsweredAtFrom := qrAnsweredAtFrom
		if qAnsweredAtFrom != "" {

			if err := r.SetQueryParam("answered_at.from", qAnsweredAtFrom); err != nil {
				return err
			}
		}
	}

	if o.AnsweredAtTo != nil {

		// query param answered_at.to
		var qrAnsweredAtTo string

		if o.AnsweredAtTo != nil {
			qrAnsweredAtTo = *o.AnsweredAtTo
		}
		qAnsweredAtTo := qrAnsweredAtTo
		if qAnsweredAtTo != "" {

			if err := r.SetQueryParam("answered_at.to", qAnsweredAtTo); err != nil {
				return err
			}
		}
	}

	if o.Cause != nil {

		// binding items for cause
		joinedCause := o.bindParamCause(reg)

		// query array param cause
		if err := r.SetQueryParam("cause", joinedCause...); err != nil {
			return err
		}
	}

	if o.CreatedAtFrom != nil {

		// query param created_at.from
		var qrCreatedAtFrom string

		if o.CreatedAtFrom != nil {
			qrCreatedAtFrom = *o.CreatedAtFrom
		}
		qCreatedAtFrom := qrCreatedAtFrom
		if qCreatedAtFrom != "" {

			if err := r.SetQueryParam("created_at.from", qCreatedAtFrom); err != nil {
				return err
			}
		}
	}

	if o.CreatedAtTo != nil {

		// query param created_at.to
		var qrCreatedAtTo string

		if o.CreatedAtTo != nil {
			qrCreatedAtTo = *o.CreatedAtTo
		}
		qCreatedAtTo := qrCreatedAtTo
		if qCreatedAtTo != "" {

			if err := r.SetQueryParam("created_at.to", qCreatedAtTo); err != nil {
				return err
			}
		}
	}

	if o.Direction != nil {

		// binding items for direction
		joinedDirection := o.bindParamDirection(reg)

		// query array param direction
		if err := r.SetQueryParam("direction", joinedDirection...); err != nil {
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

	if o.DurationFrom != nil {

		// query param duration.from
		var qrDurationFrom string

		if o.DurationFrom != nil {
			qrDurationFrom = *o.DurationFrom
		}
		qDurationFrom := qrDurationFrom
		if qDurationFrom != "" {

			if err := r.SetQueryParam("duration.from", qDurationFrom); err != nil {
				return err
			}
		}
	}

	if o.DurationTo != nil {

		// query param duration.to
		var qrDurationTo string

		if o.DurationTo != nil {
			qrDurationTo = *o.DurationTo
		}
		qDurationTo := qrDurationTo
		if qDurationTo != "" {

			if err := r.SetQueryParam("duration.to", qDurationTo); err != nil {
				return err
			}
		}
	}

	if o.ExistsFile != nil {

		// query param exists_file
		var qrExistsFile bool

		if o.ExistsFile != nil {
			qrExistsFile = *o.ExistsFile
		}
		qExistsFile := swag.FormatBool(qrExistsFile)
		if qExistsFile != "" {

			if err := r.SetQueryParam("exists_file", qExistsFile); err != nil {
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

	if o.GatewayID != nil {

		// binding items for gateway_id
		joinedGatewayID := o.bindParamGatewayID(reg)

		// query array param gateway_id
		if err := r.SetQueryParam("gateway_id", joinedGatewayID...); err != nil {
			return err
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

	if o.Missed != nil {

		// query param missed
		var qrMissed bool

		if o.Missed != nil {
			qrMissed = *o.Missed
		}
		qMissed := swag.FormatBool(qrMissed)
		if qMissed != "" {

			if err := r.SetQueryParam("missed", qMissed); err != nil {
				return err
			}
		}
	}

	if o.Number != nil {

		// query param number
		var qrNumber string

		if o.Number != nil {
			qrNumber = *o.Number
		}
		qNumber := qrNumber
		if qNumber != "" {

			if err := r.SetQueryParam("number", qNumber); err != nil {
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

	if o.ParentID != nil {

		// query param parent_id
		var qrParentID string

		if o.ParentID != nil {
			qrParentID = *o.ParentID
		}
		qParentID := qrParentID
		if qParentID != "" {

			if err := r.SetQueryParam("parent_id", qParentID); err != nil {
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

	if o.QueueID != nil {

		// binding items for queue_id
		joinedQueueID := o.bindParamQueueID(reg)

		// query array param queue_id
		if err := r.SetQueryParam("queue_id", joinedQueueID...); err != nil {
			return err
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

	if o.SkipParent != nil {

		// query param skip_parent
		var qrSkipParent bool

		if o.SkipParent != nil {
			qrSkipParent = *o.SkipParent
		}
		qSkipParent := swag.FormatBool(qrSkipParent)
		if qSkipParent != "" {

			if err := r.SetQueryParam("skip_parent", qSkipParent); err != nil {
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

	if o.StoredAtFrom != nil {

		// query param stored_at.from
		var qrStoredAtFrom string

		if o.StoredAtFrom != nil {
			qrStoredAtFrom = *o.StoredAtFrom
		}
		qStoredAtFrom := qrStoredAtFrom
		if qStoredAtFrom != "" {

			if err := r.SetQueryParam("stored_at.from", qStoredAtFrom); err != nil {
				return err
			}
		}
	}

	if o.StoredAtTo != nil {

		// query param stored_at.to
		var qrStoredAtTo string

		if o.StoredAtTo != nil {
			qrStoredAtTo = *o.StoredAtTo
		}
		qStoredAtTo := qrStoredAtTo
		if qStoredAtTo != "" {

			if err := r.SetQueryParam("stored_at.to", qStoredAtTo); err != nil {
				return err
			}
		}
	}

	if o.SupervisorID != nil {

		// binding items for supervisor_id
		joinedSupervisorID := o.bindParamSupervisorID(reg)

		// query array param supervisor_id
		if err := r.SetQueryParam("supervisor_id", joinedSupervisorID...); err != nil {
			return err
		}
	}

	if o.TeamID != nil {

		// binding items for team_id
		joinedTeamID := o.bindParamTeamID(reg)

		// query array param team_id
		if err := r.SetQueryParam("team_id", joinedTeamID...); err != nil {
			return err
		}
	}

	if o.UserID != nil {

		// binding items for user_id
		joinedUserID := o.bindParamUserID(reg)

		// query array param user_id
		if err := r.SetQueryParam("user_id", joinedUserID...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSearchActiveCall binds the parameter agent_id
func (o *SearchActiveCallParams) bindParamAgentID(formats strfmt.Registry) []string {
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

// bindParamSearchActiveCall binds the parameter cause
func (o *SearchActiveCallParams) bindParamCause(formats strfmt.Registry) []string {
	causeIR := o.Cause

	var causeIC []string
	for _, causeIIR := range causeIR { // explode []string

		causeIIV := causeIIR // string as string
		causeIC = append(causeIC, causeIIV)
	}

	// items.CollectionFormat: "multi"
	causeIS := swag.JoinByFormat(causeIC, "multi")

	return causeIS
}

// bindParamSearchActiveCall binds the parameter direction
func (o *SearchActiveCallParams) bindParamDirection(formats strfmt.Registry) []string {
	directionIR := o.Direction

	var directionIC []string
	for _, directionIIR := range directionIR { // explode []string

		directionIIV := directionIIR // string as string
		directionIC = append(directionIC, directionIIV)
	}

	// items.CollectionFormat: "multi"
	directionIS := swag.JoinByFormat(directionIC, "multi")

	return directionIS
}

// bindParamSearchActiveCall binds the parameter fields
func (o *SearchActiveCallParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamSearchActiveCall binds the parameter gateway_id
func (o *SearchActiveCallParams) bindParamGatewayID(formats strfmt.Registry) []string {
	gatewayIDIR := o.GatewayID

	var gatewayIDIC []string
	for _, gatewayIDIIR := range gatewayIDIR { // explode []string

		gatewayIDIIV := gatewayIDIIR // string as string
		gatewayIDIC = append(gatewayIDIC, gatewayIDIIV)
	}

	// items.CollectionFormat: "multi"
	gatewayIDIS := swag.JoinByFormat(gatewayIDIC, "multi")

	return gatewayIDIS
}

// bindParamSearchActiveCall binds the parameter member_id
func (o *SearchActiveCallParams) bindParamMemberID(formats strfmt.Registry) []string {
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

// bindParamSearchActiveCall binds the parameter queue_id
func (o *SearchActiveCallParams) bindParamQueueID(formats strfmt.Registry) []string {
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

// bindParamSearchActiveCall binds the parameter supervisor_id
func (o *SearchActiveCallParams) bindParamSupervisorID(formats strfmt.Registry) []string {
	supervisorIDIR := o.SupervisorID

	var supervisorIDIC []string
	for _, supervisorIDIIR := range supervisorIDIR { // explode []string

		supervisorIDIIV := supervisorIDIIR // string as string
		supervisorIDIC = append(supervisorIDIC, supervisorIDIIV)
	}

	// items.CollectionFormat: "multi"
	supervisorIDIS := swag.JoinByFormat(supervisorIDIC, "multi")

	return supervisorIDIS
}

// bindParamSearchActiveCall binds the parameter team_id
func (o *SearchActiveCallParams) bindParamTeamID(formats strfmt.Registry) []string {
	teamIDIR := o.TeamID

	var teamIDIC []string
	for _, teamIDIIR := range teamIDIR { // explode []string

		teamIDIIV := teamIDIIR // string as string
		teamIDIC = append(teamIDIC, teamIDIIV)
	}

	// items.CollectionFormat: "multi"
	teamIDIS := swag.JoinByFormat(teamIDIC, "multi")

	return teamIDIS
}

// bindParamSearchActiveCall binds the parameter user_id
func (o *SearchActiveCallParams) bindParamUserID(formats strfmt.Registry) []string {
	userIDIR := o.UserID

	var userIDIC []string
	for _, userIDIIR := range userIDIR { // explode []string

		userIDIIV := userIDIIR // string as string
		userIDIC = append(userIDIC, userIDIIV)
	}

	// items.CollectionFormat: "multi"
	userIDIS := swag.JoinByFormat(userIDIC, "multi")

	return userIDIS
}
