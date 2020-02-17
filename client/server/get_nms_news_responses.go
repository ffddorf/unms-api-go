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

// GetNmsNewsReader is a Reader for the GetNmsNews structure.
type GetNmsNewsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNmsNewsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNmsNewsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetNmsNewsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetNmsNewsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetNmsNewsNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetNmsNewsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNmsNewsOK creates a GetNmsNewsOK with default headers values
func NewGetNmsNewsOK() *GetNmsNewsOK {
	return &GetNmsNewsOK{}
}

/*GetNmsNewsOK handles this case with default header values.

Successful
*/
type GetNmsNewsOK struct {
	Payload models.NewsSchemaList
}

func (o *GetNmsNewsOK) Error() string {
	return fmt.Sprintf("[GET /nms/news][%d] getNmsNewsOK  %+v", 200, o.Payload)
}

func (o *GetNmsNewsOK) GetPayload() models.NewsSchemaList {
	return o.Payload
}

func (o *GetNmsNewsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNmsNewsUnauthorized creates a GetNmsNewsUnauthorized with default headers values
func NewGetNmsNewsUnauthorized() *GetNmsNewsUnauthorized {
	return &GetNmsNewsUnauthorized{}
}

/*GetNmsNewsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetNmsNewsUnauthorized struct {
	Payload *models.Error
}

func (o *GetNmsNewsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /nms/news][%d] getNmsNewsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetNmsNewsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNmsNewsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNmsNewsForbidden creates a GetNmsNewsForbidden with default headers values
func NewGetNmsNewsForbidden() *GetNmsNewsForbidden {
	return &GetNmsNewsForbidden{}
}

/*GetNmsNewsForbidden handles this case with default header values.

Forbidden
*/
type GetNmsNewsForbidden struct {
	Payload *models.Error
}

func (o *GetNmsNewsForbidden) Error() string {
	return fmt.Sprintf("[GET /nms/news][%d] getNmsNewsForbidden  %+v", 403, o.Payload)
}

func (o *GetNmsNewsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNmsNewsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNmsNewsNotAcceptable creates a GetNmsNewsNotAcceptable with default headers values
func NewGetNmsNewsNotAcceptable() *GetNmsNewsNotAcceptable {
	return &GetNmsNewsNotAcceptable{}
}

/*GetNmsNewsNotAcceptable handles this case with default header values.

Not Acceptable
*/
type GetNmsNewsNotAcceptable struct {
	Payload *models.Error
}

func (o *GetNmsNewsNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /nms/news][%d] getNmsNewsNotAcceptable  %+v", 406, o.Payload)
}

func (o *GetNmsNewsNotAcceptable) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNmsNewsNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNmsNewsInternalServerError creates a GetNmsNewsInternalServerError with default headers values
func NewGetNmsNewsInternalServerError() *GetNmsNewsInternalServerError {
	return &GetNmsNewsInternalServerError{}
}

/*GetNmsNewsInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetNmsNewsInternalServerError struct {
	Payload *models.Error
}

func (o *GetNmsNewsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /nms/news][%d] getNmsNewsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetNmsNewsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNmsNewsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
