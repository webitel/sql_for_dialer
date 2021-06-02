// Code generated by go-swagger; DO NOT EDIT.

package agent_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/webitel/sql_for_dialer/internal/webitelClient/models"
)

// SearchAgentReader is a Reader for the SearchAgent structure.
type SearchAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewSearchAgentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchAgentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchAgentOK creates a SearchAgentOK with default headers values
func NewSearchAgentOK() *SearchAgentOK {
	return &SearchAgentOK{}
}

/* SearchAgentOK describes a response with status code 200, with default header values.

A successful response.
*/
type SearchAgentOK struct {
	Payload *models.EngineListAgent
}

func (o *SearchAgentOK) Error() string {
	return fmt.Sprintf("[GET /call_center/agents][%d] searchAgentOK  %+v", 200, o.Payload)
}
func (o *SearchAgentOK) GetPayload() *models.EngineListAgent {
	return o.Payload
}

func (o *SearchAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineListAgent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchAgentForbidden creates a SearchAgentForbidden with default headers values
func NewSearchAgentForbidden() *SearchAgentForbidden {
	return &SearchAgentForbidden{}
}

/* SearchAgentForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type SearchAgentForbidden struct {
	Payload interface{}
}

func (o *SearchAgentForbidden) Error() string {
	return fmt.Sprintf("[GET /call_center/agents][%d] searchAgentForbidden  %+v", 403, o.Payload)
}
func (o *SearchAgentForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *SearchAgentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchAgentNotFound creates a SearchAgentNotFound with default headers values
func NewSearchAgentNotFound() *SearchAgentNotFound {
	return &SearchAgentNotFound{}
}

/* SearchAgentNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type SearchAgentNotFound struct {
	Payload string
}

func (o *SearchAgentNotFound) Error() string {
	return fmt.Sprintf("[GET /call_center/agents][%d] searchAgentNotFound  %+v", 404, o.Payload)
}
func (o *SearchAgentNotFound) GetPayload() string {
	return o.Payload
}

func (o *SearchAgentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
