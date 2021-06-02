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

// SearchRegionReader is a Reader for the SearchRegion structure.
type SearchRegionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchRegionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchRegionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewSearchRegionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchRegionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchRegionOK creates a SearchRegionOK with default headers values
func NewSearchRegionOK() *SearchRegionOK {
	return &SearchRegionOK{}
}

/* SearchRegionOK describes a response with status code 200, with default header values.

A successful response.
*/
type SearchRegionOK struct {
	Payload *models.EngineListRegion
}

func (o *SearchRegionOK) Error() string {
	return fmt.Sprintf("[GET /regions][%d] searchRegionOK  %+v", 200, o.Payload)
}
func (o *SearchRegionOK) GetPayload() *models.EngineListRegion {
	return o.Payload
}

func (o *SearchRegionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineListRegion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchRegionForbidden creates a SearchRegionForbidden with default headers values
func NewSearchRegionForbidden() *SearchRegionForbidden {
	return &SearchRegionForbidden{}
}

/* SearchRegionForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type SearchRegionForbidden struct {
	Payload interface{}
}

func (o *SearchRegionForbidden) Error() string {
	return fmt.Sprintf("[GET /regions][%d] searchRegionForbidden  %+v", 403, o.Payload)
}
func (o *SearchRegionForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *SearchRegionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchRegionNotFound creates a SearchRegionNotFound with default headers values
func NewSearchRegionNotFound() *SearchRegionNotFound {
	return &SearchRegionNotFound{}
}

/* SearchRegionNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type SearchRegionNotFound struct {
	Payload string
}

func (o *SearchRegionNotFound) Error() string {
	return fmt.Sprintf("[GET /regions][%d] searchRegionNotFound  %+v", 404, o.Payload)
}
func (o *SearchRegionNotFound) GetPayload() string {
	return o.Payload
}

func (o *SearchRegionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}