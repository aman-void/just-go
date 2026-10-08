# 06 — Pointers

The big idea: Go passes everything by value, so a function can never reach back
and change the variable you handed it. Pointers are how you hand over an address
instead of a copy — and, because Go has a garbage collector, that address is
just a note in a notebook, not a promise to free anything by hand. Think of it
as writing an address on a sticky note: the note is not the house, but you can
send someone to the house with it.

## How to run

- `go build ./...` type-checks the overview and the subexample at once.
- `go run .` is only an orientation; every topic is its own program.
- Or `make list` lists the topics and `make check` runs the chapter gate.
- Suggested order for the first pass:

| Order | Run it | What it teaches you |
| ----- | ------ | ------------------- |
| 1 | `go run .` | `&` takes an address, `*` reads the value back out |
| 2 | `go run ./call-by-value` | Every argument is a copy — and why slices and maps still surprise you |

## 1. A pointer is a sticky note with an address on it (`main.go`)

Plain English: a pointer is a variable that holds the memory address of another
value. `&x` hands you that address; `*p` reads whatever is stored there. Two
operators, and they are inverses of each other: take an address, read the value.

Real world: `x` is a crate on a shelf. `pointerToX` is not the crate, it is the
label on the crate. Hand someone the label and they can walk over and touch the
crate. Hand them a photo of the crate (`y := *pointerToX + " world"`) and they
only get a copy.

The second half of the program is the other end of the story: a pointer's zero value
is `nil`, which holds no address at all, so there is nothing to read. And
`new(T)` is the opposite — it allocates a zero `T` and hands back its address, so
that pointer is never `nil`.

```go
x := "hello"
pointerToX := &x // & is the address operator

fmt.Printf("address of x: %v\n", pointerToX)
fmt.Printf("value at pointerToX: %v\n", *pointerToX)

y := *pointerToX + " world"

var z *int
fmt.Println("z is nil:", z == nil)

// BUG: dereferencing z would panic, so it stays commented out:
// fmt.Println("dereference nil pointer:", *z)

if z == nil {
	fmt.Println("z is nil, so there is nothing to dereference")
} else {
	fmt.Println("dereference z pointer:", *z)
}

intPtr := new(int)
strPtr := new(string)
boolPtr := new(bool)

fmt.Println("new() is never nil:", intPtr != nil)
fmt.Printf("zero values: int=%d str=%q bool=%t\n", *intPtr, *strPtr, *boolPtr)
```

Output:

```text
Welcome to pointers
address of x: 0x85f6cb52070
value at pointerToX: hello
value of y: hello world
z is nil: true
z is nil, so there is nothing to dereference
new() is never nil: true
zero values: int=0 str="" bool=false
```

Beginner trap: printing a pointer prints the *address*, and that address is
meaningless to a reader and different on every run. `0x85f6cb52070` tells you
nothing; `hello` tells you everything. If you want to see data, dereference
first. The one-line tell: a pointer printed with `%v` and no `*` in front of it
is a bug waiting to happen — and the `nil` line is the same trap with the
ground pulled out from under it, since `*z` there is a runtime panic rather than
a surprise number.

The matching trap on the other side: `new(T)` is *not* a fancy way to get a `nil`
pointer. It returns a real address, so a nil check after it guards against
nothing. Reach for `new` only when you want a pointer to a zero value and
nothing else — `n := 0; p := &n` is usually clearer.

## 2. Go is always call by value (`call-by-value/`)

Plain English: when you pass a variable to a function, Go makes a copy and hands
the function the copy. Reassigning a parameter changes the copy, never the
original. Slices and maps look like exceptions, and that is where most of the
confusion lives: the *slice header* and the *map header* get copied, but both
still point at the same underlying data.

Real world: you hand a colleague a photocopy of a form. They fill in the copy;
your original is untouched. Now hand them the key to the filing cabinet instead
and everything they file is visible to you — the cabinet itself was never copied.
Slices and maps are that key; `int`, `string`, and `struct` are the photocopy.

```go
func modifyBasic(i int, s string, p person) {
	i = 1 * 2
	s = "Goodbye"
	p.name = "Bob"
}

func modSlice(s []int) {
	for i := range s {
		s[i] = s[1] * 2
	}
	s = append(s, 10)
	fmt.Println("inside modSlice:", s)
}
```

Output:

```text
========== 1. BASIC VALUES ==========
outside: 2 Hello {20 Alice}

========== 2. SLICE ==========
before: [1 2 3]
inside modSlice: [4 4 8 10]
after: [4 4 8]
```

The two halves to read carefully: `i`, `s`, and `p` come back exactly as they
started, because those parameters were copies. But `after: [4 4 8]` shows the
caller's slice *was* modified, even though only a copy of the slice value was
passed — writing to `s[i]` writes through to the shared backing array.

Beginner trap: `s = append(s, 10)` inside a function changes only the local
parameter, so the caller's slice keeps its old length and no `append` result ever
escapes. Write to the existing elements (`s[i] = v`) and the caller sees it;
reassign the slice itself and the caller does not. Same map story in reverse:
`m[k] = v` is visible to the caller, `m = map[...]{}` is not.

## Rules I want to remember

- `&x` is the address, `*p` is the value. They are inverses, not interchangeable.
- Print values, never addresses — an address is noise that changes every run.
- A nil pointer has nothing to dereference: `*p` on `p == nil` panics.
- Go has no `free`, no `delete(ptr)`, no dangling pointers. The garbage
  collector owns the memory; a pointer is just a value that gets copied around.
- Every argument is a copy. Reassigning a parameter never reaches the caller.
- Reference-typed values (slices, maps, channels, pointers) copy the header, not
  the data — so element writes are shared, and reassignment is not.
- `unsafe` exists for the cases where the type system's guarantees are in your
  way. It is rare, it is not needed for normal Go, and the compiler will not
  save you from it.
