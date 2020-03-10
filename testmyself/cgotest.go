package testmyself

// #include "cgotest.go.h"
import "C"

/**
* @ Description:
* @Author: licongfu@ringle.com
* @Date: 2020/3/6 22:30
 */
func Cgotest() {
	C.puts(C.CString("hello,world"))
	C.SayHello(C.CString("hello,writer"))
}
