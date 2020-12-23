package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	cacheMap := make(map[string]int64)
	cacheMap["next"] = 0
	cacheMap["prev"] = 0
	var result int64 = 0
	var num string
	err := errors.New("Ошибка")
	fmt.Print("Введите номер искомого числа Фибоначчи: \n")
	fmt.Fscan(os.Stdin, &num, err)
	res, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		panic(err)
	}

	fmt.Println("Число Фибоначчи: \n", fibo(res, cacheMap, result))

}

func fibo(n int64, cacheMap map[string]int64, result int64) int64 {
	if n == 0 {
		return result
	}
	if cacheMap["next"] == 0 && cacheMap["prev"] == 0 {
		cacheMap["next"] = 1
		return fibo(n-1, cacheMap, 0)
	}
	result = cacheMap["prev"] + cacheMap["next"]
	cacheMap["next"] = cacheMap["prev"]
	cacheMap["prev"] = result
	if cacheMap["prev"] > math.MaxInt64 || cacheMap["prev"] < 0 {
		panic("число вышло за допустимые разряды для типа int64")
	}
	return fibo(n-1, cacheMap, result)

}
