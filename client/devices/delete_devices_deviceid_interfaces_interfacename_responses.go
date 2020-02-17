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

// DeleteDevicesDeviceidInterfacesInterfacenameReader is a Reader for the DeleteDevicesDeviceidInterfacesInterfacename structure.
type DeleteDevicesDeviceidInterfacesInterfacenameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDevicesDeviceidInterfacesInterfacenameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDevicesDeviceidInterfacesInterfacenameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteDevicesDeviceidInterfacesInterfacenameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteDevicesDeviceidInterfacesInterfacenameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteDevicesDeviceidInterfacesInterfacenameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteDevicesDeviceidInterfacesInterfacenameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteDevicesDeviceidInterfacesInterfacenameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteDevicesDeviceidInterfacesInterfacenameOK creates a DeleteDevicesDeviceidInterfacesInterfacenameOK with default headers values
func NewDeleteDevicesDeviceidInterfacesInterfacenameOK() *DeleteDevicesDeviceidInterfacesInterfacenameOK {
	return &DeleteDevicesDeviceidInterfacesInterfacenameOK{}
}

/*DeleteDevicesDeviceidInterfacesInterfacenameOK handles this case with default header values.

Successful
*/
type DeleteDevicesDeviceidInterfacesInterfacenameOK struct {
	Payload *models.Status
}

func (o *DeleteDevicesDeviceidInterfacesInterfacenameOK) Error() string {
	return fmt.Sprintf("[DELETE /devices/{deviceId}/interfaces/{interfaceName}][%d] deleteDevicesDeviceidInterfacesInterfacenameOK  %+v", 200, o.Payload)
}

func (o *DeleteDevicesDeviceidInterfacesInterfacenameOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *DeleteDevicesDeviceidInterfacesInterfacenameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesDeviceidInterfacesInterfacenameBadRequest creates a DeleteDevicesDeviceidInterfacesInterfacenameBadRequest with default headers values
func NewDeleteDevicesDeviceidInterfacesInterfacenameBadRequest() *DeleteDevicesDeviceidInterfacesInterfacenameBadRequest {
	return &DeleteDevicesDeviceidInterfacesInterfacenameBadRequest{}
}

/*DeleteDevicesDeviceidInterfacesInterfacenameBadRequest handles this case with default header values.

Bad Request
*/
type DeleteDevicesDeviceidInterfacesInterfacenameBadRequest struct {
	Payload *models.Error
}

func (o *DeleteDevicesDeviceidInterfacesInterfacenameBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /devices/{deviceId}/interfaces/{interfaceName}][%d] deleteDevicesDeviceidInterfacesInterfacenameBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteDevicesDeviceidInterfacesInterfacenameBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesDeviceidInterfacesInterfacenameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesDeviceidInterfacesInterfacenameUnauthorized creates a DeleteDevicesDeviceidInterfacesInterfacenameUnauthorized with default headers values
func NewDeleteDevicesDeviceidInterfacesInterfacenameUnauthorized() *DeleteDevicesDeviceidInterfacesInterfacenameUnauthorized {
	return &DeleteDevicesDeviceidInterfacesInterfacenameUnauthorized{}
}

/*DeleteDevicesDeviceidInterfacesInterfacenameUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteDevicesDeviceidInterfacesInterfacenameUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteDevicesDeviceidInterfacesInterfacenameUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /devices/{deviceId}/interfaces/{interfaceName}][%d] deleteDevicesDeviceidInterfacesInterfacenameUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteDevicesDeviceidInterfacesInterfacenameUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesDeviceidInterfacesInterfacenameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesDeviceidInterfacesInterfacenameForbidden creates a DeleteDevicesDeviceidInterfacesInterfacenameForbidden with default headers values
func NewDeleteDevicesDeviceidInterfacesInterfacenameForbidden() *DeleteDevicesDeviceidInterfacesInterfacenameForbidden {
	return &DeleteDevicesDeviceidInterfacesInterfacenameForbidden{}
}

/*DeleteDevicesDeviceidInterfacesInterfacenameForbidden handles this case with default header values.

Forbidden
*/
type DeleteDevicesDeviceidInterfacesInterfacenameForbidden struct {
	Payload *models.Error
}

func (o *DeleteDevicesDeviceidInterfacesInterfacenameForbidden) Error() string {
	return fmt.Sprintf("[DELETE /devices/{deviceId}/interfaces/{interfaceName}][%d] deleteDevicesDeviceidInterfacesInterfacenameForbidden  %+v", 403, o.Payload)
}

func (o *DeleteDevicesDeviceidInterfacesInterfacenameForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesDeviceidInterfacesInterfacenameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesDeviceidInterfacesInterfacenameNotFound creates a DeleteDevicesDeviceidInterfacesInterfacenameNotFound with default headers values
func NewDeleteDevicesDeviceidInterfacesInterfacenameNotFound() *DeleteDevicesDeviceidInterfacesInterfacenameNotFound {
	return &DeleteDevicesDeviceidInterfacesInterfacenameNotFound{}
}

/*DeleteDevicesDeviceidInterfacesInterfacenameNotFound handles this case with default header values.

Not Found
*/
type DeleteDevicesDeviceidInterfacesInterfacenameNotFound struct {
	Payload *models.Error
}

func (o *DeleteDevicesDeviceidInterfacesInterfacenameNotFound) Error() string {
	return fmt.Sprintf("[DELETE /devices/{deviceId}/interfaces/{interfaceName}][%d] deleteDevicesDeviceidInterfacesInterfacenameNotFound  %+v", 404, o.Payload)
}

func (o *DeleteDevicesDeviceidInterfacesInterfacenameNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesDeviceidInterfacesInterfacenameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesDeviceidInterfacesInterfacenameInternalServerError creates a DeleteDevicesDeviceidInterfacesInterfacenameInternalServerError with default headers values
func NewDeleteDevicesDeviceidInterfacesInterfacenameInternalServerError() *DeleteDevicesDeviceidInterfacesInterfacenameInternalServerError {
	return &DeleteDevicesDeviceidInterfacesInterfacenameInternalServerError{}
}

/*DeleteDevicesDeviceidInterfacesInterfacenameInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteDevicesDeviceidInterfacesInterfacenameInternalServerError struct {
	Payload *models.Error
}

func (o *DeleteDevicesDeviceidInterfacesInterfacenameInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /devices/{deviceId}/interfaces/{interfaceName}][%d] deleteDevicesDeviceidInterfacesInterfacenameInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteDevicesDeviceidInterfacesInterfacenameInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesDeviceidInterfacesInterfacenameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
