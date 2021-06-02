// Code generated by go-swagger; DO NOT EDIT.

package list_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/webitel/sql_for_dialer/internal/webitelClient/models"
)

// ReadListCommunicationReader is a Reader for the ReadListCommunication structure.
type ReadListCommunicationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadListCommunicationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadListCommunicationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewReadListCommunicationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReadListCommunicationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadListCommunicationOK creates a ReadListCommunicationOK with default headers values
func NewReadListCommunicationOK() *ReadListCommunicationOK {
	return &ReadListCommunicationOK{}
}

/* ReadListCommunicationOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadListCommunicationOK struct {
	Payload *models.EngineListCommunication
}

func (o *ReadListCommunicationOK) Error() string {
	return fmt.Sprintf("[GET /call_center/list/{list_id}/communication/{id}][%d] readListCommunicationOK  %+v", 200, o.Payload)
}
func (o *ReadListCommunicationOK) GetPayload() *models.EngineListCommunication {
	return o.Payload
}

func (o *ReadListCommunicationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineListCommunication)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadListCommunicationForbidden creates a ReadListCommunicationForbidden with default headers values
func NewReadListCommunicationForbidden() *ReadListCommunicationForbidden {
	return &ReadListCommunicationForbidden{}
}

/* ReadListCommunicationForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type ReadListCommunicationForbidden struct {
	Payload interface{}
}

func (o *ReadListCommunicationForbidden) Error() string {
	return fmt.Sprintf("[GET /call_center/list/{list_id}/communication/{id}][%d] readListCommunicationForbidden  %+v", 403, o.Payload)
}
func (o *ReadListCommunicationForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ReadListCommunicationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadListCommunicationNotFound creates a ReadListCommunicationNotFound with default headers values
func NewReadListCommunicationNotFound() *ReadListCommunicationNotFound {
	return &ReadListCommunicationNotFound{}
}

/* ReadListCommunicationNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type ReadListCommunicationNotFound struct {
	Payload string
}

func (o *ReadListCommunicationNotFound) Error() string {
	return fmt.Sprintf("[GET /call_center/list/{list_id}/communication/{id}][%d] readListCommunicationNotFound  %+v", 404, o.Payload)
}
func (o *ReadListCommunicationNotFound) GetPayload() string {
	return o.Payload
}

func (o *ReadListCommunicationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}