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

// EngineCreateQueueRequest engine create queue request
//
// swagger:model engineCreateQueueRequest
type EngineCreateQueueRequest struct {

	// after schema
	AfterSchema *EngineLookup `json:"after_schema,omitempty"`

	// calendar
	Calendar *EngineLookup `json:"calendar,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// dnc list
	DncList *EngineLookup `json:"dnc_list,omitempty"`

	// do schema
	DoSchema *EngineLookup `json:"do_schema,omitempty"`

	// domain id
	DomainID string `json:"domain_id,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// payload
	Payload interface{} `json:"payload,omitempty"`

	// priority
	Priority int32 `json:"priority,omitempty"`

	// processing
	Processing bool `json:"processing,omitempty"`

	// processing renewal sec
	ProcessingRenewalSec int64 `json:"processing_renewal_sec,omitempty"`

	// processing sec
	ProcessingSec int64 `json:"processing_sec,omitempty"`

	// ringtone
	Ringtone *EngineLookup `json:"ringtone,omitempty"`

	// schema
	Schema *EngineLookup `json:"schema,omitempty"`

	// sec locate agent
	SecLocateAgent int32 `json:"sec_locate_agent,omitempty"`

	// sticky agent
	StickyAgent bool `json:"sticky_agent,omitempty"`

	// strategy
	Strategy string `json:"strategy,omitempty"`

	// team
	Team *EngineLookup `json:"team,omitempty"`

	// timeout
	Timeout int32 `json:"timeout,omitempty"`

	// type
	Type int32 `json:"type,omitempty"`

	// variables
	Variables map[string]string `json:"variables,omitempty"`
}

// Validate validates this engine create queue request
func (m *EngineCreateQueueRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAfterSchema(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCalendar(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDncList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDoSchema(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRingtone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchema(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeam(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineCreateQueueRequest) validateAfterSchema(formats strfmt.Registry) error {
	if swag.IsZero(m.AfterSchema) { // not required
		return nil
	}

	if m.AfterSchema != nil {
		if err := m.AfterSchema.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("after_schema")
			}
			return err
		}
	}

	return nil
}

func (m *EngineCreateQueueRequest) validateCalendar(formats strfmt.Registry) error {
	if swag.IsZero(m.Calendar) { // not required
		return nil
	}

	if m.Calendar != nil {
		if err := m.Calendar.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("calendar")
			}
			return err
		}
	}

	return nil
}

func (m *EngineCreateQueueRequest) validateDncList(formats strfmt.Registry) error {
	if swag.IsZero(m.DncList) { // not required
		return nil
	}

	if m.DncList != nil {
		if err := m.DncList.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dnc_list")
			}
			return err
		}
	}

	return nil
}

func (m *EngineCreateQueueRequest) validateDoSchema(formats strfmt.Registry) error {
	if swag.IsZero(m.DoSchema) { // not required
		return nil
	}

	if m.DoSchema != nil {
		if err := m.DoSchema.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("do_schema")
			}
			return err
		}
	}

	return nil
}

func (m *EngineCreateQueueRequest) validateRingtone(formats strfmt.Registry) error {
	if swag.IsZero(m.Ringtone) { // not required
		return nil
	}

	if m.Ringtone != nil {
		if err := m.Ringtone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ringtone")
			}
			return err
		}
	}

	return nil
}

func (m *EngineCreateQueueRequest) validateSchema(formats strfmt.Registry) error {
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

func (m *EngineCreateQueueRequest) validateTeam(formats strfmt.Registry) error {
	if swag.IsZero(m.Team) { // not required
		return nil
	}

	if m.Team != nil {
		if err := m.Team.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("team")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this engine create queue request based on the context it is used
func (m *EngineCreateQueueRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAfterSchema(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCalendar(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDncList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDoSchema(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRingtone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSchema(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTeam(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineCreateQueueRequest) contextValidateAfterSchema(ctx context.Context, formats strfmt.Registry) error {

	if m.AfterSchema != nil {
		if err := m.AfterSchema.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("after_schema")
			}
			return err
		}
	}

	return nil
}

func (m *EngineCreateQueueRequest) contextValidateCalendar(ctx context.Context, formats strfmt.Registry) error {

	if m.Calendar != nil {
		if err := m.Calendar.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("calendar")
			}
			return err
		}
	}

	return nil
}

func (m *EngineCreateQueueRequest) contextValidateDncList(ctx context.Context, formats strfmt.Registry) error {

	if m.DncList != nil {
		if err := m.DncList.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dnc_list")
			}
			return err
		}
	}

	return nil
}

func (m *EngineCreateQueueRequest) contextValidateDoSchema(ctx context.Context, formats strfmt.Registry) error {

	if m.DoSchema != nil {
		if err := m.DoSchema.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("do_schema")
			}
			return err
		}
	}

	return nil
}

func (m *EngineCreateQueueRequest) contextValidateRingtone(ctx context.Context, formats strfmt.Registry) error {

	if m.Ringtone != nil {
		if err := m.Ringtone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ringtone")
			}
			return err
		}
	}

	return nil
}

func (m *EngineCreateQueueRequest) contextValidateSchema(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EngineCreateQueueRequest) contextValidateTeam(ctx context.Context, formats strfmt.Registry) error {

	if m.Team != nil {
		if err := m.Team.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("team")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EngineCreateQueueRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineCreateQueueRequest) UnmarshalBinary(b []byte) error {
	var res EngineCreateQueueRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
