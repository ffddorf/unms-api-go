// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EditUserPreferences edit user preferences
// swagger:model EditUserPreferences
type EditUserPreferences struct {

	// alerts
	Alerts bool `json:"alerts,omitempty"`

	// map config
	MapConfig *MapConfig `json:"mapConfig,omitempty"`

	// preferences
	Preferences Preferences `json:"preferences,omitempty"`

	// table config
	TableConfig TableConfig `json:"tableConfig,omitempty"`
}

// Validate validates this edit user preferences
func (m *EditUserPreferences) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMapConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EditUserPreferences) validateMapConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.MapConfig) { // not required
		return nil
	}

	if m.MapConfig != nil {
		if err := m.MapConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mapConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EditUserPreferences) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EditUserPreferences) UnmarshalBinary(b []byte) error {
	var res EditUserPreferences
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
