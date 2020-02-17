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

// GetNmsServerconfigReader is a Reader for the GetNmsServerconfig structure.
type GetNmsServerconfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNmsServerconfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNmsServerconfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetNmsServerconfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNmsServerconfigOK creates a GetNmsServerconfigOK with default headers values
func NewGetNmsServerconfigOK() *GetNmsServerconfigOK {
	return &GetNmsServerconfigOK{}
}

/*GetNmsServerconfigOK handles this case with default header values.

Successful
*/
type GetNmsServerconfigOK struct {
	Payload *models.ServerConfig
}

func (o *GetNmsServerconfigOK) Error() string {
	return fmt.Sprintf("[GET /nms/server-config][%d] getNmsServerconfigOK  %+v", 200, o.Payload)
}

func (o *GetNmsServerconfigOK) GetPayload() *models.ServerConfig {
	return o.Payload
}

func (o *GetNmsServerconfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNmsServerconfigInternalServerError creates a GetNmsServerconfigInternalServerError with default headers values
func NewGetNmsServerconfigInternalServerError() *GetNmsServerconfigInternalServerError {
	return &GetNmsServerconfigInternalServerError{}
}

/*GetNmsServerconfigInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetNmsServerconfigInternalServerError struct {
	Payload *models.Error
}

func (o *GetNmsServerconfigInternalServerError) Error() string {
	return fmt.Sprintf("[GET /nms/server-config][%d] getNmsServerconfigInternalServerError  %+v", 500, o.Payload)
}

func (o *GetNmsServerconfigInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNmsServerconfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
