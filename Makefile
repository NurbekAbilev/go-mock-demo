run:
	CGO_ENABLED=1 go run cmd/main.go

install-mockery:
	go install github.com/vektra/mockery/v2@v2.53.3
