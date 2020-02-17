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

// PasswordStrength password strength
// swagger:model PasswordStrength
type PasswordStrength struct {

	// password
	// Required: true
	// Max Length: 64
	// Min Length: 4
	Password *string `json:"password"`
}

// Validate validates this password strength
func (m *PasswordStrength) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PasswordStrength) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	if err := validate.MinLength("password", "body", string(*m.Password), 4); err != nil {
		return err
	}

	if err := validate.MaxLength("password", "body", string(*m.Password), 64); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PasswordStrength) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PasswordStrength) UnmarshalBinary(b []byte) error {
	var res PasswordStrength
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
