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

// NewGetDiscoveryStatusDeviceidParams creates a new GetDiscoveryStatusDeviceidParams object
// with the default values initialized.
func NewGetDiscoveryStatusDeviceidParams() *GetDiscoveryStatusDeviceidParams {

	return &GetDiscoveryStatusDeviceidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDiscoveryStatusDeviceidParamsWithTimeout creates a new GetDiscoveryStatusDeviceidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDiscoveryStatusDeviceidParamsWithTimeout(timeout time.Duration) *GetDiscoveryStatusDeviceidParams {

	return &GetDiscoveryStatusDeviceidParams{

		timeout: timeout,
	}
}

// NewGetDiscoveryStatusDeviceidParamsWithContext creates a new GetDiscoveryStatusDeviceidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDiscoveryStatusDeviceidParamsWithContext(ctx context.Context) *GetDiscoveryStatusDeviceidParams {

	return &GetDiscoveryStatusDeviceidParams{

		Context: ctx,
	}
}

// NewGetDiscoveryStatusDeviceidParamsWithHTTPClient creates a new GetDiscoveryStatusDeviceidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDiscoveryStatusDeviceidParamsWithHTTPClient(client *http.Client) *GetDiscoveryStatusDeviceidParams {

	return &GetDiscoveryStatusDeviceidParams{
		HTTPClient: client,
	}
}

/*GetDiscoveryStatusDeviceidParams contains all the parameters to send to the API endpoint
for the get discovery status deviceid operation typically these are written to a http.Request
*/
type GetDiscoveryStatusDeviceidParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get discovery status deviceid params
func (o *GetDiscoveryStatusDeviceidParams) WithTimeout(timeout time.Duration) *GetDiscoveryStatusDeviceidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get discovery status deviceid params
func (o *GetDiscoveryStatusDeviceidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get discovery status deviceid params
func (o *GetDiscoveryStatusDeviceidParams) WithContext(ctx context.Context) *GetDiscoveryStatusDeviceidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get discovery status deviceid params
func (o *GetDiscoveryStatusDeviceidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get discovery status deviceid params
func (o *GetDiscoveryStatusDeviceidParams) WithHTTPClient(client *http.Client) *GetDiscoveryStatusDeviceidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get discovery status deviceid params
func (o *GetDiscoveryStatusDeviceidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetDiscoveryStatusDeviceidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
