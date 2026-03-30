test_verbose:
	go test -v ./...

test:
	go test ./...

branch_coverage:
	gobco ./sorting
	gobco ./searching

test_benchmark:
	go test -bench=. ./...

clean:
	go clean -testcache
