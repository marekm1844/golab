package integers

import "fmt"

/**
 * Add takes two integers as inputs and returns their sum.
 *
 * @param x an integer to be added.
 * @param y another integer to be added.
 * @return the sum of the two input integers.
 */
func Add(x, y int) int {
	return x + y
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}