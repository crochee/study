package arithmetic

/**
* @ Description:
* @Author:
* @Date: 2020/2/22 20:10
 */
/*func PlusOne(digits []int) []int {
	var temp = make([]int, len(digits)+1)
	for i, v := range digits {
		temp[i+1] = v
	}
	per(temp)
	if temp[0] != 1 {
		return temp[1:]
	}
	return temp
}

func per(digits []int) []int {
	if digits[len(digits)-1] == 9 {
		digits[len(digits)-1] = 0
		if len(digits)< 3 {
			digits[0] = 1
			return digits
		}
		per(digits)
	} else {
		digits[len(digits)-1]++
	}
	return digits
}*/
func PlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
			if i == 0 {
				digits = append([]int{1}, digits...)
			}
		} else {
			digits[i]++
			break
		}
	}
	return digits
}
