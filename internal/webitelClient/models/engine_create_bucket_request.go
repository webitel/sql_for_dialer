// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EngineCreateBucketRequest engine create bucket request
//
// swagger:model engineCreateBucketRequest
type EngineCreateBucketRequest struct {

	// description
	Description string `json:"description,omitempty"`

	// domain id
	DomainID string `json:"domain_id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this engine create bucket request
func (m *EngineCreateBucketRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this engine create bucket request based on context it is used
func (m *EngineCreateBucketRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EngineCreateBucketRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineCreateBucketRequest) UnmarshalBinary(b []byte) error {
	var res EngineCreateBucketRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
