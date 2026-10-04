# ─────────────────────────────────────────────────────────────────────────────
#  Just GO · repo tasks
#
#  The repo is a shelf of independent chapter modules (each with its own go.mod
#  and its own Makefile). This one fans tasks out across the whole shelf.
#
#    make            this help
#    make list       chapters and every topic they contain
#    make check      gofmt + vet + build across all chapters
#
#  Chapter work stays inside the chapter's own Makefile:
#    make -C 05-functions            help for that chapter
#    make -C 05-functions run slices
# ─────────────────────────────────────────────────────────────────────────────

SHELL := /bin/sh
.DEFAULT_GOAL := help

# A chapter is any folder that is a Go module.
CHAPTERS := $(patsubst %/go.mod,%,$(wildcard */go.mod))

C_RESET := \033[0m
C_BOLD  := \033[1m
C_DIM   := \033[2m
C_GREEN := \033[38;5;42m
C_RED   := \033[38;5;203m

.PHONY: help list check

help:
	@printf '\n'
	@printf '  %bJust GO%b  ·  learn Go by running it\n\n' '$(C_BOLD)' '$(C_RESET)'
	@printf '  %bmake list%b               chapters and their topics\n' '$(C_GREEN)' '$(C_RESET)'
	@printf '  %bmake check%b              gofmt + vet + build every chapter\n' '$(C_GREEN)' '$(C_RESET)'
	@printf '  %bmake -C <chapter> ...%b   any chapter task (run, new, list, check)\n\n' '$(C_GREEN)' '$(C_RESET)'
	@printf '  %bexample%b  make -C 05-functions run defer-func\n\n' '$(C_DIM)' '$(C_RESET)'

list:
	@for c in $(CHAPTERS); do \
	  printf '\n  %b%s%b\n' '$(C_BOLD)' "$$c" '$(C_RESET)'; \
	  $(MAKE) --no-print-directory -C $$c list; \
	done

check:
	@fail=0; \
	for c in $(CHAPTERS); do \
	  printf '\n  %b== %s ==%b\n' '$(C_BOLD)' "$$c" '$(C_RESET)'; \
	  $(MAKE) --no-print-directory -C $$c check || fail=1; \
	done; \
	if [ $$fail -ne 0 ]; then \
	  printf '\n  %bsome chapters are not clean%b\n' '$(C_RED)' '$(C_RESET)'; \
	  exit 1; \
	fi; \
	printf '\n  %ball chapters clean%b\n' '$(C_GREEN)' '$(C_RESET)'
