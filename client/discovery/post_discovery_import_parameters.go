// Code generated by go-swagger; DO NOT EDIT.

package discovery

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

// NewPostDiscoveryImportParams creates a new PostDiscoveryImportParams object
// with the default values initialized.
func NewPostDiscoveryImportParams() *PostDiscoveryImportParams {
	var ()
	return &PostDiscoveryImportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostDiscoveryImportParamsWithTimeout creates a new PostDiscoveryImportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostDiscoveryImportParamsWithTimeout(timeout time.Duration) *PostDiscoveryImportParams {
	var ()
	return &PostDiscoveryImportParams{

		timeout: timeout,
	}
}

// NewPostDiscoveryImportParamsWithContext creates a new PostDiscoveryImportParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostDiscoveryImportParamsWithContext(ctx context.Context) *PostDiscoveryImportParams {
	var ()
	return &PostDiscoveryImportParams{

		Context: ctx,
	}
}

// NewPostDiscoveryImportParamsWithHTTPClient creates a new PostDiscoveryImportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostDiscoveryImportParamsWithHTTPClient(client *http.Client) *PostDiscoveryImportParams {
	var ()
	return &PostDiscoveryImportParams{
		HTTPClient: client,
	}
}

/*PostDiscoveryImportParams contains all the parameters to send to the API endpoint
for the post discovery import operation typically these are written to a http.Request
*/
type PostDiscoveryImportParams struct {

	/*Body*/
	Body models.DiscoveryImportList

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post discovery import params
func (o *PostDiscoveryImportParams) WithTimeout(timeout time.Duration) *PostDiscoveryImportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post discovery import params
func (o *PostDiscoveryImportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post discovery import params
func (o *PostDiscoveryImportParams) WithContext(ctx context.Context) *PostDiscoveryImportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post discovery import params
func (o *PostDiscoveryImportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post discovery import params
func (o *PostDiscoveryImportParams) WithHTTPClient(client *http.Client) *PostDiscoveryImportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post discovery import params
func (o *PostDiscoveryImportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post discovery import params
func (o *PostDiscoveryImportParams) WithBody(body models.DiscoveryImportList) *PostDiscoveryImportParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post discovery import params
func (o *PostDiscoveryImportParams) SetBody(body models.DiscoveryImportList) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostDiscoveryImportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
