.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

## ------------------------------------------------- Common commands: --------------------------------------------------
## Formats the code.
format:
	${call colored,formatting is running...}
	go vet ./...
	go fmt ./...
