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

// PostDevicesOnusIDResetstatsReader is a Reader for the PostDevicesOnusIDResetstats structure.
type PostDevicesOnusIDResetstatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDevicesOnusIDResetstatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDevicesOnusIDResetstatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostDevicesOnusIDResetstatsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostDevicesOnusIDResetstatsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostDevicesOnusIDResetstatsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostDevicesOnusIDResetstatsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostDevicesOnusIDResetstatsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostDevicesOnusIDResetstatsOK creates a PostDevicesOnusIDResetstatsOK with default headers values
func NewPostDevicesOnusIDResetstatsOK() *PostDevicesOnusIDResetstatsOK {
	return &PostDevicesOnusIDResetstatsOK{}
}

/*PostDevicesOnusIDResetstatsOK handles this case with default header values.

Successful
*/
type PostDevicesOnusIDResetstatsOK struct {
	Payload *models.Status
}

func (o *PostDevicesOnusIDResetstatsOK) Error() string {
	return fmt.Sprintf("[POST /devices/onus/{id}/resetstats][%d] postDevicesOnusIdResetstatsOK  %+v", 200, o.Payload)
}

func (o *PostDevicesOnusIDResetstatsOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *PostDevicesOnusIDResetstatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesOnusIDResetstatsBadRequest creates a PostDevicesOnusIDResetstatsBadRequest with default headers values
func NewPostDevicesOnusIDResetstatsBadRequest() *PostDevicesOnusIDResetstatsBadRequest {
	return &PostDevicesOnusIDResetstatsBadRequest{}
}

/*PostDevicesOnusIDResetstatsBadRequest handles this case with default header values.

Bad Request
*/
type PostDevicesOnusIDResetstatsBadRequest struct {
	Payload *models.Error
}

func (o *PostDevicesOnusIDResetstatsBadRequest) Error() string {
	return fmt.Sprintf("[POST /devices/onus/{id}/resetstats][%d] postDevicesOnusIdResetstatsBadRequest  %+v", 400, o.Payload)
}

func (o *PostDevicesOnusIDResetstatsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesOnusIDResetstatsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesOnusIDResetstatsUnauthorized creates a PostDevicesOnusIDResetstatsUnauthorized with default headers values
func NewPostDevicesOnusIDResetstatsUnauthorized() *PostDevicesOnusIDResetstatsUnauthorized {
	return &PostDevicesOnusIDResetstatsUnauthorized{}
}

/*PostDevicesOnusIDResetstatsUnauthorized handles this case with default header values.

Unauthorized
*/
type PostDevicesOnusIDResetstatsUnauthorized struct {
	Payload *models.Error
}

func (o *PostDevicesOnusIDResetstatsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /devices/onus/{id}/resetstats][%d] postDevicesOnusIdResetstatsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostDevicesOnusIDResetstatsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesOnusIDResetstatsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesOnusIDResetstatsForbidden creates a PostDevicesOnusIDResetstatsForbidden with default headers values
func NewPostDevicesOnusIDResetstatsForbidden() *PostDevicesOnusIDResetstatsForbidden {
	return &PostDevicesOnusIDResetstatsForbidden{}
}

/*PostDevicesOnusIDResetstatsForbidden handles this case with default header values.

Forbidden
*/
type PostDevicesOnusIDResetstatsForbidden struct {
	Payload *models.Error
}

func (o *PostDevicesOnusIDResetstatsForbidden) Error() string {
	return fmt.Sprintf("[POST /devices/onus/{id}/resetstats][%d] postDevicesOnusIdResetstatsForbidden  %+v", 403, o.Payload)
}

func (o *PostDevicesOnusIDResetstatsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesOnusIDResetstatsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesOnusIDResetstatsNotFound creates a PostDevicesOnusIDResetstatsNotFound with default headers values
func NewPostDevicesOnusIDResetstatsNotFound() *PostDevicesOnusIDResetstatsNotFound {
	return &PostDevicesOnusIDResetstatsNotFound{}
}

/*PostDevicesOnusIDResetstatsNotFound handles this case with default header values.

Not Found
*/
type PostDevicesOnusIDResetstatsNotFound struct {
	Payload *models.Error
}

func (o *PostDevicesOnusIDResetstatsNotFound) Error() string {
	return fmt.Sprintf("[POST /devices/onus/{id}/resetstats][%d] postDevicesOnusIdResetstatsNotFound  %+v", 404, o.Payload)
}

func (o *PostDevicesOnusIDResetstatsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesOnusIDResetstatsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesOnusIDResetstatsInternalServerError creates a PostDevicesOnusIDResetstatsInternalServerError with default headers values
func NewPostDevicesOnusIDResetstatsInternalServerError() *PostDevicesOnusIDResetstatsInternalServerError {
	return &PostDevicesOnusIDResetstatsInternalServerError{}
}

/*PostDevicesOnusIDResetstatsInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostDevicesOnusIDResetstatsInternalServerError struct {
	Payload *models.Error
}

func (o *PostDevicesOnusIDResetstatsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /devices/onus/{id}/resetstats][%d] postDevicesOnusIdResetstatsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostDevicesOnusIDResetstatsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesOnusIDResetstatsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
