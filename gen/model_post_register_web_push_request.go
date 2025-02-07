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

type PostRegisterWebPushRequest struct {

	// WebPushのPOST先エンドポイント(ユーザー毎に異なる)
	Endpoint string `json:"endpoint"`

	// ブラウザ公開鍵
	P256dh string `json:"p256dh"`

	// WebPushの通知送信認証キー
	Auth string `json:"auth"`

	// 通知クライアント名
	Name string `json:"name"`

	// 通知レベル 1:緊急時のみ 5:タグ絵師通知のみ 9:すべて
	Level int32 `json:"level"`
}
