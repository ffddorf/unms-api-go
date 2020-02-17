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

// MacACLConfig mac ACL config
// swagger:model macACLConfig
type MacACLConfig struct {

	// clients
	Clients Clients `json:"clients,omitempty"`

	// is enabled
	// Required: true
	IsEnabled *bool `json:"isEnabled"`

	// policy
	// Required: true
	// Enum: [allow deny]
	Policy *string `json:"policy"`
}

// Validate validates this mac ACL config
func (m *MacACLConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClients(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MacACLConfig) validateClients(formats strfmt.Registry) error {

	if swag.IsZero(m.Clients) { // not required
		return nil
	}

	if err := m.Clients.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("clients")
		}
		return err
	}

	return nil
}

func (m *MacACLConfig) validateIsEnabled(formats strfmt.Registry) error {

	if err := validate.Required("isEnabled", "body", m.IsEnabled); err != nil {
		return err
	}

	return nil
}

var macAclConfigTypePolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["allow","deny"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		macAclConfigTypePolicyPropEnum = append(macAclConfigTypePolicyPropEnum, v)
	}
}

const (

	// MacACLConfigPolicyAllow captures enum value "allow"
	MacACLConfigPolicyAllow string = "allow"

	// MacACLConfigPolicyDeny captures enum value "deny"
	MacACLConfigPolicyDeny string = "deny"
)

// prop value enum
func (m *MacACLConfig) validatePolicyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, macAclConfigTypePolicyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MacACLConfig) validatePolicy(formats strfmt.Registry) error {

	if err := validate.Required("policy", "body", m.Policy); err != nil {
		return err
	}

	// value enum
	if err := m.validatePolicyEnum("policy", "body", *m.Policy); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MacACLConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MacACLConfig) UnmarshalBinary(b []byte) error {
	var res MacACLConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
