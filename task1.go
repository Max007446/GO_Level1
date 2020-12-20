//Явно использовал указатели при вводе операндов в калькуляторе
//Также явно указывал указатели при запросе вводе максимального числа при поиске простых чисел https://github.com/Max007446/GO_Level1/blob/lesson2_simple/simpleNumbers.go
//В задании с Фибоначчи тоже был запрос чисел через указатели - https://github.com/Max007446/GO_Level1/blob/lesson4_Fibonahhi/fibonahhi.go

//Не явно мы используем указатели в той же задаче Фибоначчи при использовании и передачи map func fibo(n uint64, cacheMap map[string]uint64, result uint64) uint64 {
//Не явно используется при копировании входного слайса в сортировках sortSlise := make([]int64, len(inputNums))
package main

import (
	"fmt"
	"math"
	"os"
)

var a, b, res float64
var op string

func main() {
	fmt.Print("Введите первый операнд : ")
	fmt.Scanln(&a)
	fmt.Print("Введите второй операнд : ")
	fmt.Scanln(&b)
	fmt.Print("Введите необходимую операцию на операндами : ")
	fmt.Scanln(&op)
	switch op {
	case "*":
		res = a * b
	case "+":
		res = a + b
	case "-":
		res = a + b
	case "~a":
		res = math.Sqrt(a)
	case "^":
		res = math.Pow(a, b)
	default:
		fmt.Println("Вы выбрали не валидную операцию")
		os.Exit(1)
	}

	fmt.Printf("Результат операции: %f\n ", res)
}
