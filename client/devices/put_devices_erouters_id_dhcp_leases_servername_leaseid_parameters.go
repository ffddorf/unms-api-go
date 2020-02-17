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

// NewPutDevicesEroutersIDDhcpLeasesServernameLeaseidParams creates a new PutDevicesEroutersIDDhcpLeasesServernameLeaseidParams object
// with the default values initialized.
func NewPutDevicesEroutersIDDhcpLeasesServernameLeaseidParams() *PutDevicesEroutersIDDhcpLeasesServernameLeaseidParams {
	var ()
	return &PutDevicesEroutersIDDhcpLeasesServernameLeaseidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutDevicesEroutersIDDhcpLeasesServernameLeaseidParamsWithTimeout creates a new PutDevicesEroutersIDDhcpLeasesServernameLeaseidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutDevicesEroutersIDDhcpLeasesServernameLeaseidParamsWithTimeout(timeout time.Duration) *PutDevicesEroutersIDDhcpLeasesServernameLeaseidParams {
	var ()
	return &PutDevicesEroutersIDDhcpLeasesServernameLeaseidParams{

		timeout: timeout,
	}
}

// NewPutDevicesEroutersIDDhcpLeasesServernameLeaseidParamsWithContext creates a new PutDevicesEroutersIDDhcpLeasesServernameLeaseidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutDevicesEroutersIDDhcpLeasesServernameLeaseidParamsWithContext(ctx context.Context) *PutDevicesEroutersIDDhcpLeasesServernameLeaseidParams {
	var ()
	return &PutDevicesEroutersIDDhcpLeasesServernameLeaseidParams{

		Context: ctx,
	}
}

// NewPutDevicesEroutersIDDhcpLeasesServernameLeaseidParamsWithHTTPClient creates a new PutDevicesEroutersIDDhcpLeasesServernameLeaseidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutDevicesEroutersIDDhcpLeasesServernameLeaseidParamsWithHTTPClient(client *http.Client) *PutDevicesEroutersIDDhcpLeasesServernameLeaseidParams {
	var ()
	return &PutDevicesEroutersIDDhcpLeasesServernameLeaseidParams{
		HTTPClient: client,
	}
}

/*PutDevicesEroutersIDDhcpLeasesServernameLeaseidParams contains all the parameters to send to the API endpoint
for the put devices erouters Id dhcp leases servername leaseid operation typically these are written to a http.Request
*/
type PutDevicesEroutersIDDhcpLeasesServernameLeaseidParams struct {

	/*Body*/
	Body *models.DHCPLease1
	/*ID*/
	ID string
	/*LeaseID*/
	LeaseID string
	/*ServerName*/
	ServerName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put devices erouters Id dhcp leases servername leaseid params
func (o *PutDevicesEroutersIDDhcpLeasesServernameLeaseidParams) WithTimeout(timeout time.Duration) *PutDevicesEroutersIDDhcpLeasesServernameLeaseidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put devices erouters Id dhcp leases servername leaseid params
func (o *PutDevicesEroutersIDDhcpLeasesServernameLeaseidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put devices erouters Id dhcp leases servername leaseid params
func (o *PutDevicesEroutersIDDhcpLeasesServernameLeaseidParams) WithContext(ctx context.Context) *PutDevicesEroutersIDDhcpLeasesServernameLeaseidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put devices erouters Id dhcp leases servername leaseid params
func (o *PutDevicesEroutersIDDhcpLeasesServernameLeaseidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put devices erouters Id dhcp leases servername leaseid params
func (o *PutDevicesEroutersIDDhcpLeasesServernameLeaseidParams) WithHTTPClient(client *http.Client) *PutDevicesEroutersIDDhcpLeasesServernameLeaseidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put devices erouters Id dhcp leases servername leaseid params
func (o *PutDevicesEroutersIDDhcpLeasesServernameLeaseidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put devices erouters Id dhcp leases servername leaseid params
func (o *PutDevicesEroutersIDDhcpLeasesServernameLeaseidParams) WithBody(body *models.DHCPLease1) *PutDevicesEroutersIDDhcpLeasesServernameLeaseidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put devices erouters Id dhcp leases servername leaseid params
func (o *PutDevicesEroutersIDDhcpLeasesServernameLeaseidParams) SetBody(body *models.DHCPLease1) {
	o.Body = body
}

// WithID adds the id to the put devices erouters Id dhcp leases servername leaseid params
func (o *PutDevicesEroutersIDDhcpLeasesServernameLeaseidParams) WithID(id string) *PutDevicesEroutersIDDhcpLeasesServernameLeaseidParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put devices erouters Id dhcp leases servername leaseid params
func (o *PutDevicesEroutersIDDhcpLeasesServernameLeaseidParams) SetID(id string) {
	o.ID = id
}

// WithLeaseID adds the leaseID to the put devices erouters Id dhcp leases servername leaseid params
func (o *PutDevicesEroutersIDDhcpLeasesServernameLeaseidParams) WithLeaseID(leaseID string) *PutDevicesEroutersIDDhcpLeasesServernameLeaseidParams {
	o.SetLeaseID(leaseID)
	return o
}

// SetLeaseID adds the leaseId to the put devices erouters Id dhcp leases servername leaseid params
func (o *PutDevicesEroutersIDDhcpLeasesServernameLeaseidParams) SetLeaseID(leaseID string) {
	o.LeaseID = leaseID
}

// WithServerName adds the serverName to the put devices erouters Id dhcp leases servername leaseid params
func (o *PutDevicesEroutersIDDhcpLeasesServernameLeaseidParams) WithServerName(serverName string) *PutDevicesEroutersIDDhcpLeasesServernameLeaseidParams {
	o.SetServerName(serverName)
	return o
}

// SetServerName adds the serverName to the put devices erouters Id dhcp leases servername leaseid params
func (o *PutDevicesEroutersIDDhcpLeasesServernameLeaseidParams) SetServerName(serverName string) {
	o.ServerName = serverName
}

// WriteToRequest writes these params to a swagger request
func (o *PutDevicesEroutersIDDhcpLeasesServernameLeaseidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param leaseId
	if err := r.SetPathParam("leaseId", o.LeaseID); err != nil {
		return err
	}

	// path param serverName
	if err := r.SetPathParam("serverName", o.ServerName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
