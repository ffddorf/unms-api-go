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

// DeviceSchema device schema
// swagger:model DeviceSchema
type DeviceSchema struct {

	// attributes
	Attributes *DeviceAttributes `json:"attributes,omitempty"`

	// discovery
	Discovery *Discovery `json:"discovery,omitempty"`

	// enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// firmware
	Firmware *DeviceFirmware `json:"firmware,omitempty"`

	// identification
	Identification *DeviceIdentification `json:"identification,omitempty"`

	// Custom IP address in IPv4 or IPv6 format.
	// Required: true
	IPAddress *string `json:"ipAddress"`

	// latest backup
	LatestBackup *LatestBackup `json:"latestBackup,omitempty"`

	// meta
	Meta *DeviceMeta `json:"meta,omitempty"`

	// mode
	Mode string `json:"mode,omitempty"`

	// overview
	Overview *DeviceOverview `json:"overview,omitempty"`

	// upgrade
	Upgrade *DeviceUpgrade `json:"upgrade,omitempty"`
}

// Validate validates this device schema
func (m *DeviceSchema) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiscovery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirmware(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatestBackup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMeta(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOverview(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpgrade(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceSchema) validateAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.Attributes) { // not required
		return nil
	}

	if m.Attributes != nil {
		if err := m.Attributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceSchema) validateDiscovery(formats strfmt.Registry) error {

	if swag.IsZero(m.Discovery) { // not required
		return nil
	}

	if m.Discovery != nil {
		if err := m.Discovery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("discovery")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceSchema) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *DeviceSchema) validateFirmware(formats strfmt.Registry) error {

	if swag.IsZero(m.Firmware) { // not required
		return nil
	}

	if m.Firmware != nil {
		if err := m.Firmware.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("firmware")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceSchema) validateIdentification(formats strfmt.Registry) error {

	if swag.IsZero(m.Identification) { // not required
		return nil
	}

	if m.Identification != nil {
		if err := m.Identification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identification")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceSchema) validateIPAddress(formats strfmt.Registry) error {

	if err := validate.Required("ipAddress", "body", m.IPAddress); err != nil {
		return err
	}

	return nil
}

func (m *DeviceSchema) validateLatestBackup(formats strfmt.Registry) error {

	if swag.IsZero(m.LatestBackup) { // not required
		return nil
	}

	if m.LatestBackup != nil {
		if err := m.LatestBackup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latestBackup")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceSchema) validateMeta(formats strfmt.Registry) error {

	if swag.IsZero(m.Meta) { // not required
		return nil
	}

	if m.Meta != nil {
		if err := m.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceSchema) validateOverview(formats strfmt.Registry) error {

	if swag.IsZero(m.Overview) { // not required
		return nil
	}

	if m.Overview != nil {
		if err := m.Overview.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("overview")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceSchema) validateUpgrade(formats strfmt.Registry) error {

	if swag.IsZero(m.Upgrade) { // not required
		return nil
	}

	if m.Upgrade != nil {
		if err := m.Upgrade.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("upgrade")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceSchema) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceSchema) UnmarshalBinary(b []byte) error {
	var res DeviceSchema
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
