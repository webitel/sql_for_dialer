// Code generated by go-swagger; DO NOT EDIT.

package outbound_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/webitel/sql_for_dialer/internal/webitelClient/models"
)

// UpdateOutboundResourceDisplayReader is a Reader for the UpdateOutboundResourceDisplay structure.
type UpdateOutboundResourceDisplayReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOutboundResourceDisplayReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOutboundResourceDisplayOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateOutboundResourceDisplayForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateOutboundResourceDisplayNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateOutboundResourceDisplayOK creates a UpdateOutboundResourceDisplayOK with default headers values
func NewUpdateOutboundResourceDisplayOK() *UpdateOutboundResourceDisplayOK {
	return &UpdateOutboundResourceDisplayOK{}
}

/* UpdateOutboundResourceDisplayOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateOutboundResourceDisplayOK struct {
	Payload *models.EngineResourceDisplay
}

func (o *UpdateOutboundResourceDisplayOK) Error() string {
	return fmt.Sprintf("[PUT /call_center/resources/{resource_id}/display/{id}][%d] updateOutboundResourceDisplayOK  %+v", 200, o.Payload)
}
func (o *UpdateOutboundResourceDisplayOK) GetPayload() *models.EngineResourceDisplay {
	return o.Payload
}

func (o *UpdateOutboundResourceDisplayOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineResourceDisplay)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOutboundResourceDisplayForbidden creates a UpdateOutboundResourceDisplayForbidden with default headers values
func NewUpdateOutboundResourceDisplayForbidden() *UpdateOutboundResourceDisplayForbidden {
	return &UpdateOutboundResourceDisplayForbidden{}
}

/* UpdateOutboundResourceDisplayForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type UpdateOutboundResourceDisplayForbidden struct {
	Payload interface{}
}

func (o *UpdateOutboundResourceDisplayForbidden) Error() string {
	return fmt.Sprintf("[PUT /call_center/resources/{resource_id}/display/{id}][%d] updateOutboundResourceDisplayForbidden  %+v", 403, o.Payload)
}
func (o *UpdateOutboundResourceDisplayForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateOutboundResourceDisplayForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOutboundResourceDisplayNotFound creates a UpdateOutboundResourceDisplayNotFound with default headers values
func NewUpdateOutboundResourceDisplayNotFound() *UpdateOutboundResourceDisplayNotFound {
	return &UpdateOutboundResourceDisplayNotFound{}
}

/* UpdateOutboundResourceDisplayNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type UpdateOutboundResourceDisplayNotFound struct {
	Payload string
}

func (o *UpdateOutboundResourceDisplayNotFound) Error() string {
	return fmt.Sprintf("[PUT /call_center/resources/{resource_id}/display/{id}][%d] updateOutboundResourceDisplayNotFound  %+v", 404, o.Payload)
}
func (o *UpdateOutboundResourceDisplayNotFound) GetPayload() string {
	return o.Payload
}

func (o *UpdateOutboundResourceDisplayNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
