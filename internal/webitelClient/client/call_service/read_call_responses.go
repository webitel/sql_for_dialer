// Code generated by go-swagger; DO NOT EDIT.

package call_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/webitel/sql_for_dialer/internal/webitelClient/models"
)

// ReadCallReader is a Reader for the ReadCall structure.
type ReadCallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadCallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadCallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewReadCallForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReadCallNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadCallOK creates a ReadCallOK with default headers values
func NewReadCallOK() *ReadCallOK {
	return &ReadCallOK{}
}

/* ReadCallOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadCallOK struct {
	Payload *models.EngineActiveCall
}

func (o *ReadCallOK) Error() string {
	return fmt.Sprintf("[GET /calls/active/{id}][%d] readCallOK  %+v", 200, o.Payload)
}
func (o *ReadCallOK) GetPayload() *models.EngineActiveCall {
	return o.Payload
}

func (o *ReadCallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineActiveCall)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadCallForbidden creates a ReadCallForbidden with default headers values
func NewReadCallForbidden() *ReadCallForbidden {
	return &ReadCallForbidden{}
}

/* ReadCallForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type ReadCallForbidden struct {
	Payload interface{}
}

func (o *ReadCallForbidden) Error() string {
	return fmt.Sprintf("[GET /calls/active/{id}][%d] readCallForbidden  %+v", 403, o.Payload)
}
func (o *ReadCallForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ReadCallForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadCallNotFound creates a ReadCallNotFound with default headers values
func NewReadCallNotFound() *ReadCallNotFound {
	return &ReadCallNotFound{}
}

/* ReadCallNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type ReadCallNotFound struct {
	Payload string
}

func (o *ReadCallNotFound) Error() string {
	return fmt.Sprintf("[GET /calls/active/{id}][%d] readCallNotFound  %+v", 404, o.Payload)
}
func (o *ReadCallNotFound) GetPayload() string {
	return o.Payload
}

func (o *ReadCallNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}