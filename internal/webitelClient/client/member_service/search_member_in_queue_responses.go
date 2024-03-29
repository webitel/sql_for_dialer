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

// SearchMemberInQueueReader is a Reader for the SearchMemberInQueue structure.
type SearchMemberInQueueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchMemberInQueueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchMemberInQueueOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewSearchMemberInQueueForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchMemberInQueueNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchMemberInQueueOK creates a SearchMemberInQueueOK with default headers values
func NewSearchMemberInQueueOK() *SearchMemberInQueueOK {
	return &SearchMemberInQueueOK{}
}

/* SearchMemberInQueueOK describes a response with status code 200, with default header values.

A successful response.
*/
type SearchMemberInQueueOK struct {
	Payload *models.EngineListMember
}

func (o *SearchMemberInQueueOK) Error() string {
	return fmt.Sprintf("[GET /call_center/queues/{queue_id}/members][%d] searchMemberInQueueOK  %+v", 200, o.Payload)
}
func (o *SearchMemberInQueueOK) GetPayload() *models.EngineListMember {
	return o.Payload
}

func (o *SearchMemberInQueueOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineListMember)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchMemberInQueueForbidden creates a SearchMemberInQueueForbidden with default headers values
func NewSearchMemberInQueueForbidden() *SearchMemberInQueueForbidden {
	return &SearchMemberInQueueForbidden{}
}

/* SearchMemberInQueueForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type SearchMemberInQueueForbidden struct {
	Payload interface{}
}

func (o *SearchMemberInQueueForbidden) Error() string {
	return fmt.Sprintf("[GET /call_center/queues/{queue_id}/members][%d] searchMemberInQueueForbidden  %+v", 403, o.Payload)
}
func (o *SearchMemberInQueueForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *SearchMemberInQueueForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchMemberInQueueNotFound creates a SearchMemberInQueueNotFound with default headers values
func NewSearchMemberInQueueNotFound() *SearchMemberInQueueNotFound {
	return &SearchMemberInQueueNotFound{}
}

/* SearchMemberInQueueNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type SearchMemberInQueueNotFound struct {
	Payload string
}

func (o *SearchMemberInQueueNotFound) Error() string {
	return fmt.Sprintf("[GET /call_center/queues/{queue_id}/members][%d] searchMemberInQueueNotFound  %+v", 404, o.Payload)
}
func (o *SearchMemberInQueueNotFound) GetPayload() string {
	return o.Payload
}

func (o *SearchMemberInQueueNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
