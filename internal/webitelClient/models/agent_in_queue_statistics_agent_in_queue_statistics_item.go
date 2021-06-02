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

// AgentInQueueStatisticsAgentInQueueStatisticsItem agent in queue statistics agent in queue statistics item
//
// swagger:model AgentInQueueStatisticsAgentInQueueStatisticsItem
type AgentInQueueStatisticsAgentInQueueStatisticsItem struct {

	// bucket
	Bucket *EngineLookup `json:"bucket,omitempty"`

	// member waiting
	MemberWaiting int32 `json:"member_waiting,omitempty"`

	// skill
	Skill *EngineLookup `json:"skill,omitempty"`
}

// Validate validates this agent in queue statistics agent in queue statistics item
func (m *AgentInQueueStatisticsAgentInQueueStatisticsItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBucket(formats); err != nil {
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

func (m *AgentInQueueStatisticsAgentInQueueStatisticsItem) validateBucket(formats strfmt.Registry) error {
	if swag.IsZero(m.Bucket) { // not required
		return nil
	}

	if m.Bucket != nil {
		if err := m.Bucket.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bucket")
			}
			return err
		}
	}

	return nil
}

func (m *AgentInQueueStatisticsAgentInQueueStatisticsItem) validateSkill(formats strfmt.Registry) error {
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

// ContextValidate validate this agent in queue statistics agent in queue statistics item based on the context it is used
func (m *AgentInQueueStatisticsAgentInQueueStatisticsItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBucket(ctx, formats); err != nil {
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

func (m *AgentInQueueStatisticsAgentInQueueStatisticsItem) contextValidateBucket(ctx context.Context, formats strfmt.Registry) error {

	if m.Bucket != nil {
		if err := m.Bucket.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bucket")
			}
			return err
		}
	}

	return nil
}

func (m *AgentInQueueStatisticsAgentInQueueStatisticsItem) contextValidateSkill(ctx context.Context, formats strfmt.Registry) error {

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
func (m *AgentInQueueStatisticsAgentInQueueStatisticsItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AgentInQueueStatisticsAgentInQueueStatisticsItem) UnmarshalBinary(b []byte) error {
	var res AgentInQueueStatisticsAgentInQueueStatisticsItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}