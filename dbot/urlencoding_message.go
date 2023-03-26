package dbot

import (
	"net/url"
	"regexp"
	"strings"
)

func interleaveArrays(arr1 []string, arr2 []string) string {
	var result strings.Builder
	len1, len2 := len(arr1), len(arr2)
	maxLen := len1
	if len2 > maxLen {
		maxLen = len2
	}
	for i := 0; i < maxLen; i++ {
		if i < len1 {
			result.WriteString(arr1[i])
		}
		if i < len2 {
			result.WriteString(arr2[i])
		}
	}
	return result.String()
}

func urlencoding_message(message string) string {
	re := regexp.MustCompile(`\[CQ:[^\]]*\]`)
	cqcodes := re.FindAllString(message, -1)
	others := re.Split(message, -1)
	urlencodedOthers := make([]string, len(others))
	for i, v := range others {
		urlencodedOthers[i] = url.QueryEscape(v)
	}
	urlencodedMessage := interleaveArrays(urlencodedOthers, cqcodes)
	return urlencodedMessage
}
