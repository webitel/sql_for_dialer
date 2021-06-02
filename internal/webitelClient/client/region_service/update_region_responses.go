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

// UpdateRegionReader is a Reader for the UpdateRegion structure.
type UpdateRegionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRegionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRegionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateRegionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRegionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRegionOK creates a UpdateRegionOK with default headers values
func NewUpdateRegionOK() *UpdateRegionOK {
	return &UpdateRegionOK{}
}

/* UpdateRegionOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateRegionOK struct {
	Payload *models.EngineRegion
}

func (o *UpdateRegionOK) Error() string {
	return fmt.Sprintf("[PUT /regions/{id}][%d] updateRegionOK  %+v", 200, o.Payload)
}
func (o *UpdateRegionOK) GetPayload() *models.EngineRegion {
	return o.Payload
}

func (o *UpdateRegionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineRegion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRegionForbidden creates a UpdateRegionForbidden with default headers values
func NewUpdateRegionForbidden() *UpdateRegionForbidden {
	return &UpdateRegionForbidden{}
}

/* UpdateRegionForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type UpdateRegionForbidden struct {
	Payload interface{}
}

func (o *UpdateRegionForbidden) Error() string {
	return fmt.Sprintf("[PUT /regions/{id}][%d] updateRegionForbidden  %+v", 403, o.Payload)
}
func (o *UpdateRegionForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateRegionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRegionNotFound creates a UpdateRegionNotFound with default headers values
func NewUpdateRegionNotFound() *UpdateRegionNotFound {
	return &UpdateRegionNotFound{}
}

/* UpdateRegionNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type UpdateRegionNotFound struct {
	Payload string
}

func (o *UpdateRegionNotFound) Error() string {
	return fmt.Sprintf("[PUT /regions/{id}][%d] updateRegionNotFound  %+v", 404, o.Payload)
}
func (o *UpdateRegionNotFound) GetPayload() string {
	return o.Payload
}

func (o *UpdateRegionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
