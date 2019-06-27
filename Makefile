
test:
	go test -v ./...

mod:
	go mod tidy
	go mod vendor

