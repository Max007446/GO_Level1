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
	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		inputNums = append(inputNums, num)
	}
	var sliceSortNum = make([]int64, len(inputNums))
	copy(sliceSortNum, inputNums)

	for i := 1; i < len(sliceSortNum); i++ {
		newElem := sliceSortNum[i] //в NewElem сохранем первый элемент
		j := i
		for j > 0 && sliceSortNum[j-1] > newElem {
			// все элементы больше newElem сдвигаем враво на 1
			sliceSortNum[j] = sliceSortNum[j-1]
			j--
		}
		sliceSortNum[j] = newElem
	}
	fmt.Println(sliceSortNum)
}
