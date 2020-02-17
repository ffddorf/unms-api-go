// Code generated by go-swagger; DO NOT EDIT.

package simulation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/ffddorf/unms-api-go/models"
)

// PatchSimulationDevicesIDReader is a Reader for the PatchSimulationDevicesID structure.
type PatchSimulationDevicesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchSimulationDevicesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchSimulationDevicesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchSimulationDevicesIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchSimulationDevicesIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchSimulationDevicesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchSimulationDevicesIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchSimulationDevicesIDOK creates a PatchSimulationDevicesIDOK with default headers values
func NewPatchSimulationDevicesIDOK() *PatchSimulationDevicesIDOK {
	return &PatchSimulationDevicesIDOK{}
}

/*PatchSimulationDevicesIDOK handles this case with default header values.

Successful
*/
type PatchSimulationDevicesIDOK struct {
	Payload *models.SimulationDevice
}

func (o *PatchSimulationDevicesIDOK) Error() string {
	return fmt.Sprintf("[PATCH /simulation/devices/{id}][%d] patchSimulationDevicesIdOK  %+v", 200, o.Payload)
}

func (o *PatchSimulationDevicesIDOK) GetPayload() *models.SimulationDevice {
	return o.Payload
}

func (o *PatchSimulationDevicesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SimulationDevice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSimulationDevicesIDBadRequest creates a PatchSimulationDevicesIDBadRequest with default headers values
func NewPatchSimulationDevicesIDBadRequest() *PatchSimulationDevicesIDBadRequest {
	return &PatchSimulationDevicesIDBadRequest{}
}

/*PatchSimulationDevicesIDBadRequest handles this case with default header values.

Bad Request
*/
type PatchSimulationDevicesIDBadRequest struct {
	Payload *models.Error
}

func (o *PatchSimulationDevicesIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /simulation/devices/{id}][%d] patchSimulationDevicesIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchSimulationDevicesIDBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchSimulationDevicesIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSimulationDevicesIDUnauthorized creates a PatchSimulationDevicesIDUnauthorized with default headers values
func NewPatchSimulationDevicesIDUnauthorized() *PatchSimulationDevicesIDUnauthorized {
	return &PatchSimulationDevicesIDUnauthorized{}
}

/*PatchSimulationDevicesIDUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchSimulationDevicesIDUnauthorized struct {
	Payload *models.Error
}

func (o *PatchSimulationDevicesIDUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /simulation/devices/{id}][%d] patchSimulationDevicesIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchSimulationDevicesIDUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchSimulationDevicesIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSimulationDevicesIDNotFound creates a PatchSimulationDevicesIDNotFound with default headers values
func NewPatchSimulationDevicesIDNotFound() *PatchSimulationDevicesIDNotFound {
	return &PatchSimulationDevicesIDNotFound{}
}

/*PatchSimulationDevicesIDNotFound handles this case with default header values.

Not Found
*/
type PatchSimulationDevicesIDNotFound struct {
	Payload *models.Error
}

func (o *PatchSimulationDevicesIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /simulation/devices/{id}][%d] patchSimulationDevicesIdNotFound  %+v", 404, o.Payload)
}

func (o *PatchSimulationDevicesIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchSimulationDevicesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSimulationDevicesIDInternalServerError creates a PatchSimulationDevicesIDInternalServerError with default headers values
func NewPatchSimulationDevicesIDInternalServerError() *PatchSimulationDevicesIDInternalServerError {
	return &PatchSimulationDevicesIDInternalServerError{}
}

/*PatchSimulationDevicesIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type PatchSimulationDevicesIDInternalServerError struct {
	Payload *models.Error
}

func (o *PatchSimulationDevicesIDInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /simulation/devices/{id}][%d] patchSimulationDevicesIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchSimulationDevicesIDInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchSimulationDevicesIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
