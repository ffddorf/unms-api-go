// Code generated by go-swagger; DO NOT EDIT.

package discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/ffddorf/unms-api-go/models"
)

// PostDiscoveryImportReader is a Reader for the PostDiscoveryImport structure.
type PostDiscoveryImportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDiscoveryImportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDiscoveryImportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostDiscoveryImportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostDiscoveryImportUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostDiscoveryImportInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostDiscoveryImportOK creates a PostDiscoveryImportOK with default headers values
func NewPostDiscoveryImportOK() *PostDiscoveryImportOK {
	return &PostDiscoveryImportOK{}
}

/*PostDiscoveryImportOK handles this case with default header values.

Successful
*/
type PostDiscoveryImportOK struct {
	Payload *models.Status
}

func (o *PostDiscoveryImportOK) Error() string {
	return fmt.Sprintf("[POST /discovery/import][%d] postDiscoveryImportOK  %+v", 200, o.Payload)
}

func (o *PostDiscoveryImportOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *PostDiscoveryImportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDiscoveryImportBadRequest creates a PostDiscoveryImportBadRequest with default headers values
func NewPostDiscoveryImportBadRequest() *PostDiscoveryImportBadRequest {
	return &PostDiscoveryImportBadRequest{}
}

/*PostDiscoveryImportBadRequest handles this case with default header values.

Bad Request
*/
type PostDiscoveryImportBadRequest struct {
	Payload *models.Error
}

func (o *PostDiscoveryImportBadRequest) Error() string {
	return fmt.Sprintf("[POST /discovery/import][%d] postDiscoveryImportBadRequest  %+v", 400, o.Payload)
}

func (o *PostDiscoveryImportBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDiscoveryImportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDiscoveryImportUnauthorized creates a PostDiscoveryImportUnauthorized with default headers values
func NewPostDiscoveryImportUnauthorized() *PostDiscoveryImportUnauthorized {
	return &PostDiscoveryImportUnauthorized{}
}

/*PostDiscoveryImportUnauthorized handles this case with default header values.

Unauthorized
*/
type PostDiscoveryImportUnauthorized struct {
	Payload *models.Error
}

func (o *PostDiscoveryImportUnauthorized) Error() string {
	return fmt.Sprintf("[POST /discovery/import][%d] postDiscoveryImportUnauthorized  %+v", 401, o.Payload)
}

func (o *PostDiscoveryImportUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDiscoveryImportUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDiscoveryImportInternalServerError creates a PostDiscoveryImportInternalServerError with default headers values
func NewPostDiscoveryImportInternalServerError() *PostDiscoveryImportInternalServerError {
	return &PostDiscoveryImportInternalServerError{}
}

/*PostDiscoveryImportInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostDiscoveryImportInternalServerError struct {
	Payload *models.Error
}

func (o *PostDiscoveryImportInternalServerError) Error() string {
	return fmt.Sprintf("[POST /discovery/import][%d] postDiscoveryImportInternalServerError  %+v", 500, o.Payload)
}

func (o *PostDiscoveryImportInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDiscoveryImportInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
