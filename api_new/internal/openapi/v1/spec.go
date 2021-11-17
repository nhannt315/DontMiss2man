package openapi

import (
	"net/http"
	"strings"
)

// BasePath OpenAPI URLのベースパス. OpenAPIファイルのServersに記載されているパス
// これはSpec.Serversのpath部分を動的にとっても良いが、
// Serversが複数ある場合どれをとるのかがという問題もあるので
// (Sever.Descriptionに環境を書いておいてくなども考えられるがそれも危険だし、pdevやstgのURLなどはyamlに定義しないだろうし)
// ハードコードでいいのではないか。
var BasePath = "/api/v1"

// IsV1Request V1 のリクエストかどうかを判定する
func IsV1Request(req *http.Request) bool {
	path := req.URL.Path
	return strings.HasPrefix(path, BasePath)
}
