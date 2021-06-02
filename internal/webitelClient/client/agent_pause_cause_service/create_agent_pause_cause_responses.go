// Code generated by go-swagger; DO NOT EDIT.

package agent_pause_cause_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/webitel/sql_for_dialer/internal/webitelClient/models"
)

// CreateAgentPauseCauseReader is a Reader for the CreateAgentPauseCause structure.
type CreateAgentPauseCauseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAgentPauseCauseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateAgentPauseCauseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateAgentPauseCauseForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateAgentPauseCauseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateAgentPauseCauseOK creates a CreateAgentPauseCauseOK with default headers values
func NewCreateAgentPauseCauseOK() *CreateAgentPauseCauseOK {
	return &CreateAgentPauseCauseOK{}
}

/* CreateAgentPauseCauseOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateAgentPauseCauseOK struct {
	Payload *models.EngineAgentPauseCause
}

func (o *CreateAgentPauseCauseOK) Error() string {
	return fmt.Sprintf("[POST /call_center/pause_causes][%d] createAgentPauseCauseOK  %+v", 200, o.Payload)
}
func (o *CreateAgentPauseCauseOK) GetPayload() *models.EngineAgentPauseCause {
	return o.Payload
}

func (o *CreateAgentPauseCauseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAgentPauseCause)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAgentPauseCauseForbidden creates a CreateAgentPauseCauseForbidden with default headers values
func NewCreateAgentPauseCauseForbidden() *CreateAgentPauseCauseForbidden {
	return &CreateAgentPauseCauseForbidden{}
}

/* CreateAgentPauseCauseForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type CreateAgentPauseCauseForbidden struct {
	Payload interface{}
}

func (o *CreateAgentPauseCauseForbidden) Error() string {
	return fmt.Sprintf("[POST /call_center/pause_causes][%d] createAgentPauseCauseForbidden  %+v", 403, o.Payload)
}
func (o *CreateAgentPauseCauseForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateAgentPauseCauseForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAgentPauseCauseNotFound creates a CreateAgentPauseCauseNotFound with default headers values
func NewCreateAgentPauseCauseNotFound() *CreateAgentPauseCauseNotFound {
	return &CreateAgentPauseCauseNotFound{}
}

/* CreateAgentPauseCauseNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type CreateAgentPauseCauseNotFound struct {
	Payload string
}

func (o *CreateAgentPauseCauseNotFound) Error() string {
	return fmt.Sprintf("[POST /call_center/pause_causes][%d] createAgentPauseCauseNotFound  %+v", 404, o.Payload)
}
func (o *CreateAgentPauseCauseNotFound) GetPayload() string {
	return o.Payload
}

func (o *CreateAgentPauseCauseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
