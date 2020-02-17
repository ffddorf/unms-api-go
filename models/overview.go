// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Overview overview
// swagger:model overview
type Overview struct {

	// antenna
	Antenna *Antenna `json:"antenna,omitempty"`

	// channel width
	ChannelWidth float64 `json:"channelWidth,omitempty"`
}

// Validate validates this overview
func (m *Overview) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAntenna(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Overview) validateAntenna(formats strfmt.Registry) error {

	if swag.IsZero(m.Antenna) { // not required
		return nil
	}

	if m.Antenna != nil {
		if err := m.Antenna.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("antenna")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Overview) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Overview) UnmarshalBinary(b []byte) error {
	var res Overview
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
