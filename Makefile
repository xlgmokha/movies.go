run:
	go run main.go

unit:
	go test ./test/unit/...

integration:
	go test ./test/integration/...
