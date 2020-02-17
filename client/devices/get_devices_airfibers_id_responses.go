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

// GetDevicesAirfibersIDReader is a Reader for the GetDevicesAirfibersID structure.
type GetDevicesAirfibersIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDevicesAirfibersIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDevicesAirfibersIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDevicesAirfibersIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDevicesAirfibersIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDevicesAirfibersIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDevicesAirfibersIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDevicesAirfibersIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDevicesAirfibersIDOK creates a GetDevicesAirfibersIDOK with default headers values
func NewGetDevicesAirfibersIDOK() *GetDevicesAirfibersIDOK {
	return &GetDevicesAirfibersIDOK{}
}

/*GetDevicesAirfibersIDOK handles this case with default header values.

Successful
*/
type GetDevicesAirfibersIDOK struct {
	Payload *models.AirFiber
}

func (o *GetDevicesAirfibersIDOK) Error() string {
	return fmt.Sprintf("[GET /devices/airfibers/{id}][%d] getDevicesAirfibersIdOK  %+v", 200, o.Payload)
}

func (o *GetDevicesAirfibersIDOK) GetPayload() *models.AirFiber {
	return o.Payload
}

func (o *GetDevicesAirfibersIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AirFiber)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesAirfibersIDBadRequest creates a GetDevicesAirfibersIDBadRequest with default headers values
func NewGetDevicesAirfibersIDBadRequest() *GetDevicesAirfibersIDBadRequest {
	return &GetDevicesAirfibersIDBadRequest{}
}

/*GetDevicesAirfibersIDBadRequest handles this case with default header values.

Bad Request
*/
type GetDevicesAirfibersIDBadRequest struct {
	Payload *models.Error
}

func (o *GetDevicesAirfibersIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /devices/airfibers/{id}][%d] getDevicesAirfibersIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetDevicesAirfibersIDBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesAirfibersIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesAirfibersIDUnauthorized creates a GetDevicesAirfibersIDUnauthorized with default headers values
func NewGetDevicesAirfibersIDUnauthorized() *GetDevicesAirfibersIDUnauthorized {
	return &GetDevicesAirfibersIDUnauthorized{}
}

/*GetDevicesAirfibersIDUnauthorized handles this case with default header values.

Unauthorized
*/
type GetDevicesAirfibersIDUnauthorized struct {
	Payload *models.Error
}

func (o *GetDevicesAirfibersIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /devices/airfibers/{id}][%d] getDevicesAirfibersIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetDevicesAirfibersIDUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesAirfibersIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesAirfibersIDForbidden creates a GetDevicesAirfibersIDForbidden with default headers values
func NewGetDevicesAirfibersIDForbidden() *GetDevicesAirfibersIDForbidden {
	return &GetDevicesAirfibersIDForbidden{}
}

/*GetDevicesAirfibersIDForbidden handles this case with default header values.

Forbidden
*/
type GetDevicesAirfibersIDForbidden struct {
	Payload *models.Error
}

func (o *GetDevicesAirfibersIDForbidden) Error() string {
	return fmt.Sprintf("[GET /devices/airfibers/{id}][%d] getDevicesAirfibersIdForbidden  %+v", 403, o.Payload)
}

func (o *GetDevicesAirfibersIDForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesAirfibersIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesAirfibersIDNotFound creates a GetDevicesAirfibersIDNotFound with default headers values
func NewGetDevicesAirfibersIDNotFound() *GetDevicesAirfibersIDNotFound {
	return &GetDevicesAirfibersIDNotFound{}
}

/*GetDevicesAirfibersIDNotFound handles this case with default header values.

Not Found
*/
type GetDevicesAirfibersIDNotFound struct {
	Payload *models.Error
}

func (o *GetDevicesAirfibersIDNotFound) Error() string {
	return fmt.Sprintf("[GET /devices/airfibers/{id}][%d] getDevicesAirfibersIdNotFound  %+v", 404, o.Payload)
}

func (o *GetDevicesAirfibersIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesAirfibersIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesAirfibersIDInternalServerError creates a GetDevicesAirfibersIDInternalServerError with default headers values
func NewGetDevicesAirfibersIDInternalServerError() *GetDevicesAirfibersIDInternalServerError {
	return &GetDevicesAirfibersIDInternalServerError{}
}

/*GetDevicesAirfibersIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetDevicesAirfibersIDInternalServerError struct {
	Payload *models.Error
}

func (o *GetDevicesAirfibersIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /devices/airfibers/{id}][%d] getDevicesAirfibersIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDevicesAirfibersIDInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesAirfibersIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}