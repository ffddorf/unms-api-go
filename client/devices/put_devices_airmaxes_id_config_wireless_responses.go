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

// PutDevicesAirmaxesIDConfigWirelessReader is a Reader for the PutDevicesAirmaxesIDConfigWireless structure.
type PutDevicesAirmaxesIDConfigWirelessReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutDevicesAirmaxesIDConfigWirelessReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutDevicesAirmaxesIDConfigWirelessOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutDevicesAirmaxesIDConfigWirelessBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutDevicesAirmaxesIDConfigWirelessUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutDevicesAirmaxesIDConfigWirelessForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutDevicesAirmaxesIDConfigWirelessNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutDevicesAirmaxesIDConfigWirelessInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutDevicesAirmaxesIDConfigWirelessOK creates a PutDevicesAirmaxesIDConfigWirelessOK with default headers values
func NewPutDevicesAirmaxesIDConfigWirelessOK() *PutDevicesAirmaxesIDConfigWirelessOK {
	return &PutDevicesAirmaxesIDConfigWirelessOK{}
}

/*PutDevicesAirmaxesIDConfigWirelessOK handles this case with default header values.

Successful
*/
type PutDevicesAirmaxesIDConfigWirelessOK struct {
	Payload *models.Status
}

func (o *PutDevicesAirmaxesIDConfigWirelessOK) Error() string {
	return fmt.Sprintf("[PUT /devices/airmaxes/{id}/config/wireless][%d] putDevicesAirmaxesIdConfigWirelessOK  %+v", 200, o.Payload)
}

func (o *PutDevicesAirmaxesIDConfigWirelessOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *PutDevicesAirmaxesIDConfigWirelessOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesAirmaxesIDConfigWirelessBadRequest creates a PutDevicesAirmaxesIDConfigWirelessBadRequest with default headers values
func NewPutDevicesAirmaxesIDConfigWirelessBadRequest() *PutDevicesAirmaxesIDConfigWirelessBadRequest {
	return &PutDevicesAirmaxesIDConfigWirelessBadRequest{}
}

/*PutDevicesAirmaxesIDConfigWirelessBadRequest handles this case with default header values.

Bad Request
*/
type PutDevicesAirmaxesIDConfigWirelessBadRequest struct {
	Payload *models.Error
}

func (o *PutDevicesAirmaxesIDConfigWirelessBadRequest) Error() string {
	return fmt.Sprintf("[PUT /devices/airmaxes/{id}/config/wireless][%d] putDevicesAirmaxesIdConfigWirelessBadRequest  %+v", 400, o.Payload)
}

func (o *PutDevicesAirmaxesIDConfigWirelessBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesAirmaxesIDConfigWirelessBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesAirmaxesIDConfigWirelessUnauthorized creates a PutDevicesAirmaxesIDConfigWirelessUnauthorized with default headers values
func NewPutDevicesAirmaxesIDConfigWirelessUnauthorized() *PutDevicesAirmaxesIDConfigWirelessUnauthorized {
	return &PutDevicesAirmaxesIDConfigWirelessUnauthorized{}
}

/*PutDevicesAirmaxesIDConfigWirelessUnauthorized handles this case with default header values.

Unauthorized
*/
type PutDevicesAirmaxesIDConfigWirelessUnauthorized struct {
	Payload *models.Error
}

func (o *PutDevicesAirmaxesIDConfigWirelessUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /devices/airmaxes/{id}/config/wireless][%d] putDevicesAirmaxesIdConfigWirelessUnauthorized  %+v", 401, o.Payload)
}

func (o *PutDevicesAirmaxesIDConfigWirelessUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesAirmaxesIDConfigWirelessUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesAirmaxesIDConfigWirelessForbidden creates a PutDevicesAirmaxesIDConfigWirelessForbidden with default headers values
func NewPutDevicesAirmaxesIDConfigWirelessForbidden() *PutDevicesAirmaxesIDConfigWirelessForbidden {
	return &PutDevicesAirmaxesIDConfigWirelessForbidden{}
}

/*PutDevicesAirmaxesIDConfigWirelessForbidden handles this case with default header values.

Forbidden
*/
type PutDevicesAirmaxesIDConfigWirelessForbidden struct {
	Payload *models.Error
}

func (o *PutDevicesAirmaxesIDConfigWirelessForbidden) Error() string {
	return fmt.Sprintf("[PUT /devices/airmaxes/{id}/config/wireless][%d] putDevicesAirmaxesIdConfigWirelessForbidden  %+v", 403, o.Payload)
}

func (o *PutDevicesAirmaxesIDConfigWirelessForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesAirmaxesIDConfigWirelessForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesAirmaxesIDConfigWirelessNotFound creates a PutDevicesAirmaxesIDConfigWirelessNotFound with default headers values
func NewPutDevicesAirmaxesIDConfigWirelessNotFound() *PutDevicesAirmaxesIDConfigWirelessNotFound {
	return &PutDevicesAirmaxesIDConfigWirelessNotFound{}
}

/*PutDevicesAirmaxesIDConfigWirelessNotFound handles this case with default header values.

Not Found
*/
type PutDevicesAirmaxesIDConfigWirelessNotFound struct {
	Payload *models.Error
}

func (o *PutDevicesAirmaxesIDConfigWirelessNotFound) Error() string {
	return fmt.Sprintf("[PUT /devices/airmaxes/{id}/config/wireless][%d] putDevicesAirmaxesIdConfigWirelessNotFound  %+v", 404, o.Payload)
}

func (o *PutDevicesAirmaxesIDConfigWirelessNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesAirmaxesIDConfigWirelessNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesAirmaxesIDConfigWirelessInternalServerError creates a PutDevicesAirmaxesIDConfigWirelessInternalServerError with default headers values
func NewPutDevicesAirmaxesIDConfigWirelessInternalServerError() *PutDevicesAirmaxesIDConfigWirelessInternalServerError {
	return &PutDevicesAirmaxesIDConfigWirelessInternalServerError{}
}

/*PutDevicesAirmaxesIDConfigWirelessInternalServerError handles this case with default header values.

Internal Server Error
*/
type PutDevicesAirmaxesIDConfigWirelessInternalServerError struct {
	Payload *models.Error
}

func (o *PutDevicesAirmaxesIDConfigWirelessInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /devices/airmaxes/{id}/config/wireless][%d] putDevicesAirmaxesIdConfigWirelessInternalServerError  %+v", 500, o.Payload)
}

func (o *PutDevicesAirmaxesIDConfigWirelessInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesAirmaxesIDConfigWirelessInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
