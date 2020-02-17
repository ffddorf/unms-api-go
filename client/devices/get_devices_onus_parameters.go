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

// NewGetDevicesOnusParams creates a new GetDevicesOnusParams object
// with the default values initialized.
func NewGetDevicesOnusParams() *GetDevicesOnusParams {
	var ()
	return &GetDevicesOnusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDevicesOnusParamsWithTimeout creates a new GetDevicesOnusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDevicesOnusParamsWithTimeout(timeout time.Duration) *GetDevicesOnusParams {
	var ()
	return &GetDevicesOnusParams{

		timeout: timeout,
	}
}

// NewGetDevicesOnusParamsWithContext creates a new GetDevicesOnusParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDevicesOnusParamsWithContext(ctx context.Context) *GetDevicesOnusParams {
	var ()
	return &GetDevicesOnusParams{

		Context: ctx,
	}
}

// NewGetDevicesOnusParamsWithHTTPClient creates a new GetDevicesOnusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDevicesOnusParamsWithHTTPClient(client *http.Client) *GetDevicesOnusParams {
	var ()
	return &GetDevicesOnusParams{
		HTTPClient: client,
	}
}

/*GetDevicesOnusParams contains all the parameters to send to the API endpoint
for the get devices onus operation typically these are written to a http.Request
*/
type GetDevicesOnusParams struct {

	/*ParentID*/
	ParentID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get devices onus params
func (o *GetDevicesOnusParams) WithTimeout(timeout time.Duration) *GetDevicesOnusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get devices onus params
func (o *GetDevicesOnusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get devices onus params
func (o *GetDevicesOnusParams) WithContext(ctx context.Context) *GetDevicesOnusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get devices onus params
func (o *GetDevicesOnusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get devices onus params
func (o *GetDevicesOnusParams) WithHTTPClient(client *http.Client) *GetDevicesOnusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get devices onus params
func (o *GetDevicesOnusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithParentID adds the parentID to the get devices onus params
func (o *GetDevicesOnusParams) WithParentID(parentID *string) *GetDevicesOnusParams {
	o.SetParentID(parentID)
	return o
}

// SetParentID adds the parentId to the get devices onus params
func (o *GetDevicesOnusParams) SetParentID(parentID *string) {
	o.ParentID = parentID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDevicesOnusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ParentID != nil {

		// query param parentId
		var qrParentID string
		if o.ParentID != nil {
			qrParentID = *o.ParentID
		}
		qParentID := qrParentID
		if qParentID != "" {
			if err := r.SetQueryParam("parentId", qParentID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
