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

// EngineUpdateQueueSkillRequest engine update queue skill request
//
// swagger:model engineUpdateQueueSkillRequest
type EngineUpdateQueueSkillRequest struct {

	// buckets
	Buckets []*EngineLookup `json:"buckets"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// lvl
	Lvl int32 `json:"lvl,omitempty"`

	// max capacity
	MaxCapacity int32 `json:"max_capacity,omitempty"`

	// min capacity
	MinCapacity int32 `json:"min_capacity,omitempty"`

	// queue id
	QueueID int64 `json:"queue_id,omitempty"`

	// skill
	Skill *EngineLookup `json:"skill,omitempty"`
}

// Validate validates this engine update queue skill request
func (m *EngineUpdateQueueSkillRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuckets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSkill(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineUpdateQueueSkillRequest) validateBuckets(formats strfmt.Registry) error {
	if swag.IsZero(m.Buckets) { // not required
		return nil
	}

	for i := 0; i < len(m.Buckets); i++ {
		if swag.IsZero(m.Buckets[i]) { // not required
			continue
		}

		if m.Buckets[i] != nil {
			if err := m.Buckets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("buckets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EngineUpdateQueueSkillRequest) validateSkill(formats strfmt.Registry) error {
	if swag.IsZero(m.Skill) { // not required
		return nil
	}

	if m.Skill != nil {
		if err := m.Skill.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("skill")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this engine update queue skill request based on the context it is used
func (m *EngineUpdateQueueSkillRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBuckets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSkill(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineUpdateQueueSkillRequest) contextValidateBuckets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Buckets); i++ {

		if m.Buckets[i] != nil {
			if err := m.Buckets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("buckets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EngineUpdateQueueSkillRequest) contextValidateSkill(ctx context.Context, formats strfmt.Registry) error {

	if m.Skill != nil {
		if err := m.Skill.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("skill")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EngineUpdateQueueSkillRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineUpdateQueueSkillRequest) UnmarshalBinary(b []byte) error {
	var res EngineUpdateQueueSkillRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
