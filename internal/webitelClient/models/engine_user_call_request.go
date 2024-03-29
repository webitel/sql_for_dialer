// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EngineUserCallRequest engine user call request
//
// swagger:model engineUserCallRequest
type EngineUserCallRequest struct {

	// app id
	AppID string `json:"app_id,omitempty"`

	// domain id
	DomainID string `json:"domain_id,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this engine user call request
func (m *EngineUserCallRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this engine user call request based on context it is used
func (m *EngineUserCallRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EngineUserCallRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineUserCallRequest) UnmarshalBinary(b []byte) error {
	var res EngineUserCallRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
