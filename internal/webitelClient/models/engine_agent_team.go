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

// EngineAgentTeam engine agent team
//
// swagger:model engineAgentTeam
type EngineAgentTeam struct {

	// admin
	Admin *EngineLookup `json:"admin,omitempty"`

	// call timeout
	CallTimeout int32 `json:"call_timeout,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// domain id
	DomainID string `json:"domain_id,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// max no answer
	MaxNoAnswer int32 `json:"max_no_answer,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// no answer delay time
	NoAnswerDelayTime int32 `json:"no_answer_delay_time,omitempty"`

	// strategy
	Strategy string `json:"strategy,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// wrap up time
	WrapUpTime int32 `json:"wrap_up_time,omitempty"`
}

// Validate validates this engine agent team
func (m *EngineAgentTeam) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdmin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineAgentTeam) validateAdmin(formats strfmt.Registry) error {
	if swag.IsZero(m.Admin) { // not required
		return nil
	}

	if m.Admin != nil {
		if err := m.Admin.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("admin")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this engine agent team based on the context it is used
func (m *EngineAgentTeam) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdmin(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineAgentTeam) contextValidateAdmin(ctx context.Context, formats strfmt.Registry) error {

	if m.Admin != nil {
		if err := m.Admin.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("admin")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EngineAgentTeam) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineAgentTeam) UnmarshalBinary(b []byte) error {
	var res EngineAgentTeam
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
