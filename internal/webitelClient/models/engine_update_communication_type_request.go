// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EngineUpdateCommunicationTypeRequest engine update communication type request
//
// swagger:model engineUpdateCommunicationTypeRequest
type EngineUpdateCommunicationTypeRequest struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// domain id
	DomainID string `json:"domain_id,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this engine update communication type request
func (m *EngineUpdateCommunicationTypeRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this engine update communication type request based on context it is used
func (m *EngineUpdateCommunicationTypeRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EngineUpdateCommunicationTypeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineUpdateCommunicationTypeRequest) UnmarshalBinary(b []byte) error {
	var res EngineUpdateCommunicationTypeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}