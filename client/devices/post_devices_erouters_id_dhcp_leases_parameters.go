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

// NewPostDevicesEroutersIDDhcpLeasesParams creates a new PostDevicesEroutersIDDhcpLeasesParams object
// with the default values initialized.
func NewPostDevicesEroutersIDDhcpLeasesParams() *PostDevicesEroutersIDDhcpLeasesParams {
	var ()
	return &PostDevicesEroutersIDDhcpLeasesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostDevicesEroutersIDDhcpLeasesParamsWithTimeout creates a new PostDevicesEroutersIDDhcpLeasesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostDevicesEroutersIDDhcpLeasesParamsWithTimeout(timeout time.Duration) *PostDevicesEroutersIDDhcpLeasesParams {
	var ()
	return &PostDevicesEroutersIDDhcpLeasesParams{

		timeout: timeout,
	}
}

// NewPostDevicesEroutersIDDhcpLeasesParamsWithContext creates a new PostDevicesEroutersIDDhcpLeasesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostDevicesEroutersIDDhcpLeasesParamsWithContext(ctx context.Context) *PostDevicesEroutersIDDhcpLeasesParams {
	var ()
	return &PostDevicesEroutersIDDhcpLeasesParams{

		Context: ctx,
	}
}

// NewPostDevicesEroutersIDDhcpLeasesParamsWithHTTPClient creates a new PostDevicesEroutersIDDhcpLeasesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostDevicesEroutersIDDhcpLeasesParamsWithHTTPClient(client *http.Client) *PostDevicesEroutersIDDhcpLeasesParams {
	var ()
	return &PostDevicesEroutersIDDhcpLeasesParams{
		HTTPClient: client,
	}
}

/*PostDevicesEroutersIDDhcpLeasesParams contains all the parameters to send to the API endpoint
for the post devices erouters Id dhcp leases operation typically these are written to a http.Request
*/
type PostDevicesEroutersIDDhcpLeasesParams struct {

	/*Body*/
	Body *models.DHCPLease1
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post devices erouters Id dhcp leases params
func (o *PostDevicesEroutersIDDhcpLeasesParams) WithTimeout(timeout time.Duration) *PostDevicesEroutersIDDhcpLeasesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post devices erouters Id dhcp leases params
func (o *PostDevicesEroutersIDDhcpLeasesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post devices erouters Id dhcp leases params
func (o *PostDevicesEroutersIDDhcpLeasesParams) WithContext(ctx context.Context) *PostDevicesEroutersIDDhcpLeasesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post devices erouters Id dhcp leases params
func (o *PostDevicesEroutersIDDhcpLeasesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post devices erouters Id dhcp leases params
func (o *PostDevicesEroutersIDDhcpLeasesParams) WithHTTPClient(client *http.Client) *PostDevicesEroutersIDDhcpLeasesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post devices erouters Id dhcp leases params
func (o *PostDevicesEroutersIDDhcpLeasesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post devices erouters Id dhcp leases params
func (o *PostDevicesEroutersIDDhcpLeasesParams) WithBody(body *models.DHCPLease1) *PostDevicesEroutersIDDhcpLeasesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post devices erouters Id dhcp leases params
func (o *PostDevicesEroutersIDDhcpLeasesParams) SetBody(body *models.DHCPLease1) {
	o.Body = body
}

// WithID adds the id to the post devices erouters Id dhcp leases params
func (o *PostDevicesEroutersIDDhcpLeasesParams) WithID(id string) *PostDevicesEroutersIDDhcpLeasesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post devices erouters Id dhcp leases params
func (o *PostDevicesEroutersIDDhcpLeasesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostDevicesEroutersIDDhcpLeasesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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