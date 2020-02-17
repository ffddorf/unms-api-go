// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Wifi2Ghz wifi2 ghz
// swagger:model wifi2Ghz
type Wifi2Ghz struct {

	// authentication
	// Required: true
	// Enum: [psk psk2 ent none]
	Authentication *string `json:"authentication"`

	// available
	// Required: true
	Available *bool `json:"available"`

	// channel
	// Required: true
	Channel *float64 `json:"channel"`

	// channel width
	// Required: true
	ChannelWidth *float64 `json:"channelWidth"`

	// country
	Country string `json:"country,omitempty"`

	// enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// encryption
	// Required: true
	// Enum: [wep wpa wpa2 none]
	Encryption *string `json:"encryption"`

	// frequency
	Frequency float64 `json:"frequency,omitempty"`

	// is channel auto
	IsChannelAuto bool `json:"isChannelAuto,omitempty"`

	// is w p a2 p s k enabled
	IsWPA2PSKEnabled bool `json:"isWPA2PSKEnabled,omitempty"`

	// key
	// Required: true
	Key *string `json:"key"`

	// mac
	// Pattern: ^([0-9a-fA-F][0-9a-fA-F]:){5}([0-9a-fA-F][0-9a-fA-F])$|^([0-9a-fA-F]){12}$
	Mac string `json:"mac,omitempty"`

	// mode
	// Required: true
	// Enum: [ap ap-ptp ap-ptmp ap-ptmp-airmax ap-ptmp-airmax-mixed ap-ptmp-airmax-ac sta sta-ptp sta-ptmp aprepeater repeater mesh]
	Mode *string `json:"mode"`

	// ssid
	// Required: true
	Ssid *string `json:"ssid"`

	// tx power
	// Required: true
	TxPower *float64 `json:"txPower"`
}

// Validate validates this wifi2 ghz
func (m *Wifi2Ghz) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthentication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAvailable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChannel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChannelWidth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEncryption(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMac(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSsid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTxPower(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var wifi2GhzTypeAuthenticationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["psk","psk2","ent","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		wifi2GhzTypeAuthenticationPropEnum = append(wifi2GhzTypeAuthenticationPropEnum, v)
	}
}

const (

	// Wifi2GhzAuthenticationPsk captures enum value "psk"
	Wifi2GhzAuthenticationPsk string = "psk"

	// Wifi2GhzAuthenticationPsk2 captures enum value "psk2"
	Wifi2GhzAuthenticationPsk2 string = "psk2"

	// Wifi2GhzAuthenticationEnt captures enum value "ent"
	Wifi2GhzAuthenticationEnt string = "ent"

	// Wifi2GhzAuthenticationNone captures enum value "none"
	Wifi2GhzAuthenticationNone string = "none"
)

// prop value enum
func (m *Wifi2Ghz) validateAuthenticationEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, wifi2GhzTypeAuthenticationPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Wifi2Ghz) validateAuthentication(formats strfmt.Registry) error {

	if err := validate.Required("authentication", "body", m.Authentication); err != nil {
		return err
	}

	// value enum
	if err := m.validateAuthenticationEnum("authentication", "body", *m.Authentication); err != nil {
		return err
	}

	return nil
}

func (m *Wifi2Ghz) validateAvailable(formats strfmt.Registry) error {

	if err := validate.Required("available", "body", m.Available); err != nil {
		return err
	}

	return nil
}

func (m *Wifi2Ghz) validateChannel(formats strfmt.Registry) error {

	if err := validate.Required("channel", "body", m.Channel); err != nil {
		return err
	}

	return nil
}

func (m *Wifi2Ghz) validateChannelWidth(formats strfmt.Registry) error {

	if err := validate.Required("channelWidth", "body", m.ChannelWidth); err != nil {
		return err
	}

	return nil
}

func (m *Wifi2Ghz) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

var wifi2GhzTypeEncryptionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["wep","wpa","wpa2","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		wifi2GhzTypeEncryptionPropEnum = append(wifi2GhzTypeEncryptionPropEnum, v)
	}
}

const (

	// Wifi2GhzEncryptionWep captures enum value "wep"
	Wifi2GhzEncryptionWep string = "wep"

	// Wifi2GhzEncryptionWpa captures enum value "wpa"
	Wifi2GhzEncryptionWpa string = "wpa"

	// Wifi2GhzEncryptionWpa2 captures enum value "wpa2"
	Wifi2GhzEncryptionWpa2 string = "wpa2"

	// Wifi2GhzEncryptionNone captures enum value "none"
	Wifi2GhzEncryptionNone string = "none"
)

// prop value enum
func (m *Wifi2Ghz) validateEncryptionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, wifi2GhzTypeEncryptionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Wifi2Ghz) validateEncryption(formats strfmt.Registry) error {

	if err := validate.Required("encryption", "body", m.Encryption); err != nil {
		return err
	}

	// value enum
	if err := m.validateEncryptionEnum("encryption", "body", *m.Encryption); err != nil {
		return err
	}

	return nil
}

func (m *Wifi2Ghz) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

func (m *Wifi2Ghz) validateMac(formats strfmt.Registry) error {

	if swag.IsZero(m.Mac) { // not required
		return nil
	}

	if err := validate.Pattern("mac", "body", string(m.Mac), `^([0-9a-fA-F][0-9a-fA-F]:){5}([0-9a-fA-F][0-9a-fA-F])$|^([0-9a-fA-F]){12}$`); err != nil {
		return err
	}

	return nil
}

var wifi2GhzTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ap","ap-ptp","ap-ptmp","ap-ptmp-airmax","ap-ptmp-airmax-mixed","ap-ptmp-airmax-ac","sta","sta-ptp","sta-ptmp","aprepeater","repeater","mesh"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		wifi2GhzTypeModePropEnum = append(wifi2GhzTypeModePropEnum, v)
	}
}

const (

	// Wifi2GhzModeAp captures enum value "ap"
	Wifi2GhzModeAp string = "ap"

	// Wifi2GhzModeApPtp captures enum value "ap-ptp"
	Wifi2GhzModeApPtp string = "ap-ptp"

	// Wifi2GhzModeApPtmp captures enum value "ap-ptmp"
	Wifi2GhzModeApPtmp string = "ap-ptmp"

	// Wifi2GhzModeApPtmpAirmax captures enum value "ap-ptmp-airmax"
	Wifi2GhzModeApPtmpAirmax string = "ap-ptmp-airmax"

	// Wifi2GhzModeApPtmpAirmaxMixed captures enum value "ap-ptmp-airmax-mixed"
	Wifi2GhzModeApPtmpAirmaxMixed string = "ap-ptmp-airmax-mixed"

	// Wifi2GhzModeApPtmpAirmaxAc captures enum value "ap-ptmp-airmax-ac"
	Wifi2GhzModeApPtmpAirmaxAc string = "ap-ptmp-airmax-ac"

	// Wifi2GhzModeSta captures enum value "sta"
	Wifi2GhzModeSta string = "sta"

	// Wifi2GhzModeStaPtp captures enum value "sta-ptp"
	Wifi2GhzModeStaPtp string = "sta-ptp"

	// Wifi2GhzModeStaPtmp captures enum value "sta-ptmp"
	Wifi2GhzModeStaPtmp string = "sta-ptmp"

	// Wifi2GhzModeAprepeater captures enum value "aprepeater"
	Wifi2GhzModeAprepeater string = "aprepeater"

	// Wifi2GhzModeRepeater captures enum value "repeater"
	Wifi2GhzModeRepeater string = "repeater"

	// Wifi2GhzModeMesh captures enum value "mesh"
	Wifi2GhzModeMesh string = "mesh"
)

// prop value enum
func (m *Wifi2Ghz) validateModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, wifi2GhzTypeModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Wifi2Ghz) validateMode(formats strfmt.Registry) error {

	if err := validate.Required("mode", "body", m.Mode); err != nil {
		return err
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", *m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *Wifi2Ghz) validateSsid(formats strfmt.Registry) error {

	if err := validate.Required("ssid", "body", m.Ssid); err != nil {
		return err
	}

	return nil
}

func (m *Wifi2Ghz) validateTxPower(formats strfmt.Registry) error {

	if err := validate.Required("txPower", "body", m.TxPower); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Wifi2Ghz) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Wifi2Ghz) UnmarshalBinary(b []byte) error {
	var res Wifi2Ghz
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
