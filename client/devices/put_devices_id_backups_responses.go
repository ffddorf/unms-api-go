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

// PutDevicesIDBackupsReader is a Reader for the PutDevicesIDBackups structure.
type PutDevicesIDBackupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutDevicesIDBackupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutDevicesIDBackupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutDevicesIDBackupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutDevicesIDBackupsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutDevicesIDBackupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutDevicesIDBackupsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutDevicesIDBackupsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutDevicesIDBackupsOK creates a PutDevicesIDBackupsOK with default headers values
func NewPutDevicesIDBackupsOK() *PutDevicesIDBackupsOK {
	return &PutDevicesIDBackupsOK{}
}

/*PutDevicesIDBackupsOK handles this case with default header values.

Successful
*/
type PutDevicesIDBackupsOK struct {
	Payload *models.DeviceBackup
}

func (o *PutDevicesIDBackupsOK) Error() string {
	return fmt.Sprintf("[PUT /devices/{id}/backups][%d] putDevicesIdBackupsOK  %+v", 200, o.Payload)
}

func (o *PutDevicesIDBackupsOK) GetPayload() *models.DeviceBackup {
	return o.Payload
}

func (o *PutDevicesIDBackupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceBackup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesIDBackupsBadRequest creates a PutDevicesIDBackupsBadRequest with default headers values
func NewPutDevicesIDBackupsBadRequest() *PutDevicesIDBackupsBadRequest {
	return &PutDevicesIDBackupsBadRequest{}
}

/*PutDevicesIDBackupsBadRequest handles this case with default header values.

Bad Request
*/
type PutDevicesIDBackupsBadRequest struct {
	Payload *models.Error
}

func (o *PutDevicesIDBackupsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /devices/{id}/backups][%d] putDevicesIdBackupsBadRequest  %+v", 400, o.Payload)
}

func (o *PutDevicesIDBackupsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesIDBackupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesIDBackupsUnauthorized creates a PutDevicesIDBackupsUnauthorized with default headers values
func NewPutDevicesIDBackupsUnauthorized() *PutDevicesIDBackupsUnauthorized {
	return &PutDevicesIDBackupsUnauthorized{}
}

/*PutDevicesIDBackupsUnauthorized handles this case with default header values.

Unauthorized
*/
type PutDevicesIDBackupsUnauthorized struct {
	Payload *models.Error
}

func (o *PutDevicesIDBackupsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /devices/{id}/backups][%d] putDevicesIdBackupsUnauthorized  %+v", 401, o.Payload)
}

func (o *PutDevicesIDBackupsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesIDBackupsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesIDBackupsForbidden creates a PutDevicesIDBackupsForbidden with default headers values
func NewPutDevicesIDBackupsForbidden() *PutDevicesIDBackupsForbidden {
	return &PutDevicesIDBackupsForbidden{}
}

/*PutDevicesIDBackupsForbidden handles this case with default header values.

Forbidden
*/
type PutDevicesIDBackupsForbidden struct {
	Payload *models.Error
}

func (o *PutDevicesIDBackupsForbidden) Error() string {
	return fmt.Sprintf("[PUT /devices/{id}/backups][%d] putDevicesIdBackupsForbidden  %+v", 403, o.Payload)
}

func (o *PutDevicesIDBackupsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesIDBackupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesIDBackupsNotFound creates a PutDevicesIDBackupsNotFound with default headers values
func NewPutDevicesIDBackupsNotFound() *PutDevicesIDBackupsNotFound {
	return &PutDevicesIDBackupsNotFound{}
}

/*PutDevicesIDBackupsNotFound handles this case with default header values.

Not Found
*/
type PutDevicesIDBackupsNotFound struct {
	Payload *models.Error
}

func (o *PutDevicesIDBackupsNotFound) Error() string {
	return fmt.Sprintf("[PUT /devices/{id}/backups][%d] putDevicesIdBackupsNotFound  %+v", 404, o.Payload)
}

func (o *PutDevicesIDBackupsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesIDBackupsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesIDBackupsInternalServerError creates a PutDevicesIDBackupsInternalServerError with default headers values
func NewPutDevicesIDBackupsInternalServerError() *PutDevicesIDBackupsInternalServerError {
	return &PutDevicesIDBackupsInternalServerError{}
}

/*PutDevicesIDBackupsInternalServerError handles this case with default header values.

Internal Server Error
*/
type PutDevicesIDBackupsInternalServerError struct {
	Payload *models.Error
}

func (o *PutDevicesIDBackupsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /devices/{id}/backups][%d] putDevicesIdBackupsInternalServerError  %+v", 500, o.Payload)
}

func (o *PutDevicesIDBackupsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesIDBackupsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
