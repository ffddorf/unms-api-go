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

// BulkDeleteDevices bulk delete devices
// swagger:model BulkDeleteDevices
type BulkDeleteDevices struct {

	// deleted ids
	// Required: true
	DeletedIds DeletedIds `json:"deletedIds"`

	// discovered
	Discovered bool `json:"discovered,omitempty"`

	// message
	// Required: true
	Message *string `json:"message"`

	// undeleted ids
	// Required: true
	UndeletedIds UndeletedIds `json:"undeletedIds"`
}

// Validate validates this bulk delete devices
func (m *BulkDeleteDevices) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeletedIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUndeletedIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BulkDeleteDevices) validateDeletedIds(formats strfmt.Registry) error {

	if err := validate.Required("deletedIds", "body", m.DeletedIds); err != nil {
		return err
	}

	if err := m.DeletedIds.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("deletedIds")
		}
		return err
	}

	return nil
}

func (m *BulkDeleteDevices) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *BulkDeleteDevices) validateUndeletedIds(formats strfmt.Registry) error {

	if err := validate.Required("undeletedIds", "body", m.UndeletedIds); err != nil {
		return err
	}

	if err := m.UndeletedIds.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("undeletedIds")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BulkDeleteDevices) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BulkDeleteDevices) UnmarshalBinary(b []byte) error {
	var res BulkDeleteDevices
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
