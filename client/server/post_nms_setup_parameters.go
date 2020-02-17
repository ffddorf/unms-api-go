// Code generated by go-swagger; DO NOT EDIT.

package server

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

// NewPostNmsSetupParams creates a new PostNmsSetupParams object
// with the default values initialized.
func NewPostNmsSetupParams() *PostNmsSetupParams {
	var ()
	return &PostNmsSetupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostNmsSetupParamsWithTimeout creates a new PostNmsSetupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostNmsSetupParamsWithTimeout(timeout time.Duration) *PostNmsSetupParams {
	var ()
	return &PostNmsSetupParams{

		timeout: timeout,
	}
}

// NewPostNmsSetupParamsWithContext creates a new PostNmsSetupParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostNmsSetupParamsWithContext(ctx context.Context) *PostNmsSetupParams {
	var ()
	return &PostNmsSetupParams{

		Context: ctx,
	}
}

// NewPostNmsSetupParamsWithHTTPClient creates a new PostNmsSetupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostNmsSetupParamsWithHTTPClient(client *http.Client) *PostNmsSetupParams {
	var ()
	return &PostNmsSetupParams{
		HTTPClient: client,
	}
}

/*PostNmsSetupParams contains all the parameters to send to the API endpoint
for the post nms setup operation typically these are written to a http.Request
*/
type PostNmsSetupParams struct {

	/*Body*/
	Body *models.Model45

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post nms setup params
func (o *PostNmsSetupParams) WithTimeout(timeout time.Duration) *PostNmsSetupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post nms setup params
func (o *PostNmsSetupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post nms setup params
func (o *PostNmsSetupParams) WithContext(ctx context.Context) *PostNmsSetupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post nms setup params
func (o *PostNmsSetupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post nms setup params
func (o *PostNmsSetupParams) WithHTTPClient(client *http.Client) *PostNmsSetupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post nms setup params
func (o *PostNmsSetupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post nms setup params
func (o *PostNmsSetupParams) WithBody(body *models.Model45) *PostNmsSetupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post nms setup params
func (o *PostNmsSetupParams) SetBody(body *models.Model45) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostNmsSetupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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