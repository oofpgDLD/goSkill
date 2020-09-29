package string

import "testing"

func TestIsPalindrome(t *testing.T) {
	ret := isPalindrome("ab_a")
	t.Log(ret)
}
