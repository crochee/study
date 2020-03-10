package arithmetic

/**
* @ Description:
* @Author:
* @Date: 2020/2/23 23:10
 */
func FirstUniqChar(s string) int {
	var arr = [26]int{}
	for _, v := range s {
		arr[v-'a']++
	}
	for i, v := range s {
		if arr[v-'a'] == 1 {
			return i
		}
	}
	return -1
}
