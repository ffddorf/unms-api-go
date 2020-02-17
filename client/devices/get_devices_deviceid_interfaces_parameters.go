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

// NewGetDevicesDeviceidInterfacesParams creates a new GetDevicesDeviceidInterfacesParams object
// with the default values initialized.
func NewGetDevicesDeviceidInterfacesParams() *GetDevicesDeviceidInterfacesParams {
	var ()
	return &GetDevicesDeviceidInterfacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDevicesDeviceidInterfacesParamsWithTimeout creates a new GetDevicesDeviceidInterfacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDevicesDeviceidInterfacesParamsWithTimeout(timeout time.Duration) *GetDevicesDeviceidInterfacesParams {
	var ()
	return &GetDevicesDeviceidInterfacesParams{

		timeout: timeout,
	}
}

// NewGetDevicesDeviceidInterfacesParamsWithContext creates a new GetDevicesDeviceidInterfacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDevicesDeviceidInterfacesParamsWithContext(ctx context.Context) *GetDevicesDeviceidInterfacesParams {
	var ()
	return &GetDevicesDeviceidInterfacesParams{

		Context: ctx,
	}
}

// NewGetDevicesDeviceidInterfacesParamsWithHTTPClient creates a new GetDevicesDeviceidInterfacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDevicesDeviceidInterfacesParamsWithHTTPClient(client *http.Client) *GetDevicesDeviceidInterfacesParams {
	var ()
	return &GetDevicesDeviceidInterfacesParams{
		HTTPClient: client,
	}
}

/*GetDevicesDeviceidInterfacesParams contains all the parameters to send to the API endpoint
for the get devices deviceid interfaces operation typically these are written to a http.Request
*/
type GetDevicesDeviceidInterfacesParams struct {

	/*DeviceID*/
	DeviceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get devices deviceid interfaces params
func (o *GetDevicesDeviceidInterfacesParams) WithTimeout(timeout time.Duration) *GetDevicesDeviceidInterfacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get devices deviceid interfaces params
func (o *GetDevicesDeviceidInterfacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get devices deviceid interfaces params
func (o *GetDevicesDeviceidInterfacesParams) WithContext(ctx context.Context) *GetDevicesDeviceidInterfacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get devices deviceid interfaces params
func (o *GetDevicesDeviceidInterfacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get devices deviceid interfaces params
func (o *GetDevicesDeviceidInterfacesParams) WithHTTPClient(client *http.Client) *GetDevicesDeviceidInterfacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get devices deviceid interfaces params
func (o *GetDevicesDeviceidInterfacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the get devices deviceid interfaces params
func (o *GetDevicesDeviceidInterfacesParams) WithDeviceID(deviceID string) *GetDevicesDeviceidInterfacesParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get devices deviceid interfaces params
func (o *GetDevicesDeviceidInterfacesParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDevicesDeviceidInterfacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deviceId
	if err := r.SetPathParam("deviceId", o.DeviceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
