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

// PutDevicesAircubesIDConfigSystemReader is a Reader for the PutDevicesAircubesIDConfigSystem structure.
type PutDevicesAircubesIDConfigSystemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutDevicesAircubesIDConfigSystemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutDevicesAircubesIDConfigSystemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutDevicesAircubesIDConfigSystemBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutDevicesAircubesIDConfigSystemUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutDevicesAircubesIDConfigSystemForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutDevicesAircubesIDConfigSystemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutDevicesAircubesIDConfigSystemInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutDevicesAircubesIDConfigSystemOK creates a PutDevicesAircubesIDConfigSystemOK with default headers values
func NewPutDevicesAircubesIDConfigSystemOK() *PutDevicesAircubesIDConfigSystemOK {
	return &PutDevicesAircubesIDConfigSystemOK{}
}

/*PutDevicesAircubesIDConfigSystemOK handles this case with default header values.

Successful
*/
type PutDevicesAircubesIDConfigSystemOK struct {
	Payload *models.Status
}

func (o *PutDevicesAircubesIDConfigSystemOK) Error() string {
	return fmt.Sprintf("[PUT /devices/aircubes/{id}/config/system][%d] putDevicesAircubesIdConfigSystemOK  %+v", 200, o.Payload)
}

func (o *PutDevicesAircubesIDConfigSystemOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *PutDevicesAircubesIDConfigSystemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesAircubesIDConfigSystemBadRequest creates a PutDevicesAircubesIDConfigSystemBadRequest with default headers values
func NewPutDevicesAircubesIDConfigSystemBadRequest() *PutDevicesAircubesIDConfigSystemBadRequest {
	return &PutDevicesAircubesIDConfigSystemBadRequest{}
}

/*PutDevicesAircubesIDConfigSystemBadRequest handles this case with default header values.

Bad Request
*/
type PutDevicesAircubesIDConfigSystemBadRequest struct {
	Payload *models.Error
}

func (o *PutDevicesAircubesIDConfigSystemBadRequest) Error() string {
	return fmt.Sprintf("[PUT /devices/aircubes/{id}/config/system][%d] putDevicesAircubesIdConfigSystemBadRequest  %+v", 400, o.Payload)
}

func (o *PutDevicesAircubesIDConfigSystemBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesAircubesIDConfigSystemBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesAircubesIDConfigSystemUnauthorized creates a PutDevicesAircubesIDConfigSystemUnauthorized with default headers values
func NewPutDevicesAircubesIDConfigSystemUnauthorized() *PutDevicesAircubesIDConfigSystemUnauthorized {
	return &PutDevicesAircubesIDConfigSystemUnauthorized{}
}

/*PutDevicesAircubesIDConfigSystemUnauthorized handles this case with default header values.

Unauthorized
*/
type PutDevicesAircubesIDConfigSystemUnauthorized struct {
	Payload *models.Error
}

func (o *PutDevicesAircubesIDConfigSystemUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /devices/aircubes/{id}/config/system][%d] putDevicesAircubesIdConfigSystemUnauthorized  %+v", 401, o.Payload)
}

func (o *PutDevicesAircubesIDConfigSystemUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesAircubesIDConfigSystemUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesAircubesIDConfigSystemForbidden creates a PutDevicesAircubesIDConfigSystemForbidden with default headers values
func NewPutDevicesAircubesIDConfigSystemForbidden() *PutDevicesAircubesIDConfigSystemForbidden {
	return &PutDevicesAircubesIDConfigSystemForbidden{}
}

/*PutDevicesAircubesIDConfigSystemForbidden handles this case with default header values.

Forbidden
*/
type PutDevicesAircubesIDConfigSystemForbidden struct {
	Payload *models.Error
}

func (o *PutDevicesAircubesIDConfigSystemForbidden) Error() string {
	return fmt.Sprintf("[PUT /devices/aircubes/{id}/config/system][%d] putDevicesAircubesIdConfigSystemForbidden  %+v", 403, o.Payload)
}

func (o *PutDevicesAircubesIDConfigSystemForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesAircubesIDConfigSystemForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesAircubesIDConfigSystemNotFound creates a PutDevicesAircubesIDConfigSystemNotFound with default headers values
func NewPutDevicesAircubesIDConfigSystemNotFound() *PutDevicesAircubesIDConfigSystemNotFound {
	return &PutDevicesAircubesIDConfigSystemNotFound{}
}

/*PutDevicesAircubesIDConfigSystemNotFound handles this case with default header values.

Not Found
*/
type PutDevicesAircubesIDConfigSystemNotFound struct {
	Payload *models.Error
}

func (o *PutDevicesAircubesIDConfigSystemNotFound) Error() string {
	return fmt.Sprintf("[PUT /devices/aircubes/{id}/config/system][%d] putDevicesAircubesIdConfigSystemNotFound  %+v", 404, o.Payload)
}

func (o *PutDevicesAircubesIDConfigSystemNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesAircubesIDConfigSystemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesAircubesIDConfigSystemInternalServerError creates a PutDevicesAircubesIDConfigSystemInternalServerError with default headers values
func NewPutDevicesAircubesIDConfigSystemInternalServerError() *PutDevicesAircubesIDConfigSystemInternalServerError {
	return &PutDevicesAircubesIDConfigSystemInternalServerError{}
}

/*PutDevicesAircubesIDConfigSystemInternalServerError handles this case with default header values.

Internal Server Error
*/
type PutDevicesAircubesIDConfigSystemInternalServerError struct {
	Payload *models.Error
}

func (o *PutDevicesAircubesIDConfigSystemInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /devices/aircubes/{id}/config/system][%d] putDevicesAircubesIdConfigSystemInternalServerError  %+v", 500, o.Payload)
}

func (o *PutDevicesAircubesIDConfigSystemInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesAircubesIDConfigSystemInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
