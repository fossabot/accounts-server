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

type PaginationStruct struct {

	// ページネーションタイトル(表示用)
	Title string `json:"title"`

	// ページネーション種別(表示用)
	Type string `json:"type"`

	// ヒット総数
	Count int32 `json:"count"`

	// 現在のページ
	Current int32 `json:"current"`

	// ページ数
	Pages int32 `json:"pages"`

	// ページ毎の取得数
	PerPage int32 `json:"perPage"`
}
