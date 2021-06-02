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

// SearchAgentPauseCauseReader is a Reader for the SearchAgentPauseCause structure.
type SearchAgentPauseCauseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchAgentPauseCauseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchAgentPauseCauseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewSearchAgentPauseCauseForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchAgentPauseCauseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchAgentPauseCauseOK creates a SearchAgentPauseCauseOK with default headers values
func NewSearchAgentPauseCauseOK() *SearchAgentPauseCauseOK {
	return &SearchAgentPauseCauseOK{}
}

/* SearchAgentPauseCauseOK describes a response with status code 200, with default header values.

A successful response.
*/
type SearchAgentPauseCauseOK struct {
	Payload *models.EngineListAgentPauseCause
}

func (o *SearchAgentPauseCauseOK) Error() string {
	return fmt.Sprintf("[GET /call_center/pause_causes][%d] searchAgentPauseCauseOK  %+v", 200, o.Payload)
}
func (o *SearchAgentPauseCauseOK) GetPayload() *models.EngineListAgentPauseCause {
	return o.Payload
}

func (o *SearchAgentPauseCauseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineListAgentPauseCause)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchAgentPauseCauseForbidden creates a SearchAgentPauseCauseForbidden with default headers values
func NewSearchAgentPauseCauseForbidden() *SearchAgentPauseCauseForbidden {
	return &SearchAgentPauseCauseForbidden{}
}

/* SearchAgentPauseCauseForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type SearchAgentPauseCauseForbidden struct {
	Payload interface{}
}

func (o *SearchAgentPauseCauseForbidden) Error() string {
	return fmt.Sprintf("[GET /call_center/pause_causes][%d] searchAgentPauseCauseForbidden  %+v", 403, o.Payload)
}
func (o *SearchAgentPauseCauseForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *SearchAgentPauseCauseForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchAgentPauseCauseNotFound creates a SearchAgentPauseCauseNotFound with default headers values
func NewSearchAgentPauseCauseNotFound() *SearchAgentPauseCauseNotFound {
	return &SearchAgentPauseCauseNotFound{}
}

/* SearchAgentPauseCauseNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type SearchAgentPauseCauseNotFound struct {
	Payload string
}

func (o *SearchAgentPauseCauseNotFound) Error() string {
	return fmt.Sprintf("[GET /call_center/pause_causes][%d] searchAgentPauseCauseNotFound  %+v", 404, o.Payload)
}
func (o *SearchAgentPauseCauseNotFound) GetPayload() string {
	return o.Payload
}

func (o *SearchAgentPauseCauseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}