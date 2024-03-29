// Code generated by go-swagger; DO NOT EDIT.

package queue_bucket_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/webitel/sql_for_dialer/internal/webitelClient/models"
)

// DeleteQueueBucketReader is a Reader for the DeleteQueueBucket structure.
type DeleteQueueBucketReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteQueueBucketReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteQueueBucketOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteQueueBucketForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteQueueBucketNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteQueueBucketOK creates a DeleteQueueBucketOK with default headers values
func NewDeleteQueueBucketOK() *DeleteQueueBucketOK {
	return &DeleteQueueBucketOK{}
}

/* DeleteQueueBucketOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteQueueBucketOK struct {
	Payload *models.EngineQueueBucket
}

func (o *DeleteQueueBucketOK) Error() string {
	return fmt.Sprintf("[DELETE /call_center/queues/{queue_id}/buckets/{id}][%d] deleteQueueBucketOK  %+v", 200, o.Payload)
}
func (o *DeleteQueueBucketOK) GetPayload() *models.EngineQueueBucket {
	return o.Payload
}

func (o *DeleteQueueBucketOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineQueueBucket)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteQueueBucketForbidden creates a DeleteQueueBucketForbidden with default headers values
func NewDeleteQueueBucketForbidden() *DeleteQueueBucketForbidden {
	return &DeleteQueueBucketForbidden{}
}

/* DeleteQueueBucketForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type DeleteQueueBucketForbidden struct {
	Payload interface{}
}

func (o *DeleteQueueBucketForbidden) Error() string {
	return fmt.Sprintf("[DELETE /call_center/queues/{queue_id}/buckets/{id}][%d] deleteQueueBucketForbidden  %+v", 403, o.Payload)
}
func (o *DeleteQueueBucketForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteQueueBucketForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteQueueBucketNotFound creates a DeleteQueueBucketNotFound with default headers values
func NewDeleteQueueBucketNotFound() *DeleteQueueBucketNotFound {
	return &DeleteQueueBucketNotFound{}
}

/* DeleteQueueBucketNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type DeleteQueueBucketNotFound struct {
	Payload string
}

func (o *DeleteQueueBucketNotFound) Error() string {
	return fmt.Sprintf("[DELETE /call_center/queues/{queue_id}/buckets/{id}][%d] deleteQueueBucketNotFound  %+v", 404, o.Payload)
}
func (o *DeleteQueueBucketNotFound) GetPayload() string {
	return o.Payload
}

func (o *DeleteQueueBucketNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
