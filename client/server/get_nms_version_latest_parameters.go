// Code generated by go-swagger; DO NOT EDIT.

package server

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

// NewGetNmsVersionLatestParams creates a new GetNmsVersionLatestParams object
// with the default values initialized.
func NewGetNmsVersionLatestParams() *GetNmsVersionLatestParams {

	return &GetNmsVersionLatestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNmsVersionLatestParamsWithTimeout creates a new GetNmsVersionLatestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNmsVersionLatestParamsWithTimeout(timeout time.Duration) *GetNmsVersionLatestParams {

	return &GetNmsVersionLatestParams{

		timeout: timeout,
	}
}

// NewGetNmsVersionLatestParamsWithContext creates a new GetNmsVersionLatestParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNmsVersionLatestParamsWithContext(ctx context.Context) *GetNmsVersionLatestParams {

	return &GetNmsVersionLatestParams{

		Context: ctx,
	}
}

// NewGetNmsVersionLatestParamsWithHTTPClient creates a new GetNmsVersionLatestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNmsVersionLatestParamsWithHTTPClient(client *http.Client) *GetNmsVersionLatestParams {

	return &GetNmsVersionLatestParams{
		HTTPClient: client,
	}
}

/*GetNmsVersionLatestParams contains all the parameters to send to the API endpoint
for the get nms version latest operation typically these are written to a http.Request
*/
type GetNmsVersionLatestParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get nms version latest params
func (o *GetNmsVersionLatestParams) WithTimeout(timeout time.Duration) *GetNmsVersionLatestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nms version latest params
func (o *GetNmsVersionLatestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nms version latest params
func (o *GetNmsVersionLatestParams) WithContext(ctx context.Context) *GetNmsVersionLatestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nms version latest params
func (o *GetNmsVersionLatestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nms version latest params
func (o *GetNmsVersionLatestParams) WithHTTPClient(client *http.Client) *GetNmsVersionLatestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nms version latest params
func (o *GetNmsVersionLatestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetNmsVersionLatestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
