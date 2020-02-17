// Code generated by go-swagger; DO NOT EDIT.

package sites

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/ffddorf/unms-api-go/models"
)

// PostSitesReader is a Reader for the PostSites structure.
type PostSitesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSitesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostSitesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostSitesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostSitesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostSitesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostSitesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostSitesOK creates a PostSitesOK with default headers values
func NewPostSitesOK() *PostSitesOK {
	return &PostSitesOK{}
}

/*PostSitesOK handles this case with default header values.

Successful
*/
type PostSitesOK struct {
	Payload *models.Site
}

func (o *PostSitesOK) Error() string {
	return fmt.Sprintf("[POST /sites][%d] postSitesOK  %+v", 200, o.Payload)
}

func (o *PostSitesOK) GetPayload() *models.Site {
	return o.Payload
}

func (o *PostSitesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Site)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSitesBadRequest creates a PostSitesBadRequest with default headers values
func NewPostSitesBadRequest() *PostSitesBadRequest {
	return &PostSitesBadRequest{}
}

/*PostSitesBadRequest handles this case with default header values.

Bad Request
*/
type PostSitesBadRequest struct {
	Payload *models.Error
}

func (o *PostSitesBadRequest) Error() string {
	return fmt.Sprintf("[POST /sites][%d] postSitesBadRequest  %+v", 400, o.Payload)
}

func (o *PostSitesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostSitesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSitesUnauthorized creates a PostSitesUnauthorized with default headers values
func NewPostSitesUnauthorized() *PostSitesUnauthorized {
	return &PostSitesUnauthorized{}
}

/*PostSitesUnauthorized handles this case with default header values.

Unauthorized
*/
type PostSitesUnauthorized struct {
	Payload *models.Error
}

func (o *PostSitesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /sites][%d] postSitesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostSitesUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostSitesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSitesForbidden creates a PostSitesForbidden with default headers values
func NewPostSitesForbidden() *PostSitesForbidden {
	return &PostSitesForbidden{}
}

/*PostSitesForbidden handles this case with default header values.

Forbidden
*/
type PostSitesForbidden struct {
	Payload *models.Error
}

func (o *PostSitesForbidden) Error() string {
	return fmt.Sprintf("[POST /sites][%d] postSitesForbidden  %+v", 403, o.Payload)
}

func (o *PostSitesForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostSitesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSitesInternalServerError creates a PostSitesInternalServerError with default headers values
func NewPostSitesInternalServerError() *PostSitesInternalServerError {
	return &PostSitesInternalServerError{}
}

/*PostSitesInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostSitesInternalServerError struct {
	Payload *models.Error
}

func (o *PostSitesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /sites][%d] postSitesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSitesInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostSitesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
