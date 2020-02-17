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

// PostDiscoveryConnectUbntReader is a Reader for the PostDiscoveryConnectUbnt structure.
type PostDiscoveryConnectUbntReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDiscoveryConnectUbntReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDiscoveryConnectUbntOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostDiscoveryConnectUbntBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostDiscoveryConnectUbntUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostDiscoveryConnectUbntInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostDiscoveryConnectUbntOK creates a PostDiscoveryConnectUbntOK with default headers values
func NewPostDiscoveryConnectUbntOK() *PostDiscoveryConnectUbntOK {
	return &PostDiscoveryConnectUbntOK{}
}

/*PostDiscoveryConnectUbntOK handles this case with default header values.

Successful
*/
type PostDiscoveryConnectUbntOK struct {
	Payload *models.Status
}

func (o *PostDiscoveryConnectUbntOK) Error() string {
	return fmt.Sprintf("[POST /discovery/connect/ubnt][%d] postDiscoveryConnectUbntOK  %+v", 200, o.Payload)
}

func (o *PostDiscoveryConnectUbntOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *PostDiscoveryConnectUbntOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDiscoveryConnectUbntBadRequest creates a PostDiscoveryConnectUbntBadRequest with default headers values
func NewPostDiscoveryConnectUbntBadRequest() *PostDiscoveryConnectUbntBadRequest {
	return &PostDiscoveryConnectUbntBadRequest{}
}

/*PostDiscoveryConnectUbntBadRequest handles this case with default header values.

Bad Request
*/
type PostDiscoveryConnectUbntBadRequest struct {
	Payload *models.Error
}

func (o *PostDiscoveryConnectUbntBadRequest) Error() string {
	return fmt.Sprintf("[POST /discovery/connect/ubnt][%d] postDiscoveryConnectUbntBadRequest  %+v", 400, o.Payload)
}

func (o *PostDiscoveryConnectUbntBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDiscoveryConnectUbntBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDiscoveryConnectUbntUnauthorized creates a PostDiscoveryConnectUbntUnauthorized with default headers values
func NewPostDiscoveryConnectUbntUnauthorized() *PostDiscoveryConnectUbntUnauthorized {
	return &PostDiscoveryConnectUbntUnauthorized{}
}

/*PostDiscoveryConnectUbntUnauthorized handles this case with default header values.

Unauthorized
*/
type PostDiscoveryConnectUbntUnauthorized struct {
	Payload *models.Error
}

func (o *PostDiscoveryConnectUbntUnauthorized) Error() string {
	return fmt.Sprintf("[POST /discovery/connect/ubnt][%d] postDiscoveryConnectUbntUnauthorized  %+v", 401, o.Payload)
}

func (o *PostDiscoveryConnectUbntUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDiscoveryConnectUbntUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDiscoveryConnectUbntInternalServerError creates a PostDiscoveryConnectUbntInternalServerError with default headers values
func NewPostDiscoveryConnectUbntInternalServerError() *PostDiscoveryConnectUbntInternalServerError {
	return &PostDiscoveryConnectUbntInternalServerError{}
}

/*PostDiscoveryConnectUbntInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostDiscoveryConnectUbntInternalServerError struct {
	Payload *models.Error
}

func (o *PostDiscoveryConnectUbntInternalServerError) Error() string {
	return fmt.Sprintf("[POST /discovery/connect/ubnt][%d] postDiscoveryConnectUbntInternalServerError  %+v", 500, o.Payload)
}

func (o *PostDiscoveryConnectUbntInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDiscoveryConnectUbntInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
