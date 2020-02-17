// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IPV4 LAN and WAN IPv4 settings
// swagger:model ipv4
type IPV4 struct {

	// lan
	Lan *Lan1 `json:"lan,omitempty"`

	// wan
	Wan *Wan1 `json:"wan,omitempty"`
}

// Validate validates this ipv4
func (m *IPV4) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPV4) validateLan(formats strfmt.Registry) error {

	if swag.IsZero(m.Lan) { // not required
		return nil
	}

	if m.Lan != nil {
		if err := m.Lan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lan")
			}
			return err
		}
	}

	return nil
}

func (m *IPV4) validateWan(formats strfmt.Registry) error {

	if swag.IsZero(m.Wan) { // not required
		return nil
	}

	if m.Wan != nil {
		if err := m.Wan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("wan")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPV4) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPV4) UnmarshalBinary(b []byte) error {
	var res IPV4
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
