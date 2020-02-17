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

// PostSitesSiteidUnsuspendReader is a Reader for the PostSitesSiteidUnsuspend structure.
type PostSitesSiteidUnsuspendReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSitesSiteidUnsuspendReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostSitesSiteidUnsuspendOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostSitesSiteidUnsuspendBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostSitesSiteidUnsuspendUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostSitesSiteidUnsuspendForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostSitesSiteidUnsuspendNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostSitesSiteidUnsuspendInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostSitesSiteidUnsuspendOK creates a PostSitesSiteidUnsuspendOK with default headers values
func NewPostSitesSiteidUnsuspendOK() *PostSitesSiteidUnsuspendOK {
	return &PostSitesSiteidUnsuspendOK{}
}

/*PostSitesSiteidUnsuspendOK handles this case with default header values.

Result of unsuspend, updated site
*/
type PostSitesSiteidUnsuspendOK struct {
	Payload *models.Site2
}

func (o *PostSitesSiteidUnsuspendOK) Error() string {
	return fmt.Sprintf("[POST /sites/{siteId}/unsuspend][%d] postSitesSiteidUnsuspendOK  %+v", 200, o.Payload)
}

func (o *PostSitesSiteidUnsuspendOK) GetPayload() *models.Site2 {
	return o.Payload
}

func (o *PostSitesSiteidUnsuspendOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Site2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSitesSiteidUnsuspendBadRequest creates a PostSitesSiteidUnsuspendBadRequest with default headers values
func NewPostSitesSiteidUnsuspendBadRequest() *PostSitesSiteidUnsuspendBadRequest {
	return &PostSitesSiteidUnsuspendBadRequest{}
}

/*PostSitesSiteidUnsuspendBadRequest handles this case with default header values.

Bad Request
*/
type PostSitesSiteidUnsuspendBadRequest struct {
	Payload *models.Error
}

func (o *PostSitesSiteidUnsuspendBadRequest) Error() string {
	return fmt.Sprintf("[POST /sites/{siteId}/unsuspend][%d] postSitesSiteidUnsuspendBadRequest  %+v", 400, o.Payload)
}

func (o *PostSitesSiteidUnsuspendBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostSitesSiteidUnsuspendBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSitesSiteidUnsuspendUnauthorized creates a PostSitesSiteidUnsuspendUnauthorized with default headers values
func NewPostSitesSiteidUnsuspendUnauthorized() *PostSitesSiteidUnsuspendUnauthorized {
	return &PostSitesSiteidUnsuspendUnauthorized{}
}

/*PostSitesSiteidUnsuspendUnauthorized handles this case with default header values.

Unauthorized
*/
type PostSitesSiteidUnsuspendUnauthorized struct {
	Payload *models.Error
}

func (o *PostSitesSiteidUnsuspendUnauthorized) Error() string {
	return fmt.Sprintf("[POST /sites/{siteId}/unsuspend][%d] postSitesSiteidUnsuspendUnauthorized  %+v", 401, o.Payload)
}

func (o *PostSitesSiteidUnsuspendUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostSitesSiteidUnsuspendUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSitesSiteidUnsuspendForbidden creates a PostSitesSiteidUnsuspendForbidden with default headers values
func NewPostSitesSiteidUnsuspendForbidden() *PostSitesSiteidUnsuspendForbidden {
	return &PostSitesSiteidUnsuspendForbidden{}
}

/*PostSitesSiteidUnsuspendForbidden handles this case with default header values.

Forbidden
*/
type PostSitesSiteidUnsuspendForbidden struct {
	Payload *models.Error
}

func (o *PostSitesSiteidUnsuspendForbidden) Error() string {
	return fmt.Sprintf("[POST /sites/{siteId}/unsuspend][%d] postSitesSiteidUnsuspendForbidden  %+v", 403, o.Payload)
}

func (o *PostSitesSiteidUnsuspendForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostSitesSiteidUnsuspendForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSitesSiteidUnsuspendNotFound creates a PostSitesSiteidUnsuspendNotFound with default headers values
func NewPostSitesSiteidUnsuspendNotFound() *PostSitesSiteidUnsuspendNotFound {
	return &PostSitesSiteidUnsuspendNotFound{}
}

/*PostSitesSiteidUnsuspendNotFound handles this case with default header values.

Not Found
*/
type PostSitesSiteidUnsuspendNotFound struct {
	Payload *models.Error
}

func (o *PostSitesSiteidUnsuspendNotFound) Error() string {
	return fmt.Sprintf("[POST /sites/{siteId}/unsuspend][%d] postSitesSiteidUnsuspendNotFound  %+v", 404, o.Payload)
}

func (o *PostSitesSiteidUnsuspendNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostSitesSiteidUnsuspendNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSitesSiteidUnsuspendInternalServerError creates a PostSitesSiteidUnsuspendInternalServerError with default headers values
func NewPostSitesSiteidUnsuspendInternalServerError() *PostSitesSiteidUnsuspendInternalServerError {
	return &PostSitesSiteidUnsuspendInternalServerError{}
}

/*PostSitesSiteidUnsuspendInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostSitesSiteidUnsuspendInternalServerError struct {
	Payload *models.Error
}

func (o *PostSitesSiteidUnsuspendInternalServerError) Error() string {
	return fmt.Sprintf("[POST /sites/{siteId}/unsuspend][%d] postSitesSiteidUnsuspendInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSitesSiteidUnsuspendInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostSitesSiteidUnsuspendInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
