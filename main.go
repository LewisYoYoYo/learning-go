package main

import (
	"fmt"
	"math"
	"github.com/LewisYoYoYo/project1/strutil"
)

func main() {
	// short hand
	name := "Lewis"
	// long hand
	var age int64 = 19
	var isCool = true
	fmt.Println(name, age, isCool)
	// Tells you the variable type (Int, string, boolean ect)
	fmt.Printf("%T\n", age)
	// Rounds Down
	fmt.Println(math.Floor(2.7))
	// Round up
	fmt.Println(math.Ceil(2.7))
	// Square root
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutil.Reverse(hello)))
}
