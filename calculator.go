package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

var a, b, res float64
var op string
var (
	divError = errors.New("Divide by zero error")
)

func intdiv(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, divError
	}

	return a / b, nil
}
func main() {
	fmt.Print("Введите первый операнд : ")

	_, err := fmt.Scanln(&a)
	if err != nil {
		panic(err)
	}
	fmt.Print("Введите второй операнд : ")
	if err != nil {
		panic(err)
	}
	fmt.Scanln(&b)
	if err != nil {
		panic(err)
	}
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
		res, err = intdiv(a, b)
		if err != nil {
			fmt.Println(err)

		}
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
