// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EngineAgentStatusRequest engine agent status request
//
// swagger:model engineAgentStatusRequest
type EngineAgentStatusRequest struct {

	// channels
	Channels []string `json:"channels"`

	// domain id
	DomainID string `json:"domain_id,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// on demand
	OnDemand bool `json:"on_demand,omitempty"`

	// payload
	Payload string `json:"payload,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this engine agent status request
func (m *EngineAgentStatusRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this engine agent status request based on context it is used
func (m *EngineAgentStatusRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EngineAgentStatusRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineAgentStatusRequest) UnmarshalBinary(b []byte) error {
	var res EngineAgentStatusRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}