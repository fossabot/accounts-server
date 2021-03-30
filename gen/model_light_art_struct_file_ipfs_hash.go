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

// LightArtStructFileIpfsHash - IPFSのハッシュ情報
type LightArtStructFileIpfsHash struct {

	// サムネイルハッシュ
	Thumb string `json:"thumb,omitempty"`

	// Raw画像ハッシュ
	Orig string `json:"orig,omitempty"`
}
