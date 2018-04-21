/*
 * UNMS API
 *
 * You can authorize calls with x-auth-token by clicking the button on the right, to not have to fill in authorization token for each request. Entered value in authorization token for a request, will be replaced by the one entered in authorization. To get this token use /user/login request.
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package unms

// Start and end of the interval
type NmsStatisticsInterval struct {
	// Start time as milliseconds elapsed between 1 January 1970 00:00:00 UTC and the given time
	Start float32 `json:"start,omitempty"`

	// End time as milliseconds elapsed between 1 January 1970 00:00:00 UTC and the given time
	End float32 `json:"end,omitempty"`
}
