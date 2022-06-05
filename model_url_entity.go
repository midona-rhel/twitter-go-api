/*
 * Twitter API v2
 *
 * Twitter API v2 available endpoints
 *
 * API version: 2.43
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Represent the portion of text recognized as a URL, and its start and end position within the text.
type UrlEntity struct {
	Url string `json:"url"`
	ExpandedUrl string `json:"expanded_url,omitempty"`
	// The URL as displayed in the Twitter client.
	DisplayUrl string `json:"display_url,omitempty"`
	// Fully resolved url
	UnwoundUrl string `json:"unwound_url,omitempty"`
	Status int32 `json:"status,omitempty"`
	// Title of the page the URL points to.
	Title string `json:"title,omitempty"`
	// Description of the URL landing page.
	Description string `json:"description,omitempty"`
	Images []UrlImage `json:"images,omitempty"`
	MediaKey string `json:"media_key,omitempty"`
	// Index (zero-based) at which position this entity starts.  The index is inclusive.
	Start int32 `json:"start"`
	// Index (zero-based) at which position this entity ends.  The index is exclusive.
	End int32 `json:"end"`
}
