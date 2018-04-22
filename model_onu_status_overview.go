/*
 * UNMS API
 *
 * You can authorize calls with x-auth-token by clicking the button on the right, to not have to fill in authorization token for each request. Entered value in authorization token for a request, will be replaced by the one entered in authorization. To get this token use /user/login request.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package unms

// Read-only ONU detail
type OnuStatusOverview struct {
	Identification *DeviceIdentification `json:"identification,omitempty"`

	Overview *DeviceOverview `json:"overview,omitempty"`

	Firmware *DeviceFirmware `json:"firmware,omitempty"`

	// nullable string
	ParentId string `json:"parentId,omitempty"`

	// Logical switch which determines if it makes sense to display device statistics to user
	CanDisplayStatistics bool `json:"canDisplayStatistics,omitempty"`

	Olt *OnuStatusOverviewOlt `json:"olt,omitempty"`
}
