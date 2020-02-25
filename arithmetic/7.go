package arithmetic

import "math"

/**
* @ Description:
* @Author: licongfu@ringle.com
* @Date: 2020/2/23 21:40
 */
func Reverse(x int) int {
	var y = 0
	for x != 0 {
		y = y*10 + x%10
		if math.Abs(float64(y)) > 1<<31-1 {
			return 0
		}
		x = x / 10
	}
	return y
}
