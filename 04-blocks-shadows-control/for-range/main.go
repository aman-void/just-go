// for-range is the default walk over slices, maps, and strings. It hands
// you copies, decodes runes in strings, and visits map keys in random order.
package main

import (
	"fmt"
	"sort"
)

func main() {
	stints := []int{58, 61, 59}
	for i, v := range stints {
		fmt.Println(i, v)
	}

	// Blank the index when only values matter; drop the value for key-only
	// loops (most common when a map is used as a set).
	for _, v := range stints {
		fmt.Print(v, " ")
	}
	fmt.Println()

	grid := map[string]bool{"ada": true, "bo": true, "cy": true}
	for name := range grid {
		fmt.Print(name, " ")
	}
	fmt.Println()

	// Unsorted map ranging shuffles on purpose (hash seed + jitter), so sort
	// keys for stable output. fmt.Println sorts map keys when printing maps.
	keys := make([]string, 0, len(grid))
	for k := range grid {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Print(k, " ")
	}
	fmt.Println()

	// Ranging a string yields the byte offset and the rune, decoding each
	// multibyte character. Indexing s[i] would hand back raw bytes instead.
	for i, r := range "pit_π!" {
		fmt.Println(i, r, string(r))
	}

	// The value is a copy: mutating it leaves the slice untouched.
	for _, v := range stints {
		v *= 2
	}
	fmt.Println("unchanged:", stints)

	// Since Go 1.22 each iteration gets fresh i and v, so closures started
	// in the loop no longer share one reused pair. The go line in go.mod
	// selects this behavior.
}
