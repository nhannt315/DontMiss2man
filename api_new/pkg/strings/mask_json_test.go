package strings

import (
	"testing"
)

func TestMaskJSON(t *testing.T) {

	srcJSON := `
 		{
			"token": "fadsfsaf",
			"user_id": "3",
 			"email": "test@gmail.com",
		}
    `
	expectedJSON := `
 		{
			"token": "******",
			"user_id": "3",
 			"email": "******",
		}
    `
	regexs := GenerateRegexs("token", "email")

	result := MaskJSON(srcJSON, regexs)

	if result != expectedJSON {
		t.Errorf("MaskJSON() got: %v, want %v", result, expectedJSON)
	}
}
