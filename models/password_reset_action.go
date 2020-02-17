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

// PasswordResetAction password reset action
// swagger:model PasswordResetAction
type PasswordResetAction struct {

	// New password.
	// Required: true
	// Max Length: 64
	// Min Length: 4
	Password *string `json:"password"`

	// Password reset token.
	// Required: true
	Token *string `json:"token"`
}

// Validate validates this password reset action
func (m *PasswordResetAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PasswordResetAction) validatePassword(formats strfmt.Registry) error {

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

func (m *PasswordResetAction) validateToken(formats strfmt.Registry) error {

	if err := validate.Required("token", "body", m.Token); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PasswordResetAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PasswordResetAction) UnmarshalBinary(b []byte) error {
	var res PasswordResetAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}