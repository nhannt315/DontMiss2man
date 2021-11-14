package strings

import (
	"fmt"
	"regexp"
)

type MaskJSONField struct {
	key            string
	valueToReplace string
	regexKey       *regexp.Regexp
}

// MaskJSON 指定するFieldに対してMaskしたStringを返却する
func MaskJSON(src string, regexs []*MaskJSONField) string {
	if len(regexs) == 0 {
		return src
	}

	resultStr := src
	for _, regex := range regexs {
		resultStr = regex.regexKey.ReplaceAllString(resultStr, regex.valueToReplace)
	}

	return resultStr
}

// GenerateRegexs マスクする用のRegexを生成する
func GenerateRegexs(keys ...string) []*MaskJSONField {
	regexs := make([]*MaskJSONField, 0)
	for _, key := range keys {
		regex := fmt.Sprintf(`"%s":\s*"[^"]+?([^\/"]+)"`, key)
		regexs = append(regexs, &MaskJSONField{
			key:            key,
			valueToReplace: fmt.Sprintf(`"%s": "%s"`, key, MaskedStringText),
			regexKey:       regexp.MustCompile(regex),
		})
	}
	return regexs
}
