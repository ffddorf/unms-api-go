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

// RedistributeStatic redistribute static
// swagger:model redistributeStatic
type RedistributeStatic struct {

	// enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// metric
	// Required: true
	// Maximum: 16
	// Minimum: 1
	Metric *float64 `json:"metric"`
}

// Validate validates this redistribute static
func (m *RedistributeStatic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetric(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RedistributeStatic) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *RedistributeStatic) validateMetric(formats strfmt.Registry) error {

	if err := validate.Required("metric", "body", m.Metric); err != nil {
		return err
	}

	if err := validate.Minimum("metric", "body", float64(*m.Metric), 1, false); err != nil {
		return err
	}

	if err := validate.Maximum("metric", "body", float64(*m.Metric), 16, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RedistributeStatic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RedistributeStatic) UnmarshalBinary(b []byte) error {
	var res RedistributeStatic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
