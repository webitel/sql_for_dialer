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

// SearchTimezonesReader is a Reader for the SearchTimezones structure.
type SearchTimezonesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchTimezonesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchTimezonesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewSearchTimezonesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchTimezonesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchTimezonesOK creates a SearchTimezonesOK with default headers values
func NewSearchTimezonesOK() *SearchTimezonesOK {
	return &SearchTimezonesOK{}
}

/* SearchTimezonesOK describes a response with status code 200, with default header values.

A successful response.
*/
type SearchTimezonesOK struct {
	Payload *models.EngineListTimezoneResponse
}

func (o *SearchTimezonesOK) Error() string {
	return fmt.Sprintf("[GET /calendars/timezones][%d] searchTimezonesOK  %+v", 200, o.Payload)
}
func (o *SearchTimezonesOK) GetPayload() *models.EngineListTimezoneResponse {
	return o.Payload
}

func (o *SearchTimezonesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineListTimezoneResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchTimezonesForbidden creates a SearchTimezonesForbidden with default headers values
func NewSearchTimezonesForbidden() *SearchTimezonesForbidden {
	return &SearchTimezonesForbidden{}
}

/* SearchTimezonesForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type SearchTimezonesForbidden struct {
	Payload interface{}
}

func (o *SearchTimezonesForbidden) Error() string {
	return fmt.Sprintf("[GET /calendars/timezones][%d] searchTimezonesForbidden  %+v", 403, o.Payload)
}
func (o *SearchTimezonesForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *SearchTimezonesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchTimezonesNotFound creates a SearchTimezonesNotFound with default headers values
func NewSearchTimezonesNotFound() *SearchTimezonesNotFound {
	return &SearchTimezonesNotFound{}
}

/* SearchTimezonesNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type SearchTimezonesNotFound struct {
	Payload string
}

func (o *SearchTimezonesNotFound) Error() string {
	return fmt.Sprintf("[GET /calendars/timezones][%d] searchTimezonesNotFound  %+v", 404, o.Payload)
}
func (o *SearchTimezonesNotFound) GetPayload() string {
	return o.Payload
}

func (o *SearchTimezonesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
