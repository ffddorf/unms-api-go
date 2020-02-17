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

	"github.com/ffddorf/unms-api-go/models"
)

// NewPostSitesSiteidUcrmBindParams creates a new PostSitesSiteidUcrmBindParams object
// with the default values initialized.
func NewPostSitesSiteidUcrmBindParams() *PostSitesSiteidUcrmBindParams {
	var ()
	return &PostSitesSiteidUcrmBindParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSitesSiteidUcrmBindParamsWithTimeout creates a new PostSitesSiteidUcrmBindParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSitesSiteidUcrmBindParamsWithTimeout(timeout time.Duration) *PostSitesSiteidUcrmBindParams {
	var ()
	return &PostSitesSiteidUcrmBindParams{

		timeout: timeout,
	}
}

// NewPostSitesSiteidUcrmBindParamsWithContext creates a new PostSitesSiteidUcrmBindParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSitesSiteidUcrmBindParamsWithContext(ctx context.Context) *PostSitesSiteidUcrmBindParams {
	var ()
	return &PostSitesSiteidUcrmBindParams{

		Context: ctx,
	}
}

// NewPostSitesSiteidUcrmBindParamsWithHTTPClient creates a new PostSitesSiteidUcrmBindParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSitesSiteidUcrmBindParamsWithHTTPClient(client *http.Client) *PostSitesSiteidUcrmBindParams {
	var ()
	return &PostSitesSiteidUcrmBindParams{
		HTTPClient: client,
	}
}

/*PostSitesSiteidUcrmBindParams contains all the parameters to send to the API endpoint
for the post sites siteid ucrm bind operation typically these are written to a http.Request
*/
type PostSitesSiteidUcrmBindParams struct {

	/*Body*/
	Body *models.SiteBinding
	/*SiteID*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post sites siteid ucrm bind params
func (o *PostSitesSiteidUcrmBindParams) WithTimeout(timeout time.Duration) *PostSitesSiteidUcrmBindParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post sites siteid ucrm bind params
func (o *PostSitesSiteidUcrmBindParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post sites siteid ucrm bind params
func (o *PostSitesSiteidUcrmBindParams) WithContext(ctx context.Context) *PostSitesSiteidUcrmBindParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post sites siteid ucrm bind params
func (o *PostSitesSiteidUcrmBindParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post sites siteid ucrm bind params
func (o *PostSitesSiteidUcrmBindParams) WithHTTPClient(client *http.Client) *PostSitesSiteidUcrmBindParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post sites siteid ucrm bind params
func (o *PostSitesSiteidUcrmBindParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post sites siteid ucrm bind params
func (o *PostSitesSiteidUcrmBindParams) WithBody(body *models.SiteBinding) *PostSitesSiteidUcrmBindParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post sites siteid ucrm bind params
func (o *PostSitesSiteidUcrmBindParams) SetBody(body *models.SiteBinding) {
	o.Body = body
}

// WithSiteID adds the siteID to the post sites siteid ucrm bind params
func (o *PostSitesSiteidUcrmBindParams) WithSiteID(siteID string) *PostSitesSiteidUcrmBindParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the post sites siteid ucrm bind params
func (o *PostSitesSiteidUcrmBindParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *PostSitesSiteidUcrmBindParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
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