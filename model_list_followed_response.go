/*
 * Twitter API v2
 *
 * Twitter API v2 available endpoints
 *
 * API version: 2.43
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ListFollowedResponse struct {
	Data *ListFollowedResponseData `json:"data,omitempty"`
	Errors []Problem `json:"errors,omitempty"`
}
