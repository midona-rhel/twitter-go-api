/*
 * Twitter API v2
 *
 * Twitter API v2 available endpoints
 *
 * API version: 2.43
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// You have been disconnected for operational reasons.
type OperationalDisconnectProblem struct {
	Type_ string `json:"type"`
	Title string `json:"title"`
	Detail string `json:"detail,omitempty"`
	Status int32 `json:"status,omitempty"`
	DisconnectType string `json:"disconnect_type,omitempty"`
}
