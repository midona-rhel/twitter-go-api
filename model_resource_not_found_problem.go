/*
 * Twitter API v2
 *
 * Twitter API v2 available endpoints
 *
 * API version: 2.43
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// A problem that indicates that a given Tweet, User, etc. does not exist.
type ResourceNotFoundProblem struct {
	Type_     string `json:"type"`
	Title     string `json:"title"`
	Detail    string `json:"detail,omitempty"`
	Status    int32  `json:"status,omitempty"`
	Parameter string `json:"parameter"`
	// Value will match the schema of the field.
	Value        *any   `json:"value"`
	ResourceId   string `json:"resource_id"`
	ResourceType string `json:"resource_type"`
}
