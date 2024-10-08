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

// DeleteCommunicationTypeReader is a Reader for the DeleteCommunicationType structure.
type DeleteCommunicationTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCommunicationTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCommunicationTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteCommunicationTypeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteCommunicationTypeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteCommunicationTypeOK creates a DeleteCommunicationTypeOK with default headers values
func NewDeleteCommunicationTypeOK() *DeleteCommunicationTypeOK {
	return &DeleteCommunicationTypeOK{}
}

/* DeleteCommunicationTypeOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteCommunicationTypeOK struct {
	Payload *models.EngineCommunicationType
}

func (o *DeleteCommunicationTypeOK) Error() string {
	return fmt.Sprintf("[DELETE /call_center/communication_type/{id}][%d] deleteCommunicationTypeOK  %+v", 200, o.Payload)
}
func (o *DeleteCommunicationTypeOK) GetPayload() *models.EngineCommunicationType {
	return o.Payload
}

func (o *DeleteCommunicationTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineCommunicationType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCommunicationTypeForbidden creates a DeleteCommunicationTypeForbidden with default headers values
func NewDeleteCommunicationTypeForbidden() *DeleteCommunicationTypeForbidden {
	return &DeleteCommunicationTypeForbidden{}
}

/* DeleteCommunicationTypeForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type DeleteCommunicationTypeForbidden struct {
	Payload interface{}
}

func (o *DeleteCommunicationTypeForbidden) Error() string {
	return fmt.Sprintf("[DELETE /call_center/communication_type/{id}][%d] deleteCommunicationTypeForbidden  %+v", 403, o.Payload)
}
func (o *DeleteCommunicationTypeForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteCommunicationTypeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCommunicationTypeNotFound creates a DeleteCommunicationTypeNotFound with default headers values
func NewDeleteCommunicationTypeNotFound() *DeleteCommunicationTypeNotFound {
	return &DeleteCommunicationTypeNotFound{}
}

/* DeleteCommunicationTypeNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type DeleteCommunicationTypeNotFound struct {
	Payload string
}

func (o *DeleteCommunicationTypeNotFound) Error() string {
	return fmt.Sprintf("[DELETE /call_center/communication_type/{id}][%d] deleteCommunicationTypeNotFound  %+v", 404, o.Payload)
}
func (o *DeleteCommunicationTypeNotFound) GetPayload() string {
	return o.Payload
}

func (o *DeleteCommunicationTypeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
