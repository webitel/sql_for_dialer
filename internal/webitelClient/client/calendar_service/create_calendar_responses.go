// Code generated by go-swagger; DO NOT EDIT.

package calendar_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/webitel/sql_for_dialer/internal/webitelClient/models"
)

// CreateCalendarReader is a Reader for the CreateCalendar structure.
type CreateCalendarReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCalendarReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateCalendarOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateCalendarForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateCalendarNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateCalendarOK creates a CreateCalendarOK with default headers values
func NewCreateCalendarOK() *CreateCalendarOK {
	return &CreateCalendarOK{}
}

/* CreateCalendarOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateCalendarOK struct {
	Payload *models.EngineCalendar
}

func (o *CreateCalendarOK) Error() string {
	return fmt.Sprintf("[POST /calendars][%d] createCalendarOK  %+v", 200, o.Payload)
}
func (o *CreateCalendarOK) GetPayload() *models.EngineCalendar {
	return o.Payload
}

func (o *CreateCalendarOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineCalendar)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCalendarForbidden creates a CreateCalendarForbidden with default headers values
func NewCreateCalendarForbidden() *CreateCalendarForbidden {
	return &CreateCalendarForbidden{}
}

/* CreateCalendarForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type CreateCalendarForbidden struct {
	Payload interface{}
}

func (o *CreateCalendarForbidden) Error() string {
	return fmt.Sprintf("[POST /calendars][%d] createCalendarForbidden  %+v", 403, o.Payload)
}
func (o *CreateCalendarForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateCalendarForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCalendarNotFound creates a CreateCalendarNotFound with default headers values
func NewCreateCalendarNotFound() *CreateCalendarNotFound {
	return &CreateCalendarNotFound{}
}

/* CreateCalendarNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type CreateCalendarNotFound struct {
	Payload string
}

func (o *CreateCalendarNotFound) Error() string {
	return fmt.Sprintf("[POST /calendars][%d] createCalendarNotFound  %+v", 404, o.Payload)
}
func (o *CreateCalendarNotFound) GetPayload() string {
	return o.Payload
}

func (o *CreateCalendarNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
