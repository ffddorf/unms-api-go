// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/ffddorf/unms-api-go/models"
)

// GetNmsMaintenanceBackupRestoreReader is a Reader for the GetNmsMaintenanceBackupRestore structure.
type GetNmsMaintenanceBackupRestoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNmsMaintenanceBackupRestoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNmsMaintenanceBackupRestoreOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetNmsMaintenanceBackupRestoreUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetNmsMaintenanceBackupRestoreForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetNmsMaintenanceBackupRestoreInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNmsMaintenanceBackupRestoreOK creates a GetNmsMaintenanceBackupRestoreOK with default headers values
func NewGetNmsMaintenanceBackupRestoreOK() *GetNmsMaintenanceBackupRestoreOK {
	return &GetNmsMaintenanceBackupRestoreOK{}
}

/*GetNmsMaintenanceBackupRestoreOK handles this case with default header values.

Successful
*/
type GetNmsMaintenanceBackupRestoreOK struct {
	Payload *models.Status
}

func (o *GetNmsMaintenanceBackupRestoreOK) Error() string {
	return fmt.Sprintf("[GET /nms/maintenance/backup/restore][%d] getNmsMaintenanceBackupRestoreOK  %+v", 200, o.Payload)
}

func (o *GetNmsMaintenanceBackupRestoreOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *GetNmsMaintenanceBackupRestoreOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNmsMaintenanceBackupRestoreUnauthorized creates a GetNmsMaintenanceBackupRestoreUnauthorized with default headers values
func NewGetNmsMaintenanceBackupRestoreUnauthorized() *GetNmsMaintenanceBackupRestoreUnauthorized {
	return &GetNmsMaintenanceBackupRestoreUnauthorized{}
}

/*GetNmsMaintenanceBackupRestoreUnauthorized handles this case with default header values.

Unauthorized
*/
type GetNmsMaintenanceBackupRestoreUnauthorized struct {
	Payload *models.Error
}

func (o *GetNmsMaintenanceBackupRestoreUnauthorized) Error() string {
	return fmt.Sprintf("[GET /nms/maintenance/backup/restore][%d] getNmsMaintenanceBackupRestoreUnauthorized  %+v", 401, o.Payload)
}

func (o *GetNmsMaintenanceBackupRestoreUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNmsMaintenanceBackupRestoreUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNmsMaintenanceBackupRestoreForbidden creates a GetNmsMaintenanceBackupRestoreForbidden with default headers values
func NewGetNmsMaintenanceBackupRestoreForbidden() *GetNmsMaintenanceBackupRestoreForbidden {
	return &GetNmsMaintenanceBackupRestoreForbidden{}
}

/*GetNmsMaintenanceBackupRestoreForbidden handles this case with default header values.

Forbidden
*/
type GetNmsMaintenanceBackupRestoreForbidden struct {
	Payload *models.Error
}

func (o *GetNmsMaintenanceBackupRestoreForbidden) Error() string {
	return fmt.Sprintf("[GET /nms/maintenance/backup/restore][%d] getNmsMaintenanceBackupRestoreForbidden  %+v", 403, o.Payload)
}

func (o *GetNmsMaintenanceBackupRestoreForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNmsMaintenanceBackupRestoreForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNmsMaintenanceBackupRestoreInternalServerError creates a GetNmsMaintenanceBackupRestoreInternalServerError with default headers values
func NewGetNmsMaintenanceBackupRestoreInternalServerError() *GetNmsMaintenanceBackupRestoreInternalServerError {
	return &GetNmsMaintenanceBackupRestoreInternalServerError{}
}

/*GetNmsMaintenanceBackupRestoreInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetNmsMaintenanceBackupRestoreInternalServerError struct {
	Payload *models.Error
}

func (o *GetNmsMaintenanceBackupRestoreInternalServerError) Error() string {
	return fmt.Sprintf("[GET /nms/maintenance/backup/restore][%d] getNmsMaintenanceBackupRestoreInternalServerError  %+v", 500, o.Payload)
}

func (o *GetNmsMaintenanceBackupRestoreInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNmsMaintenanceBackupRestoreInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
