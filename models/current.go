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

// Current current
// swagger:model current
type Current struct {

	// major
	// Required: true
	Major *float64 `json:"major"`

	// minor
	// Required: true
	Minor *float64 `json:"minor"`

	// order
	Order string `json:"order,omitempty"`

	// patch
	// Required: true
	Patch *float64 `json:"patch"`

	// prerelease
	Prerelease Prerelease `json:"prerelease,omitempty"`
}

// Validate validates this current
func (m *Current) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMajor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePatch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrerelease(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Current) validateMajor(formats strfmt.Registry) error {

	if err := validate.Required("major", "body", m.Major); err != nil {
		return err
	}

	return nil
}

func (m *Current) validateMinor(formats strfmt.Registry) error {

	if err := validate.Required("minor", "body", m.Minor); err != nil {
		return err
	}

	return nil
}

func (m *Current) validatePatch(formats strfmt.Registry) error {

	if err := validate.Required("patch", "body", m.Patch); err != nil {
		return err
	}

	return nil
}

func (m *Current) validatePrerelease(formats strfmt.Registry) error {

	if swag.IsZero(m.Prerelease) { // not required
		return nil
	}

	if err := m.Prerelease.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("prerelease")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Current) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Current) UnmarshalBinary(b []byte) error {
	var res Current
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}