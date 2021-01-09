package main

import (
	"sortsNew/src/BubbleSort"
	"sortsNew/src/InsertSort"
	"testing"
)

func BechmarkInsertSort(b *testing.B) {
	res := []int64{10, 25, 5, 12, 1, 125, 0, 55, 1000, 59885, 500}
	for i := 0; i < b.N; i++ {
		res = sortsNew.InsertSortT(res)
	}
}
func BechmarkBubbleSort(b *testing.B) {

	res := []int64{10, 25, 5, 12, 1, 125, 0, 55, 1000, 59885, 500}
	for i := 0; i < b.N; i++ {
		res = sortsNew.BubbleSort(res)
	}
}
