// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EngineCreateRoutingOutboundCallRequest engine create routing outbound call request
//
// swagger:model engineCreateRoutingOutboundCallRequest
type EngineCreateRoutingOutboundCallRequest struct {

	// description
	Description string `json:"description,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`

	// domain id
	DomainID string `json:"domain_id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// pattern
	Pattern string `json:"pattern,omitempty"`

	// schema
	Schema *EngineLookup `json:"schema,omitempty"`
}

// Validate validates this engine create routing outbound call request
func (m *EngineCreateRoutingOutboundCallRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSchema(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineCreateRoutingOutboundCallRequest) validateSchema(formats strfmt.Registry) error {
	if swag.IsZero(m.Schema) { // not required
		return nil
	}

	if m.Schema != nil {
		if err := m.Schema.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schema")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this engine create routing outbound call request based on the context it is used
func (m *EngineCreateRoutingOutboundCallRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSchema(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineCreateRoutingOutboundCallRequest) contextValidateSchema(ctx context.Context, formats strfmt.Registry) error {

	if m.Schema != nil {
		if err := m.Schema.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schema")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EngineCreateRoutingOutboundCallRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineCreateRoutingOutboundCallRequest) UnmarshalBinary(b []byte) error {
	var res EngineCreateRoutingOutboundCallRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
