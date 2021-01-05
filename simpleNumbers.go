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
			fmt.Println(b)
		}
		if b >= n {
			break
		}
	}
}

// наименьший делитель составного числа x не превосходит sqrt x,поэтому достаточно
//проверить только делители числа sqrt x
func SimpleN(x int) bool {
	if x < 1 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		x0 := x % i
		if x0 == 0 {
			return false
		}
	}
	return true
}
