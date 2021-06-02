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

// EngineCreateCallRequest engine create call request
//
// swagger:model engineCreateCallRequest
type EngineCreateCallRequest struct {

	// destination
	Destination string `json:"destination,omitempty"`

	// domain id
	DomainID string `json:"domain_id,omitempty"`

	// from
	From *CreateCallRequestEndpointRequest `json:"from,omitempty"`

	// params
	Params *CreateCallRequestCallSettings `json:"params,omitempty"`

	// to
	To *CreateCallRequestEndpointRequest `json:"to,omitempty"`
}

// Validate validates this engine create call request
func (m *EngineCreateCallRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineCreateCallRequest) validateFrom(formats strfmt.Registry) error {
	if swag.IsZero(m.From) { // not required
		return nil
	}

	if m.From != nil {
		if err := m.From.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("from")
			}
			return err
		}
	}

	return nil
}

func (m *EngineCreateCallRequest) validateParams(formats strfmt.Registry) error {
	if swag.IsZero(m.Params) { // not required
		return nil
	}

	if m.Params != nil {
		if err := m.Params.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("params")
			}
			return err
		}
	}

	return nil
}

func (m *EngineCreateCallRequest) validateTo(formats strfmt.Registry) error {
	if swag.IsZero(m.To) { // not required
		return nil
	}

	if m.To != nil {
		if err := m.To.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("to")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this engine create call request based on the context it is used
func (m *EngineCreateCallRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFrom(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineCreateCallRequest) contextValidateFrom(ctx context.Context, formats strfmt.Registry) error {

	if m.From != nil {
		if err := m.From.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("from")
			}
			return err
		}
	}

	return nil
}

func (m *EngineCreateCallRequest) contextValidateParams(ctx context.Context, formats strfmt.Registry) error {

	if m.Params != nil {
		if err := m.Params.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("params")
			}
			return err
		}
	}

	return nil
}

func (m *EngineCreateCallRequest) contextValidateTo(ctx context.Context, formats strfmt.Registry) error {

	if m.To != nil {
		if err := m.To.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("to")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EngineCreateCallRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineCreateCallRequest) UnmarshalBinary(b []byte) error {
	var res EngineCreateCallRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
