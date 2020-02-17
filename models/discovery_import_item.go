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

// DiscoveryImportItem discovery import item
// swagger:model DiscoveryImportItem
type DiscoveryImportItem struct {

	// https port
	// Maximum: 65535
	// Minimum: 0
	HTTPSPort *int64 `json:"httpsPort,omitempty"`

	// ip
	// Required: true
	IP *string `json:"ip"`

	// password
	// Required: true
	Password *string `json:"password"`

	// username
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this discovery import item
func (m *DiscoveryImportItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHTTPSPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DiscoveryImportItem) validateHTTPSPort(formats strfmt.Registry) error {

	if swag.IsZero(m.HTTPSPort) { // not required
		return nil
	}

	if err := validate.MinimumInt("httpsPort", "body", int64(*m.HTTPSPort), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("httpsPort", "body", int64(*m.HTTPSPort), 65535, false); err != nil {
		return err
	}

	return nil
}

func (m *DiscoveryImportItem) validateIP(formats strfmt.Registry) error {

	if err := validate.Required("ip", "body", m.IP); err != nil {
		return err
	}

	return nil
}

func (m *DiscoveryImportItem) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

func (m *DiscoveryImportItem) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DiscoveryImportItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DiscoveryImportItem) UnmarshalBinary(b []byte) error {
	var res DiscoveryImportItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
