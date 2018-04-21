/*
 * UNMS API
 *
 * You can authorize calls with x-auth-token by clicking the button on the right, to not have to fill in authorization token for each request. Entered value in authorization token for a request, will be replaced by the one entered in authorization. To get this token use /user/login request.
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package unms

type InterfaceConfigSwitchPorts struct {
	Interface_ *InterfaceIdentification `json:"interface,omitempty"`

	Enabled bool `json:"enabled,omitempty"`

	// Value range is 1 - 4087. Pvid (port vlan id), when a packet comes to eth0 from outside network with no \"vlan tag\", it will be tagged with pvid and treated as a packet with vlan tag.
	Pvid float32 `json:"pvid,omitempty"`

	Vid []float32 `json:"vid,omitempty"`
}