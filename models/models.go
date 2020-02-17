// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// Models models
// swagger:model models
type Models []string

var modelsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UF-Nano","UF-Loco","UF-Wifi","UF-Instant","UF-OLT","UF-OLT4","ER-X","ER-X-SFP","ERLite-3","ERPoe-5","ERPro-8","ER-8","ER-8-XG","ER-4","ER-6P","ER-12","ER-12P","ER-10X","EP-R8","EP-R6","EP-S16","ES-12F","ES-16-150W","ES-24-250W","ES-24-500W","ES-24-Lite","ES-48-500W","ES-48-750W","ES-48-Lite","ES-8-150W","ES-16-XG","ES-10XP","ES-10X","ES-18X","ES-26X","EP-54V-150W","EP-24V-72W","EP-54V-72W","TSW-PoE","TSW-PoE PRO","ACB-AC","ACB-ISP","ACB-LOCO","AF11FX","AF24","AF24HD","AF2X","AF3X","AF4X","AF5","AF5U","AF5X","AF-5XHD","AF-LTU","LTU-LITE","AF-LTU5","LTU-Rocket","AF60","AFLTULR","GBE-LR","GBE","R2N","R2T","R5N","R6N","R36-GPS","RM3-GPS","R2N-GPS","R5N-GPS","R9N-GPS","R5T-GPS","RM3","R36","R9N","N2N","N5N","N6N","NS3","N36","N9N","N9S","LM2","LM5","B2N","B2T","B5N","B5T","BAC","AG2","AG2-HP","AG5","AG5-HP","p2N","p5N","M25","P2B-400","P5B-300","P5B-300-ISO","P5B-400","P5B-400-ISO","P5B-620","LB5-120","LB5","N5B","N5B-16","N5B-19","N5B-300","N5B-400","N5B-Client","N2B","N2B-13","N2B-400","PAP","LAP-HP","LAP","AGW","AGW-LR","AGW-Pro","AGW-Installer","PB5","PB3","P36","PBM10","NB5","NB2","NB3","B36","NB9","SM5","WM5","IS-M5","Loco5AC","NS-5AC","R5AC-PTMP","R5AC-PTP","R5AC-Lite","R5AC-PRISM","R2AC-Prism","R2AC-Gen2","RP-5AC-Gen2","NBE-2AC-13","NBE-5AC-16","NBE-5AC-19","NBE-5AC-Gen2","PBE-5AC-300","PBE-5AC-300-ISO","PBE-5AC-400","PBE-5AC-400-ISO","PBE-5AC-500","PBE-5AC-500-ISO","PBE-5AC-620","PBE-5AC-620-ISO","PBE-2AC-400","PBE-2AC-400-ISO","PBE-5AC-X-Gen2","PBE-5AC-Gen2","PBE-5AC-ISO-Gen2","PBE-5AC-400-ISO-Gen2","LBE-5AC-16-120","LAP-120","LBE-5AC-23","LBE-5AC-Gen2","LBE-5AC-LR","LAP-GPS","IS-5AC","PS-5AC","SolarSwitch","SolarPoint","BulletAC-IP67","B-DB-AC","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		modelsItemsEnum = append(modelsItemsEnum, v)
	}
}

func (m *Models) validateModelsItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, modelsItemsEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this models
func (m Models) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {

		// value enum
		if err := m.validateModelsItemsEnum(strconv.Itoa(i), "body", m[i]); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}