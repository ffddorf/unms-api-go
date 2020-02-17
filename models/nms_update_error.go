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

// NmsUpdateError Error that caused the update to fail. Null if update succeeded or is still running.
// swagger:model NmsUpdateError
type NmsUpdateError struct {

	// Type of update error.
	// Required: true
	// Enum: [DiskSpace InstallationPackage DockerStop OldDocker Timeout PullImages Unknown]
	Error *string `json:"error"`

	// Brief description of the update error.
	// Required: true
	Message *string `json:"message"`

	// Time when the error occured.
	// Required: true
	// Format: date
	Time *strfmt.Date `json:"time"`
}

// Validate validates this nms update error
func (m *NmsUpdateError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var nmsUpdateErrorTypeErrorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DiskSpace","InstallationPackage","DockerStop","OldDocker","Timeout","PullImages","Unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nmsUpdateErrorTypeErrorPropEnum = append(nmsUpdateErrorTypeErrorPropEnum, v)
	}
}

const (

	// NmsUpdateErrorErrorDiskSpace captures enum value "DiskSpace"
	NmsUpdateErrorErrorDiskSpace string = "DiskSpace"

	// NmsUpdateErrorErrorInstallationPackage captures enum value "InstallationPackage"
	NmsUpdateErrorErrorInstallationPackage string = "InstallationPackage"

	// NmsUpdateErrorErrorDockerStop captures enum value "DockerStop"
	NmsUpdateErrorErrorDockerStop string = "DockerStop"

	// NmsUpdateErrorErrorOldDocker captures enum value "OldDocker"
	NmsUpdateErrorErrorOldDocker string = "OldDocker"

	// NmsUpdateErrorErrorTimeout captures enum value "Timeout"
	NmsUpdateErrorErrorTimeout string = "Timeout"

	// NmsUpdateErrorErrorPullImages captures enum value "PullImages"
	NmsUpdateErrorErrorPullImages string = "PullImages"

	// NmsUpdateErrorErrorUnknown captures enum value "Unknown"
	NmsUpdateErrorErrorUnknown string = "Unknown"
)

// prop value enum
func (m *NmsUpdateError) validateErrorEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, nmsUpdateErrorTypeErrorPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NmsUpdateError) validateError(formats strfmt.Registry) error {

	if err := validate.Required("error", "body", m.Error); err != nil {
		return err
	}

	// value enum
	if err := m.validateErrorEnum("error", "body", *m.Error); err != nil {
		return err
	}

	return nil
}

func (m *NmsUpdateError) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *NmsUpdateError) validateTime(formats strfmt.Registry) error {

	if err := validate.Required("time", "body", m.Time); err != nil {
		return err
	}

	if err := validate.FormatOf("time", "body", "date", m.Time.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NmsUpdateError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NmsUpdateError) UnmarshalBinary(b []byte) error {
	var res NmsUpdateError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}