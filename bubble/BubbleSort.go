package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputNums := []int64{}
	scn := bufio.NewScanner(os.Stdin)
	fmt.Printf("Введите 999 для окончания ввода данных в слайс: \n")
	for scn.Scan() {
		num, er := strconv.ParseInt(scn.Text(), 10, 64)

		if num == 999 {
			break
		}
		if er != nil {
			panic(er)
		}
		inputNums = append(inputNums, num)
	}
	t := true //флаг сортировки
	for t {
		t = false
		//перебор элементов слайса
		for j := 0; j < len(inputNums)-1; j++ {
			if inputNums[j] > inputNums[j+1] { // сравниваем текущий элемент с предыдущим
				inputNums[j], inputNums[j+1] = inputNums[j+1], inputNums[j] //обмениваемся элементами при выполнении условия
				t = true
			}
		}
	}

	fmt.Println("Отсортированный слайс: \n", inputNums)
}
