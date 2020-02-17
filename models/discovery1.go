// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Discovery1 discovery 1
// swagger:model discovery 1
type Discovery1 struct {

	// responder
	Responder *Responder `json:"responder,omitempty"`

	// scanner
	Scanner *Scanner `json:"scanner,omitempty"`
}

// Validate validates this discovery 1
func (m *Discovery1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResponder(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScanner(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Discovery1) validateResponder(formats strfmt.Registry) error {

	if swag.IsZero(m.Responder) { // not required
		return nil
	}

	if m.Responder != nil {
		if err := m.Responder.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("responder")
			}
			return err
		}
	}

	return nil
}

func (m *Discovery1) validateScanner(formats strfmt.Registry) error {

	if swag.IsZero(m.Scanner) { // not required
		return nil
	}

	if m.Scanner != nil {
		if err := m.Scanner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scanner")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Discovery1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Discovery1) UnmarshalBinary(b []byte) error {
	var res Discovery1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}