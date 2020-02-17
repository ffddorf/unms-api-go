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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetDevicesIpsParams creates a new GetDevicesIpsParams object
// with the default values initialized.
func NewGetDevicesIpsParams() *GetDevicesIpsParams {
	var ()
	return &GetDevicesIpsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDevicesIpsParamsWithTimeout creates a new GetDevicesIpsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDevicesIpsParamsWithTimeout(timeout time.Duration) *GetDevicesIpsParams {
	var ()
	return &GetDevicesIpsParams{

		timeout: timeout,
	}
}

// NewGetDevicesIpsParamsWithContext creates a new GetDevicesIpsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDevicesIpsParamsWithContext(ctx context.Context) *GetDevicesIpsParams {
	var ()
	return &GetDevicesIpsParams{

		Context: ctx,
	}
}

// NewGetDevicesIpsParamsWithHTTPClient creates a new GetDevicesIpsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDevicesIpsParamsWithHTTPClient(client *http.Client) *GetDevicesIpsParams {
	var ()
	return &GetDevicesIpsParams{
		HTTPClient: client,
	}
}

/*GetDevicesIpsParams contains all the parameters to send to the API endpoint
for the get devices ips operation typically these are written to a http.Request
*/
type GetDevicesIpsParams struct {

	/*Management
	  Return also management IPs.

	*/
	Management *bool
	/*SiteID
	  Return only IPs for this UNMS site.

	*/
	SiteID *string
	/*Suspended
	  Return only IPs which are blocked.

	*/
	Suspended *bool
	/*UcrmID
	  Return only IPs for this UCRM service.

	*/
	UcrmID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get devices ips params
func (o *GetDevicesIpsParams) WithTimeout(timeout time.Duration) *GetDevicesIpsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get devices ips params
func (o *GetDevicesIpsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get devices ips params
func (o *GetDevicesIpsParams) WithContext(ctx context.Context) *GetDevicesIpsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get devices ips params
func (o *GetDevicesIpsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get devices ips params
func (o *GetDevicesIpsParams) WithHTTPClient(client *http.Client) *GetDevicesIpsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get devices ips params
func (o *GetDevicesIpsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithManagement adds the management to the get devices ips params
func (o *GetDevicesIpsParams) WithManagement(management *bool) *GetDevicesIpsParams {
	o.SetManagement(management)
	return o
}

// SetManagement adds the management to the get devices ips params
func (o *GetDevicesIpsParams) SetManagement(management *bool) {
	o.Management = management
}

// WithSiteID adds the siteID to the get devices ips params
func (o *GetDevicesIpsParams) WithSiteID(siteID *string) *GetDevicesIpsParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the get devices ips params
func (o *GetDevicesIpsParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithSuspended adds the suspended to the get devices ips params
func (o *GetDevicesIpsParams) WithSuspended(suspended *bool) *GetDevicesIpsParams {
	o.SetSuspended(suspended)
	return o
}

// SetSuspended adds the suspended to the get devices ips params
func (o *GetDevicesIpsParams) SetSuspended(suspended *bool) {
	o.Suspended = suspended
}

// WithUcrmID adds the ucrmID to the get devices ips params
func (o *GetDevicesIpsParams) WithUcrmID(ucrmID *string) *GetDevicesIpsParams {
	o.SetUcrmID(ucrmID)
	return o
}

// SetUcrmID adds the ucrmId to the get devices ips params
func (o *GetDevicesIpsParams) SetUcrmID(ucrmID *string) {
	o.UcrmID = ucrmID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDevicesIpsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Management != nil {

		// query param management
		var qrManagement bool
		if o.Management != nil {
			qrManagement = *o.Management
		}
		qManagement := swag.FormatBool(qrManagement)
		if qManagement != "" {
			if err := r.SetQueryParam("management", qManagement); err != nil {
				return err
			}
		}

	}

	if o.SiteID != nil {

		// query param siteId
		var qrSiteID string
		if o.SiteID != nil {
			qrSiteID = *o.SiteID
		}
		qSiteID := qrSiteID
		if qSiteID != "" {
			if err := r.SetQueryParam("siteId", qSiteID); err != nil {
				return err
			}
		}

	}

	if o.Suspended != nil {

		// query param suspended
		var qrSuspended bool
		if o.Suspended != nil {
			qrSuspended = *o.Suspended
		}
		qSuspended := swag.FormatBool(qrSuspended)
		if qSuspended != "" {
			if err := r.SetQueryParam("suspended", qSuspended); err != nil {
				return err
			}
		}

	}

	if o.UcrmID != nil {

		// query param ucrmId
		var qrUcrmID string
		if o.UcrmID != nil {
			qrUcrmID = *o.UcrmID
		}
		qUcrmID := qrUcrmID
		if qUcrmID != "" {
			if err := r.SetQueryParam("ucrmId", qUcrmID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
