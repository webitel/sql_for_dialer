// Code generated by go-swagger; DO NOT EDIT.

package media_file_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/webitel/sql_for_dialer/internal/webitelClient/models"
)

// ReadMediaFileReader is a Reader for the ReadMediaFile structure.
type ReadMediaFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadMediaFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadMediaFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadMediaFileOK creates a ReadMediaFileOK with default headers values
func NewReadMediaFileOK() *ReadMediaFileOK {
	return &ReadMediaFileOK{}
}

/* ReadMediaFileOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadMediaFileOK struct {
	Payload *models.StorageMediaFile
}

func (o *ReadMediaFileOK) Error() string {
	return fmt.Sprintf("[GET /storage/media/{id}][%d] readMediaFileOK  %+v", 200, o.Payload)
}
func (o *ReadMediaFileOK) GetPayload() *models.StorageMediaFile {
	return o.Payload
}

func (o *ReadMediaFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageMediaFile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
