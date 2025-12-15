format:
    go fmt
    go mod tidy
    fd --hidden --extension yml --exec-batch yamlfmt
    fd --hidden --extension xml --exec xmllint --format --output {} {}

lint:
    golangci-lint run

test:
    go test -v ./...

test_cover:
    go test -v -cover -coverprofile=coverprofile.out ./...
