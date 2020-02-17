// Code generated by go-swagger; DO NOT EDIT.

package speed_test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/ffddorf/unms-api-go/models"
)

// GetSpeedtestsReader is a Reader for the GetSpeedtests structure.
type GetSpeedtestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSpeedtestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSpeedtestsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSpeedtestsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSpeedtestsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSpeedtestsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSpeedtestsOK creates a GetSpeedtestsOK with default headers values
func NewGetSpeedtestsOK() *GetSpeedtestsOK {
	return &GetSpeedtestsOK{}
}

/*GetSpeedtestsOK handles this case with default header values.

Successful
*/
type GetSpeedtestsOK struct {
	Payload models.ListOfSpeedTests
}

func (o *GetSpeedtestsOK) Error() string {
	return fmt.Sprintf("[GET /speed-tests][%d] getSpeedtestsOK  %+v", 200, o.Payload)
}

func (o *GetSpeedtestsOK) GetPayload() models.ListOfSpeedTests {
	return o.Payload
}

func (o *GetSpeedtestsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeedtestsUnauthorized creates a GetSpeedtestsUnauthorized with default headers values
func NewGetSpeedtestsUnauthorized() *GetSpeedtestsUnauthorized {
	return &GetSpeedtestsUnauthorized{}
}

/*GetSpeedtestsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetSpeedtestsUnauthorized struct {
	Payload *models.Error
}

func (o *GetSpeedtestsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /speed-tests][%d] getSpeedtestsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSpeedtestsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSpeedtestsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeedtestsForbidden creates a GetSpeedtestsForbidden with default headers values
func NewGetSpeedtestsForbidden() *GetSpeedtestsForbidden {
	return &GetSpeedtestsForbidden{}
}

/*GetSpeedtestsForbidden handles this case with default header values.

Forbidden
*/
type GetSpeedtestsForbidden struct {
	Payload *models.Error
}

func (o *GetSpeedtestsForbidden) Error() string {
	return fmt.Sprintf("[GET /speed-tests][%d] getSpeedtestsForbidden  %+v", 403, o.Payload)
}

func (o *GetSpeedtestsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSpeedtestsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeedtestsInternalServerError creates a GetSpeedtestsInternalServerError with default headers values
func NewGetSpeedtestsInternalServerError() *GetSpeedtestsInternalServerError {
	return &GetSpeedtestsInternalServerError{}
}

/*GetSpeedtestsInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetSpeedtestsInternalServerError struct {
	Payload *models.Error
}

func (o *GetSpeedtestsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /speed-tests][%d] getSpeedtestsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSpeedtestsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSpeedtestsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
