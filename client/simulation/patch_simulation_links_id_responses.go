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

// PatchSimulationLinksIDReader is a Reader for the PatchSimulationLinksID structure.
type PatchSimulationLinksIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchSimulationLinksIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchSimulationLinksIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchSimulationLinksIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchSimulationLinksIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchSimulationLinksIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchSimulationLinksIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchSimulationLinksIDOK creates a PatchSimulationLinksIDOK with default headers values
func NewPatchSimulationLinksIDOK() *PatchSimulationLinksIDOK {
	return &PatchSimulationLinksIDOK{}
}

/*PatchSimulationLinksIDOK handles this case with default header values.

Successful
*/
type PatchSimulationLinksIDOK struct {
	Payload *models.SimulationLink
}

func (o *PatchSimulationLinksIDOK) Error() string {
	return fmt.Sprintf("[PATCH /simulation/links/{id}][%d] patchSimulationLinksIdOK  %+v", 200, o.Payload)
}

func (o *PatchSimulationLinksIDOK) GetPayload() *models.SimulationLink {
	return o.Payload
}

func (o *PatchSimulationLinksIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SimulationLink)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSimulationLinksIDBadRequest creates a PatchSimulationLinksIDBadRequest with default headers values
func NewPatchSimulationLinksIDBadRequest() *PatchSimulationLinksIDBadRequest {
	return &PatchSimulationLinksIDBadRequest{}
}

/*PatchSimulationLinksIDBadRequest handles this case with default header values.

Bad Request
*/
type PatchSimulationLinksIDBadRequest struct {
	Payload *models.Error
}

func (o *PatchSimulationLinksIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /simulation/links/{id}][%d] patchSimulationLinksIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchSimulationLinksIDBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchSimulationLinksIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSimulationLinksIDUnauthorized creates a PatchSimulationLinksIDUnauthorized with default headers values
func NewPatchSimulationLinksIDUnauthorized() *PatchSimulationLinksIDUnauthorized {
	return &PatchSimulationLinksIDUnauthorized{}
}

/*PatchSimulationLinksIDUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchSimulationLinksIDUnauthorized struct {
	Payload *models.Error
}

func (o *PatchSimulationLinksIDUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /simulation/links/{id}][%d] patchSimulationLinksIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchSimulationLinksIDUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchSimulationLinksIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSimulationLinksIDNotFound creates a PatchSimulationLinksIDNotFound with default headers values
func NewPatchSimulationLinksIDNotFound() *PatchSimulationLinksIDNotFound {
	return &PatchSimulationLinksIDNotFound{}
}

/*PatchSimulationLinksIDNotFound handles this case with default header values.

Not Found
*/
type PatchSimulationLinksIDNotFound struct {
	Payload *models.Error
}

func (o *PatchSimulationLinksIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /simulation/links/{id}][%d] patchSimulationLinksIdNotFound  %+v", 404, o.Payload)
}

func (o *PatchSimulationLinksIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchSimulationLinksIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSimulationLinksIDInternalServerError creates a PatchSimulationLinksIDInternalServerError with default headers values
func NewPatchSimulationLinksIDInternalServerError() *PatchSimulationLinksIDInternalServerError {
	return &PatchSimulationLinksIDInternalServerError{}
}

/*PatchSimulationLinksIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type PatchSimulationLinksIDInternalServerError struct {
	Payload *models.Error
}

func (o *PatchSimulationLinksIDInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /simulation/links/{id}][%d] patchSimulationLinksIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchSimulationLinksIDInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchSimulationLinksIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
