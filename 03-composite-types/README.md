# 03 — Composite Types

Original notes on Go's array, slice, map, and struct types.

## Run

- `go build ./...` type-checks all eleven subexamples at once.
- `go run .` is only an orientation; every topic is its own program.
- In order of how much they explain: `slices`, `slice-aliasing`, `maps`,
  `structs`, then the rest.
- Or `make run slices`, `make run maps`, … — `make list` lists the topics and `make check` runs the chapter gate.

## Notes

- Arrays are almost never used directly: the length is part of the type, and
  assignment copies every element. They show up as fixed-size buffers and in
  `[N]byte` style APIs.
- A slice is a header of (pointer, len, cap). `var s []T` is nil, `[]T{}` is
  empty, `make([]T, len, cap)` preallocates. Only `make` sets the capacity.
- `append` writes into spare capacity, or allocates a bigger array and copies.
  Growth is an implementation detail — never depend on the exact cap.
- Copies of a slice share the backing array, so writing through one alias is
  visible in all of them. `a[i:j:k]` caps capacity so the next `append` must
  copy — the standard way to keep a function from mutating your data.
- Resetting: `s[:0]` reuses the array (cheapest, but old values stay reachable),
  `make([]T, 0, cap(s))` gets a fresh array, `s = nil` releases it. Zero out
  entries explicitly when they hold pointers you want collected.
- `copy` moves `min(len(dst), len(src))` elements and never grows a slice.
  `clear(s)` (Go 1.21+) zeroes elements; `clear(m)` empties a map.
- Converting a string to `[]byte` or `[]rune` copies, so the original is safe.
  Use `strings.Builder` (with `Grow`) for repeated concatenation.
- Maps: nil maps read fine but panic on write. `v, ok := m[k]` separates a
  missing key from a stored zero value. Iteration order is random — sort keys
  for deterministic output.
- `map[T]struct{}` is the set idiom; the empty struct occupies no memory.
  Deleting entries while ranging over a map is allowed.
- Maps can't be compared with `==` (only against nil). Use `maps.Equal` for
  comparable values, `reflect.DeepEqual` when values contain slices or maps.
- Structs are values: assignment copies. `==` compiles only if every field is
  comparable. Structs with identical fields convert to each other, struct tags
  aside, and anonymous structs work anywhere a named type does.
