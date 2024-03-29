// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EngineUpdateSkillRequest engine update skill request
//
// swagger:model engineUpdateSkillRequest
type EngineUpdateSkillRequest struct {

	// description
	Description string `json:"description,omitempty"`

	// domain id
	DomainID string `json:"domain_id,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this engine update skill request
func (m *EngineUpdateSkillRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this engine update skill request based on context it is used
func (m *EngineUpdateSkillRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EngineUpdateSkillRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineUpdateSkillRequest) UnmarshalBinary(b []byte) error {
	var res EngineUpdateSkillRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
