// goto jumps forward to shared cleanup. It can never skip a declaration
// nor dive into a block, which keeps it honest. Prefer labeled break or an
// early return; save goto for the rare case where both read worse.
package main

import "fmt"

func main() {
	speed := 12
	for speed < 100 {
		if speed%5 == 0 {
			goto done // both exits below need the same wind-down
		}
		speed = speed*2 + 1
	}
	fmt.Println("loop finished on its own")

done:
	fmt.Println("wind down at speed", speed)

	// Illegal jumps, left as comments so the file still builds:
	// goto skip jumps over declaration of b at ./main.go (no skipping vars)
	//   speed := 1
	//   goto skip
	//   extra := 2
	// skip:
	//   _ = extra
	// goto inner jumps into block (no diving into inner or parallel blocks)
	//   if speed > 0 {
	//       goto inner
	//   }
	//   if speed < 100 {
	//   inner:
	//       fmt.Println("unreachable legally")
	//   }
}
