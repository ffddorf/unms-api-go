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

// GetDevicesOnusIDWirelessReader is a Reader for the GetDevicesOnusIDWireless structure.
type GetDevicesOnusIDWirelessReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDevicesOnusIDWirelessReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDevicesOnusIDWirelessOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDevicesOnusIDWirelessBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDevicesOnusIDWirelessUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDevicesOnusIDWirelessForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDevicesOnusIDWirelessNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDevicesOnusIDWirelessInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewGetDevicesOnusIDWirelessNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDevicesOnusIDWirelessOK creates a GetDevicesOnusIDWirelessOK with default headers values
func NewGetDevicesOnusIDWirelessOK() *GetDevicesOnusIDWirelessOK {
	return &GetDevicesOnusIDWirelessOK{}
}

/*GetDevicesOnusIDWirelessOK handles this case with default header values.

Successful
*/
type GetDevicesOnusIDWirelessOK struct {
	Payload *models.OnuWireless
}

func (o *GetDevicesOnusIDWirelessOK) Error() string {
	return fmt.Sprintf("[GET /devices/onus/{id}/wireless][%d] getDevicesOnusIdWirelessOK  %+v", 200, o.Payload)
}

func (o *GetDevicesOnusIDWirelessOK) GetPayload() *models.OnuWireless {
	return o.Payload
}

func (o *GetDevicesOnusIDWirelessOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OnuWireless)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesOnusIDWirelessBadRequest creates a GetDevicesOnusIDWirelessBadRequest with default headers values
func NewGetDevicesOnusIDWirelessBadRequest() *GetDevicesOnusIDWirelessBadRequest {
	return &GetDevicesOnusIDWirelessBadRequest{}
}

/*GetDevicesOnusIDWirelessBadRequest handles this case with default header values.

Bad Request
*/
type GetDevicesOnusIDWirelessBadRequest struct {
	Payload *models.Error
}

func (o *GetDevicesOnusIDWirelessBadRequest) Error() string {
	return fmt.Sprintf("[GET /devices/onus/{id}/wireless][%d] getDevicesOnusIdWirelessBadRequest  %+v", 400, o.Payload)
}

func (o *GetDevicesOnusIDWirelessBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesOnusIDWirelessBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesOnusIDWirelessUnauthorized creates a GetDevicesOnusIDWirelessUnauthorized with default headers values
func NewGetDevicesOnusIDWirelessUnauthorized() *GetDevicesOnusIDWirelessUnauthorized {
	return &GetDevicesOnusIDWirelessUnauthorized{}
}

/*GetDevicesOnusIDWirelessUnauthorized handles this case with default header values.

Unauthorized
*/
type GetDevicesOnusIDWirelessUnauthorized struct {
	Payload *models.Error
}

func (o *GetDevicesOnusIDWirelessUnauthorized) Error() string {
	return fmt.Sprintf("[GET /devices/onus/{id}/wireless][%d] getDevicesOnusIdWirelessUnauthorized  %+v", 401, o.Payload)
}

func (o *GetDevicesOnusIDWirelessUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesOnusIDWirelessUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesOnusIDWirelessForbidden creates a GetDevicesOnusIDWirelessForbidden with default headers values
func NewGetDevicesOnusIDWirelessForbidden() *GetDevicesOnusIDWirelessForbidden {
	return &GetDevicesOnusIDWirelessForbidden{}
}

/*GetDevicesOnusIDWirelessForbidden handles this case with default header values.

Forbidden
*/
type GetDevicesOnusIDWirelessForbidden struct {
	Payload *models.Error
}

func (o *GetDevicesOnusIDWirelessForbidden) Error() string {
	return fmt.Sprintf("[GET /devices/onus/{id}/wireless][%d] getDevicesOnusIdWirelessForbidden  %+v", 403, o.Payload)
}

func (o *GetDevicesOnusIDWirelessForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesOnusIDWirelessForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesOnusIDWirelessNotFound creates a GetDevicesOnusIDWirelessNotFound with default headers values
func NewGetDevicesOnusIDWirelessNotFound() *GetDevicesOnusIDWirelessNotFound {
	return &GetDevicesOnusIDWirelessNotFound{}
}

/*GetDevicesOnusIDWirelessNotFound handles this case with default header values.

Not Found
*/
type GetDevicesOnusIDWirelessNotFound struct {
	Payload *models.Error
}

func (o *GetDevicesOnusIDWirelessNotFound) Error() string {
	return fmt.Sprintf("[GET /devices/onus/{id}/wireless][%d] getDevicesOnusIdWirelessNotFound  %+v", 404, o.Payload)
}

func (o *GetDevicesOnusIDWirelessNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesOnusIDWirelessNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesOnusIDWirelessInternalServerError creates a GetDevicesOnusIDWirelessInternalServerError with default headers values
func NewGetDevicesOnusIDWirelessInternalServerError() *GetDevicesOnusIDWirelessInternalServerError {
	return &GetDevicesOnusIDWirelessInternalServerError{}
}

/*GetDevicesOnusIDWirelessInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetDevicesOnusIDWirelessInternalServerError struct {
	Payload *models.Error
}

func (o *GetDevicesOnusIDWirelessInternalServerError) Error() string {
	return fmt.Sprintf("[GET /devices/onus/{id}/wireless][%d] getDevicesOnusIdWirelessInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDevicesOnusIDWirelessInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesOnusIDWirelessInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesOnusIDWirelessNotImplemented creates a GetDevicesOnusIDWirelessNotImplemented with default headers values
func NewGetDevicesOnusIDWirelessNotImplemented() *GetDevicesOnusIDWirelessNotImplemented {
	return &GetDevicesOnusIDWirelessNotImplemented{}
}

/*GetDevicesOnusIDWirelessNotImplemented handles this case with default header values.

Not Implemented
*/
type GetDevicesOnusIDWirelessNotImplemented struct {
	Payload *models.Error
}

func (o *GetDevicesOnusIDWirelessNotImplemented) Error() string {
	return fmt.Sprintf("[GET /devices/onus/{id}/wireless][%d] getDevicesOnusIdWirelessNotImplemented  %+v", 501, o.Payload)
}

func (o *GetDevicesOnusIDWirelessNotImplemented) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesOnusIDWirelessNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}