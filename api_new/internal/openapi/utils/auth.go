package utils

import (
	"net/http"
	"strings"
)

// HeaderAccessTokenKey AccessTokenがセットされるHTTPヘッダーキー
const HeaderAccessTokenKey = "Authorization"

func GetAccessTokenFromHeader(req *http.Request) string {
	authorizationContent := req.Header.Get(HeaderAccessTokenKey)
	authorizations := strings.Split(authorizationContent, " ")
	if authorizations[0] != "Bearer" {
		return ""
	}
	return authorizations[len(authorizations)-1]
}
