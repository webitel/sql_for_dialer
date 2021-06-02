// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EngineCreateRoutingSchemaRequest engine create routing schema request
//
// swagger:model engineCreateRoutingSchemaRequest
type EngineCreateRoutingSchemaRequest struct {

	// debug
	Debug bool `json:"debug,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// domain id
	DomainID string `json:"domain_id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// payload
	Payload interface{} `json:"payload,omitempty"`

	// schema
	Schema interface{} `json:"schema,omitempty"`

	// type
	Type int32 `json:"type,omitempty"`
}

// Validate validates this engine create routing schema request
func (m *EngineCreateRoutingSchemaRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this engine create routing schema request based on context it is used
func (m *EngineCreateRoutingSchemaRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EngineCreateRoutingSchemaRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineCreateRoutingSchemaRequest) UnmarshalBinary(b []byte) error {
	var res EngineCreateRoutingSchemaRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
