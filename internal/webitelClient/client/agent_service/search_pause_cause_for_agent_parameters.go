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

// NewSearchPauseCauseForAgentParams creates a new SearchPauseCauseForAgentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchPauseCauseForAgentParams() *SearchPauseCauseForAgentParams {
	return &SearchPauseCauseForAgentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchPauseCauseForAgentParamsWithTimeout creates a new SearchPauseCauseForAgentParams object
// with the ability to set a timeout on a request.
func NewSearchPauseCauseForAgentParamsWithTimeout(timeout time.Duration) *SearchPauseCauseForAgentParams {
	return &SearchPauseCauseForAgentParams{
		timeout: timeout,
	}
}

// NewSearchPauseCauseForAgentParamsWithContext creates a new SearchPauseCauseForAgentParams object
// with the ability to set a context for a request.
func NewSearchPauseCauseForAgentParamsWithContext(ctx context.Context) *SearchPauseCauseForAgentParams {
	return &SearchPauseCauseForAgentParams{
		Context: ctx,
	}
}

// NewSearchPauseCauseForAgentParamsWithHTTPClient creates a new SearchPauseCauseForAgentParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchPauseCauseForAgentParamsWithHTTPClient(client *http.Client) *SearchPauseCauseForAgentParams {
	return &SearchPauseCauseForAgentParams{
		HTTPClient: client,
	}
}

/* SearchPauseCauseForAgentParams contains all the parameters to send to the API endpoint
   for the search pause cause for agent operation.

   Typically these are written to a http.Request.
*/
type SearchPauseCauseForAgentParams struct {

	// AgentID.
	//
	// Format: int64
	AgentID string

	// AllowChange.
	//
	// Format: boolean
	AllowChange *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search pause cause for agent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchPauseCauseForAgentParams) WithDefaults() *SearchPauseCauseForAgentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search pause cause for agent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchPauseCauseForAgentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search pause cause for agent params
func (o *SearchPauseCauseForAgentParams) WithTimeout(timeout time.Duration) *SearchPauseCauseForAgentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search pause cause for agent params
func (o *SearchPauseCauseForAgentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search pause cause for agent params
func (o *SearchPauseCauseForAgentParams) WithContext(ctx context.Context) *SearchPauseCauseForAgentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search pause cause for agent params
func (o *SearchPauseCauseForAgentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search pause cause for agent params
func (o *SearchPauseCauseForAgentParams) WithHTTPClient(client *http.Client) *SearchPauseCauseForAgentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search pause cause for agent params
func (o *SearchPauseCauseForAgentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentID adds the agentID to the search pause cause for agent params
func (o *SearchPauseCauseForAgentParams) WithAgentID(agentID string) *SearchPauseCauseForAgentParams {
	o.SetAgentID(agentID)
	return o
}

// SetAgentID adds the agentId to the search pause cause for agent params
func (o *SearchPauseCauseForAgentParams) SetAgentID(agentID string) {
	o.AgentID = agentID
}

// WithAllowChange adds the allowChange to the search pause cause for agent params
func (o *SearchPauseCauseForAgentParams) WithAllowChange(allowChange *bool) *SearchPauseCauseForAgentParams {
	o.SetAllowChange(allowChange)
	return o
}

// SetAllowChange adds the allowChange to the search pause cause for agent params
func (o *SearchPauseCauseForAgentParams) SetAllowChange(allowChange *bool) {
	o.AllowChange = allowChange
}

// WriteToRequest writes these params to a swagger request
func (o *SearchPauseCauseForAgentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param agent_id
	if err := r.SetPathParam("agent_id", o.AgentID); err != nil {
		return err
	}

	if o.AllowChange != nil {

		// query param allow_change
		var qrAllowChange bool

		if o.AllowChange != nil {
			qrAllowChange = *o.AllowChange
		}
		qAllowChange := swag.FormatBool(qrAllowChange)
		if qAllowChange != "" {

			if err := r.SetQueryParam("allow_change", qAllowChange); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}