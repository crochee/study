package arithmetic

import "strings"

/**
* @ Description:
* @Author: licongfu@ringle.com
* @Date: 2020/2/24 16:39
 */
func StrStr(haystack string, needle string) int {
	if len(needle) == 0 || haystack == needle {
		return 0
	}
	if len(haystack) < len(needle) || (len(haystack) == len(needle) && haystack != needle) {
		return -1
	}
	for i := range haystack {
		if i+len(needle) <= len(haystack) && needle == haystack[i:i+len(needle)] {
			return i
		}
	}
	return -1
}

func StrStr1(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
