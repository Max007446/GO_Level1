package main

import (
	"testing"
)

// BubbleSort Сортировка пузырьком
func BubbleSort(inputNums []int64) []int64 {
	sliseSortUnSort := true //флаг сортировки
	sortSlise := make([]int64, len(inputNums))
	copy(sortSlise, inputNums)
	for sliseSortUnSort {
		sliseSortUnSort = false
		//перебор элементов слайса
		for j := 0; j < len(sortSlise)-1; j++ {
			if sortSlise[j] > sortSlise[j+1] { // сравниваем текущий элемент с предыдущим
				sortSlise[j], sortSlise[j+1] = sortSlise[j+1], sortSlise[j] //обмениваемся элементами при выполнении условия
				sliseSortUnSort = true
			}
		}
	}
	return sortSlise
}

func BechmarkbubbleSort(b *testing.B) {

	res := []int64{10, 25, 5, 12, 1, 125, 0, 55, 1000, 59885, 500}
	for i := 0; i < b.N; i++ {
		res = BubbleSort(res)
	}
}
