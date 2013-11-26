GOOS=${shell go env GOOS}
test:
ifeq (${GOOS},darwin)
	@CC=GCC-4.2 go test
else
	@go test
endif
