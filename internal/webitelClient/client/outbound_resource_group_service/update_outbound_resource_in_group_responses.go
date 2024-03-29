// Code generated by go-swagger; DO NOT EDIT.

package outbound_resource_group_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/webitel/sql_for_dialer/internal/webitelClient/models"
)

// UpdateOutboundResourceInGroupReader is a Reader for the UpdateOutboundResourceInGroup structure.
type UpdateOutboundResourceInGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOutboundResourceInGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOutboundResourceInGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateOutboundResourceInGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateOutboundResourceInGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateOutboundResourceInGroupOK creates a UpdateOutboundResourceInGroupOK with default headers values
func NewUpdateOutboundResourceInGroupOK() *UpdateOutboundResourceInGroupOK {
	return &UpdateOutboundResourceInGroupOK{}
}

/* UpdateOutboundResourceInGroupOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateOutboundResourceInGroupOK struct {
	Payload *models.EngineOutboundResourceInGroup
}

func (o *UpdateOutboundResourceInGroupOK) Error() string {
	return fmt.Sprintf("[PUT /call_center/resource_group/{group_id}/resource/{id}][%d] updateOutboundResourceInGroupOK  %+v", 200, o.Payload)
}
func (o *UpdateOutboundResourceInGroupOK) GetPayload() *models.EngineOutboundResourceInGroup {
	return o.Payload
}

func (o *UpdateOutboundResourceInGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineOutboundResourceInGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOutboundResourceInGroupForbidden creates a UpdateOutboundResourceInGroupForbidden with default headers values
func NewUpdateOutboundResourceInGroupForbidden() *UpdateOutboundResourceInGroupForbidden {
	return &UpdateOutboundResourceInGroupForbidden{}
}

/* UpdateOutboundResourceInGroupForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type UpdateOutboundResourceInGroupForbidden struct {
	Payload interface{}
}

func (o *UpdateOutboundResourceInGroupForbidden) Error() string {
	return fmt.Sprintf("[PUT /call_center/resource_group/{group_id}/resource/{id}][%d] updateOutboundResourceInGroupForbidden  %+v", 403, o.Payload)
}
func (o *UpdateOutboundResourceInGroupForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateOutboundResourceInGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOutboundResourceInGroupNotFound creates a UpdateOutboundResourceInGroupNotFound with default headers values
func NewUpdateOutboundResourceInGroupNotFound() *UpdateOutboundResourceInGroupNotFound {
	return &UpdateOutboundResourceInGroupNotFound{}
}

/* UpdateOutboundResourceInGroupNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type UpdateOutboundResourceInGroupNotFound struct {
	Payload string
}

func (o *UpdateOutboundResourceInGroupNotFound) Error() string {
	return fmt.Sprintf("[PUT /call_center/resource_group/{group_id}/resource/{id}][%d] updateOutboundResourceInGroupNotFound  %+v", 404, o.Payload)
}
func (o *UpdateOutboundResourceInGroupNotFound) GetPayload() string {
	return o.Payload
}

func (o *UpdateOutboundResourceInGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
