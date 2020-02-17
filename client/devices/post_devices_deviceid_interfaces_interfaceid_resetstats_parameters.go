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

// NewPostDevicesDeviceidInterfacesInterfaceidResetstatsParams creates a new PostDevicesDeviceidInterfacesInterfaceidResetstatsParams object
// with the default values initialized.
func NewPostDevicesDeviceidInterfacesInterfaceidResetstatsParams() *PostDevicesDeviceidInterfacesInterfaceidResetstatsParams {
	var ()
	return &PostDevicesDeviceidInterfacesInterfaceidResetstatsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostDevicesDeviceidInterfacesInterfaceidResetstatsParamsWithTimeout creates a new PostDevicesDeviceidInterfacesInterfaceidResetstatsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostDevicesDeviceidInterfacesInterfaceidResetstatsParamsWithTimeout(timeout time.Duration) *PostDevicesDeviceidInterfacesInterfaceidResetstatsParams {
	var ()
	return &PostDevicesDeviceidInterfacesInterfaceidResetstatsParams{

		timeout: timeout,
	}
}

// NewPostDevicesDeviceidInterfacesInterfaceidResetstatsParamsWithContext creates a new PostDevicesDeviceidInterfacesInterfaceidResetstatsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostDevicesDeviceidInterfacesInterfaceidResetstatsParamsWithContext(ctx context.Context) *PostDevicesDeviceidInterfacesInterfaceidResetstatsParams {
	var ()
	return &PostDevicesDeviceidInterfacesInterfaceidResetstatsParams{

		Context: ctx,
	}
}

// NewPostDevicesDeviceidInterfacesInterfaceidResetstatsParamsWithHTTPClient creates a new PostDevicesDeviceidInterfacesInterfaceidResetstatsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostDevicesDeviceidInterfacesInterfaceidResetstatsParamsWithHTTPClient(client *http.Client) *PostDevicesDeviceidInterfacesInterfaceidResetstatsParams {
	var ()
	return &PostDevicesDeviceidInterfacesInterfaceidResetstatsParams{
		HTTPClient: client,
	}
}

/*PostDevicesDeviceidInterfacesInterfaceidResetstatsParams contains all the parameters to send to the API endpoint
for the post devices deviceid interfaces interfaceid resetstats operation typically these are written to a http.Request
*/
type PostDevicesDeviceidInterfacesInterfaceidResetstatsParams struct {

	/*DeviceID*/
	DeviceID string
	/*InterfaceID*/
	InterfaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post devices deviceid interfaces interfaceid resetstats params
func (o *PostDevicesDeviceidInterfacesInterfaceidResetstatsParams) WithTimeout(timeout time.Duration) *PostDevicesDeviceidInterfacesInterfaceidResetstatsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post devices deviceid interfaces interfaceid resetstats params
func (o *PostDevicesDeviceidInterfacesInterfaceidResetstatsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post devices deviceid interfaces interfaceid resetstats params
func (o *PostDevicesDeviceidInterfacesInterfaceidResetstatsParams) WithContext(ctx context.Context) *PostDevicesDeviceidInterfacesInterfaceidResetstatsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post devices deviceid interfaces interfaceid resetstats params
func (o *PostDevicesDeviceidInterfacesInterfaceidResetstatsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post devices deviceid interfaces interfaceid resetstats params
func (o *PostDevicesDeviceidInterfacesInterfaceidResetstatsParams) WithHTTPClient(client *http.Client) *PostDevicesDeviceidInterfacesInterfaceidResetstatsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post devices deviceid interfaces interfaceid resetstats params
func (o *PostDevicesDeviceidInterfacesInterfaceidResetstatsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the post devices deviceid interfaces interfaceid resetstats params
func (o *PostDevicesDeviceidInterfacesInterfaceidResetstatsParams) WithDeviceID(deviceID string) *PostDevicesDeviceidInterfacesInterfaceidResetstatsParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the post devices deviceid interfaces interfaceid resetstats params
func (o *PostDevicesDeviceidInterfacesInterfaceidResetstatsParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WithInterfaceID adds the interfaceID to the post devices deviceid interfaces interfaceid resetstats params
func (o *PostDevicesDeviceidInterfacesInterfaceidResetstatsParams) WithInterfaceID(interfaceID string) *PostDevicesDeviceidInterfacesInterfaceidResetstatsParams {
	o.SetInterfaceID(interfaceID)
	return o
}

// SetInterfaceID adds the interfaceId to the post devices deviceid interfaces interfaceid resetstats params
func (o *PostDevicesDeviceidInterfacesInterfaceidResetstatsParams) SetInterfaceID(interfaceID string) {
	o.InterfaceID = interfaceID
}

// WriteToRequest writes these params to a swagger request
func (o *PostDevicesDeviceidInterfacesInterfaceidResetstatsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deviceId
	if err := r.SetPathParam("deviceId", o.DeviceID); err != nil {
		return err
	}

	// path param interfaceId
	if err := r.SetPathParam("interfaceId", o.InterfaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
