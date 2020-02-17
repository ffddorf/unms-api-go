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

// GetSitesSiteidTrafficReader is a Reader for the GetSitesSiteidTraffic structure.
type GetSitesSiteidTrafficReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSitesSiteidTrafficReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSitesSiteidTrafficOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSitesSiteidTrafficBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetSitesSiteidTrafficUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSitesSiteidTrafficForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSitesSiteidTrafficInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSitesSiteidTrafficOK creates a GetSitesSiteidTrafficOK with default headers values
func NewGetSitesSiteidTrafficOK() *GetSitesSiteidTrafficOK {
	return &GetSitesSiteidTrafficOK{}
}

/*GetSitesSiteidTrafficOK handles this case with default header values.

Successful
*/
type GetSitesSiteidTrafficOK struct {
	Payload models.TrafficList
}

func (o *GetSitesSiteidTrafficOK) Error() string {
	return fmt.Sprintf("[GET /sites/{siteId}/traffic][%d] getSitesSiteidTrafficOK  %+v", 200, o.Payload)
}

func (o *GetSitesSiteidTrafficOK) GetPayload() models.TrafficList {
	return o.Payload
}

func (o *GetSitesSiteidTrafficOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSitesSiteidTrafficBadRequest creates a GetSitesSiteidTrafficBadRequest with default headers values
func NewGetSitesSiteidTrafficBadRequest() *GetSitesSiteidTrafficBadRequest {
	return &GetSitesSiteidTrafficBadRequest{}
}

/*GetSitesSiteidTrafficBadRequest handles this case with default header values.

Bad Request
*/
type GetSitesSiteidTrafficBadRequest struct {
	Payload *models.Error
}

func (o *GetSitesSiteidTrafficBadRequest) Error() string {
	return fmt.Sprintf("[GET /sites/{siteId}/traffic][%d] getSitesSiteidTrafficBadRequest  %+v", 400, o.Payload)
}

func (o *GetSitesSiteidTrafficBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSitesSiteidTrafficBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSitesSiteidTrafficUnauthorized creates a GetSitesSiteidTrafficUnauthorized with default headers values
func NewGetSitesSiteidTrafficUnauthorized() *GetSitesSiteidTrafficUnauthorized {
	return &GetSitesSiteidTrafficUnauthorized{}
}

/*GetSitesSiteidTrafficUnauthorized handles this case with default header values.

Unauthorized
*/
type GetSitesSiteidTrafficUnauthorized struct {
	Payload *models.Error
}

func (o *GetSitesSiteidTrafficUnauthorized) Error() string {
	return fmt.Sprintf("[GET /sites/{siteId}/traffic][%d] getSitesSiteidTrafficUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSitesSiteidTrafficUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSitesSiteidTrafficUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSitesSiteidTrafficForbidden creates a GetSitesSiteidTrafficForbidden with default headers values
func NewGetSitesSiteidTrafficForbidden() *GetSitesSiteidTrafficForbidden {
	return &GetSitesSiteidTrafficForbidden{}
}

/*GetSitesSiteidTrafficForbidden handles this case with default header values.

Forbidden
*/
type GetSitesSiteidTrafficForbidden struct {
	Payload *models.Error
}

func (o *GetSitesSiteidTrafficForbidden) Error() string {
	return fmt.Sprintf("[GET /sites/{siteId}/traffic][%d] getSitesSiteidTrafficForbidden  %+v", 403, o.Payload)
}

func (o *GetSitesSiteidTrafficForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSitesSiteidTrafficForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSitesSiteidTrafficInternalServerError creates a GetSitesSiteidTrafficInternalServerError with default headers values
func NewGetSitesSiteidTrafficInternalServerError() *GetSitesSiteidTrafficInternalServerError {
	return &GetSitesSiteidTrafficInternalServerError{}
}

/*GetSitesSiteidTrafficInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetSitesSiteidTrafficInternalServerError struct {
	Payload *models.Error
}

func (o *GetSitesSiteidTrafficInternalServerError) Error() string {
	return fmt.Sprintf("[GET /sites/{siteId}/traffic][%d] getSitesSiteidTrafficInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSitesSiteidTrafficInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSitesSiteidTrafficInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
