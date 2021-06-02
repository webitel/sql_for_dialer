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
)

// NewSearchAgentInQueueStatisticsParams creates a new SearchAgentInQueueStatisticsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchAgentInQueueStatisticsParams() *SearchAgentInQueueStatisticsParams {
	return &SearchAgentInQueueStatisticsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchAgentInQueueStatisticsParamsWithTimeout creates a new SearchAgentInQueueStatisticsParams object
// with the ability to set a timeout on a request.
func NewSearchAgentInQueueStatisticsParamsWithTimeout(timeout time.Duration) *SearchAgentInQueueStatisticsParams {
	return &SearchAgentInQueueStatisticsParams{
		timeout: timeout,
	}
}

// NewSearchAgentInQueueStatisticsParamsWithContext creates a new SearchAgentInQueueStatisticsParams object
// with the ability to set a context for a request.
func NewSearchAgentInQueueStatisticsParamsWithContext(ctx context.Context) *SearchAgentInQueueStatisticsParams {
	return &SearchAgentInQueueStatisticsParams{
		Context: ctx,
	}
}

// NewSearchAgentInQueueStatisticsParamsWithHTTPClient creates a new SearchAgentInQueueStatisticsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchAgentInQueueStatisticsParamsWithHTTPClient(client *http.Client) *SearchAgentInQueueStatisticsParams {
	return &SearchAgentInQueueStatisticsParams{
		HTTPClient: client,
	}
}

/* SearchAgentInQueueStatisticsParams contains all the parameters to send to the API endpoint
   for the search agent in queue statistics operation.

   Typically these are written to a http.Request.
*/
type SearchAgentInQueueStatisticsParams struct {

	// AgentID.
	//
	// Format: int64
	AgentID string

	// DomainID.
	//
	// Format: int64
	DomainID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search agent in queue statistics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchAgentInQueueStatisticsParams) WithDefaults() *SearchAgentInQueueStatisticsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search agent in queue statistics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchAgentInQueueStatisticsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search agent in queue statistics params
func (o *SearchAgentInQueueStatisticsParams) WithTimeout(timeout time.Duration) *SearchAgentInQueueStatisticsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search agent in queue statistics params
func (o *SearchAgentInQueueStatisticsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search agent in queue statistics params
func (o *SearchAgentInQueueStatisticsParams) WithContext(ctx context.Context) *SearchAgentInQueueStatisticsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search agent in queue statistics params
func (o *SearchAgentInQueueStatisticsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search agent in queue statistics params
func (o *SearchAgentInQueueStatisticsParams) WithHTTPClient(client *http.Client) *SearchAgentInQueueStatisticsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search agent in queue statistics params
func (o *SearchAgentInQueueStatisticsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentID adds the agentID to the search agent in queue statistics params
func (o *SearchAgentInQueueStatisticsParams) WithAgentID(agentID string) *SearchAgentInQueueStatisticsParams {
	o.SetAgentID(agentID)
	return o
}

// SetAgentID adds the agentId to the search agent in queue statistics params
func (o *SearchAgentInQueueStatisticsParams) SetAgentID(agentID string) {
	o.AgentID = agentID
}

// WithDomainID adds the domainID to the search agent in queue statistics params
func (o *SearchAgentInQueueStatisticsParams) WithDomainID(domainID *string) *SearchAgentInQueueStatisticsParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the search agent in queue statistics params
func (o *SearchAgentInQueueStatisticsParams) SetDomainID(domainID *string) {
	o.DomainID = domainID
}

// WriteToRequest writes these params to a swagger request
func (o *SearchAgentInQueueStatisticsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}