// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/ffddorf/unms-api-go/models"
)

// PostDevicesOnusIDUnblockReader is a Reader for the PostDevicesOnusIDUnblock structure.
type PostDevicesOnusIDUnblockReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDevicesOnusIDUnblockReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDevicesOnusIDUnblockOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostDevicesOnusIDUnblockBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostDevicesOnusIDUnblockUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostDevicesOnusIDUnblockForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostDevicesOnusIDUnblockNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostDevicesOnusIDUnblockInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostDevicesOnusIDUnblockOK creates a PostDevicesOnusIDUnblockOK with default headers values
func NewPostDevicesOnusIDUnblockOK() *PostDevicesOnusIDUnblockOK {
	return &PostDevicesOnusIDUnblockOK{}
}

/*PostDevicesOnusIDUnblockOK handles this case with default header values.

Successful
*/
type PostDevicesOnusIDUnblockOK struct {
	Payload *models.Status
}

func (o *PostDevicesOnusIDUnblockOK) Error() string {
	return fmt.Sprintf("[POST /devices/onus/{id}/unblock][%d] postDevicesOnusIdUnblockOK  %+v", 200, o.Payload)
}

func (o *PostDevicesOnusIDUnblockOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *PostDevicesOnusIDUnblockOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesOnusIDUnblockBadRequest creates a PostDevicesOnusIDUnblockBadRequest with default headers values
func NewPostDevicesOnusIDUnblockBadRequest() *PostDevicesOnusIDUnblockBadRequest {
	return &PostDevicesOnusIDUnblockBadRequest{}
}

/*PostDevicesOnusIDUnblockBadRequest handles this case with default header values.

Bad Request
*/
type PostDevicesOnusIDUnblockBadRequest struct {
	Payload *models.Error
}

func (o *PostDevicesOnusIDUnblockBadRequest) Error() string {
	return fmt.Sprintf("[POST /devices/onus/{id}/unblock][%d] postDevicesOnusIdUnblockBadRequest  %+v", 400, o.Payload)
}

func (o *PostDevicesOnusIDUnblockBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesOnusIDUnblockBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesOnusIDUnblockUnauthorized creates a PostDevicesOnusIDUnblockUnauthorized with default headers values
func NewPostDevicesOnusIDUnblockUnauthorized() *PostDevicesOnusIDUnblockUnauthorized {
	return &PostDevicesOnusIDUnblockUnauthorized{}
}

/*PostDevicesOnusIDUnblockUnauthorized handles this case with default header values.

Unauthorized
*/
type PostDevicesOnusIDUnblockUnauthorized struct {
	Payload *models.Error
}

func (o *PostDevicesOnusIDUnblockUnauthorized) Error() string {
	return fmt.Sprintf("[POST /devices/onus/{id}/unblock][%d] postDevicesOnusIdUnblockUnauthorized  %+v", 401, o.Payload)
}

func (o *PostDevicesOnusIDUnblockUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesOnusIDUnblockUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesOnusIDUnblockForbidden creates a PostDevicesOnusIDUnblockForbidden with default headers values
func NewPostDevicesOnusIDUnblockForbidden() *PostDevicesOnusIDUnblockForbidden {
	return &PostDevicesOnusIDUnblockForbidden{}
}

/*PostDevicesOnusIDUnblockForbidden handles this case with default header values.

Forbidden
*/
type PostDevicesOnusIDUnblockForbidden struct {
	Payload *models.Error
}

func (o *PostDevicesOnusIDUnblockForbidden) Error() string {
	return fmt.Sprintf("[POST /devices/onus/{id}/unblock][%d] postDevicesOnusIdUnblockForbidden  %+v", 403, o.Payload)
}

func (o *PostDevicesOnusIDUnblockForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesOnusIDUnblockForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesOnusIDUnblockNotFound creates a PostDevicesOnusIDUnblockNotFound with default headers values
func NewPostDevicesOnusIDUnblockNotFound() *PostDevicesOnusIDUnblockNotFound {
	return &PostDevicesOnusIDUnblockNotFound{}
}

/*PostDevicesOnusIDUnblockNotFound handles this case with default header values.

Not Found
*/
type PostDevicesOnusIDUnblockNotFound struct {
	Payload *models.Error
}

func (o *PostDevicesOnusIDUnblockNotFound) Error() string {
	return fmt.Sprintf("[POST /devices/onus/{id}/unblock][%d] postDevicesOnusIdUnblockNotFound  %+v", 404, o.Payload)
}

func (o *PostDevicesOnusIDUnblockNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesOnusIDUnblockNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesOnusIDUnblockInternalServerError creates a PostDevicesOnusIDUnblockInternalServerError with default headers values
func NewPostDevicesOnusIDUnblockInternalServerError() *PostDevicesOnusIDUnblockInternalServerError {
	return &PostDevicesOnusIDUnblockInternalServerError{}
}

/*PostDevicesOnusIDUnblockInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostDevicesOnusIDUnblockInternalServerError struct {
	Payload *models.Error
}

func (o *PostDevicesOnusIDUnblockInternalServerError) Error() string {
	return fmt.Sprintf("[POST /devices/onus/{id}/unblock][%d] postDevicesOnusIdUnblockInternalServerError  %+v", 500, o.Payload)
}

func (o *PostDevicesOnusIDUnblockInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesOnusIDUnblockInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
