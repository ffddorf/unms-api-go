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

// PostDevicesIDRefreshReader is a Reader for the PostDevicesIDRefresh structure.
type PostDevicesIDRefreshReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDevicesIDRefreshReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDevicesIDRefreshOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostDevicesIDRefreshBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostDevicesIDRefreshUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostDevicesIDRefreshForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostDevicesIDRefreshInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostDevicesIDRefreshOK creates a PostDevicesIDRefreshOK with default headers values
func NewPostDevicesIDRefreshOK() *PostDevicesIDRefreshOK {
	return &PostDevicesIDRefreshOK{}
}

/*PostDevicesIDRefreshOK handles this case with default header values.

Successful
*/
type PostDevicesIDRefreshOK struct {
	Payload *models.Status
}

func (o *PostDevicesIDRefreshOK) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/refresh][%d] postDevicesIdRefreshOK  %+v", 200, o.Payload)
}

func (o *PostDevicesIDRefreshOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *PostDevicesIDRefreshOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesIDRefreshBadRequest creates a PostDevicesIDRefreshBadRequest with default headers values
func NewPostDevicesIDRefreshBadRequest() *PostDevicesIDRefreshBadRequest {
	return &PostDevicesIDRefreshBadRequest{}
}

/*PostDevicesIDRefreshBadRequest handles this case with default header values.

Bad Request
*/
type PostDevicesIDRefreshBadRequest struct {
	Payload *models.Error
}

func (o *PostDevicesIDRefreshBadRequest) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/refresh][%d] postDevicesIdRefreshBadRequest  %+v", 400, o.Payload)
}

func (o *PostDevicesIDRefreshBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesIDRefreshBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesIDRefreshUnauthorized creates a PostDevicesIDRefreshUnauthorized with default headers values
func NewPostDevicesIDRefreshUnauthorized() *PostDevicesIDRefreshUnauthorized {
	return &PostDevicesIDRefreshUnauthorized{}
}

/*PostDevicesIDRefreshUnauthorized handles this case with default header values.

Unauthorized
*/
type PostDevicesIDRefreshUnauthorized struct {
	Payload *models.Error
}

func (o *PostDevicesIDRefreshUnauthorized) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/refresh][%d] postDevicesIdRefreshUnauthorized  %+v", 401, o.Payload)
}

func (o *PostDevicesIDRefreshUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesIDRefreshUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesIDRefreshForbidden creates a PostDevicesIDRefreshForbidden with default headers values
func NewPostDevicesIDRefreshForbidden() *PostDevicesIDRefreshForbidden {
	return &PostDevicesIDRefreshForbidden{}
}

/*PostDevicesIDRefreshForbidden handles this case with default header values.

Forbidden
*/
type PostDevicesIDRefreshForbidden struct {
	Payload *models.Error
}

func (o *PostDevicesIDRefreshForbidden) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/refresh][%d] postDevicesIdRefreshForbidden  %+v", 403, o.Payload)
}

func (o *PostDevicesIDRefreshForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesIDRefreshForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesIDRefreshInternalServerError creates a PostDevicesIDRefreshInternalServerError with default headers values
func NewPostDevicesIDRefreshInternalServerError() *PostDevicesIDRefreshInternalServerError {
	return &PostDevicesIDRefreshInternalServerError{}
}

/*PostDevicesIDRefreshInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostDevicesIDRefreshInternalServerError struct {
	Payload *models.Error
}

func (o *PostDevicesIDRefreshInternalServerError) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/refresh][%d] postDevicesIdRefreshInternalServerError  %+v", 500, o.Payload)
}

func (o *PostDevicesIDRefreshInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesIDRefreshInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
