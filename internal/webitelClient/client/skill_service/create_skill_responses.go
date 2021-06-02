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

// CreateSkillReader is a Reader for the CreateSkill structure.
type CreateSkillReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSkillReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateSkillOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateSkillForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateSkillNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateSkillOK creates a CreateSkillOK with default headers values
func NewCreateSkillOK() *CreateSkillOK {
	return &CreateSkillOK{}
}

/* CreateSkillOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateSkillOK struct {
	Payload *models.EngineSkill
}

func (o *CreateSkillOK) Error() string {
	return fmt.Sprintf("[POST /call_center/skills][%d] createSkillOK  %+v", 200, o.Payload)
}
func (o *CreateSkillOK) GetPayload() *models.EngineSkill {
	return o.Payload
}

func (o *CreateSkillOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineSkill)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSkillForbidden creates a CreateSkillForbidden with default headers values
func NewCreateSkillForbidden() *CreateSkillForbidden {
	return &CreateSkillForbidden{}
}

/* CreateSkillForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type CreateSkillForbidden struct {
	Payload interface{}
}

func (o *CreateSkillForbidden) Error() string {
	return fmt.Sprintf("[POST /call_center/skills][%d] createSkillForbidden  %+v", 403, o.Payload)
}
func (o *CreateSkillForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateSkillForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSkillNotFound creates a CreateSkillNotFound with default headers values
func NewCreateSkillNotFound() *CreateSkillNotFound {
	return &CreateSkillNotFound{}
}

/* CreateSkillNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type CreateSkillNotFound struct {
	Payload string
}

func (o *CreateSkillNotFound) Error() string {
	return fmt.Sprintf("[POST /call_center/skills][%d] createSkillNotFound  %+v", 404, o.Payload)
}
func (o *CreateSkillNotFound) GetPayload() string {
	return o.Payload
}

func (o *CreateSkillNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}