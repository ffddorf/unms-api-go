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

// Error error
// swagger:model Error
type Error struct {

	// error
	// Required: true
	Error *string `json:"error"`

	// message
	Message string `json:"message,omitempty"`

	// status code
	// Required: true
	// Maximum: 599
	// Minimum: 400
	StatusCode *float64 `json:"statusCode"`

	// validation
	Validation Validation `json:"validation,omitempty"`
}

// Validate validates this error
func (m *Error) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Error) validateError(formats strfmt.Registry) error {

	if err := validate.Required("error", "body", m.Error); err != nil {
		return err
	}

	return nil
}

func (m *Error) validateStatusCode(formats strfmt.Registry) error {

	if err := validate.Required("statusCode", "body", m.StatusCode); err != nil {
		return err
	}

	if err := validate.Minimum("statusCode", "body", float64(*m.StatusCode), 400, false); err != nil {
		return err
	}

	if err := validate.Maximum("statusCode", "body", float64(*m.StatusCode), 599, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Error) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Error) UnmarshalBinary(b []byte) error {
	var res Error
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}