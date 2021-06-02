// Code generated by go-swagger; DO NOT EDIT.

package region_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/webitel/sql_for_dialer/internal/webitelClient/models"
)

// DeleteRegionReader is a Reader for the DeleteRegion structure.
type DeleteRegionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRegionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRegionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteRegionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRegionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRegionOK creates a DeleteRegionOK with default headers values
func NewDeleteRegionOK() *DeleteRegionOK {
	return &DeleteRegionOK{}
}

/* DeleteRegionOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteRegionOK struct {
	Payload *models.EngineRegion
}

func (o *DeleteRegionOK) Error() string {
	return fmt.Sprintf("[DELETE /regions/{id}][%d] deleteRegionOK  %+v", 200, o.Payload)
}
func (o *DeleteRegionOK) GetPayload() *models.EngineRegion {
	return o.Payload
}

func (o *DeleteRegionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineRegion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRegionForbidden creates a DeleteRegionForbidden with default headers values
func NewDeleteRegionForbidden() *DeleteRegionForbidden {
	return &DeleteRegionForbidden{}
}

/* DeleteRegionForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type DeleteRegionForbidden struct {
	Payload interface{}
}

func (o *DeleteRegionForbidden) Error() string {
	return fmt.Sprintf("[DELETE /regions/{id}][%d] deleteRegionForbidden  %+v", 403, o.Payload)
}
func (o *DeleteRegionForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteRegionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRegionNotFound creates a DeleteRegionNotFound with default headers values
func NewDeleteRegionNotFound() *DeleteRegionNotFound {
	return &DeleteRegionNotFound{}
}

/* DeleteRegionNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type DeleteRegionNotFound struct {
	Payload string
}

func (o *DeleteRegionNotFound) Error() string {
	return fmt.Sprintf("[DELETE /regions/{id}][%d] deleteRegionNotFound  %+v", 404, o.Payload)
}
func (o *DeleteRegionNotFound) GetPayload() string {
	return o.Payload
}

func (o *DeleteRegionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}