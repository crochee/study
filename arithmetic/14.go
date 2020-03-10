package arithmetic

/**
* @ Description:
* @Author:
* @Date: 2020/2/24 18:14
 */
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}
	for j := 0; j < len(strs[0]); {
		j++
		for i := 1; i < len(strs); i++ {
			if j > len(strs[i]) {
				return strs[0][:len(strs[i])]
			}
			if strs[0][:j] != strs[i][:j] {
				return strs[0][:j-1]
			}
		}
	}
	return strs[0]
}
