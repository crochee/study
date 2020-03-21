package arithmetic

import "testing"

/*
* @
* @Author:
* @Date: 2020/3/21 12:08
 */

func TestHeapSort(t *testing.T) {
	arr := []int{1, 3, 2, 6, 5, 7, 8, 9, 10, 0}
	temp := []int{0, 1, 2, 3, 5, 6, 7, 8, 9, 10}
	HeapSort(arr)
	for i, v := range arr {
		if v != temp[i] {
			t.Errorf("get %v ,unexpected %v", temp[i], v)
		}
	}
}

func BenchmarkHeapSort(b *testing.B) {
	a := []int{5, 8, 1, 6, 3, 9, 2, 1, 7}
	b.ResetTimer()
	for i := 1; i < b.N; i++ {
		HeapSort(a)
	}
}
