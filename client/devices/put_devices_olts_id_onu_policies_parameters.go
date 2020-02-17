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

// NewPutDevicesOltsIDOnuPoliciesParams creates a new PutDevicesOltsIDOnuPoliciesParams object
// with the default values initialized.
func NewPutDevicesOltsIDOnuPoliciesParams() *PutDevicesOltsIDOnuPoliciesParams {
	var ()
	return &PutDevicesOltsIDOnuPoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutDevicesOltsIDOnuPoliciesParamsWithTimeout creates a new PutDevicesOltsIDOnuPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutDevicesOltsIDOnuPoliciesParamsWithTimeout(timeout time.Duration) *PutDevicesOltsIDOnuPoliciesParams {
	var ()
	return &PutDevicesOltsIDOnuPoliciesParams{

		timeout: timeout,
	}
}

// NewPutDevicesOltsIDOnuPoliciesParamsWithContext creates a new PutDevicesOltsIDOnuPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutDevicesOltsIDOnuPoliciesParamsWithContext(ctx context.Context) *PutDevicesOltsIDOnuPoliciesParams {
	var ()
	return &PutDevicesOltsIDOnuPoliciesParams{

		Context: ctx,
	}
}

// NewPutDevicesOltsIDOnuPoliciesParamsWithHTTPClient creates a new PutDevicesOltsIDOnuPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutDevicesOltsIDOnuPoliciesParamsWithHTTPClient(client *http.Client) *PutDevicesOltsIDOnuPoliciesParams {
	var ()
	return &PutDevicesOltsIDOnuPoliciesParams{
		HTTPClient: client,
	}
}

/*PutDevicesOltsIDOnuPoliciesParams contains all the parameters to send to the API endpoint
for the put devices olts Id onu policies operation typically these are written to a http.Request
*/
type PutDevicesOltsIDOnuPoliciesParams struct {

	/*Body*/
	Body *models.OnuPolicies1
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put devices olts Id onu policies params
func (o *PutDevicesOltsIDOnuPoliciesParams) WithTimeout(timeout time.Duration) *PutDevicesOltsIDOnuPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put devices olts Id onu policies params
func (o *PutDevicesOltsIDOnuPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put devices olts Id onu policies params
func (o *PutDevicesOltsIDOnuPoliciesParams) WithContext(ctx context.Context) *PutDevicesOltsIDOnuPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put devices olts Id onu policies params
func (o *PutDevicesOltsIDOnuPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put devices olts Id onu policies params
func (o *PutDevicesOltsIDOnuPoliciesParams) WithHTTPClient(client *http.Client) *PutDevicesOltsIDOnuPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put devices olts Id onu policies params
func (o *PutDevicesOltsIDOnuPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put devices olts Id onu policies params
func (o *PutDevicesOltsIDOnuPoliciesParams) WithBody(body *models.OnuPolicies1) *PutDevicesOltsIDOnuPoliciesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put devices olts Id onu policies params
func (o *PutDevicesOltsIDOnuPoliciesParams) SetBody(body *models.OnuPolicies1) {
	o.Body = body
}

// WithID adds the id to the put devices olts Id onu policies params
func (o *PutDevicesOltsIDOnuPoliciesParams) WithID(id string) *PutDevicesOltsIDOnuPoliciesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put devices olts Id onu policies params
func (o *PutDevicesOltsIDOnuPoliciesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutDevicesOltsIDOnuPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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