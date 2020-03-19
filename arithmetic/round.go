package arithmetic

/*
* @
* @Author:
* @Date: 2020/3/19 23:40
 */

// 将n舍入为a的倍数。a必须是2的幂
func round(n, a int) int {
	return (n + a - 1) &^ (a - 1)
}
