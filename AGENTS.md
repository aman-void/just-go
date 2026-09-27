# AGENTS.md

- Root `README.md` is the codebase index + Go tooling reference; each `NN-slug/README.md` holds that chapter's notes.
- Each `NN-slug/` is an independent Go module with its own `go.mod` (`github.com/aman-void/just-go/<slug>`, Go 1.27). No root `go.mod`, no `go.work`.
- Chapter layout: root `main.go` is overview only; each subtopic is a runnable `package main` in its own folder (`go run ./<topic>`). Never add a nested `go.mod`.
- New chapters: `NN-kebab-slug/` + `go.mod` + `main.go` overview + topic subfolders + `README.md` notes. Keep every example original.
- Chapter check: `gofmt -l . && go vet ./... && go build ./...`. Root has no Go code to build.
- Session CWD defaults to `01-hello-go/`; repo root is its parent `just-go/`. `cd` to the target chapter dir before running Go commands.
- Chapter `README.md` style (human-friendly notes): big-idea intro first, then a
  run-order table (`Order | Run it | What it teaches you`), then one section per
  subtopic with plain-English idea + real-world analogy + short code/output quoted
  from the actual program + one beginner trap, ending with a rules cheat sheet.
- Chapters are numbered in learning order: `01-hello-go`, `02-predeclared-types`, `03-composite-types`, …
