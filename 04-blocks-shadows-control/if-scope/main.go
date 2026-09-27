// if can declare a variable scoped to the whole if/else chain. It keeps
// temporaries where they are used and lets each branch read them.
package main

import "fmt"

func main() {
	name := "bo"

	if n := len(name); n == 0 {
		fmt.Println("no callsign entered")
	} else if n > 5 {
		fmt.Println(name, "is a long callsign:", n)
	} else {
		fmt.Println(name, "fits the board:", n)
	}
	// fmt.Println(n) // does not compile: n exists only inside the if statement

	// The short statement shadows outer names too, so keep it to fresh names:
	cutoff := 6
	if cutoff := 4; cutoff > len(name) {
		fmt.Println("local cutoff", cutoff, "beats outer", 6)
	}
	fmt.Println("outer cutoff:", cutoff)
}
