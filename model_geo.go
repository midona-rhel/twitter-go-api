/*
 * Twitter API v2
 *
 * Twitter API v2 available endpoints
 *
 * API version: 2.43
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Geo struct {
	Type_ string `json:"type"`
	Bbox []float64 `json:"bbox"`
	Geometry *Point `json:"geometry,omitempty"`
	Properties *interface{} `json:"properties"`
}
