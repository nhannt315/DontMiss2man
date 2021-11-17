package jwk_test

import (
	"testing"

	"github.com/nhannt315/real_estate_api/pkg/jwk"
)

func TestHelper_JSONWebKeySetJSON(t *testing.T) {

	h, err := jwk.NewHelper("fullback", "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlDV2dJQkFBS0JnSElDcUk3MFNBaGEzMlJna0pxUWZNRmdoUlIxdUFXUEdGYnpsdWFxMmx0MFIreXp5RTB4CmFyRTFSMXRqY2dYbUZHbzV3enFUK0o2amVPcFNRNVV0TmxGSllib2pGTHZ1VWtaT1d3WkNJWFFyaHpUelp6Qm4KdzZldlA5UFd5OW4vZWZOLzR1TlA5K29yM0U2bUd6Q1R6Q1FvSzNYZXZrd1ZKTzJSQmlrTm9kUVJBZ01CQUFFQwpnWUJ3Nm9rZTdIbmJMSVZMMlFmZXpYT0I0cUpyWFJ2aEJaUkpxZit1ZlZDK2V6QXhFdTR5NWRxUm4vOElXRG01CjllSUtReU5FaFFFYnZUUW9mMGxJR2Z0TXU3dzU5S3Vhdzd3NHRVSFgrVTBmTkpPNnZTc0FuOUk0bFJXUlBwRTQKUW94ZmQ5VG9xMVJnVjAzVnVrNkhHQ3VrcXdEanhiMUhuMmhXV0swd040NExBUUpCQUxTeXR3dzFDWXdCZG40ZApieHlReUNpbkFCZWVieTBzSnBySkVKODdIS0dycllka1dOVVZMd0JWZVBGQ3daZDcrRlhoVjhqTFlwR0lCckMxCkhLWGl1RWtDUVFDaGhZdXRDb1A5S2E5ak41Y3lpMnFjdnJ0cy96MFhudDYrN3N0WFhFWFV6WTR4ZGVqT0Nhc0UKT0R2YlNuMWc3RDNYMmpURy9DYlJUNHRheWZRYlpZMkpBa0JIeEZjNUxQTnV2TUlBRXRhbmhNVktpQkZjVUJ1ZAoyRlE0MDdTYldWSGswQmVxbDJ0RXJoWXR6c1NySmJWWDRlL2V4QkltZU1qY1BpZFNGWXljUWZDcEFrQU0yNVduCjI3dHdEcjV3ZG45cHZhRnFBdURtcDFiVXA0Znh5UTZVMExxYVd4YWpwMUExL3AwSGcvWjMyWEVyb3dLMVNTQ0UKYXBRb2UxMkxoNklRQUVDSkFrQUt3VURJV3BRUUdoNlYwNzhiSS9oYlFuaTV5K2Q5NU93ZUVLU0M4aFB5L1ZJSQpqSzdvWEJLRlhBWm9ENnlaWlFNWFZzWU1DMW1GQU5FK2swbzFCTTlQCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0t")
	if err != nil {
		t.Errorf("Generate helper for testing, err: %v", err)
		return
	}

	expectedJson := `{"keys":[{"use":"sig","kty":"RSA","kid":"fullback","alg":"RS256","n":"cgKojvRICFrfZGCQmpB8wWCFFHW4BY8YVvOW5qraW3RH7LPITTFqsTVHW2NyBeYUajnDOpP4nqN46lJDlS02UUlhuiMUu-5SRk5bBkIhdCuHNPNnMGfDp68_09bL2f9583_i40_36ivcTqYbMJPMJCgrdd6-TBUk7ZEGKQ2h1BE","e":"AQAB"}]}`

	result, err := h.JSONWebKeySetJSON()
	if err != nil {
		t.Errorf("Generate json web key set, err: %v", err)
		return
	}

	if expectedJson != string(result) {
		t.Errorf("JSONWebKeySetJSON() got =\n %v, want \n%v", string(result), expectedJson)
	}

}
