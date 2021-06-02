// Code generated by go-swagger; DO NOT EDIT.

package member_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/webitel/sql_for_dialer/internal/webitelClient/models"
)

// UpdateMemberReader is a Reader for the UpdateMember structure.
type UpdateMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateMemberOK creates a UpdateMemberOK with default headers values
func NewUpdateMemberOK() *UpdateMemberOK {
	return &UpdateMemberOK{}
}

/* UpdateMemberOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateMemberOK struct {
	Payload *models.EngineMemberInQueue
}

func (o *UpdateMemberOK) Error() string {
	return fmt.Sprintf("[PUT /call_center/queues/{queue_id}/members/{id}][%d] updateMemberOK  %+v", 200, o.Payload)
}
func (o *UpdateMemberOK) GetPayload() *models.EngineMemberInQueue {
	return o.Payload
}

func (o *UpdateMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineMemberInQueue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMemberForbidden creates a UpdateMemberForbidden with default headers values
func NewUpdateMemberForbidden() *UpdateMemberForbidden {
	return &UpdateMemberForbidden{}
}

/* UpdateMemberForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type UpdateMemberForbidden struct {
	Payload interface{}
}

func (o *UpdateMemberForbidden) Error() string {
	return fmt.Sprintf("[PUT /call_center/queues/{queue_id}/members/{id}][%d] updateMemberForbidden  %+v", 403, o.Payload)
}
func (o *UpdateMemberForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMemberNotFound creates a UpdateMemberNotFound with default headers values
func NewUpdateMemberNotFound() *UpdateMemberNotFound {
	return &UpdateMemberNotFound{}
}

/* UpdateMemberNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type UpdateMemberNotFound struct {
	Payload string
}

func (o *UpdateMemberNotFound) Error() string {
	return fmt.Sprintf("[PUT /call_center/queues/{queue_id}/members/{id}][%d] updateMemberNotFound  %+v", 404, o.Payload)
}
func (o *UpdateMemberNotFound) GetPayload() string {
	return o.Payload
}

func (o *UpdateMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
