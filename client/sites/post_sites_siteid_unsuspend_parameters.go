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

// NewPostSitesSiteidUnsuspendParams creates a new PostSitesSiteidUnsuspendParams object
// with the default values initialized.
func NewPostSitesSiteidUnsuspendParams() *PostSitesSiteidUnsuspendParams {
	var ()
	return &PostSitesSiteidUnsuspendParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSitesSiteidUnsuspendParamsWithTimeout creates a new PostSitesSiteidUnsuspendParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSitesSiteidUnsuspendParamsWithTimeout(timeout time.Duration) *PostSitesSiteidUnsuspendParams {
	var ()
	return &PostSitesSiteidUnsuspendParams{

		timeout: timeout,
	}
}

// NewPostSitesSiteidUnsuspendParamsWithContext creates a new PostSitesSiteidUnsuspendParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSitesSiteidUnsuspendParamsWithContext(ctx context.Context) *PostSitesSiteidUnsuspendParams {
	var ()
	return &PostSitesSiteidUnsuspendParams{

		Context: ctx,
	}
}

// NewPostSitesSiteidUnsuspendParamsWithHTTPClient creates a new PostSitesSiteidUnsuspendParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSitesSiteidUnsuspendParamsWithHTTPClient(client *http.Client) *PostSitesSiteidUnsuspendParams {
	var ()
	return &PostSitesSiteidUnsuspendParams{
		HTTPClient: client,
	}
}

/*PostSitesSiteidUnsuspendParams contains all the parameters to send to the API endpoint
for the post sites siteid unsuspend operation typically these are written to a http.Request
*/
type PostSitesSiteidUnsuspendParams struct {

	/*SiteID*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post sites siteid unsuspend params
func (o *PostSitesSiteidUnsuspendParams) WithTimeout(timeout time.Duration) *PostSitesSiteidUnsuspendParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post sites siteid unsuspend params
func (o *PostSitesSiteidUnsuspendParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post sites siteid unsuspend params
func (o *PostSitesSiteidUnsuspendParams) WithContext(ctx context.Context) *PostSitesSiteidUnsuspendParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post sites siteid unsuspend params
func (o *PostSitesSiteidUnsuspendParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post sites siteid unsuspend params
func (o *PostSitesSiteidUnsuspendParams) WithHTTPClient(client *http.Client) *PostSitesSiteidUnsuspendParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post sites siteid unsuspend params
func (o *PostSitesSiteidUnsuspendParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSiteID adds the siteID to the post sites siteid unsuspend params
func (o *PostSitesSiteidUnsuspendParams) WithSiteID(siteID string) *PostSitesSiteidUnsuspendParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the post sites siteid unsuspend params
func (o *PostSitesSiteidUnsuspendParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *PostSitesSiteidUnsuspendParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param siteId
	if err := r.SetPathParam("siteId", o.SiteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}