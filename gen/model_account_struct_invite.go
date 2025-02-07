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

// AccountStructInvite - 招待情報
type AccountStructInvite struct {

	// 招待通し番号
	InviteID int32 `json:"inviteID,omitempty"`

	// 招待コード(shortuuid)
	Code string `json:"code,omitempty"`

	// 招待した人数の累計(誰を招待したかは表示されない)
	InvitedCount int32 `json:"invitedCount,omitempty"`
}
