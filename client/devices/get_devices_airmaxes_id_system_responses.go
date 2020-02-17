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

// GetDevicesAirmaxesIDSystemReader is a Reader for the GetDevicesAirmaxesIDSystem structure.
type GetDevicesAirmaxesIDSystemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDevicesAirmaxesIDSystemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDevicesAirmaxesIDSystemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDevicesAirmaxesIDSystemBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDevicesAirmaxesIDSystemUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDevicesAirmaxesIDSystemForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDevicesAirmaxesIDSystemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDevicesAirmaxesIDSystemInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDevicesAirmaxesIDSystemOK creates a GetDevicesAirmaxesIDSystemOK with default headers values
func NewGetDevicesAirmaxesIDSystemOK() *GetDevicesAirmaxesIDSystemOK {
	return &GetDevicesAirmaxesIDSystemOK{}
}

/*GetDevicesAirmaxesIDSystemOK handles this case with default header values.

Successful
*/
type GetDevicesAirmaxesIDSystemOK struct {
	Payload *models.Model23
}

func (o *GetDevicesAirmaxesIDSystemOK) Error() string {
	return fmt.Sprintf("[GET /devices/airmaxes/{id}/system][%d] getDevicesAirmaxesIdSystemOK  %+v", 200, o.Payload)
}

func (o *GetDevicesAirmaxesIDSystemOK) GetPayload() *models.Model23 {
	return o.Payload
}

func (o *GetDevicesAirmaxesIDSystemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Model23)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesAirmaxesIDSystemBadRequest creates a GetDevicesAirmaxesIDSystemBadRequest with default headers values
func NewGetDevicesAirmaxesIDSystemBadRequest() *GetDevicesAirmaxesIDSystemBadRequest {
	return &GetDevicesAirmaxesIDSystemBadRequest{}
}

/*GetDevicesAirmaxesIDSystemBadRequest handles this case with default header values.

Bad Request
*/
type GetDevicesAirmaxesIDSystemBadRequest struct {
	Payload *models.Error
}

func (o *GetDevicesAirmaxesIDSystemBadRequest) Error() string {
	return fmt.Sprintf("[GET /devices/airmaxes/{id}/system][%d] getDevicesAirmaxesIdSystemBadRequest  %+v", 400, o.Payload)
}

func (o *GetDevicesAirmaxesIDSystemBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesAirmaxesIDSystemBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesAirmaxesIDSystemUnauthorized creates a GetDevicesAirmaxesIDSystemUnauthorized with default headers values
func NewGetDevicesAirmaxesIDSystemUnauthorized() *GetDevicesAirmaxesIDSystemUnauthorized {
	return &GetDevicesAirmaxesIDSystemUnauthorized{}
}

/*GetDevicesAirmaxesIDSystemUnauthorized handles this case with default header values.

Unauthorized
*/
type GetDevicesAirmaxesIDSystemUnauthorized struct {
	Payload *models.Error
}

func (o *GetDevicesAirmaxesIDSystemUnauthorized) Error() string {
	return fmt.Sprintf("[GET /devices/airmaxes/{id}/system][%d] getDevicesAirmaxesIdSystemUnauthorized  %+v", 401, o.Payload)
}

func (o *GetDevicesAirmaxesIDSystemUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesAirmaxesIDSystemUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesAirmaxesIDSystemForbidden creates a GetDevicesAirmaxesIDSystemForbidden with default headers values
func NewGetDevicesAirmaxesIDSystemForbidden() *GetDevicesAirmaxesIDSystemForbidden {
	return &GetDevicesAirmaxesIDSystemForbidden{}
}

/*GetDevicesAirmaxesIDSystemForbidden handles this case with default header values.

Forbidden
*/
type GetDevicesAirmaxesIDSystemForbidden struct {
	Payload *models.Error
}

func (o *GetDevicesAirmaxesIDSystemForbidden) Error() string {
	return fmt.Sprintf("[GET /devices/airmaxes/{id}/system][%d] getDevicesAirmaxesIdSystemForbidden  %+v", 403, o.Payload)
}

func (o *GetDevicesAirmaxesIDSystemForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesAirmaxesIDSystemForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesAirmaxesIDSystemNotFound creates a GetDevicesAirmaxesIDSystemNotFound with default headers values
func NewGetDevicesAirmaxesIDSystemNotFound() *GetDevicesAirmaxesIDSystemNotFound {
	return &GetDevicesAirmaxesIDSystemNotFound{}
}

/*GetDevicesAirmaxesIDSystemNotFound handles this case with default header values.

Not Found
*/
type GetDevicesAirmaxesIDSystemNotFound struct {
	Payload *models.Error
}

func (o *GetDevicesAirmaxesIDSystemNotFound) Error() string {
	return fmt.Sprintf("[GET /devices/airmaxes/{id}/system][%d] getDevicesAirmaxesIdSystemNotFound  %+v", 404, o.Payload)
}

func (o *GetDevicesAirmaxesIDSystemNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesAirmaxesIDSystemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesAirmaxesIDSystemInternalServerError creates a GetDevicesAirmaxesIDSystemInternalServerError with default headers values
func NewGetDevicesAirmaxesIDSystemInternalServerError() *GetDevicesAirmaxesIDSystemInternalServerError {
	return &GetDevicesAirmaxesIDSystemInternalServerError{}
}

/*GetDevicesAirmaxesIDSystemInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetDevicesAirmaxesIDSystemInternalServerError struct {
	Payload *models.Error
}

func (o *GetDevicesAirmaxesIDSystemInternalServerError) Error() string {
	return fmt.Sprintf("[GET /devices/airmaxes/{id}/system][%d] getDevicesAirmaxesIdSystemInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDevicesAirmaxesIDSystemInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesAirmaxesIDSystemInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
