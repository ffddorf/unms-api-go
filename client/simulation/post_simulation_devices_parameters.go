// Code generated by go-swagger; DO NOT EDIT.

package simulation

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

// NewPostSimulationDevicesParams creates a new PostSimulationDevicesParams object
// with the default values initialized.
func NewPostSimulationDevicesParams() *PostSimulationDevicesParams {
	var ()
	return &PostSimulationDevicesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSimulationDevicesParamsWithTimeout creates a new PostSimulationDevicesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSimulationDevicesParamsWithTimeout(timeout time.Duration) *PostSimulationDevicesParams {
	var ()
	return &PostSimulationDevicesParams{

		timeout: timeout,
	}
}

// NewPostSimulationDevicesParamsWithContext creates a new PostSimulationDevicesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSimulationDevicesParamsWithContext(ctx context.Context) *PostSimulationDevicesParams {
	var ()
	return &PostSimulationDevicesParams{

		Context: ctx,
	}
}

// NewPostSimulationDevicesParamsWithHTTPClient creates a new PostSimulationDevicesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSimulationDevicesParamsWithHTTPClient(client *http.Client) *PostSimulationDevicesParams {
	var ()
	return &PostSimulationDevicesParams{
		HTTPClient: client,
	}
}

/*PostSimulationDevicesParams contains all the parameters to send to the API endpoint
for the post simulation devices operation typically these are written to a http.Request
*/
type PostSimulationDevicesParams struct {

	/*Body*/
	Body *models.SimulationDevicePayload
	/*XAuthToken
	  Token for user authorization.

	*/
	XAuthToken *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post simulation devices params
func (o *PostSimulationDevicesParams) WithTimeout(timeout time.Duration) *PostSimulationDevicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post simulation devices params
func (o *PostSimulationDevicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post simulation devices params
func (o *PostSimulationDevicesParams) WithContext(ctx context.Context) *PostSimulationDevicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post simulation devices params
func (o *PostSimulationDevicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post simulation devices params
func (o *PostSimulationDevicesParams) WithHTTPClient(client *http.Client) *PostSimulationDevicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post simulation devices params
func (o *PostSimulationDevicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post simulation devices params
func (o *PostSimulationDevicesParams) WithBody(body *models.SimulationDevicePayload) *PostSimulationDevicesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post simulation devices params
func (o *PostSimulationDevicesParams) SetBody(body *models.SimulationDevicePayload) {
	o.Body = body
}

// WithXAuthToken adds the xAuthToken to the post simulation devices params
func (o *PostSimulationDevicesParams) WithXAuthToken(xAuthToken *string) *PostSimulationDevicesParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the post simulation devices params
func (o *PostSimulationDevicesParams) SetXAuthToken(xAuthToken *string) {
	o.XAuthToken = xAuthToken
}

// WriteToRequest writes these params to a swagger request
func (o *PostSimulationDevicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.XAuthToken != nil {

		// header param x-auth-token
		if err := r.SetHeaderParam("x-auth-token", *o.XAuthToken); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
