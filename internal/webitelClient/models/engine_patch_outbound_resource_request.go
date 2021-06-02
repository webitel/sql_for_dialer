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

// EnginePatchOutboundResourceRequest engine patch outbound resource request
//
// swagger:model enginePatchOutboundResourceRequest
type EnginePatchOutboundResourceRequest struct {

	// domain id
	DomainID string `json:"domain_id,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// error ids
	ErrorIds []string `json:"error_ids"`

	// fields
	Fields []string `json:"fields"`

	// gateway
	Gateway *EngineLookup `json:"gateway,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// limit
	Limit int32 `json:"limit,omitempty"`

	// max successively errors
	MaxSuccessivelyErrors int32 `json:"max_successively_errors,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// number
	Number string `json:"number,omitempty"`

	// reserve
	Reserve bool `json:"reserve,omitempty"`

	// rps
	Rps int32 `json:"rps,omitempty"`

	// variables
	Variables map[string]string `json:"variables,omitempty"`
}

// Validate validates this engine patch outbound resource request
func (m *EnginePatchOutboundResourceRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGateway(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnginePatchOutboundResourceRequest) validateGateway(formats strfmt.Registry) error {
	if swag.IsZero(m.Gateway) { // not required
		return nil
	}

	if m.Gateway != nil {
		if err := m.Gateway.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gateway")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this engine patch outbound resource request based on the context it is used
func (m *EnginePatchOutboundResourceRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGateway(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnginePatchOutboundResourceRequest) contextValidateGateway(ctx context.Context, formats strfmt.Registry) error {

	if m.Gateway != nil {
		if err := m.Gateway.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gateway")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EnginePatchOutboundResourceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnginePatchOutboundResourceRequest) UnmarshalBinary(b []byte) error {
	var res EnginePatchOutboundResourceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
