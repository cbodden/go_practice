package main
import (
	"fmt"
)

func main() {
	var i1 = 5
	fmt.Printf("an int: %d, its loc in mem: %p\n", i1, &i1)

	var intP *int
	intP = &i1
	fmt.Printf("the val at mem loc %p is %d\n", intP, *intP)
}
