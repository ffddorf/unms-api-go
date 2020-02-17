// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Router1 router 1
// swagger:model router 1
type Router1 struct {

	// dhcp lease time
	DhcpLeaseTime float64 `json:"dhcpLeaseTime,omitempty"`

	// dhcp pool end
	// Required: true
	DhcpPoolEnd *string `json:"dhcpPoolEnd"`

	// dhcp pool start
	// Required: true
	DhcpPoolStart *string `json:"dhcpPoolStart"`

	// dhcp relay
	// Required: true
	DhcpRelay *string `json:"dhcpRelay"`

	// dhcp server
	// Required: true
	DhcpServer *string `json:"dhcpServer"`

	// dns proxy enable
	DNSProxyEnable bool `json:"dnsProxyEnable,omitempty"`

	// gateway
	// Required: true
	Gateway *string `json:"gateway"`

	// nat protocol ftp
	NatProtocolFtp bool `json:"natProtocolFtp,omitempty"`

	// nat protocol pptp
	NatProtocolPptp bool `json:"natProtocolPptp,omitempty"`

	// nat protocol rtsp
	NatProtocolRtsp bool `json:"natProtocolRtsp,omitempty"`

	// nat protocol sip
	NatProtocolSip bool `json:"natProtocolSip,omitempty"`

	// primary Dns
	// Required: true
	PrimaryDNS *string `json:"primaryDns"`

	// secondary Dns
	// Required: true
	SecondaryDNS *string `json:"secondaryDns"`

	// upnp enabled
	// Required: true
	UpnpEnabled *bool `json:"upnpEnabled"`

	// wan access blocked
	WanAccessBlocked bool `json:"wanAccessBlocked,omitempty"`

	// wan mode
	// Required: true
	WanMode *string `json:"wanMode"`

	// wan vlan
	// Required: true
	// Maximum: 4063
	// Minimum: 2
	WanVlan *float64 `json:"wanVlan"`
}

// Validate validates this router 1
func (m *Router1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDhcpPoolEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDhcpPoolStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDhcpRelay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDhcpServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGateway(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryDNS(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondaryDNS(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpnpEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWanMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWanVlan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Router1) validateDhcpPoolEnd(formats strfmt.Registry) error {

	if err := validate.Required("dhcpPoolEnd", "body", m.DhcpPoolEnd); err != nil {
		return err
	}

	return nil
}

func (m *Router1) validateDhcpPoolStart(formats strfmt.Registry) error {

	if err := validate.Required("dhcpPoolStart", "body", m.DhcpPoolStart); err != nil {
		return err
	}

	return nil
}

func (m *Router1) validateDhcpRelay(formats strfmt.Registry) error {

	if err := validate.Required("dhcpRelay", "body", m.DhcpRelay); err != nil {
		return err
	}

	return nil
}

func (m *Router1) validateDhcpServer(formats strfmt.Registry) error {

	if err := validate.Required("dhcpServer", "body", m.DhcpServer); err != nil {
		return err
	}

	return nil
}

func (m *Router1) validateGateway(formats strfmt.Registry) error {

	if err := validate.Required("gateway", "body", m.Gateway); err != nil {
		return err
	}

	return nil
}

func (m *Router1) validatePrimaryDNS(formats strfmt.Registry) error {

	if err := validate.Required("primaryDns", "body", m.PrimaryDNS); err != nil {
		return err
	}

	return nil
}

func (m *Router1) validateSecondaryDNS(formats strfmt.Registry) error {

	if err := validate.Required("secondaryDns", "body", m.SecondaryDNS); err != nil {
		return err
	}

	return nil
}

func (m *Router1) validateUpnpEnabled(formats strfmt.Registry) error {

	if err := validate.Required("upnpEnabled", "body", m.UpnpEnabled); err != nil {
		return err
	}

	return nil
}

func (m *Router1) validateWanMode(formats strfmt.Registry) error {

	if err := validate.Required("wanMode", "body", m.WanMode); err != nil {
		return err
	}

	return nil
}

func (m *Router1) validateWanVlan(formats strfmt.Registry) error {

	if err := validate.Required("wanVlan", "body", m.WanVlan); err != nil {
		return err
	}

	if err := validate.Minimum("wanVlan", "body", float64(*m.WanVlan), 2, false); err != nil {
		return err
	}

	if err := validate.Maximum("wanVlan", "body", float64(*m.WanVlan), 4063, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Router1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Router1) UnmarshalBinary(b []byte) error {
	var res Router1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
