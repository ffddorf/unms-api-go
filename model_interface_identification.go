/*
 * UNMS API
 *
 * You can authorize calls with x-auth-token by clicking the button on the right, to not have to fill in authorization token for each request. Entered value in authorization token for a request, will be replaced by the one entered in authorization. To get this token use /user/login request.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package unms

// read-only interface identification
type InterfaceIdentification struct {
	// physical port position
	Position float32 `json:"position,omitempty"`

	Type_ string `json:"type,omitempty"`

	// interface name
	Name string `json:"name,omitempty"`

	// computed display name from name and description
	DisplayName string `json:"displayName,omitempty"`

	// nullable string
	Description string `json:"description,omitempty"`

	Mac string `json:"mac,omitempty"`
}
