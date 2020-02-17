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
)

// NewGetDevicesOltsIDOnuPoliciesParams creates a new GetDevicesOltsIDOnuPoliciesParams object
// with the default values initialized.
func NewGetDevicesOltsIDOnuPoliciesParams() *GetDevicesOltsIDOnuPoliciesParams {
	var ()
	return &GetDevicesOltsIDOnuPoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDevicesOltsIDOnuPoliciesParamsWithTimeout creates a new GetDevicesOltsIDOnuPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDevicesOltsIDOnuPoliciesParamsWithTimeout(timeout time.Duration) *GetDevicesOltsIDOnuPoliciesParams {
	var ()
	return &GetDevicesOltsIDOnuPoliciesParams{

		timeout: timeout,
	}
}

// NewGetDevicesOltsIDOnuPoliciesParamsWithContext creates a new GetDevicesOltsIDOnuPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDevicesOltsIDOnuPoliciesParamsWithContext(ctx context.Context) *GetDevicesOltsIDOnuPoliciesParams {
	var ()
	return &GetDevicesOltsIDOnuPoliciesParams{

		Context: ctx,
	}
}

// NewGetDevicesOltsIDOnuPoliciesParamsWithHTTPClient creates a new GetDevicesOltsIDOnuPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDevicesOltsIDOnuPoliciesParamsWithHTTPClient(client *http.Client) *GetDevicesOltsIDOnuPoliciesParams {
	var ()
	return &GetDevicesOltsIDOnuPoliciesParams{
		HTTPClient: client,
	}
}

/*GetDevicesOltsIDOnuPoliciesParams contains all the parameters to send to the API endpoint
for the get devices olts Id onu policies operation typically these are written to a http.Request
*/
type GetDevicesOltsIDOnuPoliciesParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get devices olts Id onu policies params
func (o *GetDevicesOltsIDOnuPoliciesParams) WithTimeout(timeout time.Duration) *GetDevicesOltsIDOnuPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get devices olts Id onu policies params
func (o *GetDevicesOltsIDOnuPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get devices olts Id onu policies params
func (o *GetDevicesOltsIDOnuPoliciesParams) WithContext(ctx context.Context) *GetDevicesOltsIDOnuPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get devices olts Id onu policies params
func (o *GetDevicesOltsIDOnuPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get devices olts Id onu policies params
func (o *GetDevicesOltsIDOnuPoliciesParams) WithHTTPClient(client *http.Client) *GetDevicesOltsIDOnuPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get devices olts Id onu policies params
func (o *GetDevicesOltsIDOnuPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get devices olts Id onu policies params
func (o *GetDevicesOltsIDOnuPoliciesParams) WithID(id string) *GetDevicesOltsIDOnuPoliciesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get devices olts Id onu policies params
func (o *GetDevicesOltsIDOnuPoliciesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDevicesOltsIDOnuPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
