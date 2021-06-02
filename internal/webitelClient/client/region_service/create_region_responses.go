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

// CreateRegionReader is a Reader for the CreateRegion structure.
type CreateRegionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRegionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateRegionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateRegionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateRegionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRegionOK creates a CreateRegionOK with default headers values
func NewCreateRegionOK() *CreateRegionOK {
	return &CreateRegionOK{}
}

/* CreateRegionOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateRegionOK struct {
	Payload *models.EngineRegion
}

func (o *CreateRegionOK) Error() string {
	return fmt.Sprintf("[POST /regions][%d] createRegionOK  %+v", 200, o.Payload)
}
func (o *CreateRegionOK) GetPayload() *models.EngineRegion {
	return o.Payload
}

func (o *CreateRegionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineRegion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRegionForbidden creates a CreateRegionForbidden with default headers values
func NewCreateRegionForbidden() *CreateRegionForbidden {
	return &CreateRegionForbidden{}
}

/* CreateRegionForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type CreateRegionForbidden struct {
	Payload interface{}
}

func (o *CreateRegionForbidden) Error() string {
	return fmt.Sprintf("[POST /regions][%d] createRegionForbidden  %+v", 403, o.Payload)
}
func (o *CreateRegionForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateRegionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRegionNotFound creates a CreateRegionNotFound with default headers values
func NewCreateRegionNotFound() *CreateRegionNotFound {
	return &CreateRegionNotFound{}
}

/* CreateRegionNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type CreateRegionNotFound struct {
	Payload string
}

func (o *CreateRegionNotFound) Error() string {
	return fmt.Sprintf("[POST /regions][%d] createRegionNotFound  %+v", 404, o.Payload)
}
func (o *CreateRegionNotFound) GetPayload() string {
	return o.Payload
}

func (o *CreateRegionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}