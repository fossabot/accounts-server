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

type PostRegisterLineNotifyRequest struct {

	// LineNotifyのパーソナルトークン
	Token string `json:"token"`

	// 通知クライアント名
	Name string `json:"name"`

	// 通知レベル 1:緊急時のみ 5:タグ絵師通知のみ 9:すべて
	Level int32 `json:"level"`
}
