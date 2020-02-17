// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VlanListSchema vlan list schema
// swagger:model vlanListSchema
type VlanListSchema struct {

	// vlans
	Vlans Vlans1 `json:"vlans,omitempty"`
}

// Validate validates this vlan list schema
func (m *VlanListSchema) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVlans(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VlanListSchema) validateVlans(formats strfmt.Registry) error {

	if swag.IsZero(m.Vlans) { // not required
		return nil
	}

	if err := m.Vlans.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("vlans")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VlanListSchema) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VlanListSchema) UnmarshalBinary(b []byte) error {
	var res VlanListSchema
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
