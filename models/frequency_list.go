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

// FrequencyList frequency list
// swagger:model FrequencyList
type FrequencyList struct {

	// radio2 ghz frequency list
	// Required: true
	Radio2GhzFrequencyList Radio2GhzFrequencyList `json:"radio2GhzFrequencyList"`

	// radio5 ghz frequency list
	// Required: true
	Radio5GhzFrequencyList Radio5GhzFrequencyList `json:"radio5GhzFrequencyList"`
}

// Validate validates this frequency list
func (m *FrequencyList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRadio2GhzFrequencyList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRadio5GhzFrequencyList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FrequencyList) validateRadio2GhzFrequencyList(formats strfmt.Registry) error {

	if err := validate.Required("radio2GhzFrequencyList", "body", m.Radio2GhzFrequencyList); err != nil {
		return err
	}

	if err := m.Radio2GhzFrequencyList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("radio2GhzFrequencyList")
		}
		return err
	}

	return nil
}

func (m *FrequencyList) validateRadio5GhzFrequencyList(formats strfmt.Registry) error {

	if err := validate.Required("radio5GhzFrequencyList", "body", m.Radio5GhzFrequencyList); err != nil {
		return err
	}

	if err := m.Radio5GhzFrequencyList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("radio5GhzFrequencyList")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FrequencyList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FrequencyList) UnmarshalBinary(b []byte) error {
	var res FrequencyList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}