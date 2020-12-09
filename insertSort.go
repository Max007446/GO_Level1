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
	for scn.Scan() {
		num, er := strconv.ParseInt(scn.Text(), 10, 64)
		if er != nil {
			panic(er)
		}
		inputNums = append(inputNums, num)
	}

	for i := 0; i < len(inputNums); i++ {
		newElem := inputNums[i] //в NewElem сохранем первый элемент
		for j := i - 1; j >= 0 && inputNums[j] > newElem; j-- {
			// все элементы больше newElem сдвигаем враво на 1
			inputNums[j+1] = inputNums[j]
			j -= 1
			inputNums[j+1] = newElem // на освобожденное место пишем NewElem
		}
	}
	fmt.Println(inputNums)
}
