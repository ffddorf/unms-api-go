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

// PutDevicesEroutersIDRouterRoutesReader is a Reader for the PutDevicesEroutersIDRouterRoutes structure.
type PutDevicesEroutersIDRouterRoutesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutDevicesEroutersIDRouterRoutesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutDevicesEroutersIDRouterRoutesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutDevicesEroutersIDRouterRoutesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutDevicesEroutersIDRouterRoutesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutDevicesEroutersIDRouterRoutesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutDevicesEroutersIDRouterRoutesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutDevicesEroutersIDRouterRoutesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutDevicesEroutersIDRouterRoutesOK creates a PutDevicesEroutersIDRouterRoutesOK with default headers values
func NewPutDevicesEroutersIDRouterRoutesOK() *PutDevicesEroutersIDRouterRoutesOK {
	return &PutDevicesEroutersIDRouterRoutesOK{}
}

/*PutDevicesEroutersIDRouterRoutesOK handles this case with default header values.

Successful
*/
type PutDevicesEroutersIDRouterRoutesOK struct {
	Payload *models.EdgeRouterRoute
}

func (o *PutDevicesEroutersIDRouterRoutesOK) Error() string {
	return fmt.Sprintf("[PUT /devices/erouters/{id}/router/routes][%d] putDevicesEroutersIdRouterRoutesOK  %+v", 200, o.Payload)
}

func (o *PutDevicesEroutersIDRouterRoutesOK) GetPayload() *models.EdgeRouterRoute {
	return o.Payload
}

func (o *PutDevicesEroutersIDRouterRoutesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EdgeRouterRoute)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesEroutersIDRouterRoutesBadRequest creates a PutDevicesEroutersIDRouterRoutesBadRequest with default headers values
func NewPutDevicesEroutersIDRouterRoutesBadRequest() *PutDevicesEroutersIDRouterRoutesBadRequest {
	return &PutDevicesEroutersIDRouterRoutesBadRequest{}
}

/*PutDevicesEroutersIDRouterRoutesBadRequest handles this case with default header values.

Bad Request
*/
type PutDevicesEroutersIDRouterRoutesBadRequest struct {
	Payload *models.Error
}

func (o *PutDevicesEroutersIDRouterRoutesBadRequest) Error() string {
	return fmt.Sprintf("[PUT /devices/erouters/{id}/router/routes][%d] putDevicesEroutersIdRouterRoutesBadRequest  %+v", 400, o.Payload)
}

func (o *PutDevicesEroutersIDRouterRoutesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesEroutersIDRouterRoutesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesEroutersIDRouterRoutesUnauthorized creates a PutDevicesEroutersIDRouterRoutesUnauthorized with default headers values
func NewPutDevicesEroutersIDRouterRoutesUnauthorized() *PutDevicesEroutersIDRouterRoutesUnauthorized {
	return &PutDevicesEroutersIDRouterRoutesUnauthorized{}
}

/*PutDevicesEroutersIDRouterRoutesUnauthorized handles this case with default header values.

Unauthorized
*/
type PutDevicesEroutersIDRouterRoutesUnauthorized struct {
	Payload *models.Error
}

func (o *PutDevicesEroutersIDRouterRoutesUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /devices/erouters/{id}/router/routes][%d] putDevicesEroutersIdRouterRoutesUnauthorized  %+v", 401, o.Payload)
}

func (o *PutDevicesEroutersIDRouterRoutesUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesEroutersIDRouterRoutesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesEroutersIDRouterRoutesForbidden creates a PutDevicesEroutersIDRouterRoutesForbidden with default headers values
func NewPutDevicesEroutersIDRouterRoutesForbidden() *PutDevicesEroutersIDRouterRoutesForbidden {
	return &PutDevicesEroutersIDRouterRoutesForbidden{}
}

/*PutDevicesEroutersIDRouterRoutesForbidden handles this case with default header values.

Forbidden
*/
type PutDevicesEroutersIDRouterRoutesForbidden struct {
	Payload *models.Error
}

func (o *PutDevicesEroutersIDRouterRoutesForbidden) Error() string {
	return fmt.Sprintf("[PUT /devices/erouters/{id}/router/routes][%d] putDevicesEroutersIdRouterRoutesForbidden  %+v", 403, o.Payload)
}

func (o *PutDevicesEroutersIDRouterRoutesForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesEroutersIDRouterRoutesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesEroutersIDRouterRoutesNotFound creates a PutDevicesEroutersIDRouterRoutesNotFound with default headers values
func NewPutDevicesEroutersIDRouterRoutesNotFound() *PutDevicesEroutersIDRouterRoutesNotFound {
	return &PutDevicesEroutersIDRouterRoutesNotFound{}
}

/*PutDevicesEroutersIDRouterRoutesNotFound handles this case with default header values.

Not Found
*/
type PutDevicesEroutersIDRouterRoutesNotFound struct {
	Payload *models.Error
}

func (o *PutDevicesEroutersIDRouterRoutesNotFound) Error() string {
	return fmt.Sprintf("[PUT /devices/erouters/{id}/router/routes][%d] putDevicesEroutersIdRouterRoutesNotFound  %+v", 404, o.Payload)
}

func (o *PutDevicesEroutersIDRouterRoutesNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesEroutersIDRouterRoutesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesEroutersIDRouterRoutesInternalServerError creates a PutDevicesEroutersIDRouterRoutesInternalServerError with default headers values
func NewPutDevicesEroutersIDRouterRoutesInternalServerError() *PutDevicesEroutersIDRouterRoutesInternalServerError {
	return &PutDevicesEroutersIDRouterRoutesInternalServerError{}
}

/*PutDevicesEroutersIDRouterRoutesInternalServerError handles this case with default header values.

Internal Server Error
*/
type PutDevicesEroutersIDRouterRoutesInternalServerError struct {
	Payload *models.Error
}

func (o *PutDevicesEroutersIDRouterRoutesInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /devices/erouters/{id}/router/routes][%d] putDevicesEroutersIdRouterRoutesInternalServerError  %+v", 500, o.Payload)
}

func (o *PutDevicesEroutersIDRouterRoutesInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesEroutersIDRouterRoutesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
