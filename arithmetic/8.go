package arithmetic

import (
	"bytes"
	"regexp"
	"strings"
)

/**
* @ Description:
* @Author: licongfu@ringle.com
* @Date: 2020/2/24 15:09
 */
func MyAtoi(str string) int {
	var (
		out  int
		buf  bytes.Buffer
		sign = 1
	)
	for i, v := range str {
		if v >= '0' && v <= '9' {
			buf.WriteRune(v)
		} else {
			if buf.Len() == 0 {
				if (v == '-' || v == '+') && i < len(str)-1 && str[i+1] >= '0' && str[i+1] <= '9' {
					if v == '-' {
						sign = -1
					}
				} else if v == ' ' {
					continue
				} else {
					return 0
				}
			} else {
				break
			}
		}
	}
	temp := buf.String()
	for i := 0; i < buf.Len(); i++ {
		out = out*10 + sign*int(temp[i]-'0')
		if out > 1<<31-1 {
			return 1<<31 - 1
		} else if out < -1<<31 {
			return -1 << 31
		}
	}
	return out
}

func MyAtoi1(str string) int {
	temp := regexp.MustCompile(`^[\+-]?\d+`).FindString(strings.TrimSpace(str))
	var out int
	for i := 0; i < len(temp); i++ {
		if temp[0] == '-' && i != 0 {
			out = out*10 - int(temp[i]-'0')
		} else if temp[0] == '+' && i != 0 {
			out = out*10 + int(temp[i]-'0')
		} else if temp[0] != '-' && temp[0] != '+' {
			out = out*10 + int(temp[i]-'0')
		}
		if out > 1<<31-1 {
			return 1<<31 - 1
		} else if out < -1<<31 {
			return -1 << 31
		}
	}
	return out
}
