package main

import (
	"fmt"

	"stu/arithmetic"
)

/**
* @ Description:
* @Author: licongfu@ringle.com
* @Date: 2020/2/21 19:50
 */

func main() {
	presetMap := []string{
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		"X . X X X X X X X X X X X X X X X X X X X X X X X X X",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		"X X X X X X X X X X X X X X X X X X X X X X X X . X X",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
	}
	thismap := arithmetic.NewThisMap(presetMap)
	sr := arithmetic.NewSearchRoad(0, 0, 18, 10, thismap)
	if err := sr.FindoutRoad(); err != nil {
		fmt.Println(err)
	}
	sr.Print()
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
