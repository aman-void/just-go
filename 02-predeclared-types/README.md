# 02 — Predeclared Types and Declarations

Original notes on Go's built-in types and how to declare things.

## Run

- `go build ./...` type-checks the whole chapter at once.
- `go run .` is only an orientation; the real examples are the subfolders.
- Start with `go run ./zero-values`, then `./conversions` — between them they
  cover the two rules that trip up everything in the chapters after this one.

## Notes

- Every type has a zero value (`""`, `0`, `false`, nil slices/maps). Write structs
  so that zero value is immediately useful; `append` on a nil slice just works.
- `:=` is declare-and-assign inside functions only; package level uses `var`.
  Reassign existing variables with plain `=`.
- Constants without a named type are flexible: `const limit = 90_000` fits both
  `int` and `float64`. Add a type (`const name string = ...`) to lock it down.
  Unused constants compile fine; unused local variables do not.
- Integer division truncates — convert to `float64` before dividing for averages.
- `len(string)` counts bytes; `range` or `utf8.RuneCountInString` counts letters.
  Single quotes are runes (`'é'`), double quotes are strings.
- No implicit conversion anywhere: `int(x)`, `float64(n)`, `[]byte(s)` are all
  explicit, and `int` to `string` needs `strconv.Itoa`, not a cast.
- Naming: mixedCase for locals, MixedCase for exported, `const` in mixedCase
  too; short names (`i`, `r`) are fine in tiny scopes, full words elsewhere.
