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

// BlackBox black box
// swagger:model BlackBox
type BlackBox struct {

	// aircube
	Aircube *DeviceAirCube `json:"aircube,omitempty"`

	// airfiber
	Airfiber *Airfiber2 `json:"airfiber,omitempty"`

	// airmax
	Airmax *DeviceAirmax `json:"airmax,omitempty"`

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

	// interfaces
	Interfaces DeviceInterfaceListSchema `json:"interfaces,omitempty"`

	// Custom IP address in IPv4 or IPv6 format.
	// Required: true
	IPAddress *string `json:"ipAddress"`

	// latest backup
	LatestBackup *LatestBackup `json:"latestBackup,omitempty"`

	// meta
	Meta *DeviceMeta `json:"meta,omitempty"`

	// mode
	Mode string `json:"mode,omitempty"`

	// olt
	Olt *Olt `json:"olt,omitempty"`

	// onu
	Onu *DeviceOnu1 `json:"onu,omitempty"`

	// overview
	Overview *DeviceOverview `json:"overview,omitempty"`

	// upgrade
	Upgrade *DeviceUpgrade `json:"upgrade,omitempty"`
}

// Validate validates this black box
func (m *BlackBox) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAircube(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAirfiber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAirmax(formats); err != nil {
		res = append(res, err)
	}

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

	if err := m.validateInterfaces(formats); err != nil {
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

	if err := m.validateOlt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnu(formats); err != nil {
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

func (m *BlackBox) validateAircube(formats strfmt.Registry) error {

	if swag.IsZero(m.Aircube) { // not required
		return nil
	}

	if m.Aircube != nil {
		if err := m.Aircube.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aircube")
			}
			return err
		}
	}

	return nil
}

func (m *BlackBox) validateAirfiber(formats strfmt.Registry) error {

	if swag.IsZero(m.Airfiber) { // not required
		return nil
	}

	if m.Airfiber != nil {
		if err := m.Airfiber.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("airfiber")
			}
			return err
		}
	}

	return nil
}

func (m *BlackBox) validateAirmax(formats strfmt.Registry) error {

	if swag.IsZero(m.Airmax) { // not required
		return nil
	}

	if m.Airmax != nil {
		if err := m.Airmax.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("airmax")
			}
			return err
		}
	}

	return nil
}

func (m *BlackBox) validateAttributes(formats strfmt.Registry) error {

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

func (m *BlackBox) validateDiscovery(formats strfmt.Registry) error {

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

func (m *BlackBox) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *BlackBox) validateFirmware(formats strfmt.Registry) error {

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

func (m *BlackBox) validateIdentification(formats strfmt.Registry) error {

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

func (m *BlackBox) validateInterfaces(formats strfmt.Registry) error {

	if swag.IsZero(m.Interfaces) { // not required
		return nil
	}

	if err := m.Interfaces.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("interfaces")
		}
		return err
	}

	return nil
}

func (m *BlackBox) validateIPAddress(formats strfmt.Registry) error {

	if err := validate.Required("ipAddress", "body", m.IPAddress); err != nil {
		return err
	}

	return nil
}

func (m *BlackBox) validateLatestBackup(formats strfmt.Registry) error {

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

func (m *BlackBox) validateMeta(formats strfmt.Registry) error {

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

func (m *BlackBox) validateOlt(formats strfmt.Registry) error {

	if swag.IsZero(m.Olt) { // not required
		return nil
	}

	if m.Olt != nil {
		if err := m.Olt.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("olt")
			}
			return err
		}
	}

	return nil
}

func (m *BlackBox) validateOnu(formats strfmt.Registry) error {

	if swag.IsZero(m.Onu) { // not required
		return nil
	}

	if m.Onu != nil {
		if err := m.Onu.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("onu")
			}
			return err
		}
	}

	return nil
}

func (m *BlackBox) validateOverview(formats strfmt.Registry) error {

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

func (m *BlackBox) validateUpgrade(formats strfmt.Registry) error {

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
func (m *BlackBox) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BlackBox) UnmarshalBinary(b []byte) error {
	var res BlackBox
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
