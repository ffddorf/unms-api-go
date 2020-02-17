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

// NewGetDatalinksSitesSiteidParams creates a new GetDatalinksSitesSiteidParams object
// with the default values initialized.
func NewGetDatalinksSitesSiteidParams() *GetDatalinksSitesSiteidParams {
	var ()
	return &GetDatalinksSitesSiteidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDatalinksSitesSiteidParamsWithTimeout creates a new GetDatalinksSitesSiteidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDatalinksSitesSiteidParamsWithTimeout(timeout time.Duration) *GetDatalinksSitesSiteidParams {
	var ()
	return &GetDatalinksSitesSiteidParams{

		timeout: timeout,
	}
}

// NewGetDatalinksSitesSiteidParamsWithContext creates a new GetDatalinksSitesSiteidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDatalinksSitesSiteidParamsWithContext(ctx context.Context) *GetDatalinksSitesSiteidParams {
	var ()
	return &GetDatalinksSitesSiteidParams{

		Context: ctx,
	}
}

// NewGetDatalinksSitesSiteidParamsWithHTTPClient creates a new GetDatalinksSitesSiteidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDatalinksSitesSiteidParamsWithHTTPClient(client *http.Client) *GetDatalinksSitesSiteidParams {
	var ()
	return &GetDatalinksSitesSiteidParams{
		HTTPClient: client,
	}
}

/*GetDatalinksSitesSiteidParams contains all the parameters to send to the API endpoint
for the get datalinks sites siteid operation typically these are written to a http.Request
*/
type GetDatalinksSitesSiteidParams struct {

	/*SiteID*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get datalinks sites siteid params
func (o *GetDatalinksSitesSiteidParams) WithTimeout(timeout time.Duration) *GetDatalinksSitesSiteidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get datalinks sites siteid params
func (o *GetDatalinksSitesSiteidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get datalinks sites siteid params
func (o *GetDatalinksSitesSiteidParams) WithContext(ctx context.Context) *GetDatalinksSitesSiteidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get datalinks sites siteid params
func (o *GetDatalinksSitesSiteidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get datalinks sites siteid params
func (o *GetDatalinksSitesSiteidParams) WithHTTPClient(client *http.Client) *GetDatalinksSitesSiteidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get datalinks sites siteid params
func (o *GetDatalinksSitesSiteidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSiteID adds the siteID to the get datalinks sites siteid params
func (o *GetDatalinksSitesSiteidParams) WithSiteID(siteID string) *GetDatalinksSitesSiteidParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the get datalinks sites siteid params
func (o *GetDatalinksSitesSiteidParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDatalinksSitesSiteidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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