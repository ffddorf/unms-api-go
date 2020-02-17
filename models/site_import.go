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

// SiteImport site import
// swagger:model SiteImport
type SiteImport struct {

	// Address of the site.
	Address string `json:"address,omitempty"`

	// Email of the contact person.
	ContactEmail string `json:"contactEmail,omitempty"`

	// Name of the contact person.
	ContactName string `json:"contactName,omitempty"`

	// Phone number of the contact person.
	ContactPhone string `json:"contactPhone,omitempty"`

	// Site elevation without structure height.
	Elevation float64 `json:"elevation,omitempty"`

	// Site structure height.
	Height float64 `json:"height,omitempty"`

	// ip addresses
	IPAddresses IPAddresses1 `json:"ipAddresses,omitempty"`

	// Site longitude.
	// Maximum: 90
	// Minimum: -90
	Latitude *float64 `json:"latitude,omitempty"`

	// Site latitude.
	// Maximum: 180
	// Minimum: -180
	Longitude *float64 `json:"longitude,omitempty"`

	// mac addresses
	MacAddresses MacAddresses `json:"macAddresses,omitempty"`

	// Name of the site.
	// Required: true
	Name *string `json:"name"`

	// Any additional site description.
	Note string `json:"note,omitempty"`

	// ID of the parent site.
	ParentSiteID string `json:"parentSiteId,omitempty"`

	// Name of the parent site.
	ParentSiteName string `json:"parentSiteName,omitempty"`

	// ID of UCRM parent site of a client. Null if UCRM client has no parent site.
	ParentSiteUcrmID string `json:"parentSiteUcrmId,omitempty"`

	// qos
	Qos *SiteTrafficShaping1 `json:"qos,omitempty"`

	// Location of regulation.
	// Enum: [XX XY AF AX AL DZ AS AD AO AI AQ AG AR AM AW AU AT AZ BS BH BD BB BY BE BZ BJ BM BT BO BQ BA BW BV BR IO BN BG BF BI CV KH CM CA KY CF TD CL CN CX CC CO KM CK CR CI HR CW CY CZ CD DK DJ DM DO EC EG SV GQ ER EE SZ ET FK FO FJ FI FR GF PF TF GA GM GE DE GH GI GR GL GD GP GU GT GG GN GW GY HT HM HN HK HU IS IN ID IQ IE IM IL IT JM JP JE JO KZ KE KI KW KG LA LV LB LS LR LY LI LT LU MO MK MG MW MY MV ML MT MH MQ MR MU YT MX FM MD MC MN ME MS MA MZ MM NA NR NP NL NC NZ NI NE NG NU NF MP NO OM PK PW PA PG PY PE PH PN PL PT PR QA KR RS SC CG RE RO RU RW BL SH KN LC MF PM VC WS SM ST SA SN SL SG SX SK SI SB SO ZA GS SS ES LK PS SR SJ SE CH TW TJ TZ TH TL TG TK TO TT TN TR TM TC TV UG UA AE GB US UY UZ VU VA VE VN VG VI WF EH YE ZM ZW]
	RegulatoryDomain string `json:"regulatoryDomain,omitempty"`

	// If site is suspended.
	Suspended *bool `json:"suspended,omitempty"`

	// Type of the site.
	// Required: true
	// Enum: [site client]
	Type *string `json:"type"`

	// In case of client type, ID of UCRM service. In case of site type, ID of UCRM site.
	UcrmID string `json:"ucrmId,omitempty"`
}

// Validate validates this site import
func (m *SiteImport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatitude(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLongitude(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMacAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQos(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegulatoryDomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SiteImport) validateIPAddresses(formats strfmt.Registry) error {

	if swag.IsZero(m.IPAddresses) { // not required
		return nil
	}

	if err := m.IPAddresses.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ipAddresses")
		}
		return err
	}

	return nil
}

func (m *SiteImport) validateLatitude(formats strfmt.Registry) error {

	if swag.IsZero(m.Latitude) { // not required
		return nil
	}

	if err := validate.Minimum("latitude", "body", float64(*m.Latitude), -90, false); err != nil {
		return err
	}

	if err := validate.Maximum("latitude", "body", float64(*m.Latitude), 90, false); err != nil {
		return err
	}

	return nil
}

func (m *SiteImport) validateLongitude(formats strfmt.Registry) error {

	if swag.IsZero(m.Longitude) { // not required
		return nil
	}

	if err := validate.Minimum("longitude", "body", float64(*m.Longitude), -180, false); err != nil {
		return err
	}

	if err := validate.Maximum("longitude", "body", float64(*m.Longitude), 180, false); err != nil {
		return err
	}

	return nil
}

func (m *SiteImport) validateMacAddresses(formats strfmt.Registry) error {

	if swag.IsZero(m.MacAddresses) { // not required
		return nil
	}

	if err := m.MacAddresses.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("macAddresses")
		}
		return err
	}

	return nil
}

func (m *SiteImport) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *SiteImport) validateQos(formats strfmt.Registry) error {

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

var siteImportTypeRegulatoryDomainPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["XX","XY","AF","AX","AL","DZ","AS","AD","AO","AI","AQ","AG","AR","AM","AW","AU","AT","AZ","BS","BH","BD","BB","BY","BE","BZ","BJ","BM","BT","BO","BQ","BA","BW","BV","BR","IO","BN","BG","BF","BI","CV","KH","CM","CA","KY","CF","TD","CL","CN","CX","CC","CO","KM","CK","CR","CI","HR","CW","CY","CZ","CD","DK","DJ","DM","DO","EC","EG","SV","GQ","ER","EE","SZ","ET","FK","FO","FJ","FI","FR","GF","PF","TF","GA","GM","GE","DE","GH","GI","GR","GL","GD","GP","GU","GT","GG","GN","GW","GY","HT","HM","HN","HK","HU","IS","IN","ID","IQ","IE","IM","IL","IT","JM","JP","JE","JO","KZ","KE","KI","KW","KG","LA","LV","LB","LS","LR","LY","LI","LT","LU","MO","MK","MG","MW","MY","MV","ML","MT","MH","MQ","MR","MU","YT","MX","FM","MD","MC","MN","ME","MS","MA","MZ","MM","NA","NR","NP","NL","NC","NZ","NI","NE","NG","NU","NF","MP","NO","OM","PK","PW","PA","PG","PY","PE","PH","PN","PL","PT","PR","QA","KR","RS","SC","CG","RE","RO","RU","RW","BL","SH","KN","LC","MF","PM","VC","WS","SM","ST","SA","SN","SL","SG","SX","SK","SI","SB","SO","ZA","GS","SS","ES","LK","PS","SR","SJ","SE","CH","TW","TJ","TZ","TH","TL","TG","TK","TO","TT","TN","TR","TM","TC","TV","UG","UA","AE","GB","US","UY","UZ","VU","VA","VE","VN","VG","VI","WF","EH","YE","ZM","ZW"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		siteImportTypeRegulatoryDomainPropEnum = append(siteImportTypeRegulatoryDomainPropEnum, v)
	}
}

const (

	// SiteImportRegulatoryDomainXX captures enum value "XX"
	SiteImportRegulatoryDomainXX string = "XX"

	// SiteImportRegulatoryDomainXY captures enum value "XY"
	SiteImportRegulatoryDomainXY string = "XY"

	// SiteImportRegulatoryDomainAF captures enum value "AF"
	SiteImportRegulatoryDomainAF string = "AF"

	// SiteImportRegulatoryDomainAX captures enum value "AX"
	SiteImportRegulatoryDomainAX string = "AX"

	// SiteImportRegulatoryDomainAL captures enum value "AL"
	SiteImportRegulatoryDomainAL string = "AL"

	// SiteImportRegulatoryDomainDZ captures enum value "DZ"
	SiteImportRegulatoryDomainDZ string = "DZ"

	// SiteImportRegulatoryDomainAS captures enum value "AS"
	SiteImportRegulatoryDomainAS string = "AS"

	// SiteImportRegulatoryDomainAD captures enum value "AD"
	SiteImportRegulatoryDomainAD string = "AD"

	// SiteImportRegulatoryDomainAO captures enum value "AO"
	SiteImportRegulatoryDomainAO string = "AO"

	// SiteImportRegulatoryDomainAI captures enum value "AI"
	SiteImportRegulatoryDomainAI string = "AI"

	// SiteImportRegulatoryDomainAQ captures enum value "AQ"
	SiteImportRegulatoryDomainAQ string = "AQ"

	// SiteImportRegulatoryDomainAG captures enum value "AG"
	SiteImportRegulatoryDomainAG string = "AG"

	// SiteImportRegulatoryDomainAR captures enum value "AR"
	SiteImportRegulatoryDomainAR string = "AR"

	// SiteImportRegulatoryDomainAM captures enum value "AM"
	SiteImportRegulatoryDomainAM string = "AM"

	// SiteImportRegulatoryDomainAW captures enum value "AW"
	SiteImportRegulatoryDomainAW string = "AW"

	// SiteImportRegulatoryDomainAU captures enum value "AU"
	SiteImportRegulatoryDomainAU string = "AU"

	// SiteImportRegulatoryDomainAT captures enum value "AT"
	SiteImportRegulatoryDomainAT string = "AT"

	// SiteImportRegulatoryDomainAZ captures enum value "AZ"
	SiteImportRegulatoryDomainAZ string = "AZ"

	// SiteImportRegulatoryDomainBS captures enum value "BS"
	SiteImportRegulatoryDomainBS string = "BS"

	// SiteImportRegulatoryDomainBH captures enum value "BH"
	SiteImportRegulatoryDomainBH string = "BH"

	// SiteImportRegulatoryDomainBD captures enum value "BD"
	SiteImportRegulatoryDomainBD string = "BD"

	// SiteImportRegulatoryDomainBB captures enum value "BB"
	SiteImportRegulatoryDomainBB string = "BB"

	// SiteImportRegulatoryDomainBY captures enum value "BY"
	SiteImportRegulatoryDomainBY string = "BY"

	// SiteImportRegulatoryDomainBE captures enum value "BE"
	SiteImportRegulatoryDomainBE string = "BE"

	// SiteImportRegulatoryDomainBZ captures enum value "BZ"
	SiteImportRegulatoryDomainBZ string = "BZ"

	// SiteImportRegulatoryDomainBJ captures enum value "BJ"
	SiteImportRegulatoryDomainBJ string = "BJ"

	// SiteImportRegulatoryDomainBM captures enum value "BM"
	SiteImportRegulatoryDomainBM string = "BM"

	// SiteImportRegulatoryDomainBT captures enum value "BT"
	SiteImportRegulatoryDomainBT string = "BT"

	// SiteImportRegulatoryDomainBO captures enum value "BO"
	SiteImportRegulatoryDomainBO string = "BO"

	// SiteImportRegulatoryDomainBQ captures enum value "BQ"
	SiteImportRegulatoryDomainBQ string = "BQ"

	// SiteImportRegulatoryDomainBA captures enum value "BA"
	SiteImportRegulatoryDomainBA string = "BA"

	// SiteImportRegulatoryDomainBW captures enum value "BW"
	SiteImportRegulatoryDomainBW string = "BW"

	// SiteImportRegulatoryDomainBV captures enum value "BV"
	SiteImportRegulatoryDomainBV string = "BV"

	// SiteImportRegulatoryDomainBR captures enum value "BR"
	SiteImportRegulatoryDomainBR string = "BR"

	// SiteImportRegulatoryDomainIO captures enum value "IO"
	SiteImportRegulatoryDomainIO string = "IO"

	// SiteImportRegulatoryDomainBN captures enum value "BN"
	SiteImportRegulatoryDomainBN string = "BN"

	// SiteImportRegulatoryDomainBG captures enum value "BG"
	SiteImportRegulatoryDomainBG string = "BG"

	// SiteImportRegulatoryDomainBF captures enum value "BF"
	SiteImportRegulatoryDomainBF string = "BF"

	// SiteImportRegulatoryDomainBI captures enum value "BI"
	SiteImportRegulatoryDomainBI string = "BI"

	// SiteImportRegulatoryDomainCV captures enum value "CV"
	SiteImportRegulatoryDomainCV string = "CV"

	// SiteImportRegulatoryDomainKH captures enum value "KH"
	SiteImportRegulatoryDomainKH string = "KH"

	// SiteImportRegulatoryDomainCM captures enum value "CM"
	SiteImportRegulatoryDomainCM string = "CM"

	// SiteImportRegulatoryDomainCA captures enum value "CA"
	SiteImportRegulatoryDomainCA string = "CA"

	// SiteImportRegulatoryDomainKY captures enum value "KY"
	SiteImportRegulatoryDomainKY string = "KY"

	// SiteImportRegulatoryDomainCF captures enum value "CF"
	SiteImportRegulatoryDomainCF string = "CF"

	// SiteImportRegulatoryDomainTD captures enum value "TD"
	SiteImportRegulatoryDomainTD string = "TD"

	// SiteImportRegulatoryDomainCL captures enum value "CL"
	SiteImportRegulatoryDomainCL string = "CL"

	// SiteImportRegulatoryDomainCN captures enum value "CN"
	SiteImportRegulatoryDomainCN string = "CN"

	// SiteImportRegulatoryDomainCX captures enum value "CX"
	SiteImportRegulatoryDomainCX string = "CX"

	// SiteImportRegulatoryDomainCC captures enum value "CC"
	SiteImportRegulatoryDomainCC string = "CC"

	// SiteImportRegulatoryDomainCO captures enum value "CO"
	SiteImportRegulatoryDomainCO string = "CO"

	// SiteImportRegulatoryDomainKM captures enum value "KM"
	SiteImportRegulatoryDomainKM string = "KM"

	// SiteImportRegulatoryDomainCK captures enum value "CK"
	SiteImportRegulatoryDomainCK string = "CK"

	// SiteImportRegulatoryDomainCR captures enum value "CR"
	SiteImportRegulatoryDomainCR string = "CR"

	// SiteImportRegulatoryDomainCI captures enum value "CI"
	SiteImportRegulatoryDomainCI string = "CI"

	// SiteImportRegulatoryDomainHR captures enum value "HR"
	SiteImportRegulatoryDomainHR string = "HR"

	// SiteImportRegulatoryDomainCW captures enum value "CW"
	SiteImportRegulatoryDomainCW string = "CW"

	// SiteImportRegulatoryDomainCY captures enum value "CY"
	SiteImportRegulatoryDomainCY string = "CY"

	// SiteImportRegulatoryDomainCZ captures enum value "CZ"
	SiteImportRegulatoryDomainCZ string = "CZ"

	// SiteImportRegulatoryDomainCD captures enum value "CD"
	SiteImportRegulatoryDomainCD string = "CD"

	// SiteImportRegulatoryDomainDK captures enum value "DK"
	SiteImportRegulatoryDomainDK string = "DK"

	// SiteImportRegulatoryDomainDJ captures enum value "DJ"
	SiteImportRegulatoryDomainDJ string = "DJ"

	// SiteImportRegulatoryDomainDM captures enum value "DM"
	SiteImportRegulatoryDomainDM string = "DM"

	// SiteImportRegulatoryDomainDO captures enum value "DO"
	SiteImportRegulatoryDomainDO string = "DO"

	// SiteImportRegulatoryDomainEC captures enum value "EC"
	SiteImportRegulatoryDomainEC string = "EC"

	// SiteImportRegulatoryDomainEG captures enum value "EG"
	SiteImportRegulatoryDomainEG string = "EG"

	// SiteImportRegulatoryDomainSV captures enum value "SV"
	SiteImportRegulatoryDomainSV string = "SV"

	// SiteImportRegulatoryDomainGQ captures enum value "GQ"
	SiteImportRegulatoryDomainGQ string = "GQ"

	// SiteImportRegulatoryDomainER captures enum value "ER"
	SiteImportRegulatoryDomainER string = "ER"

	// SiteImportRegulatoryDomainEE captures enum value "EE"
	SiteImportRegulatoryDomainEE string = "EE"

	// SiteImportRegulatoryDomainSZ captures enum value "SZ"
	SiteImportRegulatoryDomainSZ string = "SZ"

	// SiteImportRegulatoryDomainET captures enum value "ET"
	SiteImportRegulatoryDomainET string = "ET"

	// SiteImportRegulatoryDomainFK captures enum value "FK"
	SiteImportRegulatoryDomainFK string = "FK"

	// SiteImportRegulatoryDomainFO captures enum value "FO"
	SiteImportRegulatoryDomainFO string = "FO"

	// SiteImportRegulatoryDomainFJ captures enum value "FJ"
	SiteImportRegulatoryDomainFJ string = "FJ"

	// SiteImportRegulatoryDomainFI captures enum value "FI"
	SiteImportRegulatoryDomainFI string = "FI"

	// SiteImportRegulatoryDomainFR captures enum value "FR"
	SiteImportRegulatoryDomainFR string = "FR"

	// SiteImportRegulatoryDomainGF captures enum value "GF"
	SiteImportRegulatoryDomainGF string = "GF"

	// SiteImportRegulatoryDomainPF captures enum value "PF"
	SiteImportRegulatoryDomainPF string = "PF"

	// SiteImportRegulatoryDomainTF captures enum value "TF"
	SiteImportRegulatoryDomainTF string = "TF"

	// SiteImportRegulatoryDomainGA captures enum value "GA"
	SiteImportRegulatoryDomainGA string = "GA"

	// SiteImportRegulatoryDomainGM captures enum value "GM"
	SiteImportRegulatoryDomainGM string = "GM"

	// SiteImportRegulatoryDomainGE captures enum value "GE"
	SiteImportRegulatoryDomainGE string = "GE"

	// SiteImportRegulatoryDomainDE captures enum value "DE"
	SiteImportRegulatoryDomainDE string = "DE"

	// SiteImportRegulatoryDomainGH captures enum value "GH"
	SiteImportRegulatoryDomainGH string = "GH"

	// SiteImportRegulatoryDomainGI captures enum value "GI"
	SiteImportRegulatoryDomainGI string = "GI"

	// SiteImportRegulatoryDomainGR captures enum value "GR"
	SiteImportRegulatoryDomainGR string = "GR"

	// SiteImportRegulatoryDomainGL captures enum value "GL"
	SiteImportRegulatoryDomainGL string = "GL"

	// SiteImportRegulatoryDomainGD captures enum value "GD"
	SiteImportRegulatoryDomainGD string = "GD"

	// SiteImportRegulatoryDomainGP captures enum value "GP"
	SiteImportRegulatoryDomainGP string = "GP"

	// SiteImportRegulatoryDomainGU captures enum value "GU"
	SiteImportRegulatoryDomainGU string = "GU"

	// SiteImportRegulatoryDomainGT captures enum value "GT"
	SiteImportRegulatoryDomainGT string = "GT"

	// SiteImportRegulatoryDomainGG captures enum value "GG"
	SiteImportRegulatoryDomainGG string = "GG"

	// SiteImportRegulatoryDomainGN captures enum value "GN"
	SiteImportRegulatoryDomainGN string = "GN"

	// SiteImportRegulatoryDomainGW captures enum value "GW"
	SiteImportRegulatoryDomainGW string = "GW"

	// SiteImportRegulatoryDomainGY captures enum value "GY"
	SiteImportRegulatoryDomainGY string = "GY"

	// SiteImportRegulatoryDomainHT captures enum value "HT"
	SiteImportRegulatoryDomainHT string = "HT"

	// SiteImportRegulatoryDomainHM captures enum value "HM"
	SiteImportRegulatoryDomainHM string = "HM"

	// SiteImportRegulatoryDomainHN captures enum value "HN"
	SiteImportRegulatoryDomainHN string = "HN"

	// SiteImportRegulatoryDomainHK captures enum value "HK"
	SiteImportRegulatoryDomainHK string = "HK"

	// SiteImportRegulatoryDomainHU captures enum value "HU"
	SiteImportRegulatoryDomainHU string = "HU"

	// SiteImportRegulatoryDomainIS captures enum value "IS"
	SiteImportRegulatoryDomainIS string = "IS"

	// SiteImportRegulatoryDomainIN captures enum value "IN"
	SiteImportRegulatoryDomainIN string = "IN"

	// SiteImportRegulatoryDomainID captures enum value "ID"
	SiteImportRegulatoryDomainID string = "ID"

	// SiteImportRegulatoryDomainIQ captures enum value "IQ"
	SiteImportRegulatoryDomainIQ string = "IQ"

	// SiteImportRegulatoryDomainIE captures enum value "IE"
	SiteImportRegulatoryDomainIE string = "IE"

	// SiteImportRegulatoryDomainIM captures enum value "IM"
	SiteImportRegulatoryDomainIM string = "IM"

	// SiteImportRegulatoryDomainIL captures enum value "IL"
	SiteImportRegulatoryDomainIL string = "IL"

	// SiteImportRegulatoryDomainIT captures enum value "IT"
	SiteImportRegulatoryDomainIT string = "IT"

	// SiteImportRegulatoryDomainJM captures enum value "JM"
	SiteImportRegulatoryDomainJM string = "JM"

	// SiteImportRegulatoryDomainJP captures enum value "JP"
	SiteImportRegulatoryDomainJP string = "JP"

	// SiteImportRegulatoryDomainJE captures enum value "JE"
	SiteImportRegulatoryDomainJE string = "JE"

	// SiteImportRegulatoryDomainJO captures enum value "JO"
	SiteImportRegulatoryDomainJO string = "JO"

	// SiteImportRegulatoryDomainKZ captures enum value "KZ"
	SiteImportRegulatoryDomainKZ string = "KZ"

	// SiteImportRegulatoryDomainKE captures enum value "KE"
	SiteImportRegulatoryDomainKE string = "KE"

	// SiteImportRegulatoryDomainKI captures enum value "KI"
	SiteImportRegulatoryDomainKI string = "KI"

	// SiteImportRegulatoryDomainKW captures enum value "KW"
	SiteImportRegulatoryDomainKW string = "KW"

	// SiteImportRegulatoryDomainKG captures enum value "KG"
	SiteImportRegulatoryDomainKG string = "KG"

	// SiteImportRegulatoryDomainLA captures enum value "LA"
	SiteImportRegulatoryDomainLA string = "LA"

	// SiteImportRegulatoryDomainLV captures enum value "LV"
	SiteImportRegulatoryDomainLV string = "LV"

	// SiteImportRegulatoryDomainLB captures enum value "LB"
	SiteImportRegulatoryDomainLB string = "LB"

	// SiteImportRegulatoryDomainLS captures enum value "LS"
	SiteImportRegulatoryDomainLS string = "LS"

	// SiteImportRegulatoryDomainLR captures enum value "LR"
	SiteImportRegulatoryDomainLR string = "LR"

	// SiteImportRegulatoryDomainLY captures enum value "LY"
	SiteImportRegulatoryDomainLY string = "LY"

	// SiteImportRegulatoryDomainLI captures enum value "LI"
	SiteImportRegulatoryDomainLI string = "LI"

	// SiteImportRegulatoryDomainLT captures enum value "LT"
	SiteImportRegulatoryDomainLT string = "LT"

	// SiteImportRegulatoryDomainLU captures enum value "LU"
	SiteImportRegulatoryDomainLU string = "LU"

	// SiteImportRegulatoryDomainMO captures enum value "MO"
	SiteImportRegulatoryDomainMO string = "MO"

	// SiteImportRegulatoryDomainMK captures enum value "MK"
	SiteImportRegulatoryDomainMK string = "MK"

	// SiteImportRegulatoryDomainMG captures enum value "MG"
	SiteImportRegulatoryDomainMG string = "MG"

	// SiteImportRegulatoryDomainMW captures enum value "MW"
	SiteImportRegulatoryDomainMW string = "MW"

	// SiteImportRegulatoryDomainMY captures enum value "MY"
	SiteImportRegulatoryDomainMY string = "MY"

	// SiteImportRegulatoryDomainMV captures enum value "MV"
	SiteImportRegulatoryDomainMV string = "MV"

	// SiteImportRegulatoryDomainML captures enum value "ML"
	SiteImportRegulatoryDomainML string = "ML"

	// SiteImportRegulatoryDomainMT captures enum value "MT"
	SiteImportRegulatoryDomainMT string = "MT"

	// SiteImportRegulatoryDomainMH captures enum value "MH"
	SiteImportRegulatoryDomainMH string = "MH"

	// SiteImportRegulatoryDomainMQ captures enum value "MQ"
	SiteImportRegulatoryDomainMQ string = "MQ"

	// SiteImportRegulatoryDomainMR captures enum value "MR"
	SiteImportRegulatoryDomainMR string = "MR"

	// SiteImportRegulatoryDomainMU captures enum value "MU"
	SiteImportRegulatoryDomainMU string = "MU"

	// SiteImportRegulatoryDomainYT captures enum value "YT"
	SiteImportRegulatoryDomainYT string = "YT"

	// SiteImportRegulatoryDomainMX captures enum value "MX"
	SiteImportRegulatoryDomainMX string = "MX"

	// SiteImportRegulatoryDomainFM captures enum value "FM"
	SiteImportRegulatoryDomainFM string = "FM"

	// SiteImportRegulatoryDomainMD captures enum value "MD"
	SiteImportRegulatoryDomainMD string = "MD"

	// SiteImportRegulatoryDomainMC captures enum value "MC"
	SiteImportRegulatoryDomainMC string = "MC"

	// SiteImportRegulatoryDomainMN captures enum value "MN"
	SiteImportRegulatoryDomainMN string = "MN"

	// SiteImportRegulatoryDomainME captures enum value "ME"
	SiteImportRegulatoryDomainME string = "ME"

	// SiteImportRegulatoryDomainMS captures enum value "MS"
	SiteImportRegulatoryDomainMS string = "MS"

	// SiteImportRegulatoryDomainMA captures enum value "MA"
	SiteImportRegulatoryDomainMA string = "MA"

	// SiteImportRegulatoryDomainMZ captures enum value "MZ"
	SiteImportRegulatoryDomainMZ string = "MZ"

	// SiteImportRegulatoryDomainMM captures enum value "MM"
	SiteImportRegulatoryDomainMM string = "MM"

	// SiteImportRegulatoryDomainNA captures enum value "NA"
	SiteImportRegulatoryDomainNA string = "NA"

	// SiteImportRegulatoryDomainNR captures enum value "NR"
	SiteImportRegulatoryDomainNR string = "NR"

	// SiteImportRegulatoryDomainNP captures enum value "NP"
	SiteImportRegulatoryDomainNP string = "NP"

	// SiteImportRegulatoryDomainNL captures enum value "NL"
	SiteImportRegulatoryDomainNL string = "NL"

	// SiteImportRegulatoryDomainNC captures enum value "NC"
	SiteImportRegulatoryDomainNC string = "NC"

	// SiteImportRegulatoryDomainNZ captures enum value "NZ"
	SiteImportRegulatoryDomainNZ string = "NZ"

	// SiteImportRegulatoryDomainNI captures enum value "NI"
	SiteImportRegulatoryDomainNI string = "NI"

	// SiteImportRegulatoryDomainNE captures enum value "NE"
	SiteImportRegulatoryDomainNE string = "NE"

	// SiteImportRegulatoryDomainNG captures enum value "NG"
	SiteImportRegulatoryDomainNG string = "NG"

	// SiteImportRegulatoryDomainNU captures enum value "NU"
	SiteImportRegulatoryDomainNU string = "NU"

	// SiteImportRegulatoryDomainNF captures enum value "NF"
	SiteImportRegulatoryDomainNF string = "NF"

	// SiteImportRegulatoryDomainMP captures enum value "MP"
	SiteImportRegulatoryDomainMP string = "MP"

	// SiteImportRegulatoryDomainNO captures enum value "NO"
	SiteImportRegulatoryDomainNO string = "NO"

	// SiteImportRegulatoryDomainOM captures enum value "OM"
	SiteImportRegulatoryDomainOM string = "OM"

	// SiteImportRegulatoryDomainPK captures enum value "PK"
	SiteImportRegulatoryDomainPK string = "PK"

	// SiteImportRegulatoryDomainPW captures enum value "PW"
	SiteImportRegulatoryDomainPW string = "PW"

	// SiteImportRegulatoryDomainPA captures enum value "PA"
	SiteImportRegulatoryDomainPA string = "PA"

	// SiteImportRegulatoryDomainPG captures enum value "PG"
	SiteImportRegulatoryDomainPG string = "PG"

	// SiteImportRegulatoryDomainPY captures enum value "PY"
	SiteImportRegulatoryDomainPY string = "PY"

	// SiteImportRegulatoryDomainPE captures enum value "PE"
	SiteImportRegulatoryDomainPE string = "PE"

	// SiteImportRegulatoryDomainPH captures enum value "PH"
	SiteImportRegulatoryDomainPH string = "PH"

	// SiteImportRegulatoryDomainPN captures enum value "PN"
	SiteImportRegulatoryDomainPN string = "PN"

	// SiteImportRegulatoryDomainPL captures enum value "PL"
	SiteImportRegulatoryDomainPL string = "PL"

	// SiteImportRegulatoryDomainPT captures enum value "PT"
	SiteImportRegulatoryDomainPT string = "PT"

	// SiteImportRegulatoryDomainPR captures enum value "PR"
	SiteImportRegulatoryDomainPR string = "PR"

	// SiteImportRegulatoryDomainQA captures enum value "QA"
	SiteImportRegulatoryDomainQA string = "QA"

	// SiteImportRegulatoryDomainKR captures enum value "KR"
	SiteImportRegulatoryDomainKR string = "KR"

	// SiteImportRegulatoryDomainRS captures enum value "RS"
	SiteImportRegulatoryDomainRS string = "RS"

	// SiteImportRegulatoryDomainSC captures enum value "SC"
	SiteImportRegulatoryDomainSC string = "SC"

	// SiteImportRegulatoryDomainCG captures enum value "CG"
	SiteImportRegulatoryDomainCG string = "CG"

	// SiteImportRegulatoryDomainRE captures enum value "RE"
	SiteImportRegulatoryDomainRE string = "RE"

	// SiteImportRegulatoryDomainRO captures enum value "RO"
	SiteImportRegulatoryDomainRO string = "RO"

	// SiteImportRegulatoryDomainRU captures enum value "RU"
	SiteImportRegulatoryDomainRU string = "RU"

	// SiteImportRegulatoryDomainRW captures enum value "RW"
	SiteImportRegulatoryDomainRW string = "RW"

	// SiteImportRegulatoryDomainBL captures enum value "BL"
	SiteImportRegulatoryDomainBL string = "BL"

	// SiteImportRegulatoryDomainSH captures enum value "SH"
	SiteImportRegulatoryDomainSH string = "SH"

	// SiteImportRegulatoryDomainKN captures enum value "KN"
	SiteImportRegulatoryDomainKN string = "KN"

	// SiteImportRegulatoryDomainLC captures enum value "LC"
	SiteImportRegulatoryDomainLC string = "LC"

	// SiteImportRegulatoryDomainMF captures enum value "MF"
	SiteImportRegulatoryDomainMF string = "MF"

	// SiteImportRegulatoryDomainPM captures enum value "PM"
	SiteImportRegulatoryDomainPM string = "PM"

	// SiteImportRegulatoryDomainVC captures enum value "VC"
	SiteImportRegulatoryDomainVC string = "VC"

	// SiteImportRegulatoryDomainWS captures enum value "WS"
	SiteImportRegulatoryDomainWS string = "WS"

	// SiteImportRegulatoryDomainSM captures enum value "SM"
	SiteImportRegulatoryDomainSM string = "SM"

	// SiteImportRegulatoryDomainST captures enum value "ST"
	SiteImportRegulatoryDomainST string = "ST"

	// SiteImportRegulatoryDomainSA captures enum value "SA"
	SiteImportRegulatoryDomainSA string = "SA"

	// SiteImportRegulatoryDomainSN captures enum value "SN"
	SiteImportRegulatoryDomainSN string = "SN"

	// SiteImportRegulatoryDomainSL captures enum value "SL"
	SiteImportRegulatoryDomainSL string = "SL"

	// SiteImportRegulatoryDomainSG captures enum value "SG"
	SiteImportRegulatoryDomainSG string = "SG"

	// SiteImportRegulatoryDomainSX captures enum value "SX"
	SiteImportRegulatoryDomainSX string = "SX"

	// SiteImportRegulatoryDomainSK captures enum value "SK"
	SiteImportRegulatoryDomainSK string = "SK"

	// SiteImportRegulatoryDomainSI captures enum value "SI"
	SiteImportRegulatoryDomainSI string = "SI"

	// SiteImportRegulatoryDomainSB captures enum value "SB"
	SiteImportRegulatoryDomainSB string = "SB"

	// SiteImportRegulatoryDomainSO captures enum value "SO"
	SiteImportRegulatoryDomainSO string = "SO"

	// SiteImportRegulatoryDomainZA captures enum value "ZA"
	SiteImportRegulatoryDomainZA string = "ZA"

	// SiteImportRegulatoryDomainGS captures enum value "GS"
	SiteImportRegulatoryDomainGS string = "GS"

	// SiteImportRegulatoryDomainSS captures enum value "SS"
	SiteImportRegulatoryDomainSS string = "SS"

	// SiteImportRegulatoryDomainES captures enum value "ES"
	SiteImportRegulatoryDomainES string = "ES"

	// SiteImportRegulatoryDomainLK captures enum value "LK"
	SiteImportRegulatoryDomainLK string = "LK"

	// SiteImportRegulatoryDomainPS captures enum value "PS"
	SiteImportRegulatoryDomainPS string = "PS"

	// SiteImportRegulatoryDomainSR captures enum value "SR"
	SiteImportRegulatoryDomainSR string = "SR"

	// SiteImportRegulatoryDomainSJ captures enum value "SJ"
	SiteImportRegulatoryDomainSJ string = "SJ"

	// SiteImportRegulatoryDomainSE captures enum value "SE"
	SiteImportRegulatoryDomainSE string = "SE"

	// SiteImportRegulatoryDomainCH captures enum value "CH"
	SiteImportRegulatoryDomainCH string = "CH"

	// SiteImportRegulatoryDomainTW captures enum value "TW"
	SiteImportRegulatoryDomainTW string = "TW"

	// SiteImportRegulatoryDomainTJ captures enum value "TJ"
	SiteImportRegulatoryDomainTJ string = "TJ"

	// SiteImportRegulatoryDomainTZ captures enum value "TZ"
	SiteImportRegulatoryDomainTZ string = "TZ"

	// SiteImportRegulatoryDomainTH captures enum value "TH"
	SiteImportRegulatoryDomainTH string = "TH"

	// SiteImportRegulatoryDomainTL captures enum value "TL"
	SiteImportRegulatoryDomainTL string = "TL"

	// SiteImportRegulatoryDomainTG captures enum value "TG"
	SiteImportRegulatoryDomainTG string = "TG"

	// SiteImportRegulatoryDomainTK captures enum value "TK"
	SiteImportRegulatoryDomainTK string = "TK"

	// SiteImportRegulatoryDomainTO captures enum value "TO"
	SiteImportRegulatoryDomainTO string = "TO"

	// SiteImportRegulatoryDomainTT captures enum value "TT"
	SiteImportRegulatoryDomainTT string = "TT"

	// SiteImportRegulatoryDomainTN captures enum value "TN"
	SiteImportRegulatoryDomainTN string = "TN"

	// SiteImportRegulatoryDomainTR captures enum value "TR"
	SiteImportRegulatoryDomainTR string = "TR"

	// SiteImportRegulatoryDomainTM captures enum value "TM"
	SiteImportRegulatoryDomainTM string = "TM"

	// SiteImportRegulatoryDomainTC captures enum value "TC"
	SiteImportRegulatoryDomainTC string = "TC"

	// SiteImportRegulatoryDomainTV captures enum value "TV"
	SiteImportRegulatoryDomainTV string = "TV"

	// SiteImportRegulatoryDomainUG captures enum value "UG"
	SiteImportRegulatoryDomainUG string = "UG"

	// SiteImportRegulatoryDomainUA captures enum value "UA"
	SiteImportRegulatoryDomainUA string = "UA"

	// SiteImportRegulatoryDomainAE captures enum value "AE"
	SiteImportRegulatoryDomainAE string = "AE"

	// SiteImportRegulatoryDomainGB captures enum value "GB"
	SiteImportRegulatoryDomainGB string = "GB"

	// SiteImportRegulatoryDomainUS captures enum value "US"
	SiteImportRegulatoryDomainUS string = "US"

	// SiteImportRegulatoryDomainUY captures enum value "UY"
	SiteImportRegulatoryDomainUY string = "UY"

	// SiteImportRegulatoryDomainUZ captures enum value "UZ"
	SiteImportRegulatoryDomainUZ string = "UZ"

	// SiteImportRegulatoryDomainVU captures enum value "VU"
	SiteImportRegulatoryDomainVU string = "VU"

	// SiteImportRegulatoryDomainVA captures enum value "VA"
	SiteImportRegulatoryDomainVA string = "VA"

	// SiteImportRegulatoryDomainVE captures enum value "VE"
	SiteImportRegulatoryDomainVE string = "VE"

	// SiteImportRegulatoryDomainVN captures enum value "VN"
	SiteImportRegulatoryDomainVN string = "VN"

	// SiteImportRegulatoryDomainVG captures enum value "VG"
	SiteImportRegulatoryDomainVG string = "VG"

	// SiteImportRegulatoryDomainVI captures enum value "VI"
	SiteImportRegulatoryDomainVI string = "VI"

	// SiteImportRegulatoryDomainWF captures enum value "WF"
	SiteImportRegulatoryDomainWF string = "WF"

	// SiteImportRegulatoryDomainEH captures enum value "EH"
	SiteImportRegulatoryDomainEH string = "EH"

	// SiteImportRegulatoryDomainYE captures enum value "YE"
	SiteImportRegulatoryDomainYE string = "YE"

	// SiteImportRegulatoryDomainZM captures enum value "ZM"
	SiteImportRegulatoryDomainZM string = "ZM"

	// SiteImportRegulatoryDomainZW captures enum value "ZW"
	SiteImportRegulatoryDomainZW string = "ZW"
)

// prop value enum
func (m *SiteImport) validateRegulatoryDomainEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, siteImportTypeRegulatoryDomainPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SiteImport) validateRegulatoryDomain(formats strfmt.Registry) error {

	if swag.IsZero(m.RegulatoryDomain) { // not required
		return nil
	}

	// value enum
	if err := m.validateRegulatoryDomainEnum("regulatoryDomain", "body", m.RegulatoryDomain); err != nil {
		return err
	}

	return nil
}

var siteImportTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["site","client"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		siteImportTypeTypePropEnum = append(siteImportTypeTypePropEnum, v)
	}
}

const (

	// SiteImportTypeSite captures enum value "site"
	SiteImportTypeSite string = "site"

	// SiteImportTypeClient captures enum value "client"
	SiteImportTypeClient string = "client"
)

// prop value enum
func (m *SiteImport) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, siteImportTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SiteImport) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SiteImport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SiteImport) UnmarshalBinary(b []byte) error {
	var res SiteImport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
