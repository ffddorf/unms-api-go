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

// NewPostDevicesOltsIDOnuProfilesParams creates a new PostDevicesOltsIDOnuProfilesParams object
// with the default values initialized.
func NewPostDevicesOltsIDOnuProfilesParams() *PostDevicesOltsIDOnuProfilesParams {
	var ()
	return &PostDevicesOltsIDOnuProfilesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostDevicesOltsIDOnuProfilesParamsWithTimeout creates a new PostDevicesOltsIDOnuProfilesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostDevicesOltsIDOnuProfilesParamsWithTimeout(timeout time.Duration) *PostDevicesOltsIDOnuProfilesParams {
	var ()
	return &PostDevicesOltsIDOnuProfilesParams{

		timeout: timeout,
	}
}

// NewPostDevicesOltsIDOnuProfilesParamsWithContext creates a new PostDevicesOltsIDOnuProfilesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostDevicesOltsIDOnuProfilesParamsWithContext(ctx context.Context) *PostDevicesOltsIDOnuProfilesParams {
	var ()
	return &PostDevicesOltsIDOnuProfilesParams{

		Context: ctx,
	}
}

// NewPostDevicesOltsIDOnuProfilesParamsWithHTTPClient creates a new PostDevicesOltsIDOnuProfilesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostDevicesOltsIDOnuProfilesParamsWithHTTPClient(client *http.Client) *PostDevicesOltsIDOnuProfilesParams {
	var ()
	return &PostDevicesOltsIDOnuProfilesParams{
		HTTPClient: client,
	}
}

/*PostDevicesOltsIDOnuProfilesParams contains all the parameters to send to the API endpoint
for the post devices olts Id onu profiles operation typically these are written to a http.Request
*/
type PostDevicesOltsIDOnuProfilesParams struct {

	/*Body*/
	Body *models.OnuProfile1
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post devices olts Id onu profiles params
func (o *PostDevicesOltsIDOnuProfilesParams) WithTimeout(timeout time.Duration) *PostDevicesOltsIDOnuProfilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post devices olts Id onu profiles params
func (o *PostDevicesOltsIDOnuProfilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post devices olts Id onu profiles params
func (o *PostDevicesOltsIDOnuProfilesParams) WithContext(ctx context.Context) *PostDevicesOltsIDOnuProfilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post devices olts Id onu profiles params
func (o *PostDevicesOltsIDOnuProfilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post devices olts Id onu profiles params
func (o *PostDevicesOltsIDOnuProfilesParams) WithHTTPClient(client *http.Client) *PostDevicesOltsIDOnuProfilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post devices olts Id onu profiles params
func (o *PostDevicesOltsIDOnuProfilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post devices olts Id onu profiles params
func (o *PostDevicesOltsIDOnuProfilesParams) WithBody(body *models.OnuProfile1) *PostDevicesOltsIDOnuProfilesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post devices olts Id onu profiles params
func (o *PostDevicesOltsIDOnuProfilesParams) SetBody(body *models.OnuProfile1) {
	o.Body = body
}

// WithID adds the id to the post devices olts Id onu profiles params
func (o *PostDevicesOltsIDOnuProfilesParams) WithID(id string) *PostDevicesOltsIDOnuProfilesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post devices olts Id onu profiles params
func (o *PostDevicesOltsIDOnuProfilesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostDevicesOltsIDOnuProfilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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