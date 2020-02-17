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

// OnuSystem onu system
// swagger:model OnuSystem
type OnuSystem struct {

	// Admin password
	// Max Length: 64
	// Min Length: 4
	AdminPassword string `json:"adminPassword,omitempty"`

	// Set to true if device is enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// Device name
	// Required: true
	// Min Length: 1
	Name *string `json:"name"`
}

// Validate validates this onu system
func (m *OnuSystem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdminPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OnuSystem) validateAdminPassword(formats strfmt.Registry) error {

	if swag.IsZero(m.AdminPassword) { // not required
		return nil
	}

	if err := validate.MinLength("adminPassword", "body", string(m.AdminPassword), 4); err != nil {
		return err
	}

	if err := validate.MaxLength("adminPassword", "body", string(m.AdminPassword), 64); err != nil {
		return err
	}

	return nil
}

func (m *OnuSystem) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *OnuSystem) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OnuSystem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OnuSystem) UnmarshalBinary(b []byte) error {
	var res OnuSystem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
