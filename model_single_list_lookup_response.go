/*
 * Twitter API v2
 *
 * Twitter API v2 available endpoints
 *
 * API version: 2.43
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type SingleListLookupResponse struct {
	Data *List `json:"data,omitempty"`
	Includes *SingleListLookupResponseIncludes `json:"includes,omitempty"`
	Errors []Problem `json:"errors,omitempty"`
}
