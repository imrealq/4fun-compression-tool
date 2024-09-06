package compression_sacks

import (
	"container/heap"
	"reflect"
	"testing"
)

func TestBuildHuffmanTree(t *testing.T) {
	tests := []struct {
		name    string
		freqMap map[rune]int
		want    *HuffmanNode
	}{
		{
			name:    "Single character",
			freqMap: map[rune]int{'a': 1},
			want:    &HuffmanNode{Char: 'a', Freq: 1},
		},
		{
			name: "Two characters",
			freqMap: map[rune]int{
				'a': 1,
				'b': 2,
			},
			want: &HuffmanNode{
				Freq:  3,
				Left:  &HuffmanNode{Char: 'a', Freq: 1},
				Right: &HuffmanNode{Char: 'b', Freq: 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BuildHuffmanTree(tt.freqMap)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildHuffmanTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateHuffmanCodes(t *testing.T) {
	tests := []struct {
		name string
		root *HuffmanNode
		want map[rune]string
	}{
		{
			name: "Simple tree",
			root: &HuffmanNode{
				Freq:  3,
				Left:  &HuffmanNode{Char: 'a', Freq: 1},
				Right: &HuffmanNode{Char: 'b', Freq: 2},
			},
			want: map[rune]string{
				'a': "0",
				'b': "1",
			},
		},
		{
			name: "Complex tree",
			root: &HuffmanNode{
				Freq: 6,
				Left: &HuffmanNode{
					Freq:  3,
					Left:  &HuffmanNode{Char: 'a', Freq: 1},
					Right: &HuffmanNode{Char: 'b', Freq: 2},
				},
				Right: &HuffmanNode{
					Freq:  3,
					Left:  &HuffmanNode{Char: 'c', Freq: 1},
					Right: &HuffmanNode{Char: 'd', Freq: 2},
				},
			},
			want: map[rune]string{
				'a': "00",
				'b': "01",
				'c': "10",
				'd': "11",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateHuffmanCodes(tt.root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateHuffmanCodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHuffmanHeap(t *testing.T) {
	h := &HuffmanHeap{}

	nodes := []*HuffmanNode{
		{Char: 'a', Freq: 5},
		{Char: 'b', Freq: 2},
		{Char: 'c', Freq: 8},
	}

	for _, node := range nodes {
		heap.Push(h, node)
	}

	if h.Len() != 3 {
		t.Errorf("HuffmanHeap length = %d, want 3", h.Len())
	}

	expected := []rune{'b', 'a', 'c'}
	for i := 0; i < 3; i++ {
		node := heap.Pop(h).(*HuffmanNode)
		if node.Char != expected[i] {
			t.Errorf("Pop() returned char %c, want %c", node.Char, expected[i])
		}
	}

	if h.Len() != 0 {
		t.Errorf("HuffmanHeap length after popping all elements = %d, want 0", h.Len())
	}
}
