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

// NewPutDevicesSolarbeamsIDConfigSystemParams creates a new PutDevicesSolarbeamsIDConfigSystemParams object
// with the default values initialized.
func NewPutDevicesSolarbeamsIDConfigSystemParams() *PutDevicesSolarbeamsIDConfigSystemParams {
	var ()
	return &PutDevicesSolarbeamsIDConfigSystemParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutDevicesSolarbeamsIDConfigSystemParamsWithTimeout creates a new PutDevicesSolarbeamsIDConfigSystemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutDevicesSolarbeamsIDConfigSystemParamsWithTimeout(timeout time.Duration) *PutDevicesSolarbeamsIDConfigSystemParams {
	var ()
	return &PutDevicesSolarbeamsIDConfigSystemParams{

		timeout: timeout,
	}
}

// NewPutDevicesSolarbeamsIDConfigSystemParamsWithContext creates a new PutDevicesSolarbeamsIDConfigSystemParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutDevicesSolarbeamsIDConfigSystemParamsWithContext(ctx context.Context) *PutDevicesSolarbeamsIDConfigSystemParams {
	var ()
	return &PutDevicesSolarbeamsIDConfigSystemParams{

		Context: ctx,
	}
}

// NewPutDevicesSolarbeamsIDConfigSystemParamsWithHTTPClient creates a new PutDevicesSolarbeamsIDConfigSystemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutDevicesSolarbeamsIDConfigSystemParamsWithHTTPClient(client *http.Client) *PutDevicesSolarbeamsIDConfigSystemParams {
	var ()
	return &PutDevicesSolarbeamsIDConfigSystemParams{
		HTTPClient: client,
	}
}

/*PutDevicesSolarbeamsIDConfigSystemParams contains all the parameters to send to the API endpoint
for the put devices solarbeams Id config system operation typically these are written to a http.Request
*/
type PutDevicesSolarbeamsIDConfigSystemParams struct {

	/*Body*/
	Body *models.Model37
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put devices solarbeams Id config system params
func (o *PutDevicesSolarbeamsIDConfigSystemParams) WithTimeout(timeout time.Duration) *PutDevicesSolarbeamsIDConfigSystemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put devices solarbeams Id config system params
func (o *PutDevicesSolarbeamsIDConfigSystemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put devices solarbeams Id config system params
func (o *PutDevicesSolarbeamsIDConfigSystemParams) WithContext(ctx context.Context) *PutDevicesSolarbeamsIDConfigSystemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put devices solarbeams Id config system params
func (o *PutDevicesSolarbeamsIDConfigSystemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put devices solarbeams Id config system params
func (o *PutDevicesSolarbeamsIDConfigSystemParams) WithHTTPClient(client *http.Client) *PutDevicesSolarbeamsIDConfigSystemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put devices solarbeams Id config system params
func (o *PutDevicesSolarbeamsIDConfigSystemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put devices solarbeams Id config system params
func (o *PutDevicesSolarbeamsIDConfigSystemParams) WithBody(body *models.Model37) *PutDevicesSolarbeamsIDConfigSystemParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put devices solarbeams Id config system params
func (o *PutDevicesSolarbeamsIDConfigSystemParams) SetBody(body *models.Model37) {
	o.Body = body
}

// WithID adds the id to the put devices solarbeams Id config system params
func (o *PutDevicesSolarbeamsIDConfigSystemParams) WithID(id string) *PutDevicesSolarbeamsIDConfigSystemParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put devices solarbeams Id config system params
func (o *PutDevicesSolarbeamsIDConfigSystemParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutDevicesSolarbeamsIDConfigSystemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}