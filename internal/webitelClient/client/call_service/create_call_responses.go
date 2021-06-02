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

// CreateCallReader is a Reader for the CreateCall structure.
type CreateCallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateCallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateCallForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateCallNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateCallOK creates a CreateCallOK with default headers values
func NewCreateCallOK() *CreateCallOK {
	return &CreateCallOK{}
}

/* CreateCallOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateCallOK struct {
	Payload *models.EngineCreateCallResponse
}

func (o *CreateCallOK) Error() string {
	return fmt.Sprintf("[POST /calls][%d] createCallOK  %+v", 200, o.Payload)
}
func (o *CreateCallOK) GetPayload() *models.EngineCreateCallResponse {
	return o.Payload
}

func (o *CreateCallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineCreateCallResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCallForbidden creates a CreateCallForbidden with default headers values
func NewCreateCallForbidden() *CreateCallForbidden {
	return &CreateCallForbidden{}
}

/* CreateCallForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type CreateCallForbidden struct {
	Payload interface{}
}

func (o *CreateCallForbidden) Error() string {
	return fmt.Sprintf("[POST /calls][%d] createCallForbidden  %+v", 403, o.Payload)
}
func (o *CreateCallForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateCallForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCallNotFound creates a CreateCallNotFound with default headers values
func NewCreateCallNotFound() *CreateCallNotFound {
	return &CreateCallNotFound{}
}

/* CreateCallNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type CreateCallNotFound struct {
	Payload string
}

func (o *CreateCallNotFound) Error() string {
	return fmt.Sprintf("[POST /calls][%d] createCallNotFound  %+v", 404, o.Payload)
}
func (o *CreateCallNotFound) GetPayload() string {
	return o.Payload
}

func (o *CreateCallNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}