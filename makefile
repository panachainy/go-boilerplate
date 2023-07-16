dev:
	air

i:
	go get ./cmd
	go mod tidy

t: test
test:
	go test ./...

f: fmt
fmt:
	go fmt ./...

g: generate
generate:
	go generate ./...

b: build
build:
	go build -o api ./cmd
