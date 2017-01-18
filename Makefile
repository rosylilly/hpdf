export CGO_CFLAGS := -I$(shell pwd)/tmp/libharu/include
export CGO_LDFLAGS := -L$(shell pwd)/tmp/libharu/lib

test:
	@go test -v -race -covermode=count -coverprofile=tmp/coverage.out | sed ''/PASS/s//$$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$$(printf "\033[31mFAIL\033[0m")/''
.PHONEY: test
