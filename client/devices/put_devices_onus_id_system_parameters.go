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

// NewPutDevicesOnusIDSystemParams creates a new PutDevicesOnusIDSystemParams object
// with the default values initialized.
func NewPutDevicesOnusIDSystemParams() *PutDevicesOnusIDSystemParams {
	var ()
	return &PutDevicesOnusIDSystemParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutDevicesOnusIDSystemParamsWithTimeout creates a new PutDevicesOnusIDSystemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutDevicesOnusIDSystemParamsWithTimeout(timeout time.Duration) *PutDevicesOnusIDSystemParams {
	var ()
	return &PutDevicesOnusIDSystemParams{

		timeout: timeout,
	}
}

// NewPutDevicesOnusIDSystemParamsWithContext creates a new PutDevicesOnusIDSystemParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutDevicesOnusIDSystemParamsWithContext(ctx context.Context) *PutDevicesOnusIDSystemParams {
	var ()
	return &PutDevicesOnusIDSystemParams{

		Context: ctx,
	}
}

// NewPutDevicesOnusIDSystemParamsWithHTTPClient creates a new PutDevicesOnusIDSystemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutDevicesOnusIDSystemParamsWithHTTPClient(client *http.Client) *PutDevicesOnusIDSystemParams {
	var ()
	return &PutDevicesOnusIDSystemParams{
		HTTPClient: client,
	}
}

/*PutDevicesOnusIDSystemParams contains all the parameters to send to the API endpoint
for the put devices onus Id system operation typically these are written to a http.Request
*/
type PutDevicesOnusIDSystemParams struct {

	/*Body*/
	Body *models.Model64
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put devices onus Id system params
func (o *PutDevicesOnusIDSystemParams) WithTimeout(timeout time.Duration) *PutDevicesOnusIDSystemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put devices onus Id system params
func (o *PutDevicesOnusIDSystemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put devices onus Id system params
func (o *PutDevicesOnusIDSystemParams) WithContext(ctx context.Context) *PutDevicesOnusIDSystemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put devices onus Id system params
func (o *PutDevicesOnusIDSystemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put devices onus Id system params
func (o *PutDevicesOnusIDSystemParams) WithHTTPClient(client *http.Client) *PutDevicesOnusIDSystemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put devices onus Id system params
func (o *PutDevicesOnusIDSystemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put devices onus Id system params
func (o *PutDevicesOnusIDSystemParams) WithBody(body *models.Model64) *PutDevicesOnusIDSystemParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put devices onus Id system params
func (o *PutDevicesOnusIDSystemParams) SetBody(body *models.Model64) {
	o.Body = body
}

// WithID adds the id to the put devices onus Id system params
func (o *PutDevicesOnusIDSystemParams) WithID(id string) *PutDevicesOnusIDSystemParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put devices onus Id system params
func (o *PutDevicesOnusIDSystemParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutDevicesOnusIDSystemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
