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

// PostNmsTrafficSubnetsReader is a Reader for the PostNmsTrafficSubnets structure.
type PostNmsTrafficSubnetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostNmsTrafficSubnetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostNmsTrafficSubnetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostNmsTrafficSubnetsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostNmsTrafficSubnetsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostNmsTrafficSubnetsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostNmsTrafficSubnetsOK creates a PostNmsTrafficSubnetsOK with default headers values
func NewPostNmsTrafficSubnetsOK() *PostNmsTrafficSubnetsOK {
	return &PostNmsTrafficSubnetsOK{}
}

/*PostNmsTrafficSubnetsOK handles this case with default header values.

Successful
*/
type PostNmsTrafficSubnetsOK struct {
	Payload *models.SubnetList
}

func (o *PostNmsTrafficSubnetsOK) Error() string {
	return fmt.Sprintf("[POST /nms/traffic/subnets][%d] postNmsTrafficSubnetsOK  %+v", 200, o.Payload)
}

func (o *PostNmsTrafficSubnetsOK) GetPayload() *models.SubnetList {
	return o.Payload
}

func (o *PostNmsTrafficSubnetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubnetList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNmsTrafficSubnetsUnauthorized creates a PostNmsTrafficSubnetsUnauthorized with default headers values
func NewPostNmsTrafficSubnetsUnauthorized() *PostNmsTrafficSubnetsUnauthorized {
	return &PostNmsTrafficSubnetsUnauthorized{}
}

/*PostNmsTrafficSubnetsUnauthorized handles this case with default header values.

Unauthorized
*/
type PostNmsTrafficSubnetsUnauthorized struct {
	Payload *models.Error
}

func (o *PostNmsTrafficSubnetsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /nms/traffic/subnets][%d] postNmsTrafficSubnetsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostNmsTrafficSubnetsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNmsTrafficSubnetsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNmsTrafficSubnetsForbidden creates a PostNmsTrafficSubnetsForbidden with default headers values
func NewPostNmsTrafficSubnetsForbidden() *PostNmsTrafficSubnetsForbidden {
	return &PostNmsTrafficSubnetsForbidden{}
}

/*PostNmsTrafficSubnetsForbidden handles this case with default header values.

Forbidden
*/
type PostNmsTrafficSubnetsForbidden struct {
	Payload *models.Error
}

func (o *PostNmsTrafficSubnetsForbidden) Error() string {
	return fmt.Sprintf("[POST /nms/traffic/subnets][%d] postNmsTrafficSubnetsForbidden  %+v", 403, o.Payload)
}

func (o *PostNmsTrafficSubnetsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNmsTrafficSubnetsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNmsTrafficSubnetsInternalServerError creates a PostNmsTrafficSubnetsInternalServerError with default headers values
func NewPostNmsTrafficSubnetsInternalServerError() *PostNmsTrafficSubnetsInternalServerError {
	return &PostNmsTrafficSubnetsInternalServerError{}
}

/*PostNmsTrafficSubnetsInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostNmsTrafficSubnetsInternalServerError struct {
	Payload *models.Error
}

func (o *PostNmsTrafficSubnetsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /nms/traffic/subnets][%d] postNmsTrafficSubnetsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostNmsTrafficSubnetsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNmsTrafficSubnetsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
