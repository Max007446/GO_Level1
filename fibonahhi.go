package main

import (
	"fmt"
	"os"
)

func main() {
	cacheMap := make(map[string]uint64)
	cacheMap["next"] = 0
	cacheMap["prev"] = 0
	var result uint64 = 0
	var num uint64 = 0
	fmt.Print("Введите номер искомого числа Фибоначчи: \n")
	fmt.Fscan(os.Stdin, &num)

	fmt.Println("Число Фибоначчи: \n", fibo(num, cacheMap, result))
}

func fibo(n uint64, cacheMap map[string]uint64, result uint64) uint64 {
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
	return fibo(n-1, cacheMap, result)
}
