// NOTE: a pointer is a variable that holds the memory address of another value.
// & takes the address, * reads the value back out. The syntax comes from C and
// C++, but Go's garbage collector removes the memory bookkeeping they force on
// you. The `unsafe` package goes lower still, down to raw memory and
// hand-written assembly.
package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	x := "hello"
	pointerToX := &x // & is the address operator: pointerToX now holds x's address

	fmt.Printf("address of x: %v\n", pointerToX)
	fmt.Printf("value at pointerToX: %v\n", *pointerToX)

	// WARNING: an address means nothing to a reader and changes on every run, so
	// never print one where a value belongs. Dereference first, then print.

	// Dereferencing also works anywhere a value is expected, so a pointer can be
	// read without ever naming the variable it points at.
	y := *pointerToX + " world"

	fmt.Printf("value of y: %v\n", y)

	// WARNING: a pointer's zero value is nil — it holds no address at all — and
	// dereferencing one panics, so check against nil before every dereference.
	var z *int

	fmt.Println("z is nil:", z == nil)

	// BUG: dereferencing z would panic, so it stays commented out:
	// fmt.Println("dereference nil pointer:", *z)

	// The guard that makes the dereference in the else branch safe.
	// if z == nil {
	// 	fmt.Println("z is nil, so there is nothing to dereference")
	// } else {
	// 	fmt.Println("dereference z pointer:", *z)
	// }

	// NOTE: new(T) allocates a fresh zero value of T and hands back its address,
	// so the result is never nil. Rarely needed — `n := 0; p := &n` says the same
	// thing and shows the value — but it is the built-in for "give me a pointer to
	// a zero T".
	intPtr := new(int)
	strPtr := new(string)
	boolPtr := new(bool)

	fmt.Println("new() is never nil:", intPtr != nil)

	// WARNING: this is the opposite of `var z *int` above. new(int) gives back a
	// usable pointer to a zero int, so there is nothing to guard against. Note
	// %q on the string: its zero value is "", and a bare %v prints nothing at all.
	fmt.Printf("zero values: int=%d str=%q bool=%t\n", *intPtr, *strPtr, *boolPtr)

	// The `new()` function is rarely used. For structs we use & before the struct literal to create a pointer instance.
}
