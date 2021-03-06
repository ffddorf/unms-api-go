/*
 * UNMS API
 *
 * You can authorize calls with x-auth-token by clicking the button on the right, to not have to fill in authorization token for each request. Entered value in authorization token for a request, will be replaced by the one entered in authorization. To get this token use /user/login request.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package unms

type SiteDescription struct {
	Note string `json:"note,omitempty"`

	Address string `json:"address,omitempty"`

	Location *SiteDescriptionLocation `json:"location,omitempty"`

	// Elevation on the surface of the earth
	Elevation float32 `json:"elevation,omitempty"`

	// Height of the structure
	Height float32 `json:"height,omitempty"`

	Contact *SiteDescriptionContact `json:"contact,omitempty"`

	// Endpoints are sites without endpoints
	Endpoints []SiteIdentification `json:"endpoints,omitempty"`
}
