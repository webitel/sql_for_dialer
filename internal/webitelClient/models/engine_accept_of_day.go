// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EngineAcceptOfDay engine accept of day
//
// swagger:model engineAcceptOfDay
type EngineAcceptOfDay struct {

	// day
	Day int32 `json:"day,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`

	// end time of day
	EndTimeOfDay int32 `json:"end_time_of_day,omitempty"`

	// start time of day
	StartTimeOfDay int32 `json:"start_time_of_day,omitempty"`
}

// Validate validates this engine accept of day
func (m *EngineAcceptOfDay) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this engine accept of day based on context it is used
func (m *EngineAcceptOfDay) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EngineAcceptOfDay) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineAcceptOfDay) UnmarshalBinary(b []byte) error {
	var res EngineAcceptOfDay
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
