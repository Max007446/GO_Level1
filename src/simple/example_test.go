package simpletest

import (
	"fmt"
)

func ExampleSimple() {
	fmt.Println(Simple(10))
	fmt.Println(Simple(0))
	fmt.Println(Simple(20))
}

// Output:
//[1, 2, 3, 5, 7]
//[1]
//[1, 2, 3, 5, 7, 11, 13, 17, 19]
