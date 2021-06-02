// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EngineCallFile engine call file
//
// swagger:model engineCallFile
type EngineCallFile struct {

	// id
	ID string `json:"id,omitempty"`

	// mime type
	MimeType string `json:"mime_type,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// size
	Size string `json:"size,omitempty"`
}

// Validate validates this engine call file
func (m *EngineCallFile) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this engine call file based on context it is used
func (m *EngineCallFile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EngineCallFile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineCallFile) UnmarshalBinary(b []byte) error {
	var res EngineCallFile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}