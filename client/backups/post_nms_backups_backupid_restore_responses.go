// Code generated by go-swagger; DO NOT EDIT.

package backups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/ffddorf/unms-api-go/models"
)

// PostNmsBackupsBackupidRestoreReader is a Reader for the PostNmsBackupsBackupidRestore structure.
type PostNmsBackupsBackupidRestoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostNmsBackupsBackupidRestoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostNmsBackupsBackupidRestoreOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostNmsBackupsBackupidRestoreBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostNmsBackupsBackupidRestoreUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostNmsBackupsBackupidRestoreForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostNmsBackupsBackupidRestoreNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostNmsBackupsBackupidRestoreInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostNmsBackupsBackupidRestoreOK creates a PostNmsBackupsBackupidRestoreOK with default headers values
func NewPostNmsBackupsBackupidRestoreOK() *PostNmsBackupsBackupidRestoreOK {
	return &PostNmsBackupsBackupidRestoreOK{}
}

/*PostNmsBackupsBackupidRestoreOK handles this case with default header values.

Successful
*/
type PostNmsBackupsBackupidRestoreOK struct {
	Payload *models.Status
}

func (o *PostNmsBackupsBackupidRestoreOK) Error() string {
	return fmt.Sprintf("[POST /nms/backups/{backupId}/restore][%d] postNmsBackupsBackupidRestoreOK  %+v", 200, o.Payload)
}

func (o *PostNmsBackupsBackupidRestoreOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *PostNmsBackupsBackupidRestoreOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNmsBackupsBackupidRestoreBadRequest creates a PostNmsBackupsBackupidRestoreBadRequest with default headers values
func NewPostNmsBackupsBackupidRestoreBadRequest() *PostNmsBackupsBackupidRestoreBadRequest {
	return &PostNmsBackupsBackupidRestoreBadRequest{}
}

/*PostNmsBackupsBackupidRestoreBadRequest handles this case with default header values.

Bad Request
*/
type PostNmsBackupsBackupidRestoreBadRequest struct {
	Payload *models.Error
}

func (o *PostNmsBackupsBackupidRestoreBadRequest) Error() string {
	return fmt.Sprintf("[POST /nms/backups/{backupId}/restore][%d] postNmsBackupsBackupidRestoreBadRequest  %+v", 400, o.Payload)
}

func (o *PostNmsBackupsBackupidRestoreBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNmsBackupsBackupidRestoreBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNmsBackupsBackupidRestoreUnauthorized creates a PostNmsBackupsBackupidRestoreUnauthorized with default headers values
func NewPostNmsBackupsBackupidRestoreUnauthorized() *PostNmsBackupsBackupidRestoreUnauthorized {
	return &PostNmsBackupsBackupidRestoreUnauthorized{}
}

/*PostNmsBackupsBackupidRestoreUnauthorized handles this case with default header values.

Unauthorized
*/
type PostNmsBackupsBackupidRestoreUnauthorized struct {
	Payload *models.Error
}

func (o *PostNmsBackupsBackupidRestoreUnauthorized) Error() string {
	return fmt.Sprintf("[POST /nms/backups/{backupId}/restore][%d] postNmsBackupsBackupidRestoreUnauthorized  %+v", 401, o.Payload)
}

func (o *PostNmsBackupsBackupidRestoreUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNmsBackupsBackupidRestoreUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNmsBackupsBackupidRestoreForbidden creates a PostNmsBackupsBackupidRestoreForbidden with default headers values
func NewPostNmsBackupsBackupidRestoreForbidden() *PostNmsBackupsBackupidRestoreForbidden {
	return &PostNmsBackupsBackupidRestoreForbidden{}
}

/*PostNmsBackupsBackupidRestoreForbidden handles this case with default header values.

Forbidden
*/
type PostNmsBackupsBackupidRestoreForbidden struct {
	Payload *models.Error
}

func (o *PostNmsBackupsBackupidRestoreForbidden) Error() string {
	return fmt.Sprintf("[POST /nms/backups/{backupId}/restore][%d] postNmsBackupsBackupidRestoreForbidden  %+v", 403, o.Payload)
}

func (o *PostNmsBackupsBackupidRestoreForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNmsBackupsBackupidRestoreForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNmsBackupsBackupidRestoreNotFound creates a PostNmsBackupsBackupidRestoreNotFound with default headers values
func NewPostNmsBackupsBackupidRestoreNotFound() *PostNmsBackupsBackupidRestoreNotFound {
	return &PostNmsBackupsBackupidRestoreNotFound{}
}

/*PostNmsBackupsBackupidRestoreNotFound handles this case with default header values.

Not Found
*/
type PostNmsBackupsBackupidRestoreNotFound struct {
	Payload *models.Error
}

func (o *PostNmsBackupsBackupidRestoreNotFound) Error() string {
	return fmt.Sprintf("[POST /nms/backups/{backupId}/restore][%d] postNmsBackupsBackupidRestoreNotFound  %+v", 404, o.Payload)
}

func (o *PostNmsBackupsBackupidRestoreNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNmsBackupsBackupidRestoreNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNmsBackupsBackupidRestoreInternalServerError creates a PostNmsBackupsBackupidRestoreInternalServerError with default headers values
func NewPostNmsBackupsBackupidRestoreInternalServerError() *PostNmsBackupsBackupidRestoreInternalServerError {
	return &PostNmsBackupsBackupidRestoreInternalServerError{}
}

/*PostNmsBackupsBackupidRestoreInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostNmsBackupsBackupidRestoreInternalServerError struct {
	Payload *models.Error
}

func (o *PostNmsBackupsBackupidRestoreInternalServerError) Error() string {
	return fmt.Sprintf("[POST /nms/backups/{backupId}/restore][%d] postNmsBackupsBackupidRestoreInternalServerError  %+v", 500, o.Payload)
}

func (o *PostNmsBackupsBackupidRestoreInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNmsBackupsBackupidRestoreInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
