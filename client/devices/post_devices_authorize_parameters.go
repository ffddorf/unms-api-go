// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/ffddorf/unms-api-go/models"
)

// NewPostDevicesAuthorizeParams creates a new PostDevicesAuthorizeParams object
// with the default values initialized.
func NewPostDevicesAuthorizeParams() *PostDevicesAuthorizeParams {
	var ()
	return &PostDevicesAuthorizeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostDevicesAuthorizeParamsWithTimeout creates a new PostDevicesAuthorizeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostDevicesAuthorizeParamsWithTimeout(timeout time.Duration) *PostDevicesAuthorizeParams {
	var ()
	return &PostDevicesAuthorizeParams{

		timeout: timeout,
	}
}

// NewPostDevicesAuthorizeParamsWithContext creates a new PostDevicesAuthorizeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostDevicesAuthorizeParamsWithContext(ctx context.Context) *PostDevicesAuthorizeParams {
	var ()
	return &PostDevicesAuthorizeParams{

		Context: ctx,
	}
}

// NewPostDevicesAuthorizeParamsWithHTTPClient creates a new PostDevicesAuthorizeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostDevicesAuthorizeParamsWithHTTPClient(client *http.Client) *PostDevicesAuthorizeParams {
	var ()
	return &PostDevicesAuthorizeParams{
		HTTPClient: client,
	}
}

/*PostDevicesAuthorizeParams contains all the parameters to send to the API endpoint
for the post devices authorize operation typically these are written to a http.Request
*/
type PostDevicesAuthorizeParams struct {

	/*Body*/
	Body *models.Model42

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post devices authorize params
func (o *PostDevicesAuthorizeParams) WithTimeout(timeout time.Duration) *PostDevicesAuthorizeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post devices authorize params
func (o *PostDevicesAuthorizeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post devices authorize params
func (o *PostDevicesAuthorizeParams) WithContext(ctx context.Context) *PostDevicesAuthorizeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post devices authorize params
func (o *PostDevicesAuthorizeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post devices authorize params
func (o *PostDevicesAuthorizeParams) WithHTTPClient(client *http.Client) *PostDevicesAuthorizeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post devices authorize params
func (o *PostDevicesAuthorizeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post devices authorize params
func (o *PostDevicesAuthorizeParams) WithBody(body *models.Model42) *PostDevicesAuthorizeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post devices authorize params
func (o *PostDevicesAuthorizeParams) SetBody(body *models.Model42) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostDevicesAuthorizeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}