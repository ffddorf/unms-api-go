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

// VaultCredentials1 vault credentials 1
// swagger:model VaultCredentials 1
type VaultCredentials1 struct {

	// is vault enabled
	// Required: true
	IsVaultEnabled *bool `json:"isVaultEnabled"`

	// passphrase
	Passphrase string `json:"passphrase,omitempty"`

	// regenerate pgp keys
	RegeneratePgpKeys bool `json:"regeneratePgpKeys,omitempty"`
}

// Validate validates this vault credentials 1
func (m *VaultCredentials1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIsVaultEnabled(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VaultCredentials1) validateIsVaultEnabled(formats strfmt.Registry) error {

	if err := validate.Required("isVaultEnabled", "body", m.IsVaultEnabled); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VaultCredentials1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VaultCredentials1) UnmarshalBinary(b []byte) error {
	var res VaultCredentials1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
