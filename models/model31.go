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

// Model31 model 31
// swagger:model Model 31
type Model31 struct {

	// include vlans
	IncludeVlans IncludeVlans `json:"includeVlans,omitempty"`

	// Native VLAN identification
	// Required: true
	// Maximum: 4063
	// Minimum: 1
	NativeVlan *int64 `json:"nativeVlan"`

	// Identification of port
	// Required: true
	Port *string `json:"port"`
}

// Validate validates this model 31
func (m *Model31) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIncludeVlans(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNativeVlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Model31) validateIncludeVlans(formats strfmt.Registry) error {

	if swag.IsZero(m.IncludeVlans) { // not required
		return nil
	}

	if err := m.IncludeVlans.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("includeVlans")
		}
		return err
	}

	return nil
}

func (m *Model31) validateNativeVlan(formats strfmt.Registry) error {

	if err := validate.Required("nativeVlan", "body", m.NativeVlan); err != nil {
		return err
	}

	if err := validate.MinimumInt("nativeVlan", "body", int64(*m.NativeVlan), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("nativeVlan", "body", int64(*m.NativeVlan), 4063, false); err != nil {
		return err
	}

	return nil
}

func (m *Model31) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Model31) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Model31) UnmarshalBinary(b []byte) error {
	var res Model31
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}