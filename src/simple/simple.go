package simpletest

var b int
var n int

//Simple вызов
func Simple(n int) []int {
	b = 0
	inputNums := []int{}
	for {
		b++
		if SimpleN(b) {
			inputNums = append(inputNums, b)
		}
		if b >= n {
			break
		}
	}
	return inputNums
}

// SimpleN наименьший делитель составного числа x не превосходит sqrt x,поэтому достаточно
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
