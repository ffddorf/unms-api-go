/*
 * UNMS API
 *
 * You can authorize calls with x-auth-token by clicking the button on the right, to not have to fill in authorization token for each request. Entered value in authorization token for a request, will be replaced by the one entered in authorization. To get this token use /user/login request.
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package unms

type InterfaceConfigBridgePorts struct {
	Enabled bool `json:"enabled,omitempty"`

	Interface_ *InterfaceIdentification `json:"interface,omitempty"`
}
