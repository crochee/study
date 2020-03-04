package arithmetic

import (
	"math/rand"
	"time"
)

/**
* @ Description:洗牌算法
* @Author: licongfu@ringle.com
* @Date: 2020/3/4 19:33
 */
func Shuffle(arr []int) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	var (
		n         = len(arr)
		randIndex int
	)
	for i := n; i > 0; i-- {
		randIndex = r.Intn(i)
		arr[i-1], arr[randIndex] = arr[randIndex], arr[i-1]
	}
}
