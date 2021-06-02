// Code generated by go-swagger; DO NOT EDIT.

package bucket_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/webitel/sql_for_dialer/internal/webitelClient/models"
)

// CreateBucketReader is a Reader for the CreateBucket structure.
type CreateBucketReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateBucketReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateBucketOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateBucketForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateBucketNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateBucketOK creates a CreateBucketOK with default headers values
func NewCreateBucketOK() *CreateBucketOK {
	return &CreateBucketOK{}
}

/* CreateBucketOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateBucketOK struct {
	Payload *models.EngineBucket
}

func (o *CreateBucketOK) Error() string {
	return fmt.Sprintf("[POST /call_center/buckets][%d] createBucketOK  %+v", 200, o.Payload)
}
func (o *CreateBucketOK) GetPayload() *models.EngineBucket {
	return o.Payload
}

func (o *CreateBucketOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineBucket)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBucketForbidden creates a CreateBucketForbidden with default headers values
func NewCreateBucketForbidden() *CreateBucketForbidden {
	return &CreateBucketForbidden{}
}

/* CreateBucketForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type CreateBucketForbidden struct {
	Payload interface{}
}

func (o *CreateBucketForbidden) Error() string {
	return fmt.Sprintf("[POST /call_center/buckets][%d] createBucketForbidden  %+v", 403, o.Payload)
}
func (o *CreateBucketForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateBucketForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBucketNotFound creates a CreateBucketNotFound with default headers values
func NewCreateBucketNotFound() *CreateBucketNotFound {
	return &CreateBucketNotFound{}
}

/* CreateBucketNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type CreateBucketNotFound struct {
	Payload string
}

func (o *CreateBucketNotFound) Error() string {
	return fmt.Sprintf("[POST /call_center/buckets][%d] createBucketNotFound  %+v", 404, o.Payload)
}
func (o *CreateBucketNotFound) GetPayload() string {
	return o.Payload
}

func (o *CreateBucketNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
