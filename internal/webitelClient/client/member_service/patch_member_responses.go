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

// PatchMemberReader is a Reader for the PatchMember structure.
type PatchMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPatchMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchMemberOK creates a PatchMemberOK with default headers values
func NewPatchMemberOK() *PatchMemberOK {
	return &PatchMemberOK{}
}

/* PatchMemberOK describes a response with status code 200, with default header values.

A successful response.
*/
type PatchMemberOK struct {
	Payload *models.EngineMemberInQueue
}

func (o *PatchMemberOK) Error() string {
	return fmt.Sprintf("[PATCH /call_center/queues/{queue_id}/members/{id}][%d] patchMemberOK  %+v", 200, o.Payload)
}
func (o *PatchMemberOK) GetPayload() *models.EngineMemberInQueue {
	return o.Payload
}

func (o *PatchMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineMemberInQueue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchMemberForbidden creates a PatchMemberForbidden with default headers values
func NewPatchMemberForbidden() *PatchMemberForbidden {
	return &PatchMemberForbidden{}
}

/* PatchMemberForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type PatchMemberForbidden struct {
	Payload interface{}
}

func (o *PatchMemberForbidden) Error() string {
	return fmt.Sprintf("[PATCH /call_center/queues/{queue_id}/members/{id}][%d] patchMemberForbidden  %+v", 403, o.Payload)
}
func (o *PatchMemberForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchMemberNotFound creates a PatchMemberNotFound with default headers values
func NewPatchMemberNotFound() *PatchMemberNotFound {
	return &PatchMemberNotFound{}
}

/* PatchMemberNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type PatchMemberNotFound struct {
	Payload string
}

func (o *PatchMemberNotFound) Error() string {
	return fmt.Sprintf("[PATCH /call_center/queues/{queue_id}/members/{id}][%d] patchMemberNotFound  %+v", 404, o.Payload)
}
func (o *PatchMemberNotFound) GetPayload() string {
	return o.Payload
}

func (o *PatchMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
