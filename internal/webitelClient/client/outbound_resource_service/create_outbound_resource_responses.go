// Code generated by go-swagger; DO NOT EDIT.

package outbound_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/webitel/sql_for_dialer/internal/webitelClient/models"
)

// CreateOutboundResourceReader is a Reader for the CreateOutboundResource structure.
type CreateOutboundResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOutboundResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateOutboundResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateOutboundResourceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateOutboundResourceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateOutboundResourceOK creates a CreateOutboundResourceOK with default headers values
func NewCreateOutboundResourceOK() *CreateOutboundResourceOK {
	return &CreateOutboundResourceOK{}
}

/* CreateOutboundResourceOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateOutboundResourceOK struct {
	Payload *models.EngineOutboundResource
}

func (o *CreateOutboundResourceOK) Error() string {
	return fmt.Sprintf("[POST /call_center/resources][%d] createOutboundResourceOK  %+v", 200, o.Payload)
}
func (o *CreateOutboundResourceOK) GetPayload() *models.EngineOutboundResource {
	return o.Payload
}

func (o *CreateOutboundResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineOutboundResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOutboundResourceForbidden creates a CreateOutboundResourceForbidden with default headers values
func NewCreateOutboundResourceForbidden() *CreateOutboundResourceForbidden {
	return &CreateOutboundResourceForbidden{}
}

/* CreateOutboundResourceForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type CreateOutboundResourceForbidden struct {
	Payload interface{}
}

func (o *CreateOutboundResourceForbidden) Error() string {
	return fmt.Sprintf("[POST /call_center/resources][%d] createOutboundResourceForbidden  %+v", 403, o.Payload)
}
func (o *CreateOutboundResourceForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateOutboundResourceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOutboundResourceNotFound creates a CreateOutboundResourceNotFound with default headers values
func NewCreateOutboundResourceNotFound() *CreateOutboundResourceNotFound {
	return &CreateOutboundResourceNotFound{}
}

/* CreateOutboundResourceNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type CreateOutboundResourceNotFound struct {
	Payload string
}

func (o *CreateOutboundResourceNotFound) Error() string {
	return fmt.Sprintf("[POST /call_center/resources][%d] createOutboundResourceNotFound  %+v", 404, o.Payload)
}
func (o *CreateOutboundResourceNotFound) GetPayload() string {
	return o.Payload
}

func (o *CreateOutboundResourceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}