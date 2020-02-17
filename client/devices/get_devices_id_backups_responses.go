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

// GetDevicesIDBackupsReader is a Reader for the GetDevicesIDBackups structure.
type GetDevicesIDBackupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDevicesIDBackupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDevicesIDBackupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDevicesIDBackupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDevicesIDBackupsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDevicesIDBackupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDevicesIDBackupsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDevicesIDBackupsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDevicesIDBackupsOK creates a GetDevicesIDBackupsOK with default headers values
func NewGetDevicesIDBackupsOK() *GetDevicesIDBackupsOK {
	return &GetDevicesIDBackupsOK{}
}

/*GetDevicesIDBackupsOK handles this case with default header values.

Successful
*/
type GetDevicesIDBackupsOK struct {
	Payload models.DeviceBackupList
}

func (o *GetDevicesIDBackupsOK) Error() string {
	return fmt.Sprintf("[GET /devices/{id}/backups][%d] getDevicesIdBackupsOK  %+v", 200, o.Payload)
}

func (o *GetDevicesIDBackupsOK) GetPayload() models.DeviceBackupList {
	return o.Payload
}

func (o *GetDevicesIDBackupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesIDBackupsBadRequest creates a GetDevicesIDBackupsBadRequest with default headers values
func NewGetDevicesIDBackupsBadRequest() *GetDevicesIDBackupsBadRequest {
	return &GetDevicesIDBackupsBadRequest{}
}

/*GetDevicesIDBackupsBadRequest handles this case with default header values.

Bad Request
*/
type GetDevicesIDBackupsBadRequest struct {
	Payload *models.Error
}

func (o *GetDevicesIDBackupsBadRequest) Error() string {
	return fmt.Sprintf("[GET /devices/{id}/backups][%d] getDevicesIdBackupsBadRequest  %+v", 400, o.Payload)
}

func (o *GetDevicesIDBackupsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesIDBackupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesIDBackupsUnauthorized creates a GetDevicesIDBackupsUnauthorized with default headers values
func NewGetDevicesIDBackupsUnauthorized() *GetDevicesIDBackupsUnauthorized {
	return &GetDevicesIDBackupsUnauthorized{}
}

/*GetDevicesIDBackupsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetDevicesIDBackupsUnauthorized struct {
	Payload *models.Error
}

func (o *GetDevicesIDBackupsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /devices/{id}/backups][%d] getDevicesIdBackupsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetDevicesIDBackupsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesIDBackupsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesIDBackupsForbidden creates a GetDevicesIDBackupsForbidden with default headers values
func NewGetDevicesIDBackupsForbidden() *GetDevicesIDBackupsForbidden {
	return &GetDevicesIDBackupsForbidden{}
}

/*GetDevicesIDBackupsForbidden handles this case with default header values.

Forbidden
*/
type GetDevicesIDBackupsForbidden struct {
	Payload *models.Error
}

func (o *GetDevicesIDBackupsForbidden) Error() string {
	return fmt.Sprintf("[GET /devices/{id}/backups][%d] getDevicesIdBackupsForbidden  %+v", 403, o.Payload)
}

func (o *GetDevicesIDBackupsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesIDBackupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesIDBackupsNotFound creates a GetDevicesIDBackupsNotFound with default headers values
func NewGetDevicesIDBackupsNotFound() *GetDevicesIDBackupsNotFound {
	return &GetDevicesIDBackupsNotFound{}
}

/*GetDevicesIDBackupsNotFound handles this case with default header values.

Not Found
*/
type GetDevicesIDBackupsNotFound struct {
	Payload *models.Error
}

func (o *GetDevicesIDBackupsNotFound) Error() string {
	return fmt.Sprintf("[GET /devices/{id}/backups][%d] getDevicesIdBackupsNotFound  %+v", 404, o.Payload)
}

func (o *GetDevicesIDBackupsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesIDBackupsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesIDBackupsInternalServerError creates a GetDevicesIDBackupsInternalServerError with default headers values
func NewGetDevicesIDBackupsInternalServerError() *GetDevicesIDBackupsInternalServerError {
	return &GetDevicesIDBackupsInternalServerError{}
}

/*GetDevicesIDBackupsInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetDevicesIDBackupsInternalServerError struct {
	Payload *models.Error
}

func (o *GetDevicesIDBackupsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /devices/{id}/backups][%d] getDevicesIdBackupsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDevicesIDBackupsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesIDBackupsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
