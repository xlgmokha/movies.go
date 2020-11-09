test: fmt unit

unit:
	go test ./test/unit/...

integration:
	go test ./test/integration/...

run:
	go run main.go

lint:
	golint ./...

fmt:
	gofmt -s -w -e ./
