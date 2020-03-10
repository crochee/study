package arithmetic

/**
* @ Description:
* @Author:
* @Date: 2020/2/23 21:30
 */
func ReverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}
