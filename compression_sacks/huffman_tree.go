package compression_sacks

import (
	"container/heap"
)

type HuffmanNode struct {
	Char  rune
	Freq  int
	Left  *HuffmanNode
	Right *HuffmanNode
}

// min-heap giá trị nhỏ nhất luôn ở đầu heap
// slice chứa các con trỏ struct HuffmanNode
// sử dụng con trỏ để dễ thay đổi giá trị
type HuffmanHeap []*HuffmanNode

// các interface (phương thức) của kiểu dữ liệu HuffmanHeap
// Receiver Types - kiểu dữ liệu mà interface gắn vào.
// 2 loại value receiver func (h HuffmanHeap) và pointer receiver func (h *HuffmanHeap)
func (h HuffmanHeap) Len() int           { return len(h) }
func (h HuffmanHeap) Less(i, j int) bool { return h[i].Freq < h[j].Freq }
func (h HuffmanHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *HuffmanHeap) Push(x interface{}) {
	// x.(*HuffmanNode) - type assertion
	// x là interface.
	// chuyển đổi x sang kiểu HuffmanNode
	*h = append(*h, x.(*HuffmanNode))
}

func (h *HuffmanHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func BuildHuffmanTree(freqMap map[rune]int) *HuffmanNode {
	h := &HuffmanHeap{}

	// sử dụng package "container/heap"
	// Heap sẽ tự động sắp xếp các nút này sao cho nút có tần suất thấp nhất luôn ở đầu.
	heap.Init(h)

	// range - for loop dùng cho map (hoặc slice)
	for char, freq := range freqMap {
		heap.Push(h, &HuffmanNode{Char: char, Freq: freq})
	}

	// for is Go's "while"
	for h.Len() > 1 {
		left := heap.Pop(h).(*HuffmanNode)
		right := heap.Pop(h).(*HuffmanNode)

		internalNode := &HuffmanNode{
			Freq:  left.Freq + right.Freq,
			Left:  left,
			Right: right,
		}

		heap.Push(h, internalNode)
	}

	return heap.Pop(h).(*HuffmanNode)
}
