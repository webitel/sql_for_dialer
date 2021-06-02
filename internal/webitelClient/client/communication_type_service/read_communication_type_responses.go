// Code generated by go-swagger; DO NOT EDIT.

package communication_type_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/webitel/sql_for_dialer/internal/webitelClient/models"
)

// ReadCommunicationTypeReader is a Reader for the ReadCommunicationType structure.
type ReadCommunicationTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadCommunicationTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadCommunicationTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewReadCommunicationTypeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReadCommunicationTypeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadCommunicationTypeOK creates a ReadCommunicationTypeOK with default headers values
func NewReadCommunicationTypeOK() *ReadCommunicationTypeOK {
	return &ReadCommunicationTypeOK{}
}

/* ReadCommunicationTypeOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadCommunicationTypeOK struct {
	Payload *models.EngineCommunicationType
}

func (o *ReadCommunicationTypeOK) Error() string {
	return fmt.Sprintf("[GET /call_center/communication_type/{id}][%d] readCommunicationTypeOK  %+v", 200, o.Payload)
}
func (o *ReadCommunicationTypeOK) GetPayload() *models.EngineCommunicationType {
	return o.Payload
}

func (o *ReadCommunicationTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineCommunicationType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadCommunicationTypeForbidden creates a ReadCommunicationTypeForbidden with default headers values
func NewReadCommunicationTypeForbidden() *ReadCommunicationTypeForbidden {
	return &ReadCommunicationTypeForbidden{}
}

/* ReadCommunicationTypeForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type ReadCommunicationTypeForbidden struct {
	Payload interface{}
}

func (o *ReadCommunicationTypeForbidden) Error() string {
	return fmt.Sprintf("[GET /call_center/communication_type/{id}][%d] readCommunicationTypeForbidden  %+v", 403, o.Payload)
}
func (o *ReadCommunicationTypeForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ReadCommunicationTypeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadCommunicationTypeNotFound creates a ReadCommunicationTypeNotFound with default headers values
func NewReadCommunicationTypeNotFound() *ReadCommunicationTypeNotFound {
	return &ReadCommunicationTypeNotFound{}
}

/* ReadCommunicationTypeNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type ReadCommunicationTypeNotFound struct {
	Payload string
}

func (o *ReadCommunicationTypeNotFound) Error() string {
	return fmt.Sprintf("[GET /call_center/communication_type/{id}][%d] readCommunicationTypeNotFound  %+v", 404, o.Payload)
}
func (o *ReadCommunicationTypeNotFound) GetPayload() string {
	return o.Payload
}

func (o *ReadCommunicationTypeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
