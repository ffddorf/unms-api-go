// Code generated by go-swagger; DO NOT EDIT.

package discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/ffddorf/unms-api-go/models"
)

// GetDiscoveryStatusDeviceidReader is a Reader for the GetDiscoveryStatusDeviceid structure.
type GetDiscoveryStatusDeviceidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDiscoveryStatusDeviceidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDiscoveryStatusDeviceidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDiscoveryStatusDeviceidBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDiscoveryStatusDeviceidUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDiscoveryStatusDeviceidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDiscoveryStatusDeviceidInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDiscoveryStatusDeviceidOK creates a GetDiscoveryStatusDeviceidOK with default headers values
func NewGetDiscoveryStatusDeviceidOK() *GetDiscoveryStatusDeviceidOK {
	return &GetDiscoveryStatusDeviceidOK{}
}

/*GetDiscoveryStatusDeviceidOK handles this case with default header values.

Successful
*/
type GetDiscoveryStatusDeviceidOK struct {
	Payload *models.Model16
}

func (o *GetDiscoveryStatusDeviceidOK) Error() string {
	return fmt.Sprintf("[GET /discovery/status/{deviceId}][%d] getDiscoveryStatusDeviceidOK  %+v", 200, o.Payload)
}

func (o *GetDiscoveryStatusDeviceidOK) GetPayload() *models.Model16 {
	return o.Payload
}

func (o *GetDiscoveryStatusDeviceidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Model16)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDiscoveryStatusDeviceidBadRequest creates a GetDiscoveryStatusDeviceidBadRequest with default headers values
func NewGetDiscoveryStatusDeviceidBadRequest() *GetDiscoveryStatusDeviceidBadRequest {
	return &GetDiscoveryStatusDeviceidBadRequest{}
}

/*GetDiscoveryStatusDeviceidBadRequest handles this case with default header values.

Bad Request
*/
type GetDiscoveryStatusDeviceidBadRequest struct {
	Payload *models.Error
}

func (o *GetDiscoveryStatusDeviceidBadRequest) Error() string {
	return fmt.Sprintf("[GET /discovery/status/{deviceId}][%d] getDiscoveryStatusDeviceidBadRequest  %+v", 400, o.Payload)
}

func (o *GetDiscoveryStatusDeviceidBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDiscoveryStatusDeviceidBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDiscoveryStatusDeviceidUnauthorized creates a GetDiscoveryStatusDeviceidUnauthorized with default headers values
func NewGetDiscoveryStatusDeviceidUnauthorized() *GetDiscoveryStatusDeviceidUnauthorized {
	return &GetDiscoveryStatusDeviceidUnauthorized{}
}

/*GetDiscoveryStatusDeviceidUnauthorized handles this case with default header values.

Unauthorized
*/
type GetDiscoveryStatusDeviceidUnauthorized struct {
	Payload *models.Error
}

func (o *GetDiscoveryStatusDeviceidUnauthorized) Error() string {
	return fmt.Sprintf("[GET /discovery/status/{deviceId}][%d] getDiscoveryStatusDeviceidUnauthorized  %+v", 401, o.Payload)
}

func (o *GetDiscoveryStatusDeviceidUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDiscoveryStatusDeviceidUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDiscoveryStatusDeviceidNotFound creates a GetDiscoveryStatusDeviceidNotFound with default headers values
func NewGetDiscoveryStatusDeviceidNotFound() *GetDiscoveryStatusDeviceidNotFound {
	return &GetDiscoveryStatusDeviceidNotFound{}
}

/*GetDiscoveryStatusDeviceidNotFound handles this case with default header values.

Not Found
*/
type GetDiscoveryStatusDeviceidNotFound struct {
	Payload *models.Error
}

func (o *GetDiscoveryStatusDeviceidNotFound) Error() string {
	return fmt.Sprintf("[GET /discovery/status/{deviceId}][%d] getDiscoveryStatusDeviceidNotFound  %+v", 404, o.Payload)
}

func (o *GetDiscoveryStatusDeviceidNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDiscoveryStatusDeviceidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDiscoveryStatusDeviceidInternalServerError creates a GetDiscoveryStatusDeviceidInternalServerError with default headers values
func NewGetDiscoveryStatusDeviceidInternalServerError() *GetDiscoveryStatusDeviceidInternalServerError {
	return &GetDiscoveryStatusDeviceidInternalServerError{}
}

/*GetDiscoveryStatusDeviceidInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetDiscoveryStatusDeviceidInternalServerError struct {
	Payload *models.Error
}

func (o *GetDiscoveryStatusDeviceidInternalServerError) Error() string {
	return fmt.Sprintf("[GET /discovery/status/{deviceId}][%d] getDiscoveryStatusDeviceidInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDiscoveryStatusDeviceidInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDiscoveryStatusDeviceidInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
