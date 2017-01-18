export LD_LIBRARY_PATH := $(shell pwd)/tmp/libharu/lib:$(LD_LIBRARY_PATH)
export CGO_CFLAGS := -I$(shell pwd)/tmp/libharu/include
export CGO_LDFLAGS := -L$(shell pwd)/tmp/libharu/lib

test:
	@go test -v -race -covermode=count -coverprofile=tmp/coverage.out
	@go tool cover -html=tmp/coverage.out -o tmp/coverage.html
.PHONEY: test
