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

// EngineActiveCall engine active call
//
// swagger:model engineActiveCall
type EngineActiveCall struct {

	// agent
	Agent *EngineLookup `json:"agent,omitempty"`

	// answered at
	AnsweredAt string `json:"answered_at,omitempty"`

	// app id
	AppID string `json:"app_id,omitempty"`

	// bill sec
	BillSec int32 `json:"bill_sec,omitempty"`

	// bridged at
	BridgedAt string `json:"bridged_at,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// destination
	Destination string `json:"destination,omitempty"`

	// direction
	Direction string `json:"direction,omitempty"`

	// display
	Display string `json:"display,omitempty"`

	// duration
	Duration int32 `json:"duration,omitempty"`

	// extension
	Extension string `json:"extension,omitempty"`

	// from
	From *EngineEndpoint `json:"from,omitempty"`

	// gateway
	Gateway *EngineLookup `json:"gateway,omitempty"`

	// hold sec
	HoldSec int32 `json:"hold_sec,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// joined at
	JoinedAt string `json:"joined_at,omitempty"`

	// leaving at
	LeavingAt string `json:"leaving_at,omitempty"`

	// member
	Member *EngineLookup `json:"member,omitempty"`

	// parent id
	ParentID string `json:"parent_id,omitempty"`

	// queue
	Queue *EngineLookup `json:"queue,omitempty"`

	// queue bridged at
	QueueBridgedAt string `json:"queue_bridged_at,omitempty"`

	// queue duration sec
	QueueDurationSec int32 `json:"queue_duration_sec,omitempty"`

	// queue wait sec
	QueueWaitSec int32 `json:"queue_wait_sec,omitempty"`

	// reporting at
	ReportingAt string `json:"reporting_at,omitempty"`

	// reporting sec
	ReportingSec int32 `json:"reporting_sec,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// supervisor
	Supervisor *EngineLookup `json:"supervisor,omitempty"`

	// team
	Team *EngineLookup `json:"team,omitempty"`

	// timestamp
	Timestamp string `json:"timestamp,omitempty"`

	// to
	To *EngineEndpoint `json:"to,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// user
	User *EngineLookup `json:"user,omitempty"`

	// variables
	Variables map[string]string `json:"variables,omitempty"`

	// wait sec
	WaitSec int32 `json:"wait_sec,omitempty"`
}

// Validate validates this engine active call
func (m *EngineActiveCall) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGateway(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMember(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupervisor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeam(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineActiveCall) validateAgent(formats strfmt.Registry) error {
	if swag.IsZero(m.Agent) { // not required
		return nil
	}

	if m.Agent != nil {
		if err := m.Agent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("agent")
			}
			return err
		}
	}

	return nil
}

func (m *EngineActiveCall) validateFrom(formats strfmt.Registry) error {
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

func (m *EngineActiveCall) validateGateway(formats strfmt.Registry) error {
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

func (m *EngineActiveCall) validateMember(formats strfmt.Registry) error {
	if swag.IsZero(m.Member) { // not required
		return nil
	}

	if m.Member != nil {
		if err := m.Member.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("member")
			}
			return err
		}
	}

	return nil
}

func (m *EngineActiveCall) validateQueue(formats strfmt.Registry) error {
	if swag.IsZero(m.Queue) { // not required
		return nil
	}

	if m.Queue != nil {
		if err := m.Queue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queue")
			}
			return err
		}
	}

	return nil
}

func (m *EngineActiveCall) validateSupervisor(formats strfmt.Registry) error {
	if swag.IsZero(m.Supervisor) { // not required
		return nil
	}

	if m.Supervisor != nil {
		if err := m.Supervisor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("supervisor")
			}
			return err
		}
	}

	return nil
}

func (m *EngineActiveCall) validateTeam(formats strfmt.Registry) error {
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

func (m *EngineActiveCall) validateTo(formats strfmt.Registry) error {
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

func (m *EngineActiveCall) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this engine active call based on the context it is used
func (m *EngineActiveCall) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAgent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFrom(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGateway(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMember(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQueue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSupervisor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTeam(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineActiveCall) contextValidateAgent(ctx context.Context, formats strfmt.Registry) error {

	if m.Agent != nil {
		if err := m.Agent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("agent")
			}
			return err
		}
	}

	return nil
}

func (m *EngineActiveCall) contextValidateFrom(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EngineActiveCall) contextValidateGateway(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EngineActiveCall) contextValidateMember(ctx context.Context, formats strfmt.Registry) error {

	if m.Member != nil {
		if err := m.Member.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("member")
			}
			return err
		}
	}

	return nil
}

func (m *EngineActiveCall) contextValidateQueue(ctx context.Context, formats strfmt.Registry) error {

	if m.Queue != nil {
		if err := m.Queue.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queue")
			}
			return err
		}
	}

	return nil
}

func (m *EngineActiveCall) contextValidateSupervisor(ctx context.Context, formats strfmt.Registry) error {

	if m.Supervisor != nil {
		if err := m.Supervisor.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("supervisor")
			}
			return err
		}
	}

	return nil
}

func (m *EngineActiveCall) contextValidateTeam(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EngineActiveCall) contextValidateTo(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EngineActiveCall) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if m.User != nil {
		if err := m.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EngineActiveCall) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineActiveCall) UnmarshalBinary(b []byte) error {
	var res EngineActiveCall
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}