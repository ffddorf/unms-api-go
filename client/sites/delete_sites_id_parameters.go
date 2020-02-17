// Code generated by go-swagger; DO NOT EDIT.

package sites

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

// NewDeleteSitesIDParams creates a new DeleteSitesIDParams object
// with the default values initialized.
func NewDeleteSitesIDParams() *DeleteSitesIDParams {
	var ()
	return &DeleteSitesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSitesIDParamsWithTimeout creates a new DeleteSitesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSitesIDParamsWithTimeout(timeout time.Duration) *DeleteSitesIDParams {
	var ()
	return &DeleteSitesIDParams{

		timeout: timeout,
	}
}

// NewDeleteSitesIDParamsWithContext creates a new DeleteSitesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSitesIDParamsWithContext(ctx context.Context) *DeleteSitesIDParams {
	var ()
	return &DeleteSitesIDParams{

		Context: ctx,
	}
}

// NewDeleteSitesIDParamsWithHTTPClient creates a new DeleteSitesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSitesIDParamsWithHTTPClient(client *http.Client) *DeleteSitesIDParams {
	var ()
	return &DeleteSitesIDParams{
		HTTPClient: client,
	}
}

/*DeleteSitesIDParams contains all the parameters to send to the API endpoint
for the delete sites Id operation typically these are written to a http.Request
*/
type DeleteSitesIDParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete sites Id params
func (o *DeleteSitesIDParams) WithTimeout(timeout time.Duration) *DeleteSitesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete sites Id params
func (o *DeleteSitesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete sites Id params
func (o *DeleteSitesIDParams) WithContext(ctx context.Context) *DeleteSitesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete sites Id params
func (o *DeleteSitesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete sites Id params
func (o *DeleteSitesIDParams) WithHTTPClient(client *http.Client) *DeleteSitesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete sites Id params
func (o *DeleteSitesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete sites Id params
func (o *DeleteSitesIDParams) WithID(id string) *DeleteSitesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete sites Id params
func (o *DeleteSitesIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSitesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
