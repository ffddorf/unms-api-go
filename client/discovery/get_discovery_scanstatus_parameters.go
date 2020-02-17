// Code generated by go-swagger; DO NOT EDIT.

package discovery

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

// NewGetDiscoveryScanstatusParams creates a new GetDiscoveryScanstatusParams object
// with the default values initialized.
func NewGetDiscoveryScanstatusParams() *GetDiscoveryScanstatusParams {

	return &GetDiscoveryScanstatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDiscoveryScanstatusParamsWithTimeout creates a new GetDiscoveryScanstatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDiscoveryScanstatusParamsWithTimeout(timeout time.Duration) *GetDiscoveryScanstatusParams {

	return &GetDiscoveryScanstatusParams{

		timeout: timeout,
	}
}

// NewGetDiscoveryScanstatusParamsWithContext creates a new GetDiscoveryScanstatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDiscoveryScanstatusParamsWithContext(ctx context.Context) *GetDiscoveryScanstatusParams {

	return &GetDiscoveryScanstatusParams{

		Context: ctx,
	}
}

// NewGetDiscoveryScanstatusParamsWithHTTPClient creates a new GetDiscoveryScanstatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDiscoveryScanstatusParamsWithHTTPClient(client *http.Client) *GetDiscoveryScanstatusParams {

	return &GetDiscoveryScanstatusParams{
		HTTPClient: client,
	}
}

/*GetDiscoveryScanstatusParams contains all the parameters to send to the API endpoint
for the get discovery scanstatus operation typically these are written to a http.Request
*/
type GetDiscoveryScanstatusParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get discovery scanstatus params
func (o *GetDiscoveryScanstatusParams) WithTimeout(timeout time.Duration) *GetDiscoveryScanstatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get discovery scanstatus params
func (o *GetDiscoveryScanstatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get discovery scanstatus params
func (o *GetDiscoveryScanstatusParams) WithContext(ctx context.Context) *GetDiscoveryScanstatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get discovery scanstatus params
func (o *GetDiscoveryScanstatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get discovery scanstatus params
func (o *GetDiscoveryScanstatusParams) WithHTTPClient(client *http.Client) *GetDiscoveryScanstatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get discovery scanstatus params
func (o *GetDiscoveryScanstatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetDiscoveryScanstatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
