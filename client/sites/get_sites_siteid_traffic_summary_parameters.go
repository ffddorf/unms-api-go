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

// NewGetSitesSiteidTrafficSummaryParams creates a new GetSitesSiteidTrafficSummaryParams object
// with the default values initialized.
func NewGetSitesSiteidTrafficSummaryParams() *GetSitesSiteidTrafficSummaryParams {
	var ()
	return &GetSitesSiteidTrafficSummaryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSitesSiteidTrafficSummaryParamsWithTimeout creates a new GetSitesSiteidTrafficSummaryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSitesSiteidTrafficSummaryParamsWithTimeout(timeout time.Duration) *GetSitesSiteidTrafficSummaryParams {
	var ()
	return &GetSitesSiteidTrafficSummaryParams{

		timeout: timeout,
	}
}

// NewGetSitesSiteidTrafficSummaryParamsWithContext creates a new GetSitesSiteidTrafficSummaryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSitesSiteidTrafficSummaryParamsWithContext(ctx context.Context) *GetSitesSiteidTrafficSummaryParams {
	var ()
	return &GetSitesSiteidTrafficSummaryParams{

		Context: ctx,
	}
}

// NewGetSitesSiteidTrafficSummaryParamsWithHTTPClient creates a new GetSitesSiteidTrafficSummaryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSitesSiteidTrafficSummaryParamsWithHTTPClient(client *http.Client) *GetSitesSiteidTrafficSummaryParams {
	var ()
	return &GetSitesSiteidTrafficSummaryParams{
		HTTPClient: client,
	}
}

/*GetSitesSiteidTrafficSummaryParams contains all the parameters to send to the API endpoint
for the get sites siteid traffic summary operation typically these are written to a http.Request
*/
type GetSitesSiteidTrafficSummaryParams struct {

	/*Interval
	  Interval for which to fetch traffic stats

	*/
	Interval string
	/*SiteID*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get sites siteid traffic summary params
func (o *GetSitesSiteidTrafficSummaryParams) WithTimeout(timeout time.Duration) *GetSitesSiteidTrafficSummaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sites siteid traffic summary params
func (o *GetSitesSiteidTrafficSummaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sites siteid traffic summary params
func (o *GetSitesSiteidTrafficSummaryParams) WithContext(ctx context.Context) *GetSitesSiteidTrafficSummaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sites siteid traffic summary params
func (o *GetSitesSiteidTrafficSummaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sites siteid traffic summary params
func (o *GetSitesSiteidTrafficSummaryParams) WithHTTPClient(client *http.Client) *GetSitesSiteidTrafficSummaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sites siteid traffic summary params
func (o *GetSitesSiteidTrafficSummaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInterval adds the interval to the get sites siteid traffic summary params
func (o *GetSitesSiteidTrafficSummaryParams) WithInterval(interval string) *GetSitesSiteidTrafficSummaryParams {
	o.SetInterval(interval)
	return o
}

// SetInterval adds the interval to the get sites siteid traffic summary params
func (o *GetSitesSiteidTrafficSummaryParams) SetInterval(interval string) {
	o.Interval = interval
}

// WithSiteID adds the siteID to the get sites siteid traffic summary params
func (o *GetSitesSiteidTrafficSummaryParams) WithSiteID(siteID string) *GetSitesSiteidTrafficSummaryParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the get sites siteid traffic summary params
func (o *GetSitesSiteidTrafficSummaryParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSitesSiteidTrafficSummaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param interval
	qrInterval := o.Interval
	qInterval := qrInterval
	if qInterval != "" {
		if err := r.SetQueryParam("interval", qInterval); err != nil {
			return err
		}
	}

	// path param siteId
	if err := r.SetPathParam("siteId", o.SiteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
