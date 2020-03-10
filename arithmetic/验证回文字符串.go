package arithmetic

import (
	"bytes"
)

/**
* @ Description:
* @Author:
* @Date: 2020/2/24 14:00
 */
func IsPalindrome(s string) bool {
	var buf bytes.Buffer
	for _, v := range s {
		if (v >= 'a' && v <= 'z') || (v >= '0' && v <= '9') {
			buf.WriteRune(v)
		} else if v >= 'A' && v <= 'Z' {
			buf.WriteRune(v - ('A' - 'a'))
		}
	}
	var (
		i    = 0
		j    = buf.Len() - 1
		temp = buf.String()
	)
	for i < j {
		if temp[i] == temp[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}
