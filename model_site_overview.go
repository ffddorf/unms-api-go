/*
 * UNMS API
 *
 * You can authorize calls with x-auth-token by clicking the button on the right, to not have to fill in authorization token for each request. Entered value in authorization token for a request, will be replaced by the one entered in authorization. To get this token use /user/login request.
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package unms

// Gallery is created with site. There is only one gallery per site. Site devices is possible to load with /devices resource.
type SiteOverview struct {
	Identification *SiteIdentification `json:"identification,omitempty"`

	Description *SiteDescription `json:"description,omitempty"`

	Notifications *SiteOverviewNotifications `json:"notifications,omitempty"`
}