// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NetworkStatisticInterval network statistic interval
// swagger:model NetworkStatisticInterval
type NetworkStatisticInterval struct {

	// end
	End string `json:"end,omitempty"`

	// start
	Start string `json:"start,omitempty"`
}

// Validate validates this network statistic interval
func (m *NetworkStatisticInterval) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkStatisticInterval) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkStatisticInterval) UnmarshalBinary(b []byte) error {
	var res NetworkStatisticInterval
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
