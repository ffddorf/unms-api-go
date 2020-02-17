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

// NewClientSchema new client schema
// swagger:model NewClientSchema
type NewClientSchema struct {

	// Address of the Client site.
	Address string `json:"address,omitempty"`

	// Client's e-mail.
	Email string `json:"email,omitempty"`

	// First name of the Client.
	// Required: true
	FirstName *string `json:"firstName"`

	// Last name of the Client.
	// Required: true
	LastName *string `json:"lastName"`

	// location
	Location *NewClientLocation `json:"location,omitempty"`

	// Custom note.
	Note string `json:"note,omitempty"`

	// Client's phone number.
	Phone string `json:"phone,omitempty"`

	// ID of the CRM Service plan that should be assigned to the new Client
	// Required: true
	ServicePlanID *string `json:"servicePlanId"`
}

// Validate validates this new client schema
func (m *NewClientSchema) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFirstName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServicePlanID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewClientSchema) validateFirstName(formats strfmt.Registry) error {

	if err := validate.Required("firstName", "body", m.FirstName); err != nil {
		return err
	}

	return nil
}

func (m *NewClientSchema) validateLastName(formats strfmt.Registry) error {

	if err := validate.Required("lastName", "body", m.LastName); err != nil {
		return err
	}

	return nil
}

func (m *NewClientSchema) validateLocation(formats strfmt.Registry) error {

	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *NewClientSchema) validateServicePlanID(formats strfmt.Registry) error {

	if err := validate.Required("servicePlanId", "body", m.ServicePlanID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewClientSchema) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewClientSchema) UnmarshalBinary(b []byte) error {
	var res NewClientSchema
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
