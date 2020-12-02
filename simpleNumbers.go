package main

import (
	"fmt"
)

var b int
var n int

func main() {

	fmt.Print("Введите предельное число поиска простых чисел : ")
	fmt.Scanln(&n)

	for {
		b++

		if SimpleN(b) {
			println(b)
		}
		if b >= n {
			break
		}
	}
}
func SimpleN(x int) bool {
	if x < 1 {
		return false
	}
	result := true

	for i := 2; i*i <= x; i++ {
		if x%i == 0 {

			result := false
			return result
		}
	}
	return result
}
