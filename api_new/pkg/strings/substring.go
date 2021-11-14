package strings

// SubString returns sub string of given str, starting at index startIdx,
// and having maximum length by maxLength, without invalid character error.
// (In case of splitting in middle of Japanese string, for example).
// maxLength will be counted per rune, not byte.
func SubString(str string, startIdx, maxLength uint) string {
	if str == "" {
		return str
	}

	runes := []rune(str)
	if startIdx == 0 && len(runes) <= int(maxLength) { // avoid creating new string
		return str
	}

	numRunes := uint(len(runes))
	if startIdx > numRunes {
		return ""
	}
	endIdx := startIdx + maxLength
	if endIdx > numRunes {
		endIdx = numRunes
	}

	return string(runes[startIdx:endIdx])
}
