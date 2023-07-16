dev:
	air

t: test
test:
	go test ./...

f: fmt
fmt:
	go fmt ./...

g: generate
generate:
	go generate ./...
