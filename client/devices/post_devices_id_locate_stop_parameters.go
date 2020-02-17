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

// NewPostDevicesIDLocateStopParams creates a new PostDevicesIDLocateStopParams object
// with the default values initialized.
func NewPostDevicesIDLocateStopParams() *PostDevicesIDLocateStopParams {
	var ()
	return &PostDevicesIDLocateStopParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostDevicesIDLocateStopParamsWithTimeout creates a new PostDevicesIDLocateStopParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostDevicesIDLocateStopParamsWithTimeout(timeout time.Duration) *PostDevicesIDLocateStopParams {
	var ()
	return &PostDevicesIDLocateStopParams{

		timeout: timeout,
	}
}

// NewPostDevicesIDLocateStopParamsWithContext creates a new PostDevicesIDLocateStopParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostDevicesIDLocateStopParamsWithContext(ctx context.Context) *PostDevicesIDLocateStopParams {
	var ()
	return &PostDevicesIDLocateStopParams{

		Context: ctx,
	}
}

// NewPostDevicesIDLocateStopParamsWithHTTPClient creates a new PostDevicesIDLocateStopParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostDevicesIDLocateStopParamsWithHTTPClient(client *http.Client) *PostDevicesIDLocateStopParams {
	var ()
	return &PostDevicesIDLocateStopParams{
		HTTPClient: client,
	}
}

/*PostDevicesIDLocateStopParams contains all the parameters to send to the API endpoint
for the post devices Id locate stop operation typically these are written to a http.Request
*/
type PostDevicesIDLocateStopParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post devices Id locate stop params
func (o *PostDevicesIDLocateStopParams) WithTimeout(timeout time.Duration) *PostDevicesIDLocateStopParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post devices Id locate stop params
func (o *PostDevicesIDLocateStopParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post devices Id locate stop params
func (o *PostDevicesIDLocateStopParams) WithContext(ctx context.Context) *PostDevicesIDLocateStopParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post devices Id locate stop params
func (o *PostDevicesIDLocateStopParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post devices Id locate stop params
func (o *PostDevicesIDLocateStopParams) WithHTTPClient(client *http.Client) *PostDevicesIDLocateStopParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post devices Id locate stop params
func (o *PostDevicesIDLocateStopParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post devices Id locate stop params
func (o *PostDevicesIDLocateStopParams) WithID(id string) *PostDevicesIDLocateStopParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post devices Id locate stop params
func (o *PostDevicesIDLocateStopParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostDevicesIDLocateStopParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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