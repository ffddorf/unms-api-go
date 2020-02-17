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

// NewPutSitesSiteidQosParams creates a new PutSitesSiteidQosParams object
// with the default values initialized.
func NewPutSitesSiteidQosParams() *PutSitesSiteidQosParams {
	var ()
	return &PutSitesSiteidQosParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutSitesSiteidQosParamsWithTimeout creates a new PutSitesSiteidQosParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutSitesSiteidQosParamsWithTimeout(timeout time.Duration) *PutSitesSiteidQosParams {
	var ()
	return &PutSitesSiteidQosParams{

		timeout: timeout,
	}
}

// NewPutSitesSiteidQosParamsWithContext creates a new PutSitesSiteidQosParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutSitesSiteidQosParamsWithContext(ctx context.Context) *PutSitesSiteidQosParams {
	var ()
	return &PutSitesSiteidQosParams{

		Context: ctx,
	}
}

// NewPutSitesSiteidQosParamsWithHTTPClient creates a new PutSitesSiteidQosParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutSitesSiteidQosParamsWithHTTPClient(client *http.Client) *PutSitesSiteidQosParams {
	var ()
	return &PutSitesSiteidQosParams{
		HTTPClient: client,
	}
}

/*PutSitesSiteidQosParams contains all the parameters to send to the API endpoint
for the put sites siteid qos operation typically these are written to a http.Request
*/
type PutSitesSiteidQosParams struct {

	/*Body*/
	Body *models.SiteTrafficShaping1
	/*SiteID*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put sites siteid qos params
func (o *PutSitesSiteidQosParams) WithTimeout(timeout time.Duration) *PutSitesSiteidQosParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put sites siteid qos params
func (o *PutSitesSiteidQosParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put sites siteid qos params
func (o *PutSitesSiteidQosParams) WithContext(ctx context.Context) *PutSitesSiteidQosParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put sites siteid qos params
func (o *PutSitesSiteidQosParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put sites siteid qos params
func (o *PutSitesSiteidQosParams) WithHTTPClient(client *http.Client) *PutSitesSiteidQosParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put sites siteid qos params
func (o *PutSitesSiteidQosParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put sites siteid qos params
func (o *PutSitesSiteidQosParams) WithBody(body *models.SiteTrafficShaping1) *PutSitesSiteidQosParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put sites siteid qos params
func (o *PutSitesSiteidQosParams) SetBody(body *models.SiteTrafficShaping1) {
	o.Body = body
}

// WithSiteID adds the siteID to the put sites siteid qos params
func (o *PutSitesSiteidQosParams) WithSiteID(siteID string) *PutSitesSiteidQosParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the put sites siteid qos params
func (o *PutSitesSiteidQosParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *PutSitesSiteidQosParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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