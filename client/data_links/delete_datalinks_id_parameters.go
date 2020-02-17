// Code generated by go-swagger; DO NOT EDIT.

package data_links

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

// NewDeleteDatalinksIDParams creates a new DeleteDatalinksIDParams object
// with the default values initialized.
func NewDeleteDatalinksIDParams() *DeleteDatalinksIDParams {
	var ()
	return &DeleteDatalinksIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDatalinksIDParamsWithTimeout creates a new DeleteDatalinksIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDatalinksIDParamsWithTimeout(timeout time.Duration) *DeleteDatalinksIDParams {
	var ()
	return &DeleteDatalinksIDParams{

		timeout: timeout,
	}
}

// NewDeleteDatalinksIDParamsWithContext creates a new DeleteDatalinksIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDatalinksIDParamsWithContext(ctx context.Context) *DeleteDatalinksIDParams {
	var ()
	return &DeleteDatalinksIDParams{

		Context: ctx,
	}
}

// NewDeleteDatalinksIDParamsWithHTTPClient creates a new DeleteDatalinksIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDatalinksIDParamsWithHTTPClient(client *http.Client) *DeleteDatalinksIDParams {
	var ()
	return &DeleteDatalinksIDParams{
		HTTPClient: client,
	}
}

/*DeleteDatalinksIDParams contains all the parameters to send to the API endpoint
for the delete datalinks Id operation typically these are written to a http.Request
*/
type DeleteDatalinksIDParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete datalinks Id params
func (o *DeleteDatalinksIDParams) WithTimeout(timeout time.Duration) *DeleteDatalinksIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete datalinks Id params
func (o *DeleteDatalinksIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete datalinks Id params
func (o *DeleteDatalinksIDParams) WithContext(ctx context.Context) *DeleteDatalinksIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete datalinks Id params
func (o *DeleteDatalinksIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete datalinks Id params
func (o *DeleteDatalinksIDParams) WithHTTPClient(client *http.Client) *DeleteDatalinksIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete datalinks Id params
func (o *DeleteDatalinksIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete datalinks Id params
func (o *DeleteDatalinksIDParams) WithID(id string) *DeleteDatalinksIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete datalinks Id params
func (o *DeleteDatalinksIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDatalinksIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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