package arithmetic

/**
* @ Description:
* @Author: licongfu@ringle.com
* @Date: 2020/2/24 11:37
 */
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var temp = make(map[rune]int)
	for _, v := range s {
		temp[v]++
	}
	for _, v := range t {
		if num, ok := temp[v]; ok && num > 0 {
			temp[v]--
		} else {
			return false
		}
	}
	return true
}
