package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputNums := []int64{}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Введите 999 для окончания ввода данных в слайс: \n")
	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)

		if num == 999 {
			break
		}
		if err != nil {
			panic(err)
		}
		inputNums = append(inputNums, num)
	}
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
	fmt.Println("Входной слайс: \n", inputNums)
	fmt.Println("Отсортированный слайс: \n", sortSlise)
}
