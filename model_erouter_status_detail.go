/*
 * UNMS API
 *
 * You can authorize calls with x-auth-token by clicking the button on the right, to not have to fill in authorization token for each request. Entered value in authorization token for a request, will be replaced by the one entered in authorization. To get this token use /user/login request.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package unms

// Read-only erouter detail
type ErouterStatusDetail struct {
	Identification *DeviceIdentification `json:"identification,omitempty"`

	Overview *DeviceOverview `json:"overview,omitempty"`

	Firmware *DeviceFirmware `json:"firmware,omitempty"`

	Upgrade *DeviceUpgradeStatus `json:"upgrade,omitempty"`

	Gateway string `json:"gateway,omitempty"`
}
