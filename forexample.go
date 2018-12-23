// `for` is Go's only looping construct. Here are
// three basic types of `for` loops.

package main

import (
	"fmt"
	"time"
)

func main() {

	// The most basic type, with a single condition.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// A classic initial/condition/after `for` loop.
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// `for` without a condition will loop repeatedly
	// until you `break` out of the loop or `return` from
	// the enclosing function.
	var counter int = 0
	for {
		if counter < 10 {
			fmt.Println("loop")
			time.Sleep(2 * time.Second)
			fmt.Println("through")
			time.Sleep(2 * time.Second)
			counter++
		} else {
			break
		}
	}

	// You can also `continue` to the next iteration of
	// the loop.
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
