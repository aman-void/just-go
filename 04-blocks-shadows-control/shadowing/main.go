// := always declares in the current block, so it can quietly shadow an
// outer variable instead of updating it. Use = when you mean to reuse.
package main

import "fmt"

func main() {
	fuel := 10
	if fuel > 5 {
		fmt.Println("before:", fuel)
		fuel := 5 // shadow: a new fuel that lives only inside this if body
		fmt.Println("inside:", fuel)
	}
	fmt.Println("after:", fuel) // outer fuel never changed

	if fuel > 5 {
		// := reuses only names from the same block, so fuel shadows again
		// even though pit is the only truly new variable here.
		fuel, pit := 5, 20
		fmt.Println("shadowed pair:", fuel, pit)
	}
	fmt.Println("outer still:", fuel)

	// The fix is plain assignment when the variable already exists:
	fuel = 5
	fmt.Println("reassigned:", fuel)

	// The same trap applies to predeclared names (true, len, ...) and to
	// imports: a local called fmt or len would hide the package or builtin
	// for the rest of its block, so never reuse those names.
}
