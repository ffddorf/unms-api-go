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

// Vlp3 vlp3
// swagger:model vlp3
type Vlp3 struct {

	// enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// ports
	Ports Ports `json:"ports,omitempty"`

	// tagged
	// Required: true
	Tagged *bool `json:"tagged"`

	// vlan Id
	// Required: true
	// Maximum: 4096
	// Minimum: 1
	VlanID *int64 `json:"vlanId"`
}

// Validate validates this vlp3
func (m *Vlp3) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTagged(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlanID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Vlp3) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *Vlp3) validatePorts(formats strfmt.Registry) error {

	if swag.IsZero(m.Ports) { // not required
		return nil
	}

	if err := m.Ports.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ports")
		}
		return err
	}

	return nil
}

func (m *Vlp3) validateTagged(formats strfmt.Registry) error {

	if err := validate.Required("tagged", "body", m.Tagged); err != nil {
		return err
	}

	return nil
}

func (m *Vlp3) validateVlanID(formats strfmt.Registry) error {

	if err := validate.Required("vlanId", "body", m.VlanID); err != nil {
		return err
	}

	if err := validate.MinimumInt("vlanId", "body", int64(*m.VlanID), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("vlanId", "body", int64(*m.VlanID), 4096, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Vlp3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Vlp3) UnmarshalBinary(b []byte) error {
	var res Vlp3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}