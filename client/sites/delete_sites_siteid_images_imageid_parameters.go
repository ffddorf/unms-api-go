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

// NewDeleteSitesSiteidImagesImageidParams creates a new DeleteSitesSiteidImagesImageidParams object
// with the default values initialized.
func NewDeleteSitesSiteidImagesImageidParams() *DeleteSitesSiteidImagesImageidParams {
	var ()
	return &DeleteSitesSiteidImagesImageidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSitesSiteidImagesImageidParamsWithTimeout creates a new DeleteSitesSiteidImagesImageidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSitesSiteidImagesImageidParamsWithTimeout(timeout time.Duration) *DeleteSitesSiteidImagesImageidParams {
	var ()
	return &DeleteSitesSiteidImagesImageidParams{

		timeout: timeout,
	}
}

// NewDeleteSitesSiteidImagesImageidParamsWithContext creates a new DeleteSitesSiteidImagesImageidParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSitesSiteidImagesImageidParamsWithContext(ctx context.Context) *DeleteSitesSiteidImagesImageidParams {
	var ()
	return &DeleteSitesSiteidImagesImageidParams{

		Context: ctx,
	}
}

// NewDeleteSitesSiteidImagesImageidParamsWithHTTPClient creates a new DeleteSitesSiteidImagesImageidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSitesSiteidImagesImageidParamsWithHTTPClient(client *http.Client) *DeleteSitesSiteidImagesImageidParams {
	var ()
	return &DeleteSitesSiteidImagesImageidParams{
		HTTPClient: client,
	}
}

/*DeleteSitesSiteidImagesImageidParams contains all the parameters to send to the API endpoint
for the delete sites siteid images imageid operation typically these are written to a http.Request
*/
type DeleteSitesSiteidImagesImageidParams struct {

	/*ImageID*/
	ImageID string
	/*SiteID*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete sites siteid images imageid params
func (o *DeleteSitesSiteidImagesImageidParams) WithTimeout(timeout time.Duration) *DeleteSitesSiteidImagesImageidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete sites siteid images imageid params
func (o *DeleteSitesSiteidImagesImageidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete sites siteid images imageid params
func (o *DeleteSitesSiteidImagesImageidParams) WithContext(ctx context.Context) *DeleteSitesSiteidImagesImageidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete sites siteid images imageid params
func (o *DeleteSitesSiteidImagesImageidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete sites siteid images imageid params
func (o *DeleteSitesSiteidImagesImageidParams) WithHTTPClient(client *http.Client) *DeleteSitesSiteidImagesImageidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete sites siteid images imageid params
func (o *DeleteSitesSiteidImagesImageidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImageID adds the imageID to the delete sites siteid images imageid params
func (o *DeleteSitesSiteidImagesImageidParams) WithImageID(imageID string) *DeleteSitesSiteidImagesImageidParams {
	o.SetImageID(imageID)
	return o
}

// SetImageID adds the imageId to the delete sites siteid images imageid params
func (o *DeleteSitesSiteidImagesImageidParams) SetImageID(imageID string) {
	o.ImageID = imageID
}

// WithSiteID adds the siteID to the delete sites siteid images imageid params
func (o *DeleteSitesSiteidImagesImageidParams) WithSiteID(siteID string) *DeleteSitesSiteidImagesImageidParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the delete sites siteid images imageid params
func (o *DeleteSitesSiteidImagesImageidParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSitesSiteidImagesImageidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param imageId
	if err := r.SetPathParam("imageId", o.ImageID); err != nil {
		return err
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
