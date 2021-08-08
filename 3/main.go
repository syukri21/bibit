package refractor

import "strings"

func findFirstStringInBrackets(s string) string {
	firstIndex := strings.IndexByte(s, '(')
	if firstIndex < 0 {
		return ""
	}
	firstIndex++
	lastIndex := strings.IndexByte(s[firstIndex:], ')')
	if lastIndex < 0 {
		return ""
	}
	return s[firstIndex : firstIndex+lastIndex]
}
