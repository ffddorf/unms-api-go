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

// VisibleBy visible by
// swagger:model visibleBy
type VisibleBy struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// model
	// Required: true
	// Enum: [UF-Nano UF-Loco UF-Wifi UF-Instant UF-OLT UF-OLT4 ER-X ER-X-SFP ERLite-3 ERPoe-5 ERPro-8 ER-8 ER-8-XG ER-4 ER-6P ER-12 ER-12P ER-10X EP-R8 EP-R6 EP-S16 ES-12F ES-16-150W ES-24-250W ES-24-500W ES-24-Lite ES-48-500W ES-48-750W ES-48-Lite ES-8-150W ES-16-XG ES-10XP ES-10X ES-18X ES-26X EP-54V-150W EP-24V-72W EP-54V-72W TSW-PoE TSW-PoE PRO ACB-AC ACB-ISP ACB-LOCO AF11FX AF24 AF24HD AF2X AF3X AF4X AF5 AF5U AF5X AF-5XHD AF-LTU LTU-LITE AF-LTU5 LTU-Rocket AF60 AFLTULR GBE-LR GBE R2N R2T R5N R6N R36-GPS RM3-GPS R2N-GPS R5N-GPS R9N-GPS R5T-GPS RM3 R36 R9N N2N N5N N6N NS3 N36 N9N N9S LM2 LM5 B2N B2T B5N B5T BAC AG2 AG2-HP AG5 AG5-HP p2N p5N M25 P2B-400 P5B-300 P5B-300-ISO P5B-400 P5B-400-ISO P5B-620 LB5-120 LB5 N5B N5B-16 N5B-19 N5B-300 N5B-400 N5B-Client N2B N2B-13 N2B-400 PAP LAP-HP LAP AGW AGW-LR AGW-Pro AGW-Installer PB5 PB3 P36 PBM10 NB5 NB2 NB3 B36 NB9 SM5 WM5 IS-M5 Loco5AC NS-5AC R5AC-PTMP R5AC-PTP R5AC-Lite R5AC-PRISM R2AC-Prism R2AC-Gen2 RP-5AC-Gen2 NBE-2AC-13 NBE-5AC-16 NBE-5AC-19 NBE-5AC-Gen2 PBE-5AC-300 PBE-5AC-300-ISO PBE-5AC-400 PBE-5AC-400-ISO PBE-5AC-500 PBE-5AC-500-ISO PBE-5AC-620 PBE-5AC-620-ISO PBE-2AC-400 PBE-2AC-400-ISO PBE-5AC-X-Gen2 PBE-5AC-Gen2 PBE-5AC-ISO-Gen2 PBE-5AC-400-ISO-Gen2 LBE-5AC-16-120 LAP-120 LBE-5AC-23 LBE-5AC-Gen2 LBE-5AC-LR LAP-GPS IS-5AC PS-5AC SolarSwitch SolarPoint BulletAC-IP67 B-DB-AC UNKNOWN]
	Model *string `json:"model"`

	// name
	// Required: true
	Name *string `json:"name"`

	// type
	// Required: true
	// Enum: [onu olt erouter eswitch epower airCube airMax airFiber toughSwitch solarBeam blackBox]
	Type *string `json:"type"`
}

// Validate validates this visible by
func (m *VisibleBy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
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

func (m *VisibleBy) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

var visibleByTypeModelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UF-Nano","UF-Loco","UF-Wifi","UF-Instant","UF-OLT","UF-OLT4","ER-X","ER-X-SFP","ERLite-3","ERPoe-5","ERPro-8","ER-8","ER-8-XG","ER-4","ER-6P","ER-12","ER-12P","ER-10X","EP-R8","EP-R6","EP-S16","ES-12F","ES-16-150W","ES-24-250W","ES-24-500W","ES-24-Lite","ES-48-500W","ES-48-750W","ES-48-Lite","ES-8-150W","ES-16-XG","ES-10XP","ES-10X","ES-18X","ES-26X","EP-54V-150W","EP-24V-72W","EP-54V-72W","TSW-PoE","TSW-PoE PRO","ACB-AC","ACB-ISP","ACB-LOCO","AF11FX","AF24","AF24HD","AF2X","AF3X","AF4X","AF5","AF5U","AF5X","AF-5XHD","AF-LTU","LTU-LITE","AF-LTU5","LTU-Rocket","AF60","AFLTULR","GBE-LR","GBE","R2N","R2T","R5N","R6N","R36-GPS","RM3-GPS","R2N-GPS","R5N-GPS","R9N-GPS","R5T-GPS","RM3","R36","R9N","N2N","N5N","N6N","NS3","N36","N9N","N9S","LM2","LM5","B2N","B2T","B5N","B5T","BAC","AG2","AG2-HP","AG5","AG5-HP","p2N","p5N","M25","P2B-400","P5B-300","P5B-300-ISO","P5B-400","P5B-400-ISO","P5B-620","LB5-120","LB5","N5B","N5B-16","N5B-19","N5B-300","N5B-400","N5B-Client","N2B","N2B-13","N2B-400","PAP","LAP-HP","LAP","AGW","AGW-LR","AGW-Pro","AGW-Installer","PB5","PB3","P36","PBM10","NB5","NB2","NB3","B36","NB9","SM5","WM5","IS-M5","Loco5AC","NS-5AC","R5AC-PTMP","R5AC-PTP","R5AC-Lite","R5AC-PRISM","R2AC-Prism","R2AC-Gen2","RP-5AC-Gen2","NBE-2AC-13","NBE-5AC-16","NBE-5AC-19","NBE-5AC-Gen2","PBE-5AC-300","PBE-5AC-300-ISO","PBE-5AC-400","PBE-5AC-400-ISO","PBE-5AC-500","PBE-5AC-500-ISO","PBE-5AC-620","PBE-5AC-620-ISO","PBE-2AC-400","PBE-2AC-400-ISO","PBE-5AC-X-Gen2","PBE-5AC-Gen2","PBE-5AC-ISO-Gen2","PBE-5AC-400-ISO-Gen2","LBE-5AC-16-120","LAP-120","LBE-5AC-23","LBE-5AC-Gen2","LBE-5AC-LR","LAP-GPS","IS-5AC","PS-5AC","SolarSwitch","SolarPoint","BulletAC-IP67","B-DB-AC","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		visibleByTypeModelPropEnum = append(visibleByTypeModelPropEnum, v)
	}
}

const (

	// VisibleByModelUFNano captures enum value "UF-Nano"
	VisibleByModelUFNano string = "UF-Nano"

	// VisibleByModelUFLoco captures enum value "UF-Loco"
	VisibleByModelUFLoco string = "UF-Loco"

	// VisibleByModelUFWifi captures enum value "UF-Wifi"
	VisibleByModelUFWifi string = "UF-Wifi"

	// VisibleByModelUFInstant captures enum value "UF-Instant"
	VisibleByModelUFInstant string = "UF-Instant"

	// VisibleByModelUFOLT captures enum value "UF-OLT"
	VisibleByModelUFOLT string = "UF-OLT"

	// VisibleByModelUFOLT4 captures enum value "UF-OLT4"
	VisibleByModelUFOLT4 string = "UF-OLT4"

	// VisibleByModelERX captures enum value "ER-X"
	VisibleByModelERX string = "ER-X"

	// VisibleByModelERXSFP captures enum value "ER-X-SFP"
	VisibleByModelERXSFP string = "ER-X-SFP"

	// VisibleByModelERLite3 captures enum value "ERLite-3"
	VisibleByModelERLite3 string = "ERLite-3"

	// VisibleByModelERPoe5 captures enum value "ERPoe-5"
	VisibleByModelERPoe5 string = "ERPoe-5"

	// VisibleByModelERPro8 captures enum value "ERPro-8"
	VisibleByModelERPro8 string = "ERPro-8"

	// VisibleByModelER8 captures enum value "ER-8"
	VisibleByModelER8 string = "ER-8"

	// VisibleByModelER8XG captures enum value "ER-8-XG"
	VisibleByModelER8XG string = "ER-8-XG"

	// VisibleByModelER4 captures enum value "ER-4"
	VisibleByModelER4 string = "ER-4"

	// VisibleByModelER6P captures enum value "ER-6P"
	VisibleByModelER6P string = "ER-6P"

	// VisibleByModelER12 captures enum value "ER-12"
	VisibleByModelER12 string = "ER-12"

	// VisibleByModelER12P captures enum value "ER-12P"
	VisibleByModelER12P string = "ER-12P"

	// VisibleByModelER10X captures enum value "ER-10X"
	VisibleByModelER10X string = "ER-10X"

	// VisibleByModelEPR8 captures enum value "EP-R8"
	VisibleByModelEPR8 string = "EP-R8"

	// VisibleByModelEPR6 captures enum value "EP-R6"
	VisibleByModelEPR6 string = "EP-R6"

	// VisibleByModelEPS16 captures enum value "EP-S16"
	VisibleByModelEPS16 string = "EP-S16"

	// VisibleByModelES12F captures enum value "ES-12F"
	VisibleByModelES12F string = "ES-12F"

	// VisibleByModelES16150W captures enum value "ES-16-150W"
	VisibleByModelES16150W string = "ES-16-150W"

	// VisibleByModelES24250W captures enum value "ES-24-250W"
	VisibleByModelES24250W string = "ES-24-250W"

	// VisibleByModelES24500W captures enum value "ES-24-500W"
	VisibleByModelES24500W string = "ES-24-500W"

	// VisibleByModelES24Lite captures enum value "ES-24-Lite"
	VisibleByModelES24Lite string = "ES-24-Lite"

	// VisibleByModelES48500W captures enum value "ES-48-500W"
	VisibleByModelES48500W string = "ES-48-500W"

	// VisibleByModelES48750W captures enum value "ES-48-750W"
	VisibleByModelES48750W string = "ES-48-750W"

	// VisibleByModelES48Lite captures enum value "ES-48-Lite"
	VisibleByModelES48Lite string = "ES-48-Lite"

	// VisibleByModelES8150W captures enum value "ES-8-150W"
	VisibleByModelES8150W string = "ES-8-150W"

	// VisibleByModelES16XG captures enum value "ES-16-XG"
	VisibleByModelES16XG string = "ES-16-XG"

	// VisibleByModelES10XP captures enum value "ES-10XP"
	VisibleByModelES10XP string = "ES-10XP"

	// VisibleByModelES10X captures enum value "ES-10X"
	VisibleByModelES10X string = "ES-10X"

	// VisibleByModelES18X captures enum value "ES-18X"
	VisibleByModelES18X string = "ES-18X"

	// VisibleByModelES26X captures enum value "ES-26X"
	VisibleByModelES26X string = "ES-26X"

	// VisibleByModelEP54V150W captures enum value "EP-54V-150W"
	VisibleByModelEP54V150W string = "EP-54V-150W"

	// VisibleByModelEP24V72W captures enum value "EP-24V-72W"
	VisibleByModelEP24V72W string = "EP-24V-72W"

	// VisibleByModelEP54V72W captures enum value "EP-54V-72W"
	VisibleByModelEP54V72W string = "EP-54V-72W"

	// VisibleByModelTSWPoE captures enum value "TSW-PoE"
	VisibleByModelTSWPoE string = "TSW-PoE"

	// VisibleByModelTSWPoEPRO captures enum value "TSW-PoE PRO"
	VisibleByModelTSWPoEPRO string = "TSW-PoE PRO"

	// VisibleByModelACBAC captures enum value "ACB-AC"
	VisibleByModelACBAC string = "ACB-AC"

	// VisibleByModelACBISP captures enum value "ACB-ISP"
	VisibleByModelACBISP string = "ACB-ISP"

	// VisibleByModelACBLOCO captures enum value "ACB-LOCO"
	VisibleByModelACBLOCO string = "ACB-LOCO"

	// VisibleByModelAF11FX captures enum value "AF11FX"
	VisibleByModelAF11FX string = "AF11FX"

	// VisibleByModelAF24 captures enum value "AF24"
	VisibleByModelAF24 string = "AF24"

	// VisibleByModelAF24HD captures enum value "AF24HD"
	VisibleByModelAF24HD string = "AF24HD"

	// VisibleByModelAF2X captures enum value "AF2X"
	VisibleByModelAF2X string = "AF2X"

	// VisibleByModelAF3X captures enum value "AF3X"
	VisibleByModelAF3X string = "AF3X"

	// VisibleByModelAF4X captures enum value "AF4X"
	VisibleByModelAF4X string = "AF4X"

	// VisibleByModelAF5 captures enum value "AF5"
	VisibleByModelAF5 string = "AF5"

	// VisibleByModelAF5U captures enum value "AF5U"
	VisibleByModelAF5U string = "AF5U"

	// VisibleByModelAF5X captures enum value "AF5X"
	VisibleByModelAF5X string = "AF5X"

	// VisibleByModelAF5XHD captures enum value "AF-5XHD"
	VisibleByModelAF5XHD string = "AF-5XHD"

	// VisibleByModelAFLTU captures enum value "AF-LTU"
	VisibleByModelAFLTU string = "AF-LTU"

	// VisibleByModelLTULITE captures enum value "LTU-LITE"
	VisibleByModelLTULITE string = "LTU-LITE"

	// VisibleByModelAFLTU5 captures enum value "AF-LTU5"
	VisibleByModelAFLTU5 string = "AF-LTU5"

	// VisibleByModelLTURocket captures enum value "LTU-Rocket"
	VisibleByModelLTURocket string = "LTU-Rocket"

	// VisibleByModelAF60 captures enum value "AF60"
	VisibleByModelAF60 string = "AF60"

	// VisibleByModelAFLTULR captures enum value "AFLTULR"
	VisibleByModelAFLTULR string = "AFLTULR"

	// VisibleByModelGBELR captures enum value "GBE-LR"
	VisibleByModelGBELR string = "GBE-LR"

	// VisibleByModelGBE captures enum value "GBE"
	VisibleByModelGBE string = "GBE"

	// VisibleByModelR2N captures enum value "R2N"
	VisibleByModelR2N string = "R2N"

	// VisibleByModelR2T captures enum value "R2T"
	VisibleByModelR2T string = "R2T"

	// VisibleByModelR5N captures enum value "R5N"
	VisibleByModelR5N string = "R5N"

	// VisibleByModelR6N captures enum value "R6N"
	VisibleByModelR6N string = "R6N"

	// VisibleByModelR36GPS captures enum value "R36-GPS"
	VisibleByModelR36GPS string = "R36-GPS"

	// VisibleByModelRM3GPS captures enum value "RM3-GPS"
	VisibleByModelRM3GPS string = "RM3-GPS"

	// VisibleByModelR2NGPS captures enum value "R2N-GPS"
	VisibleByModelR2NGPS string = "R2N-GPS"

	// VisibleByModelR5NGPS captures enum value "R5N-GPS"
	VisibleByModelR5NGPS string = "R5N-GPS"

	// VisibleByModelR9NGPS captures enum value "R9N-GPS"
	VisibleByModelR9NGPS string = "R9N-GPS"

	// VisibleByModelR5TGPS captures enum value "R5T-GPS"
	VisibleByModelR5TGPS string = "R5T-GPS"

	// VisibleByModelRM3 captures enum value "RM3"
	VisibleByModelRM3 string = "RM3"

	// VisibleByModelR36 captures enum value "R36"
	VisibleByModelR36 string = "R36"

	// VisibleByModelR9N captures enum value "R9N"
	VisibleByModelR9N string = "R9N"

	// VisibleByModelN2N captures enum value "N2N"
	VisibleByModelN2N string = "N2N"

	// VisibleByModelN5N captures enum value "N5N"
	VisibleByModelN5N string = "N5N"

	// VisibleByModelN6N captures enum value "N6N"
	VisibleByModelN6N string = "N6N"

	// VisibleByModelNS3 captures enum value "NS3"
	VisibleByModelNS3 string = "NS3"

	// VisibleByModelN36 captures enum value "N36"
	VisibleByModelN36 string = "N36"

	// VisibleByModelN9N captures enum value "N9N"
	VisibleByModelN9N string = "N9N"

	// VisibleByModelN9S captures enum value "N9S"
	VisibleByModelN9S string = "N9S"

	// VisibleByModelLM2 captures enum value "LM2"
	VisibleByModelLM2 string = "LM2"

	// VisibleByModelLM5 captures enum value "LM5"
	VisibleByModelLM5 string = "LM5"

	// VisibleByModelB2N captures enum value "B2N"
	VisibleByModelB2N string = "B2N"

	// VisibleByModelB2T captures enum value "B2T"
	VisibleByModelB2T string = "B2T"

	// VisibleByModelB5N captures enum value "B5N"
	VisibleByModelB5N string = "B5N"

	// VisibleByModelB5T captures enum value "B5T"
	VisibleByModelB5T string = "B5T"

	// VisibleByModelBAC captures enum value "BAC"
	VisibleByModelBAC string = "BAC"

	// VisibleByModelAG2 captures enum value "AG2"
	VisibleByModelAG2 string = "AG2"

	// VisibleByModelAG2HP captures enum value "AG2-HP"
	VisibleByModelAG2HP string = "AG2-HP"

	// VisibleByModelAG5 captures enum value "AG5"
	VisibleByModelAG5 string = "AG5"

	// VisibleByModelAG5HP captures enum value "AG5-HP"
	VisibleByModelAG5HP string = "AG5-HP"

	// VisibleByModelP2N captures enum value "p2N"
	VisibleByModelP2N string = "p2N"

	// VisibleByModelP5N captures enum value "p5N"
	VisibleByModelP5N string = "p5N"

	// VisibleByModelM25 captures enum value "M25"
	VisibleByModelM25 string = "M25"

	// VisibleByModelP2B400 captures enum value "P2B-400"
	VisibleByModelP2B400 string = "P2B-400"

	// VisibleByModelP5B300 captures enum value "P5B-300"
	VisibleByModelP5B300 string = "P5B-300"

	// VisibleByModelP5B300ISO captures enum value "P5B-300-ISO"
	VisibleByModelP5B300ISO string = "P5B-300-ISO"

	// VisibleByModelP5B400 captures enum value "P5B-400"
	VisibleByModelP5B400 string = "P5B-400"

	// VisibleByModelP5B400ISO captures enum value "P5B-400-ISO"
	VisibleByModelP5B400ISO string = "P5B-400-ISO"

	// VisibleByModelP5B620 captures enum value "P5B-620"
	VisibleByModelP5B620 string = "P5B-620"

	// VisibleByModelLB5120 captures enum value "LB5-120"
	VisibleByModelLB5120 string = "LB5-120"

	// VisibleByModelLB5 captures enum value "LB5"
	VisibleByModelLB5 string = "LB5"

	// VisibleByModelN5B captures enum value "N5B"
	VisibleByModelN5B string = "N5B"

	// VisibleByModelN5B16 captures enum value "N5B-16"
	VisibleByModelN5B16 string = "N5B-16"

	// VisibleByModelN5B19 captures enum value "N5B-19"
	VisibleByModelN5B19 string = "N5B-19"

	// VisibleByModelN5B300 captures enum value "N5B-300"
	VisibleByModelN5B300 string = "N5B-300"

	// VisibleByModelN5B400 captures enum value "N5B-400"
	VisibleByModelN5B400 string = "N5B-400"

	// VisibleByModelN5BClient captures enum value "N5B-Client"
	VisibleByModelN5BClient string = "N5B-Client"

	// VisibleByModelN2B captures enum value "N2B"
	VisibleByModelN2B string = "N2B"

	// VisibleByModelN2B13 captures enum value "N2B-13"
	VisibleByModelN2B13 string = "N2B-13"

	// VisibleByModelN2B400 captures enum value "N2B-400"
	VisibleByModelN2B400 string = "N2B-400"

	// VisibleByModelPAP captures enum value "PAP"
	VisibleByModelPAP string = "PAP"

	// VisibleByModelLAPHP captures enum value "LAP-HP"
	VisibleByModelLAPHP string = "LAP-HP"

	// VisibleByModelLAP captures enum value "LAP"
	VisibleByModelLAP string = "LAP"

	// VisibleByModelAGW captures enum value "AGW"
	VisibleByModelAGW string = "AGW"

	// VisibleByModelAGWLR captures enum value "AGW-LR"
	VisibleByModelAGWLR string = "AGW-LR"

	// VisibleByModelAGWPro captures enum value "AGW-Pro"
	VisibleByModelAGWPro string = "AGW-Pro"

	// VisibleByModelAGWInstaller captures enum value "AGW-Installer"
	VisibleByModelAGWInstaller string = "AGW-Installer"

	// VisibleByModelPB5 captures enum value "PB5"
	VisibleByModelPB5 string = "PB5"

	// VisibleByModelPB3 captures enum value "PB3"
	VisibleByModelPB3 string = "PB3"

	// VisibleByModelP36 captures enum value "P36"
	VisibleByModelP36 string = "P36"

	// VisibleByModelPBM10 captures enum value "PBM10"
	VisibleByModelPBM10 string = "PBM10"

	// VisibleByModelNB5 captures enum value "NB5"
	VisibleByModelNB5 string = "NB5"

	// VisibleByModelNB2 captures enum value "NB2"
	VisibleByModelNB2 string = "NB2"

	// VisibleByModelNB3 captures enum value "NB3"
	VisibleByModelNB3 string = "NB3"

	// VisibleByModelB36 captures enum value "B36"
	VisibleByModelB36 string = "B36"

	// VisibleByModelNB9 captures enum value "NB9"
	VisibleByModelNB9 string = "NB9"

	// VisibleByModelSM5 captures enum value "SM5"
	VisibleByModelSM5 string = "SM5"

	// VisibleByModelWM5 captures enum value "WM5"
	VisibleByModelWM5 string = "WM5"

	// VisibleByModelISM5 captures enum value "IS-M5"
	VisibleByModelISM5 string = "IS-M5"

	// VisibleByModelLoco5AC captures enum value "Loco5AC"
	VisibleByModelLoco5AC string = "Loco5AC"

	// VisibleByModelNS5AC captures enum value "NS-5AC"
	VisibleByModelNS5AC string = "NS-5AC"

	// VisibleByModelR5ACPTMP captures enum value "R5AC-PTMP"
	VisibleByModelR5ACPTMP string = "R5AC-PTMP"

	// VisibleByModelR5ACPTP captures enum value "R5AC-PTP"
	VisibleByModelR5ACPTP string = "R5AC-PTP"

	// VisibleByModelR5ACLite captures enum value "R5AC-Lite"
	VisibleByModelR5ACLite string = "R5AC-Lite"

	// VisibleByModelR5ACPRISM captures enum value "R5AC-PRISM"
	VisibleByModelR5ACPRISM string = "R5AC-PRISM"

	// VisibleByModelR2ACPrism captures enum value "R2AC-Prism"
	VisibleByModelR2ACPrism string = "R2AC-Prism"

	// VisibleByModelR2ACGen2 captures enum value "R2AC-Gen2"
	VisibleByModelR2ACGen2 string = "R2AC-Gen2"

	// VisibleByModelRP5ACGen2 captures enum value "RP-5AC-Gen2"
	VisibleByModelRP5ACGen2 string = "RP-5AC-Gen2"

	// VisibleByModelNBE2AC13 captures enum value "NBE-2AC-13"
	VisibleByModelNBE2AC13 string = "NBE-2AC-13"

	// VisibleByModelNBE5AC16 captures enum value "NBE-5AC-16"
	VisibleByModelNBE5AC16 string = "NBE-5AC-16"

	// VisibleByModelNBE5AC19 captures enum value "NBE-5AC-19"
	VisibleByModelNBE5AC19 string = "NBE-5AC-19"

	// VisibleByModelNBE5ACGen2 captures enum value "NBE-5AC-Gen2"
	VisibleByModelNBE5ACGen2 string = "NBE-5AC-Gen2"

	// VisibleByModelPBE5AC300 captures enum value "PBE-5AC-300"
	VisibleByModelPBE5AC300 string = "PBE-5AC-300"

	// VisibleByModelPBE5AC300ISO captures enum value "PBE-5AC-300-ISO"
	VisibleByModelPBE5AC300ISO string = "PBE-5AC-300-ISO"

	// VisibleByModelPBE5AC400 captures enum value "PBE-5AC-400"
	VisibleByModelPBE5AC400 string = "PBE-5AC-400"

	// VisibleByModelPBE5AC400ISO captures enum value "PBE-5AC-400-ISO"
	VisibleByModelPBE5AC400ISO string = "PBE-5AC-400-ISO"

	// VisibleByModelPBE5AC500 captures enum value "PBE-5AC-500"
	VisibleByModelPBE5AC500 string = "PBE-5AC-500"

	// VisibleByModelPBE5AC500ISO captures enum value "PBE-5AC-500-ISO"
	VisibleByModelPBE5AC500ISO string = "PBE-5AC-500-ISO"

	// VisibleByModelPBE5AC620 captures enum value "PBE-5AC-620"
	VisibleByModelPBE5AC620 string = "PBE-5AC-620"

	// VisibleByModelPBE5AC620ISO captures enum value "PBE-5AC-620-ISO"
	VisibleByModelPBE5AC620ISO string = "PBE-5AC-620-ISO"

	// VisibleByModelPBE2AC400 captures enum value "PBE-2AC-400"
	VisibleByModelPBE2AC400 string = "PBE-2AC-400"

	// VisibleByModelPBE2AC400ISO captures enum value "PBE-2AC-400-ISO"
	VisibleByModelPBE2AC400ISO string = "PBE-2AC-400-ISO"

	// VisibleByModelPBE5ACXGen2 captures enum value "PBE-5AC-X-Gen2"
	VisibleByModelPBE5ACXGen2 string = "PBE-5AC-X-Gen2"

	// VisibleByModelPBE5ACGen2 captures enum value "PBE-5AC-Gen2"
	VisibleByModelPBE5ACGen2 string = "PBE-5AC-Gen2"

	// VisibleByModelPBE5ACISOGen2 captures enum value "PBE-5AC-ISO-Gen2"
	VisibleByModelPBE5ACISOGen2 string = "PBE-5AC-ISO-Gen2"

	// VisibleByModelPBE5AC400ISOGen2 captures enum value "PBE-5AC-400-ISO-Gen2"
	VisibleByModelPBE5AC400ISOGen2 string = "PBE-5AC-400-ISO-Gen2"

	// VisibleByModelLBE5AC16120 captures enum value "LBE-5AC-16-120"
	VisibleByModelLBE5AC16120 string = "LBE-5AC-16-120"

	// VisibleByModelLAP120 captures enum value "LAP-120"
	VisibleByModelLAP120 string = "LAP-120"

	// VisibleByModelLBE5AC23 captures enum value "LBE-5AC-23"
	VisibleByModelLBE5AC23 string = "LBE-5AC-23"

	// VisibleByModelLBE5ACGen2 captures enum value "LBE-5AC-Gen2"
	VisibleByModelLBE5ACGen2 string = "LBE-5AC-Gen2"

	// VisibleByModelLBE5ACLR captures enum value "LBE-5AC-LR"
	VisibleByModelLBE5ACLR string = "LBE-5AC-LR"

	// VisibleByModelLAPGPS captures enum value "LAP-GPS"
	VisibleByModelLAPGPS string = "LAP-GPS"

	// VisibleByModelIS5AC captures enum value "IS-5AC"
	VisibleByModelIS5AC string = "IS-5AC"

	// VisibleByModelPS5AC captures enum value "PS-5AC"
	VisibleByModelPS5AC string = "PS-5AC"

	// VisibleByModelSolarSwitch captures enum value "SolarSwitch"
	VisibleByModelSolarSwitch string = "SolarSwitch"

	// VisibleByModelSolarPoint captures enum value "SolarPoint"
	VisibleByModelSolarPoint string = "SolarPoint"

	// VisibleByModelBulletACIP67 captures enum value "BulletAC-IP67"
	VisibleByModelBulletACIP67 string = "BulletAC-IP67"

	// VisibleByModelBDBAC captures enum value "B-DB-AC"
	VisibleByModelBDBAC string = "B-DB-AC"

	// VisibleByModelUNKNOWN captures enum value "UNKNOWN"
	VisibleByModelUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (m *VisibleBy) validateModelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, visibleByTypeModelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VisibleBy) validateModel(formats strfmt.Registry) error {

	if err := validate.Required("model", "body", m.Model); err != nil {
		return err
	}

	// value enum
	if err := m.validateModelEnum("model", "body", *m.Model); err != nil {
		return err
	}

	return nil
}

func (m *VisibleBy) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var visibleByTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["onu","olt","erouter","eswitch","epower","airCube","airMax","airFiber","toughSwitch","solarBeam","blackBox"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		visibleByTypeTypePropEnum = append(visibleByTypeTypePropEnum, v)
	}
}

const (

	// VisibleByTypeOnu captures enum value "onu"
	VisibleByTypeOnu string = "onu"

	// VisibleByTypeOlt captures enum value "olt"
	VisibleByTypeOlt string = "olt"

	// VisibleByTypeErouter captures enum value "erouter"
	VisibleByTypeErouter string = "erouter"

	// VisibleByTypeEswitch captures enum value "eswitch"
	VisibleByTypeEswitch string = "eswitch"

	// VisibleByTypeEpower captures enum value "epower"
	VisibleByTypeEpower string = "epower"

	// VisibleByTypeAirCube captures enum value "airCube"
	VisibleByTypeAirCube string = "airCube"

	// VisibleByTypeAirMax captures enum value "airMax"
	VisibleByTypeAirMax string = "airMax"

	// VisibleByTypeAirFiber captures enum value "airFiber"
	VisibleByTypeAirFiber string = "airFiber"

	// VisibleByTypeToughSwitch captures enum value "toughSwitch"
	VisibleByTypeToughSwitch string = "toughSwitch"

	// VisibleByTypeSolarBeam captures enum value "solarBeam"
	VisibleByTypeSolarBeam string = "solarBeam"

	// VisibleByTypeBlackBox captures enum value "blackBox"
	VisibleByTypeBlackBox string = "blackBox"
)

// prop value enum
func (m *VisibleBy) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, visibleByTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VisibleBy) validateType(formats strfmt.Registry) error {

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
func (m *VisibleBy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VisibleBy) UnmarshalBinary(b []byte) error {
	var res VisibleBy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}