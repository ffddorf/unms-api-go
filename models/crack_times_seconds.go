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

// CrackTimesSeconds crack times seconds
// swagger:model crack_times_seconds
type CrackTimesSeconds struct {

	// offline fast hashing 1e10 per second
	// Required: true
	OfflineFastHashing1e10PerSecond *float64 `json:"offline_fast_hashing_1e10_per_second"`

	// offline slow hashing 1e4 per second
	// Required: true
	OfflineSlowHashing1e4PerSecond *float64 `json:"offline_slow_hashing_1e4_per_second"`

	// online no throttling 10 per second
	// Required: true
	OnlineNoThrottling10PerSecond *float64 `json:"online_no_throttling_10_per_second"`

	// online throttling 100 per hour
	// Required: true
	OnlineThrottling100PerHour *float64 `json:"online_throttling_100_per_hour"`
}

// Validate validates this crack times seconds
func (m *CrackTimesSeconds) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOfflineFastHashing1e10PerSecond(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOfflineSlowHashing1e4PerSecond(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnlineNoThrottling10PerSecond(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnlineThrottling100PerHour(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CrackTimesSeconds) validateOfflineFastHashing1e10PerSecond(formats strfmt.Registry) error {

	if err := validate.Required("offline_fast_hashing_1e10_per_second", "body", m.OfflineFastHashing1e10PerSecond); err != nil {
		return err
	}

	return nil
}

func (m *CrackTimesSeconds) validateOfflineSlowHashing1e4PerSecond(formats strfmt.Registry) error {

	if err := validate.Required("offline_slow_hashing_1e4_per_second", "body", m.OfflineSlowHashing1e4PerSecond); err != nil {
		return err
	}

	return nil
}

func (m *CrackTimesSeconds) validateOnlineNoThrottling10PerSecond(formats strfmt.Registry) error {

	if err := validate.Required("online_no_throttling_10_per_second", "body", m.OnlineNoThrottling10PerSecond); err != nil {
		return err
	}

	return nil
}

func (m *CrackTimesSeconds) validateOnlineThrottling100PerHour(formats strfmt.Registry) error {

	if err := validate.Required("online_throttling_100_per_hour", "body", m.OnlineThrottling100PerHour); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CrackTimesSeconds) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CrackTimesSeconds) UnmarshalBinary(b []byte) error {
	var res CrackTimesSeconds
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
