test:
	@go test -v -race -covermode=count -coverprofile=tmp/coverage.out | sed ''/PASS/s//$$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$$(printf "\033[31mFAIL\033[0m")/''
.PHONEY: test
