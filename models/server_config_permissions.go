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

// ServerConfigPermissions server config permissions
// swagger:model ServerConfigPermissions
type ServerConfigPermissions struct {

	// Whether SSL certificate can be configured or not.
	// Required: true
	CanConfigureCertificate *bool `json:"canConfigureCertificate"`

	// Whether device profiles can be configured or not.
	// Required: true
	CanConfigureDeviceProfiles *bool `json:"canConfigureDeviceProfiles"`

	// Whether UNMS setup wizard can show the First device screen.
	// Required: true
	CanConfigureFirstDeviceInSetup *bool `json:"canConfigureFirstDeviceInSetup"`

	// Whether hostname can be configured or not.
	// Required: true
	CanConfigureHostname *bool `json:"canConfigureHostname"`

	// Whether map engine is configurable or not.
	// Required: true
	CanConfigureMaps *bool `json:"canConfigureMaps"`

	// Whether netflow can be configured or not.
	// Required: true
	CanConfigureNetflow *bool `json:"canConfigureNetflow"`

	// Whether SMTP can be configured or not.
	// Required: true
	CanConfigureSMTP *bool `json:"canConfigureSmtp"`

	// Whether CRM can be configured from UNMS UI or not.
	// Required: true
	CanConfigureUcrm *bool `json:"canConfigureUcrm"`

	// Whether support info can be downloaded or not.
	// Required: true
	CanDownloadSupportInfo *bool `json:"canDownloadSupportInfo"`

	// Whether UNMS can discover devices in local network.
	// Required: true
	CanRunLocalDiscovery *bool `json:"canRunLocalDiscovery"`

	// Whether UNMS setup wizard can be completed without authentication.
	// Required: true
	CanSetupWithoutAuthentication *bool `json:"canSetupWithoutAuthentication"`

	// Whether UNMS setup can skip privacy policy agreement screen.
	// Required: true
	CanSkipPrivacyPolicyAgreement *bool `json:"canSkipPrivacyPolicyAgreement"`

	// Whether UNMS can be updated from UNMS UI or not.
	// Required: true
	CanUpdateUnms *bool `json:"canUpdateUnms"`
}

// Validate validates this server config permissions
func (m *ServerConfigPermissions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCanConfigureCertificate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCanConfigureDeviceProfiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCanConfigureFirstDeviceInSetup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCanConfigureHostname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCanConfigureMaps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCanConfigureNetflow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCanConfigureSMTP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCanConfigureUcrm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCanDownloadSupportInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCanRunLocalDiscovery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCanSetupWithoutAuthentication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCanSkipPrivacyPolicyAgreement(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCanUpdateUnms(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServerConfigPermissions) validateCanConfigureCertificate(formats strfmt.Registry) error {

	if err := validate.Required("canConfigureCertificate", "body", m.CanConfigureCertificate); err != nil {
		return err
	}

	return nil
}

func (m *ServerConfigPermissions) validateCanConfigureDeviceProfiles(formats strfmt.Registry) error {

	if err := validate.Required("canConfigureDeviceProfiles", "body", m.CanConfigureDeviceProfiles); err != nil {
		return err
	}

	return nil
}

func (m *ServerConfigPermissions) validateCanConfigureFirstDeviceInSetup(formats strfmt.Registry) error {

	if err := validate.Required("canConfigureFirstDeviceInSetup", "body", m.CanConfigureFirstDeviceInSetup); err != nil {
		return err
	}

	return nil
}

func (m *ServerConfigPermissions) validateCanConfigureHostname(formats strfmt.Registry) error {

	if err := validate.Required("canConfigureHostname", "body", m.CanConfigureHostname); err != nil {
		return err
	}

	return nil
}

func (m *ServerConfigPermissions) validateCanConfigureMaps(formats strfmt.Registry) error {

	if err := validate.Required("canConfigureMaps", "body", m.CanConfigureMaps); err != nil {
		return err
	}

	return nil
}

func (m *ServerConfigPermissions) validateCanConfigureNetflow(formats strfmt.Registry) error {

	if err := validate.Required("canConfigureNetflow", "body", m.CanConfigureNetflow); err != nil {
		return err
	}

	return nil
}

func (m *ServerConfigPermissions) validateCanConfigureSMTP(formats strfmt.Registry) error {

	if err := validate.Required("canConfigureSmtp", "body", m.CanConfigureSMTP); err != nil {
		return err
	}

	return nil
}

func (m *ServerConfigPermissions) validateCanConfigureUcrm(formats strfmt.Registry) error {

	if err := validate.Required("canConfigureUcrm", "body", m.CanConfigureUcrm); err != nil {
		return err
	}

	return nil
}

func (m *ServerConfigPermissions) validateCanDownloadSupportInfo(formats strfmt.Registry) error {

	if err := validate.Required("canDownloadSupportInfo", "body", m.CanDownloadSupportInfo); err != nil {
		return err
	}

	return nil
}

func (m *ServerConfigPermissions) validateCanRunLocalDiscovery(formats strfmt.Registry) error {

	if err := validate.Required("canRunLocalDiscovery", "body", m.CanRunLocalDiscovery); err != nil {
		return err
	}

	return nil
}

func (m *ServerConfigPermissions) validateCanSetupWithoutAuthentication(formats strfmt.Registry) error {

	if err := validate.Required("canSetupWithoutAuthentication", "body", m.CanSetupWithoutAuthentication); err != nil {
		return err
	}

	return nil
}

func (m *ServerConfigPermissions) validateCanSkipPrivacyPolicyAgreement(formats strfmt.Registry) error {

	if err := validate.Required("canSkipPrivacyPolicyAgreement", "body", m.CanSkipPrivacyPolicyAgreement); err != nil {
		return err
	}

	return nil
}

func (m *ServerConfigPermissions) validateCanUpdateUnms(formats strfmt.Registry) error {

	if err := validate.Required("canUpdateUnms", "body", m.CanUpdateUnms); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServerConfigPermissions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerConfigPermissions) UnmarshalBinary(b []byte) error {
	var res ServerConfigPermissions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}