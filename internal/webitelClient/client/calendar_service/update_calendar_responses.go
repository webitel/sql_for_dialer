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

// UpdateCalendarReader is a Reader for the UpdateCalendar structure.
type UpdateCalendarReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCalendarReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateCalendarOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateCalendarForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateCalendarNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateCalendarOK creates a UpdateCalendarOK with default headers values
func NewUpdateCalendarOK() *UpdateCalendarOK {
	return &UpdateCalendarOK{}
}

/* UpdateCalendarOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateCalendarOK struct {
	Payload *models.EngineCalendar
}

func (o *UpdateCalendarOK) Error() string {
	return fmt.Sprintf("[PUT /calendars/{id}][%d] updateCalendarOK  %+v", 200, o.Payload)
}
func (o *UpdateCalendarOK) GetPayload() *models.EngineCalendar {
	return o.Payload
}

func (o *UpdateCalendarOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineCalendar)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCalendarForbidden creates a UpdateCalendarForbidden with default headers values
func NewUpdateCalendarForbidden() *UpdateCalendarForbidden {
	return &UpdateCalendarForbidden{}
}

/* UpdateCalendarForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type UpdateCalendarForbidden struct {
	Payload interface{}
}

func (o *UpdateCalendarForbidden) Error() string {
	return fmt.Sprintf("[PUT /calendars/{id}][%d] updateCalendarForbidden  %+v", 403, o.Payload)
}
func (o *UpdateCalendarForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateCalendarForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCalendarNotFound creates a UpdateCalendarNotFound with default headers values
func NewUpdateCalendarNotFound() *UpdateCalendarNotFound {
	return &UpdateCalendarNotFound{}
}

/* UpdateCalendarNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type UpdateCalendarNotFound struct {
	Payload string
}

func (o *UpdateCalendarNotFound) Error() string {
	return fmt.Sprintf("[PUT /calendars/{id}][%d] updateCalendarNotFound  %+v", 404, o.Payload)
}
func (o *UpdateCalendarNotFound) GetPayload() string {
	return o.Payload
}

func (o *UpdateCalendarNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}