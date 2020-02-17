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

// PostSitesSiteidImagesImageidRotateleftReader is a Reader for the PostSitesSiteidImagesImageidRotateleft structure.
type PostSitesSiteidImagesImageidRotateleftReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSitesSiteidImagesImageidRotateleftReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostSitesSiteidImagesImageidRotateleftOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostSitesSiteidImagesImageidRotateleftBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostSitesSiteidImagesImageidRotateleftUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostSitesSiteidImagesImageidRotateleftForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostSitesSiteidImagesImageidRotateleftInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostSitesSiteidImagesImageidRotateleftOK creates a PostSitesSiteidImagesImageidRotateleftOK with default headers values
func NewPostSitesSiteidImagesImageidRotateleftOK() *PostSitesSiteidImagesImageidRotateleftOK {
	return &PostSitesSiteidImagesImageidRotateleftOK{}
}

/*PostSitesSiteidImagesImageidRotateleftOK handles this case with default header values.

Successful
*/
type PostSitesSiteidImagesImageidRotateleftOK struct {
	Payload *models.Status
}

func (o *PostSitesSiteidImagesImageidRotateleftOK) Error() string {
	return fmt.Sprintf("[POST /sites/{siteId}/images/{imageId}/rotateleft][%d] postSitesSiteidImagesImageidRotateleftOK  %+v", 200, o.Payload)
}

func (o *PostSitesSiteidImagesImageidRotateleftOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *PostSitesSiteidImagesImageidRotateleftOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSitesSiteidImagesImageidRotateleftBadRequest creates a PostSitesSiteidImagesImageidRotateleftBadRequest with default headers values
func NewPostSitesSiteidImagesImageidRotateleftBadRequest() *PostSitesSiteidImagesImageidRotateleftBadRequest {
	return &PostSitesSiteidImagesImageidRotateleftBadRequest{}
}

/*PostSitesSiteidImagesImageidRotateleftBadRequest handles this case with default header values.

Bad Request
*/
type PostSitesSiteidImagesImageidRotateleftBadRequest struct {
	Payload *models.Error
}

func (o *PostSitesSiteidImagesImageidRotateleftBadRequest) Error() string {
	return fmt.Sprintf("[POST /sites/{siteId}/images/{imageId}/rotateleft][%d] postSitesSiteidImagesImageidRotateleftBadRequest  %+v", 400, o.Payload)
}

func (o *PostSitesSiteidImagesImageidRotateleftBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostSitesSiteidImagesImageidRotateleftBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSitesSiteidImagesImageidRotateleftUnauthorized creates a PostSitesSiteidImagesImageidRotateleftUnauthorized with default headers values
func NewPostSitesSiteidImagesImageidRotateleftUnauthorized() *PostSitesSiteidImagesImageidRotateleftUnauthorized {
	return &PostSitesSiteidImagesImageidRotateleftUnauthorized{}
}

/*PostSitesSiteidImagesImageidRotateleftUnauthorized handles this case with default header values.

Unauthorized
*/
type PostSitesSiteidImagesImageidRotateleftUnauthorized struct {
	Payload *models.Error
}

func (o *PostSitesSiteidImagesImageidRotateleftUnauthorized) Error() string {
	return fmt.Sprintf("[POST /sites/{siteId}/images/{imageId}/rotateleft][%d] postSitesSiteidImagesImageidRotateleftUnauthorized  %+v", 401, o.Payload)
}

func (o *PostSitesSiteidImagesImageidRotateleftUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostSitesSiteidImagesImageidRotateleftUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSitesSiteidImagesImageidRotateleftForbidden creates a PostSitesSiteidImagesImageidRotateleftForbidden with default headers values
func NewPostSitesSiteidImagesImageidRotateleftForbidden() *PostSitesSiteidImagesImageidRotateleftForbidden {
	return &PostSitesSiteidImagesImageidRotateleftForbidden{}
}

/*PostSitesSiteidImagesImageidRotateleftForbidden handles this case with default header values.

Forbidden
*/
type PostSitesSiteidImagesImageidRotateleftForbidden struct {
	Payload *models.Error
}

func (o *PostSitesSiteidImagesImageidRotateleftForbidden) Error() string {
	return fmt.Sprintf("[POST /sites/{siteId}/images/{imageId}/rotateleft][%d] postSitesSiteidImagesImageidRotateleftForbidden  %+v", 403, o.Payload)
}

func (o *PostSitesSiteidImagesImageidRotateleftForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostSitesSiteidImagesImageidRotateleftForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSitesSiteidImagesImageidRotateleftInternalServerError creates a PostSitesSiteidImagesImageidRotateleftInternalServerError with default headers values
func NewPostSitesSiteidImagesImageidRotateleftInternalServerError() *PostSitesSiteidImagesImageidRotateleftInternalServerError {
	return &PostSitesSiteidImagesImageidRotateleftInternalServerError{}
}

/*PostSitesSiteidImagesImageidRotateleftInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostSitesSiteidImagesImageidRotateleftInternalServerError struct {
	Payload *models.Error
}

func (o *PostSitesSiteidImagesImageidRotateleftInternalServerError) Error() string {
	return fmt.Sprintf("[POST /sites/{siteId}/images/{imageId}/rotateleft][%d] postSitesSiteidImagesImageidRotateleftInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSitesSiteidImagesImageidRotateleftInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostSitesSiteidImagesImageidRotateleftInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
