// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EngineAgentChannel engine agent channel
//
// swagger:model engineAgentChannel
type EngineAgentChannel struct {

	// channel
	Channel string `json:"channel,omitempty"`

	// joined at
	JoinedAt string `json:"joined_at,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// timeout
	Timeout string `json:"timeout,omitempty"`
}

// Validate validates this engine agent channel
func (m *EngineAgentChannel) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this engine agent channel based on context it is used
func (m *EngineAgentChannel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EngineAgentChannel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineAgentChannel) UnmarshalBinary(b []byte) error {
	var res EngineAgentChannel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}