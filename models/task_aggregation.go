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

// TaskAggregation task aggregation
// swagger:model TaskAggregation
type TaskAggregation struct {

	// canceled
	// Minimum: 0
	Canceled *int64 `json:"canceled,omitempty"`

	// failed
	// Minimum: 0
	Failed *int64 `json:"failed,omitempty"`

	// in progress
	// Minimum: 0
	InProgress *int64 `json:"in-progress,omitempty"`

	// queued
	// Minimum: 0
	Queued *int64 `json:"queued,omitempty"`

	// success
	// Minimum: 0
	Success *int64 `json:"success,omitempty"`
}

// Validate validates this task aggregation
func (m *TaskAggregation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCanceled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFailed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInProgress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueued(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuccess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskAggregation) validateCanceled(formats strfmt.Registry) error {

	if swag.IsZero(m.Canceled) { // not required
		return nil
	}

	if err := validate.MinimumInt("canceled", "body", int64(*m.Canceled), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *TaskAggregation) validateFailed(formats strfmt.Registry) error {

	if swag.IsZero(m.Failed) { // not required
		return nil
	}

	if err := validate.MinimumInt("failed", "body", int64(*m.Failed), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *TaskAggregation) validateInProgress(formats strfmt.Registry) error {

	if swag.IsZero(m.InProgress) { // not required
		return nil
	}

	if err := validate.MinimumInt("in-progress", "body", int64(*m.InProgress), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *TaskAggregation) validateQueued(formats strfmt.Registry) error {

	if swag.IsZero(m.Queued) { // not required
		return nil
	}

	if err := validate.MinimumInt("queued", "body", int64(*m.Queued), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *TaskAggregation) validateSuccess(formats strfmt.Registry) error {

	if swag.IsZero(m.Success) { // not required
		return nil
	}

	if err := validate.MinimumInt("success", "body", int64(*m.Success), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaskAggregation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskAggregation) UnmarshalBinary(b []byte) error {
	var res TaskAggregation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
