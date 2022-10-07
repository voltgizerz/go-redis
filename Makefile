run:
	go run *.go

gen-mock:
	mockery --name=RedisInterface --recursive
	mockery --name=BookInterface --recursive

test:
	go test -cover ./...
