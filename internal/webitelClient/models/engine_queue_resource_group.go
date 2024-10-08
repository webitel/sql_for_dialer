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

// EngineQueueResourceGroup engine queue resource group
//
// swagger:model engineQueueResourceGroup
type EngineQueueResourceGroup struct {

	// id
	ID string `json:"id,omitempty"`

	// resource group
	ResourceGroup *EngineLookup `json:"resource_group,omitempty"`
}

// Validate validates this engine queue resource group
func (m *EngineQueueResourceGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourceGroup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineQueueResourceGroup) validateResourceGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceGroup) { // not required
		return nil
	}

	if m.ResourceGroup != nil {
		if err := m.ResourceGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource_group")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this engine queue resource group based on the context it is used
func (m *EngineQueueResourceGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResourceGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineQueueResourceGroup) contextValidateResourceGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.ResourceGroup != nil {
		if err := m.ResourceGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource_group")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EngineQueueResourceGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineQueueResourceGroup) UnmarshalBinary(b []byte) error {
	var res EngineQueueResourceGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
