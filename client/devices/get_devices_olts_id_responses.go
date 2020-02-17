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

// GetDevicesOltsIDReader is a Reader for the GetDevicesOltsID structure.
type GetDevicesOltsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDevicesOltsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDevicesOltsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDevicesOltsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDevicesOltsIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDevicesOltsIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDevicesOltsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDevicesOltsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDevicesOltsIDOK creates a GetDevicesOltsIDOK with default headers values
func NewGetDevicesOltsIDOK() *GetDevicesOltsIDOK {
	return &GetDevicesOltsIDOK{}
}

/*GetDevicesOltsIDOK handles this case with default header values.

Successful
*/
type GetDevicesOltsIDOK struct {
	Payload *models.Olt
}

func (o *GetDevicesOltsIDOK) Error() string {
	return fmt.Sprintf("[GET /devices/olts/{id}][%d] getDevicesOltsIdOK  %+v", 200, o.Payload)
}

func (o *GetDevicesOltsIDOK) GetPayload() *models.Olt {
	return o.Payload
}

func (o *GetDevicesOltsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Olt)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesOltsIDBadRequest creates a GetDevicesOltsIDBadRequest with default headers values
func NewGetDevicesOltsIDBadRequest() *GetDevicesOltsIDBadRequest {
	return &GetDevicesOltsIDBadRequest{}
}

/*GetDevicesOltsIDBadRequest handles this case with default header values.

Bad Request
*/
type GetDevicesOltsIDBadRequest struct {
	Payload *models.Error
}

func (o *GetDevicesOltsIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /devices/olts/{id}][%d] getDevicesOltsIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetDevicesOltsIDBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesOltsIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesOltsIDUnauthorized creates a GetDevicesOltsIDUnauthorized with default headers values
func NewGetDevicesOltsIDUnauthorized() *GetDevicesOltsIDUnauthorized {
	return &GetDevicesOltsIDUnauthorized{}
}

/*GetDevicesOltsIDUnauthorized handles this case with default header values.

Unauthorized
*/
type GetDevicesOltsIDUnauthorized struct {
	Payload *models.Error
}

func (o *GetDevicesOltsIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /devices/olts/{id}][%d] getDevicesOltsIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetDevicesOltsIDUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesOltsIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesOltsIDForbidden creates a GetDevicesOltsIDForbidden with default headers values
func NewGetDevicesOltsIDForbidden() *GetDevicesOltsIDForbidden {
	return &GetDevicesOltsIDForbidden{}
}

/*GetDevicesOltsIDForbidden handles this case with default header values.

Forbidden
*/
type GetDevicesOltsIDForbidden struct {
	Payload *models.Error
}

func (o *GetDevicesOltsIDForbidden) Error() string {
	return fmt.Sprintf("[GET /devices/olts/{id}][%d] getDevicesOltsIdForbidden  %+v", 403, o.Payload)
}

func (o *GetDevicesOltsIDForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesOltsIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesOltsIDNotFound creates a GetDevicesOltsIDNotFound with default headers values
func NewGetDevicesOltsIDNotFound() *GetDevicesOltsIDNotFound {
	return &GetDevicesOltsIDNotFound{}
}

/*GetDevicesOltsIDNotFound handles this case with default header values.

Not Found
*/
type GetDevicesOltsIDNotFound struct {
	Payload *models.Error
}

func (o *GetDevicesOltsIDNotFound) Error() string {
	return fmt.Sprintf("[GET /devices/olts/{id}][%d] getDevicesOltsIdNotFound  %+v", 404, o.Payload)
}

func (o *GetDevicesOltsIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesOltsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesOltsIDInternalServerError creates a GetDevicesOltsIDInternalServerError with default headers values
func NewGetDevicesOltsIDInternalServerError() *GetDevicesOltsIDInternalServerError {
	return &GetDevicesOltsIDInternalServerError{}
}

/*GetDevicesOltsIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetDevicesOltsIDInternalServerError struct {
	Payload *models.Error
}

func (o *GetDevicesOltsIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /devices/olts/{id}][%d] getDevicesOltsIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDevicesOltsIDInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesOltsIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
