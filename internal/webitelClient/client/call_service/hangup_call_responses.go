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

// HangupCallReader is a Reader for the HangupCall structure.
type HangupCallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HangupCallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHangupCallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewHangupCallForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewHangupCallNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewHangupCallOK creates a HangupCallOK with default headers values
func NewHangupCallOK() *HangupCallOK {
	return &HangupCallOK{}
}

/* HangupCallOK describes a response with status code 200, with default header values.

A successful response.
*/
type HangupCallOK struct {
	Payload models.EngineHangupCallResponse
}

func (o *HangupCallOK) Error() string {
	return fmt.Sprintf("[DELETE /calls/active/{id}][%d] hangupCallOK  %+v", 200, o.Payload)
}
func (o *HangupCallOK) GetPayload() models.EngineHangupCallResponse {
	return o.Payload
}

func (o *HangupCallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHangupCallForbidden creates a HangupCallForbidden with default headers values
func NewHangupCallForbidden() *HangupCallForbidden {
	return &HangupCallForbidden{}
}

/* HangupCallForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type HangupCallForbidden struct {
	Payload interface{}
}

func (o *HangupCallForbidden) Error() string {
	return fmt.Sprintf("[DELETE /calls/active/{id}][%d] hangupCallForbidden  %+v", 403, o.Payload)
}
func (o *HangupCallForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *HangupCallForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHangupCallNotFound creates a HangupCallNotFound with default headers values
func NewHangupCallNotFound() *HangupCallNotFound {
	return &HangupCallNotFound{}
}

/* HangupCallNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type HangupCallNotFound struct {
	Payload string
}

func (o *HangupCallNotFound) Error() string {
	return fmt.Sprintf("[DELETE /calls/active/{id}][%d] hangupCallNotFound  %+v", 404, o.Payload)
}
func (o *HangupCallNotFound) GetPayload() string {
	return o.Payload
}

func (o *HangupCallNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}