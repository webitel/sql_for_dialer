// Code generated by go-swagger; DO NOT EDIT.

package agent_service

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

// NewAgentStateHistoryParams creates a new AgentStateHistoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAgentStateHistoryParams() *AgentStateHistoryParams {
	return &AgentStateHistoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAgentStateHistoryParamsWithTimeout creates a new AgentStateHistoryParams object
// with the ability to set a timeout on a request.
func NewAgentStateHistoryParamsWithTimeout(timeout time.Duration) *AgentStateHistoryParams {
	return &AgentStateHistoryParams{
		timeout: timeout,
	}
}

// NewAgentStateHistoryParamsWithContext creates a new AgentStateHistoryParams object
// with the ability to set a context for a request.
func NewAgentStateHistoryParamsWithContext(ctx context.Context) *AgentStateHistoryParams {
	return &AgentStateHistoryParams{
		Context: ctx,
	}
}

// NewAgentStateHistoryParamsWithHTTPClient creates a new AgentStateHistoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewAgentStateHistoryParamsWithHTTPClient(client *http.Client) *AgentStateHistoryParams {
	return &AgentStateHistoryParams{
		HTTPClient: client,
	}
}

/* AgentStateHistoryParams contains all the parameters to send to the API endpoint
   for the agent state history operation.

   Typically these are written to a http.Request.
*/
type AgentStateHistoryParams struct {

	// AgentID.
	//
	// Format: int64
	AgentID string

	// DomainID.
	//
	// Format: int64
	DomainID *string

	// Page.
	//
	// Format: int32
	Page *int32

	// Q.
	Q *string

	// Size.
	//
	// Format: int32
	Size *int32

	// TimeFrom.
	//
	// Format: int64
	TimeFrom *string

	// TimeTo.
	//
	// Format: int64
	TimeTo *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the agent state history params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AgentStateHistoryParams) WithDefaults() *AgentStateHistoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the agent state history params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AgentStateHistoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the agent state history params
func (o *AgentStateHistoryParams) WithTimeout(timeout time.Duration) *AgentStateHistoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the agent state history params
func (o *AgentStateHistoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the agent state history params
func (o *AgentStateHistoryParams) WithContext(ctx context.Context) *AgentStateHistoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the agent state history params
func (o *AgentStateHistoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the agent state history params
func (o *AgentStateHistoryParams) WithHTTPClient(client *http.Client) *AgentStateHistoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the agent state history params
func (o *AgentStateHistoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentID adds the agentID to the agent state history params
func (o *AgentStateHistoryParams) WithAgentID(agentID string) *AgentStateHistoryParams {
	o.SetAgentID(agentID)
	return o
}

// SetAgentID adds the agentId to the agent state history params
func (o *AgentStateHistoryParams) SetAgentID(agentID string) {
	o.AgentID = agentID
}

// WithDomainID adds the domainID to the agent state history params
func (o *AgentStateHistoryParams) WithDomainID(domainID *string) *AgentStateHistoryParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the agent state history params
func (o *AgentStateHistoryParams) SetDomainID(domainID *string) {
	o.DomainID = domainID
}

// WithPage adds the page to the agent state history params
func (o *AgentStateHistoryParams) WithPage(page *int32) *AgentStateHistoryParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the agent state history params
func (o *AgentStateHistoryParams) SetPage(page *int32) {
	o.Page = page
}

// WithQ adds the q to the agent state history params
func (o *AgentStateHistoryParams) WithQ(q *string) *AgentStateHistoryParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the agent state history params
func (o *AgentStateHistoryParams) SetQ(q *string) {
	o.Q = q
}

// WithSize adds the size to the agent state history params
func (o *AgentStateHistoryParams) WithSize(size *int32) *AgentStateHistoryParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the agent state history params
func (o *AgentStateHistoryParams) SetSize(size *int32) {
	o.Size = size
}

// WithTimeFrom adds the timeFrom to the agent state history params
func (o *AgentStateHistoryParams) WithTimeFrom(timeFrom *string) *AgentStateHistoryParams {
	o.SetTimeFrom(timeFrom)
	return o
}

// SetTimeFrom adds the timeFrom to the agent state history params
func (o *AgentStateHistoryParams) SetTimeFrom(timeFrom *string) {
	o.TimeFrom = timeFrom
}

// WithTimeTo adds the timeTo to the agent state history params
func (o *AgentStateHistoryParams) WithTimeTo(timeTo *string) *AgentStateHistoryParams {
	o.SetTimeTo(timeTo)
	return o
}

// SetTimeTo adds the timeTo to the agent state history params
func (o *AgentStateHistoryParams) SetTimeTo(timeTo *string) {
	o.TimeTo = timeTo
}

// WriteToRequest writes these params to a swagger request
func (o *AgentStateHistoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param agent_id
	if err := r.SetPathParam("agent_id", o.AgentID); err != nil {
		return err
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

	if o.TimeFrom != nil {

		// query param time_from
		var qrTimeFrom string

		if o.TimeFrom != nil {
			qrTimeFrom = *o.TimeFrom
		}
		qTimeFrom := qrTimeFrom
		if qTimeFrom != "" {

			if err := r.SetQueryParam("time_from", qTimeFrom); err != nil {
				return err
			}
		}
	}

	if o.TimeTo != nil {

		// query param time_to
		var qrTimeTo string

		if o.TimeTo != nil {
			qrTimeTo = *o.TimeTo
		}
		qTimeTo := qrTimeTo
		if qTimeTo != "" {

			if err := r.SetQueryParam("time_to", qTimeTo); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}