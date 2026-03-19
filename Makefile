.PHONY: build
build:
	go build -ldflags "-s -w" -o ./bin/mlist-backend ./cmd

.PHONY: test
test:
	go test `go list ./...`
