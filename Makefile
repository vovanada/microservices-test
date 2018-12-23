run-tests:
	TEST_URL=http://0.0.0.0:8888 go test -v -count=1 ./test-pack/...

run-linter:
	golangci-lint run -v

.PHONY: run-tests run-linter
