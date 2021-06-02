// Code generated by go-swagger; DO NOT EDIT.

package agent_skill_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/webitel/sql_for_dialer/internal/webitelClient/models"
)

// UpdateAgentSkillReader is a Reader for the UpdateAgentSkill structure.
type UpdateAgentSkillReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAgentSkillReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAgentSkillOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateAgentSkillForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateAgentSkillNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateAgentSkillOK creates a UpdateAgentSkillOK with default headers values
func NewUpdateAgentSkillOK() *UpdateAgentSkillOK {
	return &UpdateAgentSkillOK{}
}

/* UpdateAgentSkillOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateAgentSkillOK struct {
	Payload *models.EngineAgentSkill
}

func (o *UpdateAgentSkillOK) Error() string {
	return fmt.Sprintf("[PUT /call_center/agents/{agent_id}/skills/{id}][%d] updateAgentSkillOK  %+v", 200, o.Payload)
}
func (o *UpdateAgentSkillOK) GetPayload() *models.EngineAgentSkill {
	return o.Payload
}

func (o *UpdateAgentSkillOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAgentSkill)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAgentSkillForbidden creates a UpdateAgentSkillForbidden with default headers values
func NewUpdateAgentSkillForbidden() *UpdateAgentSkillForbidden {
	return &UpdateAgentSkillForbidden{}
}

/* UpdateAgentSkillForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type UpdateAgentSkillForbidden struct {
	Payload interface{}
}

func (o *UpdateAgentSkillForbidden) Error() string {
	return fmt.Sprintf("[PUT /call_center/agents/{agent_id}/skills/{id}][%d] updateAgentSkillForbidden  %+v", 403, o.Payload)
}
func (o *UpdateAgentSkillForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateAgentSkillForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAgentSkillNotFound creates a UpdateAgentSkillNotFound with default headers values
func NewUpdateAgentSkillNotFound() *UpdateAgentSkillNotFound {
	return &UpdateAgentSkillNotFound{}
}

/* UpdateAgentSkillNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type UpdateAgentSkillNotFound struct {
	Payload string
}

func (o *UpdateAgentSkillNotFound) Error() string {
	return fmt.Sprintf("[PUT /call_center/agents/{agent_id}/skills/{id}][%d] updateAgentSkillNotFound  %+v", 404, o.Payload)
}
func (o *UpdateAgentSkillNotFound) GetPayload() string {
	return o.Payload
}

func (o *UpdateAgentSkillNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
