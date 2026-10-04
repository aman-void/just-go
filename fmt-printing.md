# Printing with `fmt`: a practical guide

Go prints through one package, `fmt`, and twelve functions that all work the
same way. Once you see the pattern, every function feels familiar.

Official reference: https://pkg.go.dev/fmt#hdr-Printing (Go 1.27).
The explanation below is my own; the examples are original.

## The big picture

Each function name is built from two parts: **where the text goes** and
**how it is shaped**.

**Where** — the beginning of the name:

- No letter (`Print`): onto the screen (standard output).
- `S` (`Sprint`): back to you as a string. Nothing is displayed.
- `F` (`Fprint`): into a destination you provide — a file, the screen, an
  in-memory buffer. Go calls such destinations "writers".
- `A` (`Append`): onto the end of a byte list (`[]byte`).

**How** — the ending of the name:

- No ending (`Print`): list your values with commas; each is shown sensibly.
- `ln` (`Println`): same, but values are always separated by spaces and a
  new line is added at the end. The easy option.
- `f` (`Printf`): you write a template with placeholders (`%d` for a number,
  `%s` for a string) and supply the values. Precise control.

That gives the full family:

| Destination   | Plain    | Easy (`ln`) | Template (`f`) |
| ------------- | -------- | ----------- | -------------- |
| Screen        | `Print`  | `Println`   | `Printf`       |
| As a string   | `Sprint` | `Sprintln`  | `Sprintf`      |
| Into a writer | `Fprint` | `Fprintln`  | `Fprintf`      |
| Onto bytes    | `Append` | `Appendln`  | `Appendf`      |

Below, each family is shown the way the official docs present it —
signature, what it does, and a worked example.

## Printing to the screen

```go
func Print(a ...any) (n int, err error)
func Println(a ...any) (n int, err error)
func Printf(format string, a ...any) (n int, err error)
```

`Print` writes each value in its default form. `Println` adds spaces
between all values and a newline at the end. `Printf` formats by template.
All three return the byte count and any write error.

```go
laps := 5
best := 58.4

fmt.Print("laps: ", laps, " best: ", best, "\n")
// laps: 5 best: 58.4
// No newline is added for you, hence the explicit "\n".

fmt.Println("laps:", laps, "best:", best)
// laps: 5 best: 58.4
// Spaces and the newline come free — best for quick output and debugging.

fmt.Printf("laps: %d best: %.1f\n", laps, best)
// laps: 5 best: 58.4
// %d marks an integer slot, %.1f a decimal rounded to one place.
```

## Returning a string

```go
func Sprint(a ...any) string
func Sprintln(a ...any) string
func Sprintf(format string, a ...any) string
```

Identical shaping rules, but nothing is displayed: the text is handed back
as a `string`. Reach for these when a function needs to _build_ a message —
an error, a log line, a test fixture — rather than print it.

```go
msg := fmt.Sprint("lap ", 3, " done")       // "lap 3 done"
line := fmt.Sprintln("lap", 3, "done")      // "lap 3 done\n"
report := fmt.Sprintf("lap %d of %d", 3, 5) // "lap 3 of 5"
fmt.Print(report, " | ", msg, " | ", line)
```

## Writing to a destination you choose

```go
func Fprint(w io.Writer, a ...any) (n int, err error)
func Fprintln(w io.Writer, a ...any) (n int, err error)
func Fprintf(w io.Writer, format string, a ...any) (n int, err error)
```

The first argument is any writer: `os.Stdout`, `os.Stderr`, an open file, a
network connection, or a `bytes.Buffer` (an in-memory notepad, ideal for
practice and tests). The returned error matters for real destinations —
files and networks can fail — so check it outside of examples.

```go
var buf bytes.Buffer // in-memory notepad; needs `import "bytes"`
fmt.Fprint(&buf, "lap ", 3)
fmt.Fprintln(&buf, "done")
fmt.Fprintf(&buf, "best: %.1f\n", 58.4)
fmt.Print(buf.String())
// lap 3done
// best: 58.4

// Problems go to Stderr so they never mix with real program output:
fmt.Fprintf(os.Stderr, "warn: lap %d over cutoff\n", 6)
```

## Appending to a byte list

```go
func Append(b []byte, a ...any) []byte
func Appendln(b []byte, a ...any) []byte
func Appendf(b []byte, format string, a ...any) []byte
```

These glue formatted text onto the end of a `[]byte` and hand the longer
list back — handy when assembling bytes directly, e.g. network payloads,
without round-tripping through strings. One habit to build: always keep
the result (`b = fmt.Append(b, ...)`), because Go may reuse the underlying
storage and discarding it loses your text.

```go
var b []byte // starts empty
b = fmt.Append(b, "lap ", 3)
b = fmt.Appendln(b, "done")
b = fmt.Appendf(b, "best: %.1f", 58.4)
fmt.Printf("%s\n", b)
// lap 3done
// best: 58.4
```

## Placeholders, the ones you'll actually reach for

In template (`f`) functions, each `%` code marks a slot for the next value.
A small set covers nearly all daily use:

| Code          | Meaning                                            | Example                      | Prints               |
| ------------- | -------------------------------------------------- | ---------------------------- | -------------------- |
| `%v`          | the value, plainly                                 | `fmt.Printf("%v", 42)`       | `42`                 |
| `%+v`         | struct with field names — your debugging friend    | `fmt.Printf("%+v", p)`       | `{name:Ada score:9}` |
| `%#v`         | complete Go-style dump                             | `fmt.Printf("%#v", "x")`     | `"x"`                |
| `%T`          | the type, not the value                            | `fmt.Printf("%T", 3)`        | `int`                |
| `%%`          | a literal percent sign                             | `fmt.Printf("70%%")`         | `70%`                |
| `%t`          | a boolean                                          | `fmt.Printf("%t", true)`     | `true`               |
| `%d`          | a whole number (`%b` binary, `%o` octal, `%x` hex) | `fmt.Printf("%d", 42)`       | `42`                 |
| `%c`          | the character for a code point                     | `fmt.Printf("%c", 65)`       | `A`                  |
| `%q`          | quoted, safe to reuse                              | `fmt.Printf("%q", "café")`   | `"café"`             |
| `%f` / `%.2f` | decimal / rounded to 2 places                      | `fmt.Printf("%.2f", 58.456)` | `58.46`              |
| `%s`          | a string as-is                                     | `fmt.Printf("%s", "pit")`    | `pit`                |
| `%p`          | a pointer's address                                | `fmt.Printf("%p", &laps)`    | `0xc000...`          |

And for lining values into neat columns:

| Code          | Meaning                                   | Example                              | Prints        |
| ------------- | ----------------------------------------- | ------------------------------------ | ------------- |
| `%9.2f`       | at least 9 characters wide, 2 decimals    | `fmt.Printf("%9.2f", 58.4)`          | `    58.40`   |
| `%-9d`        | align to the left                         | `fmt.Printf("%-9d\|", 42)`           | `42       \|` |
| `%05d`        | pad with zeros                            | `fmt.Printf("%05d", 42)`             | `00042`       |
| `%#x`         | prefix the base (`0x`, `0b`, `0`)         | `fmt.Printf("%#x", 255)`             | `0xff`        |
| `%[2]d %[1]d` | pick values by position, reuse or reorder | `fmt.Sprintf("%[2]d %[1]d", 11, 22)` | `22 11`       |

## Common stumbles

1. **Output glued to the next line.** `Print` adds no newline; add `"\n"`
   yourself or switch to `Println`.
2. **Missing spaces.** `Print` joins strings with no gap: `Print("a", "b")`
   gives `ab`, `Print("a", 1)` gives `a1` — a space appears only between
   two non-strings. `Println` always separates, so prefer it when unsure.
3. **A strange `%!d(string=hi)` in the output.** That is Go telling you a
   placeholder got the wrong kind of value — here a number slot received a
   string. It never crashes; run `go vet ./...` and it points these out.
4. **A struct printed as `{Ada 9}`.** Which field is which? Use `%+v` while
   debugging to see `{name:Ada score:9}`.
5. **A mystery blank line.** The `ln` functions already end with a newline,
   so `Println("hi\n")` prints an extra empty line. Add newlines yourself
   only with the plain and `f` forms.
6. **Errors mixed into results.** Report problems on `os.Stderr`
   (`Fprintf(os.Stderr, ...)`), keeping them apart from real output when
   someone pipes your program into a file or another tool.
