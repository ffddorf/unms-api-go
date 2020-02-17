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

// SiteTrafficShaping1 site traffic shaping 1
// swagger:model SiteTrafficShaping 1
type SiteTrafficShaping1 struct {

	//
	//       Number indicating how many times the guaranteed rate is less than the maximum allowed bandwidth.
	//       Null for auto calculation by UNMS.
	//
	// Required: true
	// Maximum: 10000
	// Minimum: 1
	Aggregation *int64 `json:"aggregation"`

	// Amount of bytes that can be sent at downloadSpeed in excess of the guaranteed rate.
	// Required: true
	// Maximum: 1e+11
	// Enum: [0]
	DownloadBurstSize *int64 `json:"downloadBurstSize"`

	// Download speed limit in bps.
	// Required: true
	// Maximum: 1e+11
	// Enum: [0]
	DownloadSpeed *int64 `json:"downloadSpeed"`

	// Set to TRUE if optional Traffic Shaping queue are activated.
	// Required: true
	Enabled *bool `json:"enabled"`

	//
	//       Whether or not set Traffic Shaping on Client devices like, CPEs and ONUs.
	//       Enabling this will lower your traffic in local network.
	//
	// Enum: [all gateway]
	Propagation string `json:"propagation,omitempty"`

	// Amount of bytes that can be sent at uploadSpeed in excess of the guaranteed rate.
	// Required: true
	// Maximum: 1e+11
	// Enum: [0]
	UploadBurstSize *int64 `json:"uploadBurstSize"`

	// Upload speed limit in bps.
	// Required: true
	// Maximum: 1e+11
	// Enum: [0]
	UploadSpeed *int64 `json:"uploadSpeed"`
}

// Validate validates this site traffic shaping 1
func (m *SiteTrafficShaping1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAggregation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDownloadBurstSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDownloadSpeed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePropagation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUploadBurstSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUploadSpeed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SiteTrafficShaping1) validateAggregation(formats strfmt.Registry) error {

	if err := validate.Required("aggregation", "body", m.Aggregation); err != nil {
		return err
	}

	if err := validate.MinimumInt("aggregation", "body", int64(*m.Aggregation), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("aggregation", "body", int64(*m.Aggregation), 10000, false); err != nil {
		return err
	}

	return nil
}

var siteTrafficShaping1TypeDownloadBurstSizePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[0]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		siteTrafficShaping1TypeDownloadBurstSizePropEnum = append(siteTrafficShaping1TypeDownloadBurstSizePropEnum, v)
	}
}

// prop value enum
func (m *SiteTrafficShaping1) validateDownloadBurstSizeEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, siteTrafficShaping1TypeDownloadBurstSizePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SiteTrafficShaping1) validateDownloadBurstSize(formats strfmt.Registry) error {

	if err := validate.Required("downloadBurstSize", "body", m.DownloadBurstSize); err != nil {
		return err
	}

	if err := validate.MaximumInt("downloadBurstSize", "body", int64(*m.DownloadBurstSize), 1e+11, false); err != nil {
		return err
	}

	// value enum
	if err := m.validateDownloadBurstSizeEnum("downloadBurstSize", "body", *m.DownloadBurstSize); err != nil {
		return err
	}

	return nil
}

var siteTrafficShaping1TypeDownloadSpeedPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[0]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		siteTrafficShaping1TypeDownloadSpeedPropEnum = append(siteTrafficShaping1TypeDownloadSpeedPropEnum, v)
	}
}

// prop value enum
func (m *SiteTrafficShaping1) validateDownloadSpeedEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, siteTrafficShaping1TypeDownloadSpeedPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SiteTrafficShaping1) validateDownloadSpeed(formats strfmt.Registry) error {

	if err := validate.Required("downloadSpeed", "body", m.DownloadSpeed); err != nil {
		return err
	}

	if err := validate.MaximumInt("downloadSpeed", "body", int64(*m.DownloadSpeed), 1e+11, false); err != nil {
		return err
	}

	// value enum
	if err := m.validateDownloadSpeedEnum("downloadSpeed", "body", *m.DownloadSpeed); err != nil {
		return err
	}

	return nil
}

func (m *SiteTrafficShaping1) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

var siteTrafficShaping1TypePropagationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["all","gateway"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		siteTrafficShaping1TypePropagationPropEnum = append(siteTrafficShaping1TypePropagationPropEnum, v)
	}
}

const (

	// SiteTrafficShaping1PropagationAll captures enum value "all"
	SiteTrafficShaping1PropagationAll string = "all"

	// SiteTrafficShaping1PropagationGateway captures enum value "gateway"
	SiteTrafficShaping1PropagationGateway string = "gateway"
)

// prop value enum
func (m *SiteTrafficShaping1) validatePropagationEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, siteTrafficShaping1TypePropagationPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SiteTrafficShaping1) validatePropagation(formats strfmt.Registry) error {

	if swag.IsZero(m.Propagation) { // not required
		return nil
	}

	// value enum
	if err := m.validatePropagationEnum("propagation", "body", m.Propagation); err != nil {
		return err
	}

	return nil
}

var siteTrafficShaping1TypeUploadBurstSizePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[0]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		siteTrafficShaping1TypeUploadBurstSizePropEnum = append(siteTrafficShaping1TypeUploadBurstSizePropEnum, v)
	}
}

// prop value enum
func (m *SiteTrafficShaping1) validateUploadBurstSizeEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, siteTrafficShaping1TypeUploadBurstSizePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SiteTrafficShaping1) validateUploadBurstSize(formats strfmt.Registry) error {

	if err := validate.Required("uploadBurstSize", "body", m.UploadBurstSize); err != nil {
		return err
	}

	if err := validate.MaximumInt("uploadBurstSize", "body", int64(*m.UploadBurstSize), 1e+11, false); err != nil {
		return err
	}

	// value enum
	if err := m.validateUploadBurstSizeEnum("uploadBurstSize", "body", *m.UploadBurstSize); err != nil {
		return err
	}

	return nil
}

var siteTrafficShaping1TypeUploadSpeedPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[0]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		siteTrafficShaping1TypeUploadSpeedPropEnum = append(siteTrafficShaping1TypeUploadSpeedPropEnum, v)
	}
}

// prop value enum
func (m *SiteTrafficShaping1) validateUploadSpeedEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, siteTrafficShaping1TypeUploadSpeedPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SiteTrafficShaping1) validateUploadSpeed(formats strfmt.Registry) error {

	if err := validate.Required("uploadSpeed", "body", m.UploadSpeed); err != nil {
		return err
	}

	if err := validate.MaximumInt("uploadSpeed", "body", int64(*m.UploadSpeed), 1e+11, false); err != nil {
		return err
	}

	// value enum
	if err := m.validateUploadSpeedEnum("uploadSpeed", "body", *m.UploadSpeed); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SiteTrafficShaping1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SiteTrafficShaping1) UnmarshalBinary(b []byte) error {
	var res SiteTrafficShaping1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}