// Code generated by go-swagger; DO NOT EDIT.

package sites

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/ffddorf/unms-api-go/models"
)

// GetSitesIDImagesReader is a Reader for the GetSitesIDImages structure.
type GetSitesIDImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSitesIDImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSitesIDImagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSitesIDImagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetSitesIDImagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSitesIDImagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSitesIDImagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSitesIDImagesOK creates a GetSitesIDImagesOK with default headers values
func NewGetSitesIDImagesOK() *GetSitesIDImagesOK {
	return &GetSitesIDImagesOK{}
}

/*GetSitesIDImagesOK handles this case with default header values.

Successful
*/
type GetSitesIDImagesOK struct {
	Payload models.ListOfImages
}

func (o *GetSitesIDImagesOK) Error() string {
	return fmt.Sprintf("[GET /sites/{id}/images][%d] getSitesIdImagesOK  %+v", 200, o.Payload)
}

func (o *GetSitesIDImagesOK) GetPayload() models.ListOfImages {
	return o.Payload
}

func (o *GetSitesIDImagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSitesIDImagesBadRequest creates a GetSitesIDImagesBadRequest with default headers values
func NewGetSitesIDImagesBadRequest() *GetSitesIDImagesBadRequest {
	return &GetSitesIDImagesBadRequest{}
}

/*GetSitesIDImagesBadRequest handles this case with default header values.

Bad Request
*/
type GetSitesIDImagesBadRequest struct {
	Payload *models.Error
}

func (o *GetSitesIDImagesBadRequest) Error() string {
	return fmt.Sprintf("[GET /sites/{id}/images][%d] getSitesIdImagesBadRequest  %+v", 400, o.Payload)
}

func (o *GetSitesIDImagesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSitesIDImagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSitesIDImagesUnauthorized creates a GetSitesIDImagesUnauthorized with default headers values
func NewGetSitesIDImagesUnauthorized() *GetSitesIDImagesUnauthorized {
	return &GetSitesIDImagesUnauthorized{}
}

/*GetSitesIDImagesUnauthorized handles this case with default header values.

Unauthorized
*/
type GetSitesIDImagesUnauthorized struct {
	Payload *models.Error
}

func (o *GetSitesIDImagesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /sites/{id}/images][%d] getSitesIdImagesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSitesIDImagesUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSitesIDImagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSitesIDImagesForbidden creates a GetSitesIDImagesForbidden with default headers values
func NewGetSitesIDImagesForbidden() *GetSitesIDImagesForbidden {
	return &GetSitesIDImagesForbidden{}
}

/*GetSitesIDImagesForbidden handles this case with default header values.

Forbidden
*/
type GetSitesIDImagesForbidden struct {
	Payload *models.Error
}

func (o *GetSitesIDImagesForbidden) Error() string {
	return fmt.Sprintf("[GET /sites/{id}/images][%d] getSitesIdImagesForbidden  %+v", 403, o.Payload)
}

func (o *GetSitesIDImagesForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSitesIDImagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSitesIDImagesInternalServerError creates a GetSitesIDImagesInternalServerError with default headers values
func NewGetSitesIDImagesInternalServerError() *GetSitesIDImagesInternalServerError {
	return &GetSitesIDImagesInternalServerError{}
}

/*GetSitesIDImagesInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetSitesIDImagesInternalServerError struct {
	Payload *models.Error
}

func (o *GetSitesIDImagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /sites/{id}/images][%d] getSitesIdImagesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSitesIDImagesInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSitesIDImagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}