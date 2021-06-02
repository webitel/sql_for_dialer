// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EngineListReportGeneral engine list report general
//
// swagger:model engineListReportGeneral
type EngineListReportGeneral struct {

	// aggs
	Aggs *EngineQueueReportGeneralAgentStatus `json:"aggs,omitempty"`

	// items
	Items []*EngineQueueReportGeneral `json:"items"`

	// next
	Next bool `json:"next,omitempty"`
}

// Validate validates this engine list report general
func (m *EngineListReportGeneral) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAggs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineListReportGeneral) validateAggs(formats strfmt.Registry) error {
	if swag.IsZero(m.Aggs) { // not required
		return nil
	}

	if m.Aggs != nil {
		if err := m.Aggs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggs")
			}
			return err
		}
	}

	return nil
}

func (m *EngineListReportGeneral) validateItems(formats strfmt.Registry) error {
	if swag.IsZero(m.Items) { // not required
		return nil
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this engine list report general based on the context it is used
func (m *EngineListReportGeneral) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAggs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineListReportGeneral) contextValidateAggs(ctx context.Context, formats strfmt.Registry) error {

	if m.Aggs != nil {
		if err := m.Aggs.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggs")
			}
			return err
		}
	}

	return nil
}

func (m *EngineListReportGeneral) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Items); i++ {

		if m.Items[i] != nil {
			if err := m.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EngineListReportGeneral) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineListReportGeneral) UnmarshalBinary(b []byte) error {
	var res EngineListReportGeneral
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
