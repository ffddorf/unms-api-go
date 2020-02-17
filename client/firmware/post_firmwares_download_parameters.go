// Code generated by go-swagger; DO NOT EDIT.

package firmware

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

// NewPostFirmwaresDownloadParams creates a new PostFirmwaresDownloadParams object
// with the default values initialized.
func NewPostFirmwaresDownloadParams() *PostFirmwaresDownloadParams {

	return &PostFirmwaresDownloadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostFirmwaresDownloadParamsWithTimeout creates a new PostFirmwaresDownloadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostFirmwaresDownloadParamsWithTimeout(timeout time.Duration) *PostFirmwaresDownloadParams {

	return &PostFirmwaresDownloadParams{

		timeout: timeout,
	}
}

// NewPostFirmwaresDownloadParamsWithContext creates a new PostFirmwaresDownloadParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostFirmwaresDownloadParamsWithContext(ctx context.Context) *PostFirmwaresDownloadParams {

	return &PostFirmwaresDownloadParams{

		Context: ctx,
	}
}

// NewPostFirmwaresDownloadParamsWithHTTPClient creates a new PostFirmwaresDownloadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostFirmwaresDownloadParamsWithHTTPClient(client *http.Client) *PostFirmwaresDownloadParams {

	return &PostFirmwaresDownloadParams{
		HTTPClient: client,
	}
}

/*PostFirmwaresDownloadParams contains all the parameters to send to the API endpoint
for the post firmwares download operation typically these are written to a http.Request
*/
type PostFirmwaresDownloadParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post firmwares download params
func (o *PostFirmwaresDownloadParams) WithTimeout(timeout time.Duration) *PostFirmwaresDownloadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post firmwares download params
func (o *PostFirmwaresDownloadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post firmwares download params
func (o *PostFirmwaresDownloadParams) WithContext(ctx context.Context) *PostFirmwaresDownloadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post firmwares download params
func (o *PostFirmwaresDownloadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post firmwares download params
func (o *PostFirmwaresDownloadParams) WithHTTPClient(client *http.Client) *PostFirmwaresDownloadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post firmwares download params
func (o *PostFirmwaresDownloadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PostFirmwaresDownloadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
