// Code generated by go-swagger; DO NOT EDIT.

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/ffddorf/unms-api-go/models"
)

// PostUserLastreleasenotesseenReader is a Reader for the PostUserLastreleasenotesseen structure.
type PostUserLastreleasenotesseenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUserLastreleasenotesseenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostUserLastreleasenotesseenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostUserLastreleasenotesseenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostUserLastreleasenotesseenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostUserLastreleasenotesseenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostUserLastreleasenotesseenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostUserLastreleasenotesseenOK creates a PostUserLastreleasenotesseenOK with default headers values
func NewPostUserLastreleasenotesseenOK() *PostUserLastreleasenotesseenOK {
	return &PostUserLastreleasenotesseenOK{}
}

/*PostUserLastreleasenotesseenOK handles this case with default header values.

Successful
*/
type PostUserLastreleasenotesseenOK struct {
	Payload *models.Status
}

func (o *PostUserLastreleasenotesseenOK) Error() string {
	return fmt.Sprintf("[POST /user/last-release-notes-seen][%d] postUserLastreleasenotesseenOK  %+v", 200, o.Payload)
}

func (o *PostUserLastreleasenotesseenOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *PostUserLastreleasenotesseenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserLastreleasenotesseenBadRequest creates a PostUserLastreleasenotesseenBadRequest with default headers values
func NewPostUserLastreleasenotesseenBadRequest() *PostUserLastreleasenotesseenBadRequest {
	return &PostUserLastreleasenotesseenBadRequest{}
}

/*PostUserLastreleasenotesseenBadRequest handles this case with default header values.

Bad Request
*/
type PostUserLastreleasenotesseenBadRequest struct {
	Payload *models.Error
}

func (o *PostUserLastreleasenotesseenBadRequest) Error() string {
	return fmt.Sprintf("[POST /user/last-release-notes-seen][%d] postUserLastreleasenotesseenBadRequest  %+v", 400, o.Payload)
}

func (o *PostUserLastreleasenotesseenBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostUserLastreleasenotesseenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserLastreleasenotesseenUnauthorized creates a PostUserLastreleasenotesseenUnauthorized with default headers values
func NewPostUserLastreleasenotesseenUnauthorized() *PostUserLastreleasenotesseenUnauthorized {
	return &PostUserLastreleasenotesseenUnauthorized{}
}

/*PostUserLastreleasenotesseenUnauthorized handles this case with default header values.

Unauthorized
*/
type PostUserLastreleasenotesseenUnauthorized struct {
	Payload *models.Error
}

func (o *PostUserLastreleasenotesseenUnauthorized) Error() string {
	return fmt.Sprintf("[POST /user/last-release-notes-seen][%d] postUserLastreleasenotesseenUnauthorized  %+v", 401, o.Payload)
}

func (o *PostUserLastreleasenotesseenUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostUserLastreleasenotesseenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserLastreleasenotesseenForbidden creates a PostUserLastreleasenotesseenForbidden with default headers values
func NewPostUserLastreleasenotesseenForbidden() *PostUserLastreleasenotesseenForbidden {
	return &PostUserLastreleasenotesseenForbidden{}
}

/*PostUserLastreleasenotesseenForbidden handles this case with default header values.

Forbidden
*/
type PostUserLastreleasenotesseenForbidden struct {
	Payload *models.Error
}

func (o *PostUserLastreleasenotesseenForbidden) Error() string {
	return fmt.Sprintf("[POST /user/last-release-notes-seen][%d] postUserLastreleasenotesseenForbidden  %+v", 403, o.Payload)
}

func (o *PostUserLastreleasenotesseenForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostUserLastreleasenotesseenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserLastreleasenotesseenInternalServerError creates a PostUserLastreleasenotesseenInternalServerError with default headers values
func NewPostUserLastreleasenotesseenInternalServerError() *PostUserLastreleasenotesseenInternalServerError {
	return &PostUserLastreleasenotesseenInternalServerError{}
}

/*PostUserLastreleasenotesseenInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostUserLastreleasenotesseenInternalServerError struct {
	Payload *models.Error
}

func (o *PostUserLastreleasenotesseenInternalServerError) Error() string {
	return fmt.Sprintf("[POST /user/last-release-notes-seen][%d] postUserLastreleasenotesseenInternalServerError  %+v", 500, o.Payload)
}

func (o *PostUserLastreleasenotesseenInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostUserLastreleasenotesseenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}