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

// SiteBindings site bindings
// swagger:model SiteBindings
type SiteBindings struct {

	// site Id
	// Required: true
	SiteID *string `json:"siteId"`

	// ucrm Id
	// Required: true
	UcrmID *string `json:"ucrmId"`
}

// Validate validates this site bindings
func (m *SiteBindings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSiteID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUcrmID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SiteBindings) validateSiteID(formats strfmt.Registry) error {

	if err := validate.Required("siteId", "body", m.SiteID); err != nil {
		return err
	}

	return nil
}

func (m *SiteBindings) validateUcrmID(formats strfmt.Registry) error {

	if err := validate.Required("ucrmId", "body", m.UcrmID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SiteBindings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SiteBindings) UnmarshalBinary(b []byte) error {
	var res SiteBindings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
