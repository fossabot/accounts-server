/*
 * UsagiBooru Accounts API
 *
 * アカウント関連API
 *
 * API version: 2.0
 * Contact: dsgamer777@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gen

type GetTimelineFollowingResponse struct {

	Pagination PaginationStruct `json:"pagination,omitempty"`

	Follows []LightArtistStruct `json:"follows,omitempty"`
}
