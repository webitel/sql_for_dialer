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

// UpdateQueueSkillReader is a Reader for the UpdateQueueSkill structure.
type UpdateQueueSkillReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateQueueSkillReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateQueueSkillOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateQueueSkillForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateQueueSkillNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateQueueSkillOK creates a UpdateQueueSkillOK with default headers values
func NewUpdateQueueSkillOK() *UpdateQueueSkillOK {
	return &UpdateQueueSkillOK{}
}

/* UpdateQueueSkillOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateQueueSkillOK struct {
	Payload *models.EngineQueueSkill
}

func (o *UpdateQueueSkillOK) Error() string {
	return fmt.Sprintf("[PUT /call_center/queues/{queue_id}/skills/{id}][%d] updateQueueSkillOK  %+v", 200, o.Payload)
}
func (o *UpdateQueueSkillOK) GetPayload() *models.EngineQueueSkill {
	return o.Payload
}

func (o *UpdateQueueSkillOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineQueueSkill)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateQueueSkillForbidden creates a UpdateQueueSkillForbidden with default headers values
func NewUpdateQueueSkillForbidden() *UpdateQueueSkillForbidden {
	return &UpdateQueueSkillForbidden{}
}

/* UpdateQueueSkillForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type UpdateQueueSkillForbidden struct {
	Payload interface{}
}

func (o *UpdateQueueSkillForbidden) Error() string {
	return fmt.Sprintf("[PUT /call_center/queues/{queue_id}/skills/{id}][%d] updateQueueSkillForbidden  %+v", 403, o.Payload)
}
func (o *UpdateQueueSkillForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateQueueSkillForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateQueueSkillNotFound creates a UpdateQueueSkillNotFound with default headers values
func NewUpdateQueueSkillNotFound() *UpdateQueueSkillNotFound {
	return &UpdateQueueSkillNotFound{}
}

/* UpdateQueueSkillNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type UpdateQueueSkillNotFound struct {
	Payload string
}

func (o *UpdateQueueSkillNotFound) Error() string {
	return fmt.Sprintf("[PUT /call_center/queues/{queue_id}/skills/{id}][%d] updateQueueSkillNotFound  %+v", 404, o.Payload)
}
func (o *UpdateQueueSkillNotFound) GetPayload() string {
	return o.Payload
}

func (o *UpdateQueueSkillNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
