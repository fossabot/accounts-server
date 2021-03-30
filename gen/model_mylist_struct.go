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

import (
	"time"
)

type MylistStruct struct {

	// マイリストID
	MylistID int32 `json:"mylistID,omitempty"`

	// マイリスト名
	Name string `json:"name,omitempty"`

	// マイリスト説明文
	Description string `json:"description,omitempty"`

	// マイリスト作成日時
	CreatedDate time.Time `json:"createdDate,omitempty"`

	// マイリスト更新日時
	UpdatedDate time.Time `json:"updatedDate,omitempty"`

	// 公開/非公開
	Private bool `json:"private,omitempty"`

	// イラストID一覧
	Arts []LightArtStruct `json:"arts,omitempty"`

	Owner LightAccountStruct `json:"owner,omitempty"`
}
