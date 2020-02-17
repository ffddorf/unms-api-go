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

// Station station
// swagger:model Station
type Station struct {

	// connected
	Connected bool `json:"connected,omitempty"`

	// device identification
	DeviceIdentification *DeviceIdentification1 `json:"deviceIdentification,omitempty"`

	// distance
	Distance int64 `json:"distance,omitempty"`

	// downlink capacity
	DownlinkCapacity int64 `json:"downlinkCapacity,omitempty"`

	// Custom IP address in IPv4 or IPv6 format.
	IPAddress string `json:"ipAddress,omitempty"`

	// latency
	Latency int64 `json:"latency,omitempty"`

	// mac
	// Pattern: ^([0-9a-fA-F][0-9a-fA-F]:){5}([0-9a-fA-F][0-9a-fA-F])$|^([0-9a-fA-F]){12}$
	Mac string `json:"mac,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// radio
	Radio string `json:"radio,omitempty"`

	// rx bytes
	RxBytes int64 `json:"rxBytes,omitempty"`

	// rx modulation
	RxModulation string `json:"rxModulation,omitempty"`

	// rx rate
	RxRate int64 `json:"rxRate,omitempty"`

	// rx signal
	RxSignal int64 `json:"rxSignal,omitempty"`

	// timestamp
	// Format: date
	Timestamp strfmt.Date `json:"timestamp,omitempty"`

	// tx bytes
	TxBytes int64 `json:"txBytes,omitempty"`

	// tx modulation
	TxModulation string `json:"txModulation,omitempty"`

	// tx rate
	TxRate int64 `json:"txRate,omitempty"`

	// tx signal
	TxSignal int64 `json:"txSignal,omitempty"`

	// uplink capacity
	UplinkCapacity int64 `json:"uplinkCapacity,omitempty"`

	// uptime
	Uptime int64 `json:"uptime,omitempty"`

	// vendor
	Vendor string `json:"vendor,omitempty"`
}

// Validate validates this station
func (m *Station) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeviceIdentification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMac(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Station) validateDeviceIdentification(formats strfmt.Registry) error {

	if swag.IsZero(m.DeviceIdentification) { // not required
		return nil
	}

	if m.DeviceIdentification != nil {
		if err := m.DeviceIdentification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deviceIdentification")
			}
			return err
		}
	}

	return nil
}

func (m *Station) validateMac(formats strfmt.Registry) error {

	if swag.IsZero(m.Mac) { // not required
		return nil
	}

	if err := validate.Pattern("mac", "body", string(m.Mac), `^([0-9a-fA-F][0-9a-fA-F]:){5}([0-9a-fA-F][0-9a-fA-F])$|^([0-9a-fA-F]){12}$`); err != nil {
		return err
	}

	return nil
}

func (m *Station) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Station) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Station) UnmarshalBinary(b []byte) error {
	var res Station
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
