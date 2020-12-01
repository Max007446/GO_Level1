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
	case "/":
		res = a / b
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
