// := declares something new, and only inside a function. Package level, and any
// variable that already exists, wants var or a plain =.
package main

import "fmt"

var serverPort = 8080 // no := up here

func main() {
	laps := 3               // short decl, type inferred as int
	var best float64 = 58.4 // explicit type when it matters
	fmt.Println("laps:", laps, "best:", best, "port:", serverPort)

	laps = 4 // plain assignment for existing variables
	fmt.Println("updated laps:", laps)
}
