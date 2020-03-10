package arithmetic

import (
	"bytes"
	"strconv"
)

/**
* @ Description:
* @Author:
* @Date: 2020/2/24 17:11
 */
func CountAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	var (
		buf  bytes.Buffer
		temp = "1"
	)
	for i := 1; i < n; i++ {
		var temps = new(st)
		for j, v := range temp {
			if j == 0 { // 初始化
				temps = &st{
					count: 1,
					value: v,
					out:   "",
				}
			} else {
				if temps.value != v {
					buf.WriteString(strconv.Itoa(temps.count))
					buf.WriteByte(byte(temps.value))
					temps.out += buf.String()
					buf.Reset()
					temps.value = v
					temps.count = 1
				} else {
					temps.count++
				}
			}
		}
		buf.WriteString(strconv.Itoa(temps.count))
		buf.WriteByte(byte(temps.value))
		temps.out += buf.String()
		buf.Reset()
		temp = temps.out
	}
	return temp
}

type st struct {
	count int
	value int32
	out   string
}
