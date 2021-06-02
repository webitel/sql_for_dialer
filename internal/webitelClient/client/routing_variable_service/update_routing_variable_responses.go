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

// UpdateRoutingVariableReader is a Reader for the UpdateRoutingVariable structure.
type UpdateRoutingVariableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRoutingVariableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRoutingVariableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateRoutingVariableForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRoutingVariableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRoutingVariableOK creates a UpdateRoutingVariableOK with default headers values
func NewUpdateRoutingVariableOK() *UpdateRoutingVariableOK {
	return &UpdateRoutingVariableOK{}
}

/* UpdateRoutingVariableOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateRoutingVariableOK struct {
	Payload *models.EngineRoutingVariable
}

func (o *UpdateRoutingVariableOK) Error() string {
	return fmt.Sprintf("[PUT /routing/variables/{id}][%d] updateRoutingVariableOK  %+v", 200, o.Payload)
}
func (o *UpdateRoutingVariableOK) GetPayload() *models.EngineRoutingVariable {
	return o.Payload
}

func (o *UpdateRoutingVariableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineRoutingVariable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRoutingVariableForbidden creates a UpdateRoutingVariableForbidden with default headers values
func NewUpdateRoutingVariableForbidden() *UpdateRoutingVariableForbidden {
	return &UpdateRoutingVariableForbidden{}
}

/* UpdateRoutingVariableForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type UpdateRoutingVariableForbidden struct {
	Payload interface{}
}

func (o *UpdateRoutingVariableForbidden) Error() string {
	return fmt.Sprintf("[PUT /routing/variables/{id}][%d] updateRoutingVariableForbidden  %+v", 403, o.Payload)
}
func (o *UpdateRoutingVariableForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateRoutingVariableForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRoutingVariableNotFound creates a UpdateRoutingVariableNotFound with default headers values
func NewUpdateRoutingVariableNotFound() *UpdateRoutingVariableNotFound {
	return &UpdateRoutingVariableNotFound{}
}

/* UpdateRoutingVariableNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type UpdateRoutingVariableNotFound struct {
	Payload string
}

func (o *UpdateRoutingVariableNotFound) Error() string {
	return fmt.Sprintf("[PUT /routing/variables/{id}][%d] updateRoutingVariableNotFound  %+v", 404, o.Payload)
}
func (o *UpdateRoutingVariableNotFound) GetPayload() string {
	return o.Payload
}

func (o *UpdateRoutingVariableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
