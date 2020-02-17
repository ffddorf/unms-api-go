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

	strfmt "github.com/go-openapi/strfmt"

	"github.com/ffddorf/unms-api-go/models"
)

// NewPostDevicesMaintenanceDisableParams creates a new PostDevicesMaintenanceDisableParams object
// with the default values initialized.
func NewPostDevicesMaintenanceDisableParams() *PostDevicesMaintenanceDisableParams {
	var ()
	return &PostDevicesMaintenanceDisableParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostDevicesMaintenanceDisableParamsWithTimeout creates a new PostDevicesMaintenanceDisableParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostDevicesMaintenanceDisableParamsWithTimeout(timeout time.Duration) *PostDevicesMaintenanceDisableParams {
	var ()
	return &PostDevicesMaintenanceDisableParams{

		timeout: timeout,
	}
}

// NewPostDevicesMaintenanceDisableParamsWithContext creates a new PostDevicesMaintenanceDisableParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostDevicesMaintenanceDisableParamsWithContext(ctx context.Context) *PostDevicesMaintenanceDisableParams {
	var ()
	return &PostDevicesMaintenanceDisableParams{

		Context: ctx,
	}
}

// NewPostDevicesMaintenanceDisableParamsWithHTTPClient creates a new PostDevicesMaintenanceDisableParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostDevicesMaintenanceDisableParamsWithHTTPClient(client *http.Client) *PostDevicesMaintenanceDisableParams {
	var ()
	return &PostDevicesMaintenanceDisableParams{
		HTTPClient: client,
	}
}

/*PostDevicesMaintenanceDisableParams contains all the parameters to send to the API endpoint
for the post devices maintenance disable operation typically these are written to a http.Request
*/
type PostDevicesMaintenanceDisableParams struct {

	/*Body*/
	Body *models.MaintenanceModeSchema

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post devices maintenance disable params
func (o *PostDevicesMaintenanceDisableParams) WithTimeout(timeout time.Duration) *PostDevicesMaintenanceDisableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post devices maintenance disable params
func (o *PostDevicesMaintenanceDisableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post devices maintenance disable params
func (o *PostDevicesMaintenanceDisableParams) WithContext(ctx context.Context) *PostDevicesMaintenanceDisableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post devices maintenance disable params
func (o *PostDevicesMaintenanceDisableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post devices maintenance disable params
func (o *PostDevicesMaintenanceDisableParams) WithHTTPClient(client *http.Client) *PostDevicesMaintenanceDisableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post devices maintenance disable params
func (o *PostDevicesMaintenanceDisableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post devices maintenance disable params
func (o *PostDevicesMaintenanceDisableParams) WithBody(body *models.MaintenanceModeSchema) *PostDevicesMaintenanceDisableParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post devices maintenance disable params
func (o *PostDevicesMaintenanceDisableParams) SetBody(body *models.MaintenanceModeSchema) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostDevicesMaintenanceDisableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
