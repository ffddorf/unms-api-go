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

// NewPostSitesIDImagesParams creates a new PostSitesIDImagesParams object
// with the default values initialized.
func NewPostSitesIDImagesParams() *PostSitesIDImagesParams {
	var ()
	return &PostSitesIDImagesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSitesIDImagesParamsWithTimeout creates a new PostSitesIDImagesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSitesIDImagesParamsWithTimeout(timeout time.Duration) *PostSitesIDImagesParams {
	var ()
	return &PostSitesIDImagesParams{

		timeout: timeout,
	}
}

// NewPostSitesIDImagesParamsWithContext creates a new PostSitesIDImagesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSitesIDImagesParamsWithContext(ctx context.Context) *PostSitesIDImagesParams {
	var ()
	return &PostSitesIDImagesParams{

		Context: ctx,
	}
}

// NewPostSitesIDImagesParamsWithHTTPClient creates a new PostSitesIDImagesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSitesIDImagesParamsWithHTTPClient(client *http.Client) *PostSitesIDImagesParams {
	var ()
	return &PostSitesIDImagesParams{
		HTTPClient: client,
	}
}

/*PostSitesIDImagesParams contains all the parameters to send to the API endpoint
for the post sites Id images operation typically these are written to a http.Request
*/
type PostSitesIDImagesParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post sites Id images params
func (o *PostSitesIDImagesParams) WithTimeout(timeout time.Duration) *PostSitesIDImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post sites Id images params
func (o *PostSitesIDImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post sites Id images params
func (o *PostSitesIDImagesParams) WithContext(ctx context.Context) *PostSitesIDImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post sites Id images params
func (o *PostSitesIDImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post sites Id images params
func (o *PostSitesIDImagesParams) WithHTTPClient(client *http.Client) *PostSitesIDImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post sites Id images params
func (o *PostSitesIDImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post sites Id images params
func (o *PostSitesIDImagesParams) WithID(id string) *PostSitesIDImagesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post sites Id images params
func (o *PostSitesIDImagesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostSitesIDImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
