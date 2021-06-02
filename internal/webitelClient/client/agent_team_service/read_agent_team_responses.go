// Code generated by go-swagger; DO NOT EDIT.

package agent_team_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/webitel/sql_for_dialer/internal/webitelClient/models"
)

// ReadAgentTeamReader is a Reader for the ReadAgentTeam structure.
type ReadAgentTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadAgentTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadAgentTeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewReadAgentTeamForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReadAgentTeamNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadAgentTeamOK creates a ReadAgentTeamOK with default headers values
func NewReadAgentTeamOK() *ReadAgentTeamOK {
	return &ReadAgentTeamOK{}
}

/* ReadAgentTeamOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadAgentTeamOK struct {
	Payload *models.EngineAgentTeam
}

func (o *ReadAgentTeamOK) Error() string {
	return fmt.Sprintf("[GET /call_center/teams/{id}][%d] readAgentTeamOK  %+v", 200, o.Payload)
}
func (o *ReadAgentTeamOK) GetPayload() *models.EngineAgentTeam {
	return o.Payload
}

func (o *ReadAgentTeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAgentTeam)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadAgentTeamForbidden creates a ReadAgentTeamForbidden with default headers values
func NewReadAgentTeamForbidden() *ReadAgentTeamForbidden {
	return &ReadAgentTeamForbidden{}
}

/* ReadAgentTeamForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type ReadAgentTeamForbidden struct {
	Payload interface{}
}

func (o *ReadAgentTeamForbidden) Error() string {
	return fmt.Sprintf("[GET /call_center/teams/{id}][%d] readAgentTeamForbidden  %+v", 403, o.Payload)
}
func (o *ReadAgentTeamForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ReadAgentTeamForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadAgentTeamNotFound creates a ReadAgentTeamNotFound with default headers values
func NewReadAgentTeamNotFound() *ReadAgentTeamNotFound {
	return &ReadAgentTeamNotFound{}
}

/* ReadAgentTeamNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type ReadAgentTeamNotFound struct {
	Payload string
}

func (o *ReadAgentTeamNotFound) Error() string {
	return fmt.Sprintf("[GET /call_center/teams/{id}][%d] readAgentTeamNotFound  %+v", 404, o.Payload)
}
func (o *ReadAgentTeamNotFound) GetPayload() string {
	return o.Payload
}

func (o *ReadAgentTeamNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}