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
)

// NewDeleteSimulationDevicesIDParams creates a new DeleteSimulationDevicesIDParams object
// with the default values initialized.
func NewDeleteSimulationDevicesIDParams() *DeleteSimulationDevicesIDParams {
	var ()
	return &DeleteSimulationDevicesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSimulationDevicesIDParamsWithTimeout creates a new DeleteSimulationDevicesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSimulationDevicesIDParamsWithTimeout(timeout time.Duration) *DeleteSimulationDevicesIDParams {
	var ()
	return &DeleteSimulationDevicesIDParams{

		timeout: timeout,
	}
}

// NewDeleteSimulationDevicesIDParamsWithContext creates a new DeleteSimulationDevicesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSimulationDevicesIDParamsWithContext(ctx context.Context) *DeleteSimulationDevicesIDParams {
	var ()
	return &DeleteSimulationDevicesIDParams{

		Context: ctx,
	}
}

// NewDeleteSimulationDevicesIDParamsWithHTTPClient creates a new DeleteSimulationDevicesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSimulationDevicesIDParamsWithHTTPClient(client *http.Client) *DeleteSimulationDevicesIDParams {
	var ()
	return &DeleteSimulationDevicesIDParams{
		HTTPClient: client,
	}
}

/*DeleteSimulationDevicesIDParams contains all the parameters to send to the API endpoint
for the delete simulation devices Id operation typically these are written to a http.Request
*/
type DeleteSimulationDevicesIDParams struct {

	/*ID*/
	ID string
	/*XAuthToken
	  Token for user authorization.

	*/
	XAuthToken *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete simulation devices Id params
func (o *DeleteSimulationDevicesIDParams) WithTimeout(timeout time.Duration) *DeleteSimulationDevicesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete simulation devices Id params
func (o *DeleteSimulationDevicesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete simulation devices Id params
func (o *DeleteSimulationDevicesIDParams) WithContext(ctx context.Context) *DeleteSimulationDevicesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete simulation devices Id params
func (o *DeleteSimulationDevicesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete simulation devices Id params
func (o *DeleteSimulationDevicesIDParams) WithHTTPClient(client *http.Client) *DeleteSimulationDevicesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete simulation devices Id params
func (o *DeleteSimulationDevicesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete simulation devices Id params
func (o *DeleteSimulationDevicesIDParams) WithID(id string) *DeleteSimulationDevicesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete simulation devices Id params
func (o *DeleteSimulationDevicesIDParams) SetID(id string) {
	o.ID = id
}

// WithXAuthToken adds the xAuthToken to the delete simulation devices Id params
func (o *DeleteSimulationDevicesIDParams) WithXAuthToken(xAuthToken *string) *DeleteSimulationDevicesIDParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the delete simulation devices Id params
func (o *DeleteSimulationDevicesIDParams) SetXAuthToken(xAuthToken *string) {
	o.XAuthToken = xAuthToken
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSimulationDevicesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
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