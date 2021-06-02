// Code generated by go-swagger; DO NOT EDIT.

package agent_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/webitel/sql_for_dialer/internal/webitelClient/models"
)

// SearchLookupUsersAgentNotExistsReader is a Reader for the SearchLookupUsersAgentNotExists structure.
type SearchLookupUsersAgentNotExistsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchLookupUsersAgentNotExistsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchLookupUsersAgentNotExistsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewSearchLookupUsersAgentNotExistsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchLookupUsersAgentNotExistsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchLookupUsersAgentNotExistsOK creates a SearchLookupUsersAgentNotExistsOK with default headers values
func NewSearchLookupUsersAgentNotExistsOK() *SearchLookupUsersAgentNotExistsOK {
	return &SearchLookupUsersAgentNotExistsOK{}
}

/* SearchLookupUsersAgentNotExistsOK describes a response with status code 200, with default header values.

A successful response.
*/
type SearchLookupUsersAgentNotExistsOK struct {
	Payload *models.EngineListAgentUser
}

func (o *SearchLookupUsersAgentNotExistsOK) Error() string {
	return fmt.Sprintf("[GET /call_center/lookups/agents/users][%d] searchLookupUsersAgentNotExistsOK  %+v", 200, o.Payload)
}
func (o *SearchLookupUsersAgentNotExistsOK) GetPayload() *models.EngineListAgentUser {
	return o.Payload
}

func (o *SearchLookupUsersAgentNotExistsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineListAgentUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchLookupUsersAgentNotExistsForbidden creates a SearchLookupUsersAgentNotExistsForbidden with default headers values
func NewSearchLookupUsersAgentNotExistsForbidden() *SearchLookupUsersAgentNotExistsForbidden {
	return &SearchLookupUsersAgentNotExistsForbidden{}
}

/* SearchLookupUsersAgentNotExistsForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type SearchLookupUsersAgentNotExistsForbidden struct {
	Payload interface{}
}

func (o *SearchLookupUsersAgentNotExistsForbidden) Error() string {
	return fmt.Sprintf("[GET /call_center/lookups/agents/users][%d] searchLookupUsersAgentNotExistsForbidden  %+v", 403, o.Payload)
}
func (o *SearchLookupUsersAgentNotExistsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *SearchLookupUsersAgentNotExistsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchLookupUsersAgentNotExistsNotFound creates a SearchLookupUsersAgentNotExistsNotFound with default headers values
func NewSearchLookupUsersAgentNotExistsNotFound() *SearchLookupUsersAgentNotExistsNotFound {
	return &SearchLookupUsersAgentNotExistsNotFound{}
}

/* SearchLookupUsersAgentNotExistsNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type SearchLookupUsersAgentNotExistsNotFound struct {
	Payload string
}

func (o *SearchLookupUsersAgentNotExistsNotFound) Error() string {
	return fmt.Sprintf("[GET /call_center/lookups/agents/users][%d] searchLookupUsersAgentNotExistsNotFound  %+v", 404, o.Payload)
}
func (o *SearchLookupUsersAgentNotExistsNotFound) GetPayload() string {
	return o.Payload
}

func (o *SearchLookupUsersAgentNotExistsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
