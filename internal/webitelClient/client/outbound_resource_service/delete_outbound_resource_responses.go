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

// DeleteOutboundResourceReader is a Reader for the DeleteOutboundResource structure.
type DeleteOutboundResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOutboundResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteOutboundResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteOutboundResourceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteOutboundResourceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteOutboundResourceOK creates a DeleteOutboundResourceOK with default headers values
func NewDeleteOutboundResourceOK() *DeleteOutboundResourceOK {
	return &DeleteOutboundResourceOK{}
}

/* DeleteOutboundResourceOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteOutboundResourceOK struct {
	Payload *models.EngineOutboundResource
}

func (o *DeleteOutboundResourceOK) Error() string {
	return fmt.Sprintf("[DELETE /call_center/resources/{id}][%d] deleteOutboundResourceOK  %+v", 200, o.Payload)
}
func (o *DeleteOutboundResourceOK) GetPayload() *models.EngineOutboundResource {
	return o.Payload
}

func (o *DeleteOutboundResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineOutboundResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundResourceForbidden creates a DeleteOutboundResourceForbidden with default headers values
func NewDeleteOutboundResourceForbidden() *DeleteOutboundResourceForbidden {
	return &DeleteOutboundResourceForbidden{}
}

/* DeleteOutboundResourceForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type DeleteOutboundResourceForbidden struct {
	Payload interface{}
}

func (o *DeleteOutboundResourceForbidden) Error() string {
	return fmt.Sprintf("[DELETE /call_center/resources/{id}][%d] deleteOutboundResourceForbidden  %+v", 403, o.Payload)
}
func (o *DeleteOutboundResourceForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteOutboundResourceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundResourceNotFound creates a DeleteOutboundResourceNotFound with default headers values
func NewDeleteOutboundResourceNotFound() *DeleteOutboundResourceNotFound {
	return &DeleteOutboundResourceNotFound{}
}

/* DeleteOutboundResourceNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type DeleteOutboundResourceNotFound struct {
	Payload string
}

func (o *DeleteOutboundResourceNotFound) Error() string {
	return fmt.Sprintf("[DELETE /call_center/resources/{id}][%d] deleteOutboundResourceNotFound  %+v", 404, o.Payload)
}
func (o *DeleteOutboundResourceNotFound) GetPayload() string {
	return o.Payload
}

func (o *DeleteOutboundResourceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
