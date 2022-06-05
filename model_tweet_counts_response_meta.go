/*
 * Twitter API v2
 *
 * Twitter API v2 available endpoints
 *
 * API version: 2.43
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type TweetCountsResponseMeta struct {
	// This value is used to get the next 'page' of results by providing it to the next_token parameter.
	NextToken string `json:"next_token,omitempty"`
	// Sum of search query count results
	TotalTweetCount int32 `json:"total_tweet_count,omitempty"`
}
