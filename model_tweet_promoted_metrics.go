/*
 * Twitter API v2
 *
 * Twitter API v2 available endpoints
 *
 * API version: 2.43
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Promoted nonpublic engagement metrics for the Tweet at the time of the request.
type TweetPromotedMetrics struct {
	// Number of times this Tweet has been viewed.
	ImpressionCount int32 `json:"impression_count,omitempty"`
	// Number of times this Tweet has been liked.
	LikeCount int32 `json:"like_count,omitempty"`
	// Number of times this Tweet has been replied to.
	ReplyCount int32 `json:"reply_count,omitempty"`
	// Number of times this Tweet has been Retweeted.
	RetweetCount int32 `json:"retweet_count,omitempty"`
}
