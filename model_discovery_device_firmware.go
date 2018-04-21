/*
 * UNMS API
 *
 * You can authorize calls with x-auth-token by clicking the button on the right, to not have to fill in authorization token for each request. Entered value in authorization token for a request, will be replaced by the one entered in authorization. To get this token use /user/login request.
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package unms

// Firmware version overview
type DiscoveryDeviceFirmware struct {
	// Current FW version
	Current string `json:"current,omitempty"`

	// Latest available FW version
	Latest string `json:"latest,omitempty"`

	// True if version >= minimumVersion
	Compatible bool `json:"compatible,omitempty"`
}