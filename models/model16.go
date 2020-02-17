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

// Model16 model 16
// swagger:model Model 16
type Model16 struct {

	// connected
	// Required: true
	Connected *bool `json:"connected"`

	// device Id
	// Required: true
	DeviceID *string `json:"deviceId"`

	// duplicate of
	// Required: true
	DuplicateOf *string `json:"duplicateOf"`

	// error
	// Required: true
	Error *string `json:"error"`

	// name
	// Required: true
	Name *string `json:"name"`

	// protocol
	// Required: true
	Protocol *string `json:"protocol"`

	// snmp community
	// Required: true
	SnmpCommunity *string `json:"snmpCommunity"`

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this model 16
func (m *Model16) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnected(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDuplicateOf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnmpCommunity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Model16) validateConnected(formats strfmt.Registry) error {

	if err := validate.Required("connected", "body", m.Connected); err != nil {
		return err
	}

	return nil
}

func (m *Model16) validateDeviceID(formats strfmt.Registry) error {

	if err := validate.Required("deviceId", "body", m.DeviceID); err != nil {
		return err
	}

	return nil
}

func (m *Model16) validateDuplicateOf(formats strfmt.Registry) error {

	if err := validate.Required("duplicateOf", "body", m.DuplicateOf); err != nil {
		return err
	}

	return nil
}

func (m *Model16) validateError(formats strfmt.Registry) error {

	if err := validate.Required("error", "body", m.Error); err != nil {
		return err
	}

	return nil
}

func (m *Model16) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Model16) validateProtocol(formats strfmt.Registry) error {

	if err := validate.Required("protocol", "body", m.Protocol); err != nil {
		return err
	}

	return nil
}

func (m *Model16) validateSnmpCommunity(formats strfmt.Registry) error {

	if err := validate.Required("snmpCommunity", "body", m.SnmpCommunity); err != nil {
		return err
	}

	return nil
}

func (m *Model16) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Model16) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Model16) UnmarshalBinary(b []byte) error {
	var res Model16
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}