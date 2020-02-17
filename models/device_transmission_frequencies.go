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

// DeviceTransmissionFrequencies device transmission frequencies
// swagger:model deviceTransmissionFrequencies
type DeviceTransmissionFrequencies struct {

	// ap
	// Required: true
	// Enum: [minimal low medium high realtime]
	Ap *string `json:"ap"`

	// gpon
	// Required: true
	// Enum: [minimal low medium high realtime]
	Gpon *string `json:"gpon"`

	// other
	// Required: true
	// Enum: [minimal low medium high realtime]
	Other *string `json:"other"`

	// router
	// Required: true
	// Enum: [minimal low medium high realtime]
	Router *string `json:"router"`

	// station
	// Required: true
	// Enum: [minimal low medium high realtime]
	Station *string `json:"station"`

	// switch
	// Required: true
	// Enum: [minimal low medium high realtime]
	Switch *string `json:"switch"`

	// ups
	// Required: true
	// Enum: [minimal low medium high realtime]
	Ups *string `json:"ups"`
}

// Validate validates this device transmission frequencies
func (m *DeviceTransmissionFrequencies) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGpon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOther(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSwitch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var deviceTransmissionFrequenciesTypeApPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["minimal","low","medium","high","realtime"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviceTransmissionFrequenciesTypeApPropEnum = append(deviceTransmissionFrequenciesTypeApPropEnum, v)
	}
}

const (

	// DeviceTransmissionFrequenciesApMinimal captures enum value "minimal"
	DeviceTransmissionFrequenciesApMinimal string = "minimal"

	// DeviceTransmissionFrequenciesApLow captures enum value "low"
	DeviceTransmissionFrequenciesApLow string = "low"

	// DeviceTransmissionFrequenciesApMedium captures enum value "medium"
	DeviceTransmissionFrequenciesApMedium string = "medium"

	// DeviceTransmissionFrequenciesApHigh captures enum value "high"
	DeviceTransmissionFrequenciesApHigh string = "high"

	// DeviceTransmissionFrequenciesApRealtime captures enum value "realtime"
	DeviceTransmissionFrequenciesApRealtime string = "realtime"
)

// prop value enum
func (m *DeviceTransmissionFrequencies) validateApEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, deviceTransmissionFrequenciesTypeApPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DeviceTransmissionFrequencies) validateAp(formats strfmt.Registry) error {

	if err := validate.Required("ap", "body", m.Ap); err != nil {
		return err
	}

	// value enum
	if err := m.validateApEnum("ap", "body", *m.Ap); err != nil {
		return err
	}

	return nil
}

var deviceTransmissionFrequenciesTypeGponPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["minimal","low","medium","high","realtime"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviceTransmissionFrequenciesTypeGponPropEnum = append(deviceTransmissionFrequenciesTypeGponPropEnum, v)
	}
}

const (

	// DeviceTransmissionFrequenciesGponMinimal captures enum value "minimal"
	DeviceTransmissionFrequenciesGponMinimal string = "minimal"

	// DeviceTransmissionFrequenciesGponLow captures enum value "low"
	DeviceTransmissionFrequenciesGponLow string = "low"

	// DeviceTransmissionFrequenciesGponMedium captures enum value "medium"
	DeviceTransmissionFrequenciesGponMedium string = "medium"

	// DeviceTransmissionFrequenciesGponHigh captures enum value "high"
	DeviceTransmissionFrequenciesGponHigh string = "high"

	// DeviceTransmissionFrequenciesGponRealtime captures enum value "realtime"
	DeviceTransmissionFrequenciesGponRealtime string = "realtime"
)

// prop value enum
func (m *DeviceTransmissionFrequencies) validateGponEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, deviceTransmissionFrequenciesTypeGponPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DeviceTransmissionFrequencies) validateGpon(formats strfmt.Registry) error {

	if err := validate.Required("gpon", "body", m.Gpon); err != nil {
		return err
	}

	// value enum
	if err := m.validateGponEnum("gpon", "body", *m.Gpon); err != nil {
		return err
	}

	return nil
}

var deviceTransmissionFrequenciesTypeOtherPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["minimal","low","medium","high","realtime"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviceTransmissionFrequenciesTypeOtherPropEnum = append(deviceTransmissionFrequenciesTypeOtherPropEnum, v)
	}
}

const (

	// DeviceTransmissionFrequenciesOtherMinimal captures enum value "minimal"
	DeviceTransmissionFrequenciesOtherMinimal string = "minimal"

	// DeviceTransmissionFrequenciesOtherLow captures enum value "low"
	DeviceTransmissionFrequenciesOtherLow string = "low"

	// DeviceTransmissionFrequenciesOtherMedium captures enum value "medium"
	DeviceTransmissionFrequenciesOtherMedium string = "medium"

	// DeviceTransmissionFrequenciesOtherHigh captures enum value "high"
	DeviceTransmissionFrequenciesOtherHigh string = "high"

	// DeviceTransmissionFrequenciesOtherRealtime captures enum value "realtime"
	DeviceTransmissionFrequenciesOtherRealtime string = "realtime"
)

// prop value enum
func (m *DeviceTransmissionFrequencies) validateOtherEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, deviceTransmissionFrequenciesTypeOtherPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DeviceTransmissionFrequencies) validateOther(formats strfmt.Registry) error {

	if err := validate.Required("other", "body", m.Other); err != nil {
		return err
	}

	// value enum
	if err := m.validateOtherEnum("other", "body", *m.Other); err != nil {
		return err
	}

	return nil
}

var deviceTransmissionFrequenciesTypeRouterPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["minimal","low","medium","high","realtime"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviceTransmissionFrequenciesTypeRouterPropEnum = append(deviceTransmissionFrequenciesTypeRouterPropEnum, v)
	}
}

const (

	// DeviceTransmissionFrequenciesRouterMinimal captures enum value "minimal"
	DeviceTransmissionFrequenciesRouterMinimal string = "minimal"

	// DeviceTransmissionFrequenciesRouterLow captures enum value "low"
	DeviceTransmissionFrequenciesRouterLow string = "low"

	// DeviceTransmissionFrequenciesRouterMedium captures enum value "medium"
	DeviceTransmissionFrequenciesRouterMedium string = "medium"

	// DeviceTransmissionFrequenciesRouterHigh captures enum value "high"
	DeviceTransmissionFrequenciesRouterHigh string = "high"

	// DeviceTransmissionFrequenciesRouterRealtime captures enum value "realtime"
	DeviceTransmissionFrequenciesRouterRealtime string = "realtime"
)

// prop value enum
func (m *DeviceTransmissionFrequencies) validateRouterEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, deviceTransmissionFrequenciesTypeRouterPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DeviceTransmissionFrequencies) validateRouter(formats strfmt.Registry) error {

	if err := validate.Required("router", "body", m.Router); err != nil {
		return err
	}

	// value enum
	if err := m.validateRouterEnum("router", "body", *m.Router); err != nil {
		return err
	}

	return nil
}

var deviceTransmissionFrequenciesTypeStationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["minimal","low","medium","high","realtime"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviceTransmissionFrequenciesTypeStationPropEnum = append(deviceTransmissionFrequenciesTypeStationPropEnum, v)
	}
}

const (

	// DeviceTransmissionFrequenciesStationMinimal captures enum value "minimal"
	DeviceTransmissionFrequenciesStationMinimal string = "minimal"

	// DeviceTransmissionFrequenciesStationLow captures enum value "low"
	DeviceTransmissionFrequenciesStationLow string = "low"

	// DeviceTransmissionFrequenciesStationMedium captures enum value "medium"
	DeviceTransmissionFrequenciesStationMedium string = "medium"

	// DeviceTransmissionFrequenciesStationHigh captures enum value "high"
	DeviceTransmissionFrequenciesStationHigh string = "high"

	// DeviceTransmissionFrequenciesStationRealtime captures enum value "realtime"
	DeviceTransmissionFrequenciesStationRealtime string = "realtime"
)

// prop value enum
func (m *DeviceTransmissionFrequencies) validateStationEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, deviceTransmissionFrequenciesTypeStationPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DeviceTransmissionFrequencies) validateStation(formats strfmt.Registry) error {

	if err := validate.Required("station", "body", m.Station); err != nil {
		return err
	}

	// value enum
	if err := m.validateStationEnum("station", "body", *m.Station); err != nil {
		return err
	}

	return nil
}

var deviceTransmissionFrequenciesTypeSwitchPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["minimal","low","medium","high","realtime"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviceTransmissionFrequenciesTypeSwitchPropEnum = append(deviceTransmissionFrequenciesTypeSwitchPropEnum, v)
	}
}

const (

	// DeviceTransmissionFrequenciesSwitchMinimal captures enum value "minimal"
	DeviceTransmissionFrequenciesSwitchMinimal string = "minimal"

	// DeviceTransmissionFrequenciesSwitchLow captures enum value "low"
	DeviceTransmissionFrequenciesSwitchLow string = "low"

	// DeviceTransmissionFrequenciesSwitchMedium captures enum value "medium"
	DeviceTransmissionFrequenciesSwitchMedium string = "medium"

	// DeviceTransmissionFrequenciesSwitchHigh captures enum value "high"
	DeviceTransmissionFrequenciesSwitchHigh string = "high"

	// DeviceTransmissionFrequenciesSwitchRealtime captures enum value "realtime"
	DeviceTransmissionFrequenciesSwitchRealtime string = "realtime"
)

// prop value enum
func (m *DeviceTransmissionFrequencies) validateSwitchEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, deviceTransmissionFrequenciesTypeSwitchPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DeviceTransmissionFrequencies) validateSwitch(formats strfmt.Registry) error {

	if err := validate.Required("switch", "body", m.Switch); err != nil {
		return err
	}

	// value enum
	if err := m.validateSwitchEnum("switch", "body", *m.Switch); err != nil {
		return err
	}

	return nil
}

var deviceTransmissionFrequenciesTypeUpsPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["minimal","low","medium","high","realtime"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviceTransmissionFrequenciesTypeUpsPropEnum = append(deviceTransmissionFrequenciesTypeUpsPropEnum, v)
	}
}

const (

	// DeviceTransmissionFrequenciesUpsMinimal captures enum value "minimal"
	DeviceTransmissionFrequenciesUpsMinimal string = "minimal"

	// DeviceTransmissionFrequenciesUpsLow captures enum value "low"
	DeviceTransmissionFrequenciesUpsLow string = "low"

	// DeviceTransmissionFrequenciesUpsMedium captures enum value "medium"
	DeviceTransmissionFrequenciesUpsMedium string = "medium"

	// DeviceTransmissionFrequenciesUpsHigh captures enum value "high"
	DeviceTransmissionFrequenciesUpsHigh string = "high"

	// DeviceTransmissionFrequenciesUpsRealtime captures enum value "realtime"
	DeviceTransmissionFrequenciesUpsRealtime string = "realtime"
)

// prop value enum
func (m *DeviceTransmissionFrequencies) validateUpsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, deviceTransmissionFrequenciesTypeUpsPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DeviceTransmissionFrequencies) validateUps(formats strfmt.Registry) error {

	if err := validate.Required("ups", "body", m.Ups); err != nil {
		return err
	}

	// value enum
	if err := m.validateUpsEnum("ups", "body", *m.Ups); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceTransmissionFrequencies) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceTransmissionFrequencies) UnmarshalBinary(b []byte) error {
	var res DeviceTransmissionFrequencies
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}