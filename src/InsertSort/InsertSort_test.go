package main

import (
	"testing"
)

//InsertSortT метод вставками
func InsertSortT(inputNums []int64) []int64 {
	for i := 1; i < len(inputNums); i++ {
		newElem := inputNums[i] //в NewElem сохранем первый элемент
		j := i
		for j > 0 && inputNums[j-1] > newElem {
			// все элементы больше newElem сдвигаем враво на 1
			inputNums[j] = inputNums[j-1]
			j--
		}
		inputNums[j] = newElem // на освобожденное место пишем NewElem
	}
	return inputNums
}

func BechmarkInsertSortT(b *testing.B) {
	res := []int64{10, 25, 5, 12, 1, 125, 0, 55, 1000, 59885, 500}
	for i := 0; i < b.N; i++ {
		res = InsertSortT(res)
	}
}
