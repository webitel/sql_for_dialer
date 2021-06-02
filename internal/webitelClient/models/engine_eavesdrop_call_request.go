// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EngineEavesdropCallRequest engine eavesdrop call request
//
// swagger:model engineEavesdropCallRequest
type EngineEavesdropCallRequest struct {

	// app id
	AppID string `json:"app_id,omitempty"`

	// control
	Control bool `json:"control,omitempty"`

	// domain id
	DomainID string `json:"domain_id,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// listen a
	Listena bool `json:"listen_a,omitempty"`

	// listen b
	Listenb bool `json:"listen_b,omitempty"`

	// whisper a
	Whispera bool `json:"whisper_a,omitempty"`

	// whisper b
	Whisperb bool `json:"whisper_b,omitempty"`
}

// Validate validates this engine eavesdrop call request
func (m *EngineEavesdropCallRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this engine eavesdrop call request based on context it is used
func (m *EngineEavesdropCallRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EngineEavesdropCallRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineEavesdropCallRequest) UnmarshalBinary(b []byte) error {
	var res EngineEavesdropCallRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}