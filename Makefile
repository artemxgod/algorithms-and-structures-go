all: run

.PHONY: run
run:
	go run main.go

DEFAULT_GOAL: run