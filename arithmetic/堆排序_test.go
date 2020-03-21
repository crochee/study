package arithmetic

import "testing"

/*
* @
* @Author:
* @Date: 2020/3/21 12:08
 */
func BenchmarkHeapSort(b *testing.B) {
	a := []int{5, 8, 1, 6, 3, 9, 2, 1, 7}
	b.ResetTimer()
	for i := 1; i < b.N; i++ {
		HeapSort(a)
	}
}
