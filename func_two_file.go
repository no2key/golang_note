/*функция в новом файле*/

main.go
********************
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	z := my_test(10, 20)
	x := (rand.Intn(10))
	fmt.Println("Work", x, "test", "now", z)
}

**************
test.go
**************
package main

func my_test(x, y int) int {
	return (x + y)
}