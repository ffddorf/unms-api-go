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

// NewGetNmsStatisticsParams creates a new GetNmsStatisticsParams object
// with the default values initialized.
func NewGetNmsStatisticsParams() *GetNmsStatisticsParams {

	return &GetNmsStatisticsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNmsStatisticsParamsWithTimeout creates a new GetNmsStatisticsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNmsStatisticsParamsWithTimeout(timeout time.Duration) *GetNmsStatisticsParams {

	return &GetNmsStatisticsParams{

		timeout: timeout,
	}
}

// NewGetNmsStatisticsParamsWithContext creates a new GetNmsStatisticsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNmsStatisticsParamsWithContext(ctx context.Context) *GetNmsStatisticsParams {

	return &GetNmsStatisticsParams{

		Context: ctx,
	}
}

// NewGetNmsStatisticsParamsWithHTTPClient creates a new GetNmsStatisticsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNmsStatisticsParamsWithHTTPClient(client *http.Client) *GetNmsStatisticsParams {

	return &GetNmsStatisticsParams{
		HTTPClient: client,
	}
}

/*GetNmsStatisticsParams contains all the parameters to send to the API endpoint
for the get nms statistics operation typically these are written to a http.Request
*/
type GetNmsStatisticsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get nms statistics params
func (o *GetNmsStatisticsParams) WithTimeout(timeout time.Duration) *GetNmsStatisticsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nms statistics params
func (o *GetNmsStatisticsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nms statistics params
func (o *GetNmsStatisticsParams) WithContext(ctx context.Context) *GetNmsStatisticsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nms statistics params
func (o *GetNmsStatisticsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nms statistics params
func (o *GetNmsStatisticsParams) WithHTTPClient(client *http.Client) *GetNmsStatisticsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nms statistics params
func (o *GetNmsStatisticsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetNmsStatisticsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}