// Code generated by go-swagger; DO NOT EDIT.

package routing_outbound_call_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/webitel/sql_for_dialer/internal/webitelClient/models"
)

// UpdateRoutingOutboundCallReader is a Reader for the UpdateRoutingOutboundCall structure.
type UpdateRoutingOutboundCallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRoutingOutboundCallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRoutingOutboundCallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateRoutingOutboundCallForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRoutingOutboundCallNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRoutingOutboundCallOK creates a UpdateRoutingOutboundCallOK with default headers values
func NewUpdateRoutingOutboundCallOK() *UpdateRoutingOutboundCallOK {
	return &UpdateRoutingOutboundCallOK{}
}

/* UpdateRoutingOutboundCallOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateRoutingOutboundCallOK struct {
	Payload *models.EngineRoutingOutboundCall
}

func (o *UpdateRoutingOutboundCallOK) Error() string {
	return fmt.Sprintf("[PUT /routing/outbound/calls/{id}][%d] updateRoutingOutboundCallOK  %+v", 200, o.Payload)
}
func (o *UpdateRoutingOutboundCallOK) GetPayload() *models.EngineRoutingOutboundCall {
	return o.Payload
}

func (o *UpdateRoutingOutboundCallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineRoutingOutboundCall)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRoutingOutboundCallForbidden creates a UpdateRoutingOutboundCallForbidden with default headers values
func NewUpdateRoutingOutboundCallForbidden() *UpdateRoutingOutboundCallForbidden {
	return &UpdateRoutingOutboundCallForbidden{}
}

/* UpdateRoutingOutboundCallForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type UpdateRoutingOutboundCallForbidden struct {
	Payload interface{}
}

func (o *UpdateRoutingOutboundCallForbidden) Error() string {
	return fmt.Sprintf("[PUT /routing/outbound/calls/{id}][%d] updateRoutingOutboundCallForbidden  %+v", 403, o.Payload)
}
func (o *UpdateRoutingOutboundCallForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateRoutingOutboundCallForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRoutingOutboundCallNotFound creates a UpdateRoutingOutboundCallNotFound with default headers values
func NewUpdateRoutingOutboundCallNotFound() *UpdateRoutingOutboundCallNotFound {
	return &UpdateRoutingOutboundCallNotFound{}
}

/* UpdateRoutingOutboundCallNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type UpdateRoutingOutboundCallNotFound struct {
	Payload string
}

func (o *UpdateRoutingOutboundCallNotFound) Error() string {
	return fmt.Sprintf("[PUT /routing/outbound/calls/{id}][%d] updateRoutingOutboundCallNotFound  %+v", 404, o.Payload)
}
func (o *UpdateRoutingOutboundCallNotFound) GetPayload() string {
	return o.Payload
}

func (o *UpdateRoutingOutboundCallNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
