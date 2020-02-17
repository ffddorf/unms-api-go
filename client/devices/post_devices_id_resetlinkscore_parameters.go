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

// NewPostDevicesIDResetlinkscoreParams creates a new PostDevicesIDResetlinkscoreParams object
// with the default values initialized.
func NewPostDevicesIDResetlinkscoreParams() *PostDevicesIDResetlinkscoreParams {
	var ()
	return &PostDevicesIDResetlinkscoreParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostDevicesIDResetlinkscoreParamsWithTimeout creates a new PostDevicesIDResetlinkscoreParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostDevicesIDResetlinkscoreParamsWithTimeout(timeout time.Duration) *PostDevicesIDResetlinkscoreParams {
	var ()
	return &PostDevicesIDResetlinkscoreParams{

		timeout: timeout,
	}
}

// NewPostDevicesIDResetlinkscoreParamsWithContext creates a new PostDevicesIDResetlinkscoreParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostDevicesIDResetlinkscoreParamsWithContext(ctx context.Context) *PostDevicesIDResetlinkscoreParams {
	var ()
	return &PostDevicesIDResetlinkscoreParams{

		Context: ctx,
	}
}

// NewPostDevicesIDResetlinkscoreParamsWithHTTPClient creates a new PostDevicesIDResetlinkscoreParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostDevicesIDResetlinkscoreParamsWithHTTPClient(client *http.Client) *PostDevicesIDResetlinkscoreParams {
	var ()
	return &PostDevicesIDResetlinkscoreParams{
		HTTPClient: client,
	}
}

/*PostDevicesIDResetlinkscoreParams contains all the parameters to send to the API endpoint
for the post devices Id resetlinkscore operation typically these are written to a http.Request
*/
type PostDevicesIDResetlinkscoreParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post devices Id resetlinkscore params
func (o *PostDevicesIDResetlinkscoreParams) WithTimeout(timeout time.Duration) *PostDevicesIDResetlinkscoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post devices Id resetlinkscore params
func (o *PostDevicesIDResetlinkscoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post devices Id resetlinkscore params
func (o *PostDevicesIDResetlinkscoreParams) WithContext(ctx context.Context) *PostDevicesIDResetlinkscoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post devices Id resetlinkscore params
func (o *PostDevicesIDResetlinkscoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post devices Id resetlinkscore params
func (o *PostDevicesIDResetlinkscoreParams) WithHTTPClient(client *http.Client) *PostDevicesIDResetlinkscoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post devices Id resetlinkscore params
func (o *PostDevicesIDResetlinkscoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post devices Id resetlinkscore params
func (o *PostDevicesIDResetlinkscoreParams) WithID(id string) *PostDevicesIDResetlinkscoreParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post devices Id resetlinkscore params
func (o *PostDevicesIDResetlinkscoreParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostDevicesIDResetlinkscoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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