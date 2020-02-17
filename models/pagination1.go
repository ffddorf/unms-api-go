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

// Pagination1 pagination 1
// swagger:model Pagination 1
type Pagination1 struct {

	// selected items
	// Minimum: 0
	Count *int64 `json:"count,omitempty"`

	// actual page
	// Minimum: 1
	Page int64 `json:"page,omitempty"`

	// pages count
	// Minimum: 1
	Pages int64 `json:"pages,omitempty"`

	// count of found items
	// Minimum: 0
	Total *int64 `json:"total,omitempty"`
}

// Validate validates this pagination 1
func (m *Pagination1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Pagination1) validateCount(formats strfmt.Registry) error {

	if swag.IsZero(m.Count) { // not required
		return nil
	}

	if err := validate.MinimumInt("count", "body", int64(*m.Count), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Pagination1) validatePage(formats strfmt.Registry) error {

	if swag.IsZero(m.Page) { // not required
		return nil
	}

	if err := validate.MinimumInt("page", "body", int64(m.Page), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *Pagination1) validatePages(formats strfmt.Registry) error {

	if swag.IsZero(m.Pages) { // not required
		return nil
	}

	if err := validate.MinimumInt("pages", "body", int64(m.Pages), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *Pagination1) validateTotal(formats strfmt.Registry) error {

	if swag.IsZero(m.Total) { // not required
		return nil
	}

	if err := validate.MinimumInt("total", "body", int64(*m.Total), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Pagination1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Pagination1) UnmarshalBinary(b []byte) error {
	var res Pagination1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}