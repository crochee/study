package main

import (
	"fmt"
	"github.com/spaolacci/murmur3"
)

/**
* @ Description:
* @Author: licongfu@ringle.com
* @Date: 2020/2/21 19:50
 */

func main() {
	var (
		index  uint64
		bucket = make(map[uint64]int)
	)
	for i := 15000000; i < 15000000+15000000; i++ {
		index = murmur3.Sum64([]byte(fmt.Sprint(i))) % 10
		bucket[index]++
	}
	fmt.Println(bucket)
	// nums1 := " u34 ui "
	// fmt.Println(arithmetic.IsValidBST(nums1))
}
