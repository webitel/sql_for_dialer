// Code generated by go-swagger; DO NOT EDIT.

package routing_schema_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/webitel/sql_for_dialer/internal/webitelClient/models"
)

// CreateRoutingSchemaReader is a Reader for the CreateRoutingSchema structure.
type CreateRoutingSchemaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRoutingSchemaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateRoutingSchemaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateRoutingSchemaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateRoutingSchemaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRoutingSchemaOK creates a CreateRoutingSchemaOK with default headers values
func NewCreateRoutingSchemaOK() *CreateRoutingSchemaOK {
	return &CreateRoutingSchemaOK{}
}

/* CreateRoutingSchemaOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateRoutingSchemaOK struct {
	Payload *models.EngineRoutingSchema
}

func (o *CreateRoutingSchemaOK) Error() string {
	return fmt.Sprintf("[POST /routing/schema][%d] createRoutingSchemaOK  %+v", 200, o.Payload)
}
func (o *CreateRoutingSchemaOK) GetPayload() *models.EngineRoutingSchema {
	return o.Payload
}

func (o *CreateRoutingSchemaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineRoutingSchema)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRoutingSchemaForbidden creates a CreateRoutingSchemaForbidden with default headers values
func NewCreateRoutingSchemaForbidden() *CreateRoutingSchemaForbidden {
	return &CreateRoutingSchemaForbidden{}
}

/* CreateRoutingSchemaForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type CreateRoutingSchemaForbidden struct {
	Payload interface{}
}

func (o *CreateRoutingSchemaForbidden) Error() string {
	return fmt.Sprintf("[POST /routing/schema][%d] createRoutingSchemaForbidden  %+v", 403, o.Payload)
}
func (o *CreateRoutingSchemaForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateRoutingSchemaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRoutingSchemaNotFound creates a CreateRoutingSchemaNotFound with default headers values
func NewCreateRoutingSchemaNotFound() *CreateRoutingSchemaNotFound {
	return &CreateRoutingSchemaNotFound{}
}

/* CreateRoutingSchemaNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type CreateRoutingSchemaNotFound struct {
	Payload string
}

func (o *CreateRoutingSchemaNotFound) Error() string {
	return fmt.Sprintf("[POST /routing/schema][%d] createRoutingSchemaNotFound  %+v", 404, o.Payload)
}
func (o *CreateRoutingSchemaNotFound) GetPayload() string {
	return o.Payload
}

func (o *CreateRoutingSchemaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}