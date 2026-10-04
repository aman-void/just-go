# AGENTS.md

- Root `README.md` is the codebase index + Go tooling reference; each `NN-slug/README.md` holds that chapter's notes.
- Each `NN-slug/` is an independent Go module with its own `go.mod` (`github.com/aman-void/just-go/<slug>`, Go 1.27). No root `go.mod`, no `go.work`.
- Chapter layout: root `main.go` is overview only; each subtopic is a runnable `package main` in its own folder (`go run ./<topic>`). Never add a nested `go.mod`.
- Every chapter ships an identical `Makefile` (copy it verbatim to new chapters — it derives the chapter name from its folder and the topics from `*/main.go`, so never hand-edit it). Targets: `help`, `list`, `run <topic> [args]`, `new <topic>`, `build`, `vet`, `test`, `fmt`, `check`, `tidy`; `make run <topic> ANIM=0` skips the banner. The root `Makefile` fans `list`/`check` across chapters.
- New chapters: `NN-kebab-slug/` + `go.mod` + `main.go` overview + topic subfolders + `README.md` notes + the chapter `Makefile`. Keep every example original.
- Chapter check: `make check` (same as `gofmt -l . && go vet ./... && go build ./...`). Root has no Go code to build.
- Session CWD defaults to `01-hello-go/`; repo root is its parent `just-go/`. `cd` to the target chapter dir before running Go commands.
- Chapter `README.md` style (human-friendly notes): big-idea intro first, then a
  run-order table (`Order | Run it | What it teaches you`), then one section per
  subtopic with plain-English idea + real-world analogy + short code/output quoted
  from the actual program + one beginner trap, ending with a rules cheat sheet.
- Comment tags (Zed `zed-comment` extension, already installed): tag the **first
  line only** of a comment block; continuation lines stay untagged. Tag meanings:
  `NOTE:` file header + ordinary explanation, `WARNING:` beginner traps, `BUG:`
  commented-out code that panics or fails to compile, `PERF:` capacity growth /
  map ordering / copy and GC costs, `TODO:` genuinely unfinished. Extension also
  accepts `?:`, `*:`, `!:`, `#:` and `NAME(user):`, but this repo uses the word
  tags above. Never tag commented-out code itself — it is code, not prose.
- Chapters are numbered in learning order: `01-hello-go`, `02-predeclared-types`, `03-composite-types`, …
