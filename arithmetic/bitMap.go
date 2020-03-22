package arithmetic

/*
* @
* @Author:
* @Date: 2020/3/22 12:31
 */

type BitMap struct {
	capacity int
	words    []uint64
}

func NewBitMap(cap int) *BitMap {
	return &BitMap{
		capacity: cap,
		words:    make([]uint64, 0),
	}
}

func (bm *BitMap) GetBit(id int) bool {
	if id > bm.capacity-1 {
		panic("超出BitMap有效范围")
	}
	// 定位是那个word
	wordIndex, bit := bm.GetIndex(id)
	return bm.words[wordIndex]&(1<<bit) != 0
}

func (bm *BitMap) GetIndex(id int) (int, int) {
	// 右移6位，相当于除以64 id>>6
	return id / 64, id % 64
}

func (bm *BitMap) SetBit(id int) {
	if id > bm.capacity-1 {
		panic("超出BitMap有效范围")
	}
	wordIndex, bit := bm.GetIndex(id)
	for wordIndex >= len(bm.words) {
		bm.words = append(bm.words, 0)
	}
	bm.words[wordIndex] |= 1 << bit
}
