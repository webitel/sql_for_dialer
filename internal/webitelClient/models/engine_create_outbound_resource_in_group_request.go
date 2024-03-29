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

// EngineCreateOutboundResourceInGroupRequest engine create outbound resource in group request
//
// swagger:model engineCreateOutboundResourceInGroupRequest
type EngineCreateOutboundResourceInGroupRequest struct {

	// domain id
	DomainID string `json:"domain_id,omitempty"`

	// group id
	GroupID string `json:"group_id,omitempty"`

	// resource
	Resource *EngineLookup `json:"resource,omitempty"`
}

// Validate validates this engine create outbound resource in group request
func (m *EngineCreateOutboundResourceInGroupRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineCreateOutboundResourceInGroupRequest) validateResource(formats strfmt.Registry) error {
	if swag.IsZero(m.Resource) { // not required
		return nil
	}

	if m.Resource != nil {
		if err := m.Resource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this engine create outbound resource in group request based on the context it is used
func (m *EngineCreateOutboundResourceInGroupRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineCreateOutboundResourceInGroupRequest) contextValidateResource(ctx context.Context, formats strfmt.Registry) error {

	if m.Resource != nil {
		if err := m.Resource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EngineCreateOutboundResourceInGroupRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineCreateOutboundResourceInGroupRequest) UnmarshalBinary(b []byte) error {
	var res EngineCreateOutboundResourceInGroupRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
