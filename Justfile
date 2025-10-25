format:
    go fmt
    go mod tidy
    fd --hidden --extension yml --exec-batch yamlfmt

lint:
    golangci-lint run

test:
    go test -v ./...
