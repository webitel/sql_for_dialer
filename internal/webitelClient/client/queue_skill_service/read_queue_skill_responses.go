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

// ReadQueueSkillReader is a Reader for the ReadQueueSkill structure.
type ReadQueueSkillReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadQueueSkillReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadQueueSkillOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewReadQueueSkillForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReadQueueSkillNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadQueueSkillOK creates a ReadQueueSkillOK with default headers values
func NewReadQueueSkillOK() *ReadQueueSkillOK {
	return &ReadQueueSkillOK{}
}

/* ReadQueueSkillOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadQueueSkillOK struct {
	Payload *models.EngineQueueSkill
}

func (o *ReadQueueSkillOK) Error() string {
	return fmt.Sprintf("[GET /call_center/queues/{queue_id}/skills/{id}][%d] readQueueSkillOK  %+v", 200, o.Payload)
}
func (o *ReadQueueSkillOK) GetPayload() *models.EngineQueueSkill {
	return o.Payload
}

func (o *ReadQueueSkillOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineQueueSkill)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadQueueSkillForbidden creates a ReadQueueSkillForbidden with default headers values
func NewReadQueueSkillForbidden() *ReadQueueSkillForbidden {
	return &ReadQueueSkillForbidden{}
}

/* ReadQueueSkillForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type ReadQueueSkillForbidden struct {
	Payload interface{}
}

func (o *ReadQueueSkillForbidden) Error() string {
	return fmt.Sprintf("[GET /call_center/queues/{queue_id}/skills/{id}][%d] readQueueSkillForbidden  %+v", 403, o.Payload)
}
func (o *ReadQueueSkillForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ReadQueueSkillForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadQueueSkillNotFound creates a ReadQueueSkillNotFound with default headers values
func NewReadQueueSkillNotFound() *ReadQueueSkillNotFound {
	return &ReadQueueSkillNotFound{}
}

/* ReadQueueSkillNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type ReadQueueSkillNotFound struct {
	Payload string
}

func (o *ReadQueueSkillNotFound) Error() string {
	return fmt.Sprintf("[GET /call_center/queues/{queue_id}/skills/{id}][%d] readQueueSkillNotFound  %+v", 404, o.Payload)
}
func (o *ReadQueueSkillNotFound) GetPayload() string {
	return o.Payload
}

func (o *ReadQueueSkillNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
