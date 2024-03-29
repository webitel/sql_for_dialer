// Code generated by go-swagger; DO NOT EDIT.

package queue_skill_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/webitel/sql_for_dialer/internal/webitelClient/models"
)

// PatchQueueSkillReader is a Reader for the PatchQueueSkill structure.
type PatchQueueSkillReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchQueueSkillReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchQueueSkillOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPatchQueueSkillForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchQueueSkillNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchQueueSkillOK creates a PatchQueueSkillOK with default headers values
func NewPatchQueueSkillOK() *PatchQueueSkillOK {
	return &PatchQueueSkillOK{}
}

/* PatchQueueSkillOK describes a response with status code 200, with default header values.

A successful response.
*/
type PatchQueueSkillOK struct {
	Payload *models.EngineQueueSkill
}

func (o *PatchQueueSkillOK) Error() string {
	return fmt.Sprintf("[PATCH /call_center/queues/{queue_id}/skills/{id}][%d] patchQueueSkillOK  %+v", 200, o.Payload)
}
func (o *PatchQueueSkillOK) GetPayload() *models.EngineQueueSkill {
	return o.Payload
}

func (o *PatchQueueSkillOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineQueueSkill)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchQueueSkillForbidden creates a PatchQueueSkillForbidden with default headers values
func NewPatchQueueSkillForbidden() *PatchQueueSkillForbidden {
	return &PatchQueueSkillForbidden{}
}

/* PatchQueueSkillForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type PatchQueueSkillForbidden struct {
	Payload interface{}
}

func (o *PatchQueueSkillForbidden) Error() string {
	return fmt.Sprintf("[PATCH /call_center/queues/{queue_id}/skills/{id}][%d] patchQueueSkillForbidden  %+v", 403, o.Payload)
}
func (o *PatchQueueSkillForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchQueueSkillForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchQueueSkillNotFound creates a PatchQueueSkillNotFound with default headers values
func NewPatchQueueSkillNotFound() *PatchQueueSkillNotFound {
	return &PatchQueueSkillNotFound{}
}

/* PatchQueueSkillNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type PatchQueueSkillNotFound struct {
	Payload string
}

func (o *PatchQueueSkillNotFound) Error() string {
	return fmt.Sprintf("[PATCH /call_center/queues/{queue_id}/skills/{id}][%d] patchQueueSkillNotFound  %+v", 404, o.Payload)
}
func (o *PatchQueueSkillNotFound) GetPayload() string {
	return o.Payload
}

func (o *PatchQueueSkillNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
