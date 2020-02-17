// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OnuPolicies onu policies
// swagger:model OnuPolicies
type OnuPolicies struct {

	// default state
	DefaultState string `json:"defaultState,omitempty"`

	// dhcp option82
	DhcpOption82 bool `json:"dhcpOption82,omitempty"`
}

// Validate validates this onu policies
func (m *OnuPolicies) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OnuPolicies) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OnuPolicies) UnmarshalBinary(b []byte) error {
	var res OnuPolicies
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}