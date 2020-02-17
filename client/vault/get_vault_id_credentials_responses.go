// Code generated by go-swagger; DO NOT EDIT.

package vault

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/ffddorf/unms-api-go/models"
)

// GetVaultIDCredentialsReader is a Reader for the GetVaultIDCredentials structure.
type GetVaultIDCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVaultIDCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVaultIDCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVaultIDCredentialsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetVaultIDCredentialsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetVaultIDCredentialsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetVaultIDCredentialsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetVaultIDCredentialsOK creates a GetVaultIDCredentialsOK with default headers values
func NewGetVaultIDCredentialsOK() *GetVaultIDCredentialsOK {
	return &GetVaultIDCredentialsOK{}
}

/*GetVaultIDCredentialsOK handles this case with default header values.

Successful
*/
type GetVaultIDCredentialsOK struct {
	Payload *models.VaultCredentials
}

func (o *GetVaultIDCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /vault/{id}/credentials][%d] getVaultIdCredentialsOK  %+v", 200, o.Payload)
}

func (o *GetVaultIDCredentialsOK) GetPayload() *models.VaultCredentials {
	return o.Payload
}

func (o *GetVaultIDCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VaultCredentials)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVaultIDCredentialsBadRequest creates a GetVaultIDCredentialsBadRequest with default headers values
func NewGetVaultIDCredentialsBadRequest() *GetVaultIDCredentialsBadRequest {
	return &GetVaultIDCredentialsBadRequest{}
}

/*GetVaultIDCredentialsBadRequest handles this case with default header values.

Bad Request
*/
type GetVaultIDCredentialsBadRequest struct {
	Payload *models.Error
}

func (o *GetVaultIDCredentialsBadRequest) Error() string {
	return fmt.Sprintf("[GET /vault/{id}/credentials][%d] getVaultIdCredentialsBadRequest  %+v", 400, o.Payload)
}

func (o *GetVaultIDCredentialsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVaultIDCredentialsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVaultIDCredentialsUnauthorized creates a GetVaultIDCredentialsUnauthorized with default headers values
func NewGetVaultIDCredentialsUnauthorized() *GetVaultIDCredentialsUnauthorized {
	return &GetVaultIDCredentialsUnauthorized{}
}

/*GetVaultIDCredentialsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetVaultIDCredentialsUnauthorized struct {
	Payload *models.Error
}

func (o *GetVaultIDCredentialsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /vault/{id}/credentials][%d] getVaultIdCredentialsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetVaultIDCredentialsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVaultIDCredentialsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVaultIDCredentialsForbidden creates a GetVaultIDCredentialsForbidden with default headers values
func NewGetVaultIDCredentialsForbidden() *GetVaultIDCredentialsForbidden {
	return &GetVaultIDCredentialsForbidden{}
}

/*GetVaultIDCredentialsForbidden handles this case with default header values.

Forbidden
*/
type GetVaultIDCredentialsForbidden struct {
	Payload *models.Error
}

func (o *GetVaultIDCredentialsForbidden) Error() string {
	return fmt.Sprintf("[GET /vault/{id}/credentials][%d] getVaultIdCredentialsForbidden  %+v", 403, o.Payload)
}

func (o *GetVaultIDCredentialsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVaultIDCredentialsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVaultIDCredentialsInternalServerError creates a GetVaultIDCredentialsInternalServerError with default headers values
func NewGetVaultIDCredentialsInternalServerError() *GetVaultIDCredentialsInternalServerError {
	return &GetVaultIDCredentialsInternalServerError{}
}

/*GetVaultIDCredentialsInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetVaultIDCredentialsInternalServerError struct {
	Payload *models.Error
}

func (o *GetVaultIDCredentialsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /vault/{id}/credentials][%d] getVaultIdCredentialsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetVaultIDCredentialsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVaultIDCredentialsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}