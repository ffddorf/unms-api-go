// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Site2 Result of unsuspend, updated site
// swagger:model Site 2
type Site2 struct {

	// description
	Description *SiteDescription `json:"description,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// identification
	Identification *SiteIdentificationDetail `json:"identification,omitempty"`

	// notifications
	Notifications *SiteNotifications `json:"notifications,omitempty"`

	// qos
	Qos *SiteTrafficShaping `json:"qos,omitempty"`

	// ucrm
	Ucrm *SiteUcrmDescription `json:"ucrm,omitempty"`
}

// Validate validates this site 2
func (m *Site2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotifications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQos(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUcrm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Site2) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if m.Description != nil {
		if err := m.Description.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("description")
			}
			return err
		}
	}

	return nil
}

func (m *Site2) validateIdentification(formats strfmt.Registry) error {

	if swag.IsZero(m.Identification) { // not required
		return nil
	}

	if m.Identification != nil {
		if err := m.Identification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identification")
			}
			return err
		}
	}

	return nil
}

func (m *Site2) validateNotifications(formats strfmt.Registry) error {

	if swag.IsZero(m.Notifications) { // not required
		return nil
	}

	if m.Notifications != nil {
		if err := m.Notifications.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notifications")
			}
			return err
		}
	}

	return nil
}

func (m *Site2) validateQos(formats strfmt.Registry) error {

	if swag.IsZero(m.Qos) { // not required
		return nil
	}

	if m.Qos != nil {
		if err := m.Qos.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("qos")
			}
			return err
		}
	}

	return nil
}

func (m *Site2) validateUcrm(formats strfmt.Registry) error {

	if swag.IsZero(m.Ucrm) { // not required
		return nil
	}

	if m.Ucrm != nil {
		if err := m.Ucrm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ucrm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Site2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Site2) UnmarshalBinary(b []byte) error {
	var res Site2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}