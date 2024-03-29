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

// SearchCalendarReader is a Reader for the SearchCalendar structure.
type SearchCalendarReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchCalendarReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchCalendarOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewSearchCalendarForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchCalendarNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchCalendarOK creates a SearchCalendarOK with default headers values
func NewSearchCalendarOK() *SearchCalendarOK {
	return &SearchCalendarOK{}
}

/* SearchCalendarOK describes a response with status code 200, with default header values.

A successful response.
*/
type SearchCalendarOK struct {
	Payload *models.EngineListCalendar
}

func (o *SearchCalendarOK) Error() string {
	return fmt.Sprintf("[GET /calendars][%d] searchCalendarOK  %+v", 200, o.Payload)
}
func (o *SearchCalendarOK) GetPayload() *models.EngineListCalendar {
	return o.Payload
}

func (o *SearchCalendarOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineListCalendar)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchCalendarForbidden creates a SearchCalendarForbidden with default headers values
func NewSearchCalendarForbidden() *SearchCalendarForbidden {
	return &SearchCalendarForbidden{}
}

/* SearchCalendarForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type SearchCalendarForbidden struct {
	Payload interface{}
}

func (o *SearchCalendarForbidden) Error() string {
	return fmt.Sprintf("[GET /calendars][%d] searchCalendarForbidden  %+v", 403, o.Payload)
}
func (o *SearchCalendarForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *SearchCalendarForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchCalendarNotFound creates a SearchCalendarNotFound with default headers values
func NewSearchCalendarNotFound() *SearchCalendarNotFound {
	return &SearchCalendarNotFound{}
}

/* SearchCalendarNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type SearchCalendarNotFound struct {
	Payload string
}

func (o *SearchCalendarNotFound) Error() string {
	return fmt.Sprintf("[GET /calendars][%d] searchCalendarNotFound  %+v", 404, o.Payload)
}
func (o *SearchCalendarNotFound) GetPayload() string {
	return o.Payload
}

func (o *SearchCalendarNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
