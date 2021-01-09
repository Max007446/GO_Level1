package main

import (
	"testing"
)

//InsertSortT метод вставками
func InsertSortT(inputNums []int64) []int64 {
	for i := 0; i < len(inputNums); i++ {
		newElem := inputNums[i] //в NewElem сохранем первый элемент
		for j := i - 1; j >= 0 && inputNums[j] > newElem; j-- {
			// все элементы больше newElem сдвигаем враво на 1
			inputNums[j+1] = inputNums[j]
			j--
			inputNums[j+1] = newElem // на освобожденное место пишем NewElem
		}
	}
	return inputNums
}

func BechmarkInsertSortT(b *testing.B) {
	res := []int64{10, 25, 5, 12, 1, 125, 0, 55, 1000, 59885, 500}
	for i := 0; i < b.N; i++ {
		res = InsertSortT(res)
	}
}
