GO_BENCH_ARGS ?= -benchtime=30s
	
format:
	gofmt -w .
benchmark:
	@go test -bench=. $(GO_BENCH_ARGS) --run=$(run) -benchmem -tags=benchmark ./... 2>&1 | grep -E '^(Benchmark|PASS|FAIL|ok|\s+[0-9]+)' | sort > bench.txt
b1x:
	make benchmark GO_BENCH_ARGS="-benchtime=1x"
b10x:
	make benchmark GO_BENCH_ARGS="-benchtime=10x"
b1s:
	make benchmark GO_BENCH_ARGS="-benchtime=1s"
b10s:
	make benchmark GO_BENCH_ARGS="-benchtime=10s"
tidy:
	@echo "\033[0;32mRunning go mod tidy...\033[0m"
	go mod tidy -v
	@echo "\033[0;32mVerifying packages...\033[0m"
	go mod verify