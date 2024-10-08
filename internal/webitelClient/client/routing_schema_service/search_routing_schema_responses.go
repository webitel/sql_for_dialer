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

// SearchRoutingSchemaReader is a Reader for the SearchRoutingSchema structure.
type SearchRoutingSchemaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchRoutingSchemaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchRoutingSchemaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewSearchRoutingSchemaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchRoutingSchemaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchRoutingSchemaOK creates a SearchRoutingSchemaOK with default headers values
func NewSearchRoutingSchemaOK() *SearchRoutingSchemaOK {
	return &SearchRoutingSchemaOK{}
}

/* SearchRoutingSchemaOK describes a response with status code 200, with default header values.

A successful response.
*/
type SearchRoutingSchemaOK struct {
	Payload *models.EngineListRoutingSchema
}

func (o *SearchRoutingSchemaOK) Error() string {
	return fmt.Sprintf("[GET /routing/schema][%d] searchRoutingSchemaOK  %+v", 200, o.Payload)
}
func (o *SearchRoutingSchemaOK) GetPayload() *models.EngineListRoutingSchema {
	return o.Payload
}

func (o *SearchRoutingSchemaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineListRoutingSchema)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchRoutingSchemaForbidden creates a SearchRoutingSchemaForbidden with default headers values
func NewSearchRoutingSchemaForbidden() *SearchRoutingSchemaForbidden {
	return &SearchRoutingSchemaForbidden{}
}

/* SearchRoutingSchemaForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type SearchRoutingSchemaForbidden struct {
	Payload interface{}
}

func (o *SearchRoutingSchemaForbidden) Error() string {
	return fmt.Sprintf("[GET /routing/schema][%d] searchRoutingSchemaForbidden  %+v", 403, o.Payload)
}
func (o *SearchRoutingSchemaForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *SearchRoutingSchemaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchRoutingSchemaNotFound creates a SearchRoutingSchemaNotFound with default headers values
func NewSearchRoutingSchemaNotFound() *SearchRoutingSchemaNotFound {
	return &SearchRoutingSchemaNotFound{}
}

/* SearchRoutingSchemaNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type SearchRoutingSchemaNotFound struct {
	Payload string
}

func (o *SearchRoutingSchemaNotFound) Error() string {
	return fmt.Sprintf("[GET /routing/schema][%d] searchRoutingSchemaNotFound  %+v", 404, o.Payload)
}
func (o *SearchRoutingSchemaNotFound) GetPayload() string {
	return o.Payload
}

func (o *SearchRoutingSchemaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
