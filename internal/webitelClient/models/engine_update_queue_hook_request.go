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

// EngineUpdateQueueHookRequest engine update queue hook request
//
// swagger:model engineUpdateQueueHookRequest
type EngineUpdateQueueHookRequest struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// event
	Event string `json:"event,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// properties
	Properties []string `json:"properties"`

	// queue id
	QueueID int64 `json:"queue_id,omitempty"`

	// schema
	Schema *EngineLookup `json:"schema,omitempty"`
}

// Validate validates this engine update queue hook request
func (m *EngineUpdateQueueHookRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSchema(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineUpdateQueueHookRequest) validateSchema(formats strfmt.Registry) error {
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

// ContextValidate validate this engine update queue hook request based on the context it is used
func (m *EngineUpdateQueueHookRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSchema(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineUpdateQueueHookRequest) contextValidateSchema(ctx context.Context, formats strfmt.Registry) error {

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
func (m *EngineUpdateQueueHookRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineUpdateQueueHookRequest) UnmarshalBinary(b []byte) error {
	var res EngineUpdateQueueHookRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
