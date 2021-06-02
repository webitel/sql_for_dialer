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

// PatchAgentSkillReader is a Reader for the PatchAgentSkill structure.
type PatchAgentSkillReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAgentSkillReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchAgentSkillOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPatchAgentSkillForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchAgentSkillNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchAgentSkillOK creates a PatchAgentSkillOK with default headers values
func NewPatchAgentSkillOK() *PatchAgentSkillOK {
	return &PatchAgentSkillOK{}
}

/* PatchAgentSkillOK describes a response with status code 200, with default header values.

A successful response.
*/
type PatchAgentSkillOK struct {
	Payload *models.EngineAgentSkill
}

func (o *PatchAgentSkillOK) Error() string {
	return fmt.Sprintf("[PATCH /call_center/agents/{agent_id}/skills/{id}][%d] patchAgentSkillOK  %+v", 200, o.Payload)
}
func (o *PatchAgentSkillOK) GetPayload() *models.EngineAgentSkill {
	return o.Payload
}

func (o *PatchAgentSkillOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAgentSkill)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAgentSkillForbidden creates a PatchAgentSkillForbidden with default headers values
func NewPatchAgentSkillForbidden() *PatchAgentSkillForbidden {
	return &PatchAgentSkillForbidden{}
}

/* PatchAgentSkillForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type PatchAgentSkillForbidden struct {
	Payload interface{}
}

func (o *PatchAgentSkillForbidden) Error() string {
	return fmt.Sprintf("[PATCH /call_center/agents/{agent_id}/skills/{id}][%d] patchAgentSkillForbidden  %+v", 403, o.Payload)
}
func (o *PatchAgentSkillForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchAgentSkillForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAgentSkillNotFound creates a PatchAgentSkillNotFound with default headers values
func NewPatchAgentSkillNotFound() *PatchAgentSkillNotFound {
	return &PatchAgentSkillNotFound{}
}

/* PatchAgentSkillNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type PatchAgentSkillNotFound struct {
	Payload string
}

func (o *PatchAgentSkillNotFound) Error() string {
	return fmt.Sprintf("[PATCH /call_center/agents/{agent_id}/skills/{id}][%d] patchAgentSkillNotFound  %+v", 404, o.Payload)
}
func (o *PatchAgentSkillNotFound) GetPayload() string {
	return o.Payload
}

func (o *PatchAgentSkillNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}