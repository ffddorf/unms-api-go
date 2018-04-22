/*
 * UNMS API
 *
 * You can authorize calls with x-auth-token by clicking the button on the right, to not have to fill in authorization token for each request. Entered value in authorization token for a request, will be replaced by the one entered in authorization. To get this token use /user/login request.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package unms

// If status is Unconnected these flags indicate the problem
type DiscoveryDeviceFlags struct {
	CanAuthenticate bool `json:"canAuthenticate,omitempty"`

	CanConnect bool `json:"canConnect,omitempty"`

	UnsupportedDevice bool `json:"unsupportedDevice,omitempty"`

	MissingCredentials bool `json:"missingCredentials,omitempty"`

	UnsupportedFirmware bool `json:"unsupportedFirmware,omitempty"`

	Error_ bool `json:"error,omitempty"`
}
