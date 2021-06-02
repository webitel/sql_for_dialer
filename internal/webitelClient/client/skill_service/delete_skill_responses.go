// Code generated by go-swagger; DO NOT EDIT.

package skill_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/webitel/sql_for_dialer/internal/webitelClient/models"
)

// DeleteSkillReader is a Reader for the DeleteSkill structure.
type DeleteSkillReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSkillReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSkillOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteSkillForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteSkillNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteSkillOK creates a DeleteSkillOK with default headers values
func NewDeleteSkillOK() *DeleteSkillOK {
	return &DeleteSkillOK{}
}

/* DeleteSkillOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteSkillOK struct {
	Payload *models.EngineSkill
}

func (o *DeleteSkillOK) Error() string {
	return fmt.Sprintf("[DELETE /call_center/skills/{id}][%d] deleteSkillOK  %+v", 200, o.Payload)
}
func (o *DeleteSkillOK) GetPayload() *models.EngineSkill {
	return o.Payload
}

func (o *DeleteSkillOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineSkill)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSkillForbidden creates a DeleteSkillForbidden with default headers values
func NewDeleteSkillForbidden() *DeleteSkillForbidden {
	return &DeleteSkillForbidden{}
}

/* DeleteSkillForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type DeleteSkillForbidden struct {
	Payload interface{}
}

func (o *DeleteSkillForbidden) Error() string {
	return fmt.Sprintf("[DELETE /call_center/skills/{id}][%d] deleteSkillForbidden  %+v", 403, o.Payload)
}
func (o *DeleteSkillForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteSkillForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSkillNotFound creates a DeleteSkillNotFound with default headers values
func NewDeleteSkillNotFound() *DeleteSkillNotFound {
	return &DeleteSkillNotFound{}
}

/* DeleteSkillNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type DeleteSkillNotFound struct {
	Payload string
}

func (o *DeleteSkillNotFound) Error() string {
	return fmt.Sprintf("[DELETE /call_center/skills/{id}][%d] deleteSkillNotFound  %+v", 404, o.Payload)
}
func (o *DeleteSkillNotFound) GetPayload() string {
	return o.Payload
}

func (o *DeleteSkillNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}