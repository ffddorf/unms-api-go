// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Aggregation aggregation
// swagger:model aggregation
type Aggregation struct {

	// all count
	AllCount float64 `json:"allCount,omitempty"`

	// outage count
	OutageCount float64 `json:"outageCount,omitempty"`

	// unreachable count
	UnreachableCount float64 `json:"unreachableCount,omitempty"`
}

// Validate validates this aggregation
func (m *Aggregation) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Aggregation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Aggregation) UnmarshalBinary(b []byte) error {
	var res Aggregation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}