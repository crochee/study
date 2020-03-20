package main

import (
	"fmt"
	"golangStu/arithmetic"
)

/**
* @ Description:
* @Author: licongfu@ringle.com
* @Date: 2020/2/21 19:50
 */

func main() {
	a := []int{5, 8, 1, 6, 3, 9, 2, 1, 7}
	fmt.Println(a)
	arithmetic.QuickSort(a, 0, len(a)-1)
	fmt.Println(a)
	/*// 初始化
	gtk.Init(&os.Args)
	// 用户初始化
	// 1)创建窗口
	win := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	// 2)设置属性
	win.SetTitle("first")
	win.SetSizeRequest(480, 320)
	win.Show()
	// 主事件循环
	gtk.Main()*/
	/*var (
		index  uint64
		bucket = make(map[uint64]int)
	)
	for i := 15000000; i < 15000000+15000000; i++ {
		index = murmur3.Sum64([]byte(fmt.Sprint(i))) % 10
		bucket[index]++
	}
	fmt.Printf("%p\n", &index)
	fmt.Println(bucket, "\nthis:", unsafe.Pointer(&index))*/
	// nums1 := " u34 ui "
	// fmt.Println(arithmetic.IsValidBST(nums1))
}
