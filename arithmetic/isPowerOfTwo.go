package arithmetic

/*
* @
* @Author:
* @Date: 2020/3/19 21:11
 */

func countOne(x int) int {
	var count int
	for x > 0 {
		x &= x - 1 // x&（x-1） 从最高位开始将所有1置0
		count++
	}
	return count
}

func isPowerOfTwo(x uintptr) bool { // true时,x必为2的指数，即二进制数必须是1开头，其他全为0的数
	return x&(x-1) == 0
}
