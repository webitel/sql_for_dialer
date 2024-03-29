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

// DeleteAgentTeamReader is a Reader for the DeleteAgentTeam structure.
type DeleteAgentTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAgentTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAgentTeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteAgentTeamForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAgentTeamNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAgentTeamOK creates a DeleteAgentTeamOK with default headers values
func NewDeleteAgentTeamOK() *DeleteAgentTeamOK {
	return &DeleteAgentTeamOK{}
}

/* DeleteAgentTeamOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteAgentTeamOK struct {
	Payload *models.EngineAgentTeam
}

func (o *DeleteAgentTeamOK) Error() string {
	return fmt.Sprintf("[DELETE /call_center/teams/{id}][%d] deleteAgentTeamOK  %+v", 200, o.Payload)
}
func (o *DeleteAgentTeamOK) GetPayload() *models.EngineAgentTeam {
	return o.Payload
}

func (o *DeleteAgentTeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAgentTeam)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAgentTeamForbidden creates a DeleteAgentTeamForbidden with default headers values
func NewDeleteAgentTeamForbidden() *DeleteAgentTeamForbidden {
	return &DeleteAgentTeamForbidden{}
}

/* DeleteAgentTeamForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type DeleteAgentTeamForbidden struct {
	Payload interface{}
}

func (o *DeleteAgentTeamForbidden) Error() string {
	return fmt.Sprintf("[DELETE /call_center/teams/{id}][%d] deleteAgentTeamForbidden  %+v", 403, o.Payload)
}
func (o *DeleteAgentTeamForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteAgentTeamForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAgentTeamNotFound creates a DeleteAgentTeamNotFound with default headers values
func NewDeleteAgentTeamNotFound() *DeleteAgentTeamNotFound {
	return &DeleteAgentTeamNotFound{}
}

/* DeleteAgentTeamNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type DeleteAgentTeamNotFound struct {
	Payload string
}

func (o *DeleteAgentTeamNotFound) Error() string {
	return fmt.Sprintf("[DELETE /call_center/teams/{id}][%d] deleteAgentTeamNotFound  %+v", 404, o.Payload)
}
func (o *DeleteAgentTeamNotFound) GetPayload() string {
	return o.Payload
}

func (o *DeleteAgentTeamNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
