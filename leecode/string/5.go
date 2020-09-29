package string

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	reg := regexp.MustCompile(`[^A-Za-z0-9]+`)
	ret := reg.ReplaceAll([]byte(s), []byte(""))
	str := strings.ToLower(string(ret))

	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}