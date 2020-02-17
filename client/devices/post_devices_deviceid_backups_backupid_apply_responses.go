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

// PostDevicesDeviceidBackupsBackupidApplyReader is a Reader for the PostDevicesDeviceidBackupsBackupidApply structure.
type PostDevicesDeviceidBackupsBackupidApplyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDevicesDeviceidBackupsBackupidApplyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDevicesDeviceidBackupsBackupidApplyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostDevicesDeviceidBackupsBackupidApplyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostDevicesDeviceidBackupsBackupidApplyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostDevicesDeviceidBackupsBackupidApplyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostDevicesDeviceidBackupsBackupidApplyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostDevicesDeviceidBackupsBackupidApplyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostDevicesDeviceidBackupsBackupidApplyOK creates a PostDevicesDeviceidBackupsBackupidApplyOK with default headers values
func NewPostDevicesDeviceidBackupsBackupidApplyOK() *PostDevicesDeviceidBackupsBackupidApplyOK {
	return &PostDevicesDeviceidBackupsBackupidApplyOK{}
}

/*PostDevicesDeviceidBackupsBackupidApplyOK handles this case with default header values.

Successful
*/
type PostDevicesDeviceidBackupsBackupidApplyOK struct {
	Payload *models.Status
}

func (o *PostDevicesDeviceidBackupsBackupidApplyOK) Error() string {
	return fmt.Sprintf("[POST /devices/{deviceId}/backups/{backupId}/apply][%d] postDevicesDeviceidBackupsBackupidApplyOK  %+v", 200, o.Payload)
}

func (o *PostDevicesDeviceidBackupsBackupidApplyOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *PostDevicesDeviceidBackupsBackupidApplyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesDeviceidBackupsBackupidApplyBadRequest creates a PostDevicesDeviceidBackupsBackupidApplyBadRequest with default headers values
func NewPostDevicesDeviceidBackupsBackupidApplyBadRequest() *PostDevicesDeviceidBackupsBackupidApplyBadRequest {
	return &PostDevicesDeviceidBackupsBackupidApplyBadRequest{}
}

/*PostDevicesDeviceidBackupsBackupidApplyBadRequest handles this case with default header values.

Bad Request
*/
type PostDevicesDeviceidBackupsBackupidApplyBadRequest struct {
	Payload *models.Error
}

func (o *PostDevicesDeviceidBackupsBackupidApplyBadRequest) Error() string {
	return fmt.Sprintf("[POST /devices/{deviceId}/backups/{backupId}/apply][%d] postDevicesDeviceidBackupsBackupidApplyBadRequest  %+v", 400, o.Payload)
}

func (o *PostDevicesDeviceidBackupsBackupidApplyBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesDeviceidBackupsBackupidApplyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesDeviceidBackupsBackupidApplyUnauthorized creates a PostDevicesDeviceidBackupsBackupidApplyUnauthorized with default headers values
func NewPostDevicesDeviceidBackupsBackupidApplyUnauthorized() *PostDevicesDeviceidBackupsBackupidApplyUnauthorized {
	return &PostDevicesDeviceidBackupsBackupidApplyUnauthorized{}
}

/*PostDevicesDeviceidBackupsBackupidApplyUnauthorized handles this case with default header values.

Unauthorized
*/
type PostDevicesDeviceidBackupsBackupidApplyUnauthorized struct {
	Payload *models.Error
}

func (o *PostDevicesDeviceidBackupsBackupidApplyUnauthorized) Error() string {
	return fmt.Sprintf("[POST /devices/{deviceId}/backups/{backupId}/apply][%d] postDevicesDeviceidBackupsBackupidApplyUnauthorized  %+v", 401, o.Payload)
}

func (o *PostDevicesDeviceidBackupsBackupidApplyUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesDeviceidBackupsBackupidApplyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesDeviceidBackupsBackupidApplyForbidden creates a PostDevicesDeviceidBackupsBackupidApplyForbidden with default headers values
func NewPostDevicesDeviceidBackupsBackupidApplyForbidden() *PostDevicesDeviceidBackupsBackupidApplyForbidden {
	return &PostDevicesDeviceidBackupsBackupidApplyForbidden{}
}

/*PostDevicesDeviceidBackupsBackupidApplyForbidden handles this case with default header values.

Forbidden
*/
type PostDevicesDeviceidBackupsBackupidApplyForbidden struct {
	Payload *models.Error
}

func (o *PostDevicesDeviceidBackupsBackupidApplyForbidden) Error() string {
	return fmt.Sprintf("[POST /devices/{deviceId}/backups/{backupId}/apply][%d] postDevicesDeviceidBackupsBackupidApplyForbidden  %+v", 403, o.Payload)
}

func (o *PostDevicesDeviceidBackupsBackupidApplyForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesDeviceidBackupsBackupidApplyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesDeviceidBackupsBackupidApplyNotFound creates a PostDevicesDeviceidBackupsBackupidApplyNotFound with default headers values
func NewPostDevicesDeviceidBackupsBackupidApplyNotFound() *PostDevicesDeviceidBackupsBackupidApplyNotFound {
	return &PostDevicesDeviceidBackupsBackupidApplyNotFound{}
}

/*PostDevicesDeviceidBackupsBackupidApplyNotFound handles this case with default header values.

Not Found
*/
type PostDevicesDeviceidBackupsBackupidApplyNotFound struct {
	Payload *models.Error
}

func (o *PostDevicesDeviceidBackupsBackupidApplyNotFound) Error() string {
	return fmt.Sprintf("[POST /devices/{deviceId}/backups/{backupId}/apply][%d] postDevicesDeviceidBackupsBackupidApplyNotFound  %+v", 404, o.Payload)
}

func (o *PostDevicesDeviceidBackupsBackupidApplyNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesDeviceidBackupsBackupidApplyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesDeviceidBackupsBackupidApplyInternalServerError creates a PostDevicesDeviceidBackupsBackupidApplyInternalServerError with default headers values
func NewPostDevicesDeviceidBackupsBackupidApplyInternalServerError() *PostDevicesDeviceidBackupsBackupidApplyInternalServerError {
	return &PostDevicesDeviceidBackupsBackupidApplyInternalServerError{}
}

/*PostDevicesDeviceidBackupsBackupidApplyInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostDevicesDeviceidBackupsBackupidApplyInternalServerError struct {
	Payload *models.Error
}

func (o *PostDevicesDeviceidBackupsBackupidApplyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /devices/{deviceId}/backups/{backupId}/apply][%d] postDevicesDeviceidBackupsBackupidApplyInternalServerError  %+v", 500, o.Payload)
}

func (o *PostDevicesDeviceidBackupsBackupidApplyInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesDeviceidBackupsBackupidApplyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}