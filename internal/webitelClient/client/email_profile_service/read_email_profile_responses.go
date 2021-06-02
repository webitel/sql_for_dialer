// Code generated by go-swagger; DO NOT EDIT.

package email_profile_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/webitel/sql_for_dialer/internal/webitelClient/models"
)

// ReadEmailProfileReader is a Reader for the ReadEmailProfile structure.
type ReadEmailProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadEmailProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadEmailProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewReadEmailProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReadEmailProfileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadEmailProfileOK creates a ReadEmailProfileOK with default headers values
func NewReadEmailProfileOK() *ReadEmailProfileOK {
	return &ReadEmailProfileOK{}
}

/* ReadEmailProfileOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadEmailProfileOK struct {
	Payload *models.EngineEmailProfile
}

func (o *ReadEmailProfileOK) Error() string {
	return fmt.Sprintf("[GET /email/profile/{id}][%d] readEmailProfileOK  %+v", 200, o.Payload)
}
func (o *ReadEmailProfileOK) GetPayload() *models.EngineEmailProfile {
	return o.Payload
}

func (o *ReadEmailProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineEmailProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadEmailProfileForbidden creates a ReadEmailProfileForbidden with default headers values
func NewReadEmailProfileForbidden() *ReadEmailProfileForbidden {
	return &ReadEmailProfileForbidden{}
}

/* ReadEmailProfileForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type ReadEmailProfileForbidden struct {
	Payload interface{}
}

func (o *ReadEmailProfileForbidden) Error() string {
	return fmt.Sprintf("[GET /email/profile/{id}][%d] readEmailProfileForbidden  %+v", 403, o.Payload)
}
func (o *ReadEmailProfileForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ReadEmailProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadEmailProfileNotFound creates a ReadEmailProfileNotFound with default headers values
func NewReadEmailProfileNotFound() *ReadEmailProfileNotFound {
	return &ReadEmailProfileNotFound{}
}

/* ReadEmailProfileNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type ReadEmailProfileNotFound struct {
	Payload string
}

func (o *ReadEmailProfileNotFound) Error() string {
	return fmt.Sprintf("[GET /email/profile/{id}][%d] readEmailProfileNotFound  %+v", 404, o.Payload)
}
func (o *ReadEmailProfileNotFound) GetPayload() string {
	return o.Payload
}

func (o *ReadEmailProfileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
