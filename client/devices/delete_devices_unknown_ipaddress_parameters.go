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

// NewDeleteDevicesUnknownIpaddressParams creates a new DeleteDevicesUnknownIpaddressParams object
// with the default values initialized.
func NewDeleteDevicesUnknownIpaddressParams() *DeleteDevicesUnknownIpaddressParams {
	var ()
	return &DeleteDevicesUnknownIpaddressParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDevicesUnknownIpaddressParamsWithTimeout creates a new DeleteDevicesUnknownIpaddressParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDevicesUnknownIpaddressParamsWithTimeout(timeout time.Duration) *DeleteDevicesUnknownIpaddressParams {
	var ()
	return &DeleteDevicesUnknownIpaddressParams{

		timeout: timeout,
	}
}

// NewDeleteDevicesUnknownIpaddressParamsWithContext creates a new DeleteDevicesUnknownIpaddressParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDevicesUnknownIpaddressParamsWithContext(ctx context.Context) *DeleteDevicesUnknownIpaddressParams {
	var ()
	return &DeleteDevicesUnknownIpaddressParams{

		Context: ctx,
	}
}

// NewDeleteDevicesUnknownIpaddressParamsWithHTTPClient creates a new DeleteDevicesUnknownIpaddressParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDevicesUnknownIpaddressParamsWithHTTPClient(client *http.Client) *DeleteDevicesUnknownIpaddressParams {
	var ()
	return &DeleteDevicesUnknownIpaddressParams{
		HTTPClient: client,
	}
}

/*DeleteDevicesUnknownIpaddressParams contains all the parameters to send to the API endpoint
for the delete devices unknown ipaddress operation typically these are written to a http.Request
*/
type DeleteDevicesUnknownIpaddressParams struct {

	/*IPAddress*/
	IPAddress string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete devices unknown ipaddress params
func (o *DeleteDevicesUnknownIpaddressParams) WithTimeout(timeout time.Duration) *DeleteDevicesUnknownIpaddressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete devices unknown ipaddress params
func (o *DeleteDevicesUnknownIpaddressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete devices unknown ipaddress params
func (o *DeleteDevicesUnknownIpaddressParams) WithContext(ctx context.Context) *DeleteDevicesUnknownIpaddressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete devices unknown ipaddress params
func (o *DeleteDevicesUnknownIpaddressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete devices unknown ipaddress params
func (o *DeleteDevicesUnknownIpaddressParams) WithHTTPClient(client *http.Client) *DeleteDevicesUnknownIpaddressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete devices unknown ipaddress params
func (o *DeleteDevicesUnknownIpaddressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIPAddress adds the iPAddress to the delete devices unknown ipaddress params
func (o *DeleteDevicesUnknownIpaddressParams) WithIPAddress(iPAddress string) *DeleteDevicesUnknownIpaddressParams {
	o.SetIPAddress(iPAddress)
	return o
}

// SetIPAddress adds the ipAddress to the delete devices unknown ipaddress params
func (o *DeleteDevicesUnknownIpaddressParams) SetIPAddress(iPAddress string) {
	o.IPAddress = iPAddress
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDevicesUnknownIpaddressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ipAddress
	if err := r.SetPathParam("ipAddress", o.IPAddress); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
