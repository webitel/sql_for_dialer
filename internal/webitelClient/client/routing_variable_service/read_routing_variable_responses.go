// Code generated by go-swagger; DO NOT EDIT.

package routing_variable_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/webitel/sql_for_dialer/internal/webitelClient/models"
)

// ReadRoutingVariableReader is a Reader for the ReadRoutingVariable structure.
type ReadRoutingVariableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadRoutingVariableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadRoutingVariableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewReadRoutingVariableForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReadRoutingVariableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadRoutingVariableOK creates a ReadRoutingVariableOK with default headers values
func NewReadRoutingVariableOK() *ReadRoutingVariableOK {
	return &ReadRoutingVariableOK{}
}

/* ReadRoutingVariableOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadRoutingVariableOK struct {
	Payload *models.EngineRoutingVariable
}

func (o *ReadRoutingVariableOK) Error() string {
	return fmt.Sprintf("[GET /routing/variables/{id}][%d] readRoutingVariableOK  %+v", 200, o.Payload)
}
func (o *ReadRoutingVariableOK) GetPayload() *models.EngineRoutingVariable {
	return o.Payload
}

func (o *ReadRoutingVariableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineRoutingVariable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadRoutingVariableForbidden creates a ReadRoutingVariableForbidden with default headers values
func NewReadRoutingVariableForbidden() *ReadRoutingVariableForbidden {
	return &ReadRoutingVariableForbidden{}
}

/* ReadRoutingVariableForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type ReadRoutingVariableForbidden struct {
	Payload interface{}
}

func (o *ReadRoutingVariableForbidden) Error() string {
	return fmt.Sprintf("[GET /routing/variables/{id}][%d] readRoutingVariableForbidden  %+v", 403, o.Payload)
}
func (o *ReadRoutingVariableForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ReadRoutingVariableForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadRoutingVariableNotFound creates a ReadRoutingVariableNotFound with default headers values
func NewReadRoutingVariableNotFound() *ReadRoutingVariableNotFound {
	return &ReadRoutingVariableNotFound{}
}

/* ReadRoutingVariableNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type ReadRoutingVariableNotFound struct {
	Payload string
}

func (o *ReadRoutingVariableNotFound) Error() string {
	return fmt.Sprintf("[GET /routing/variables/{id}][%d] readRoutingVariableNotFound  %+v", 404, o.Payload)
}
func (o *ReadRoutingVariableNotFound) GetPayload() string {
	return o.Payload
}

func (o *ReadRoutingVariableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
