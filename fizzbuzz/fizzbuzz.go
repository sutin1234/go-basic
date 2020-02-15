package fizzbuzz

// RUN TEST
// filename fizzbuzz_test.go
// run test normal => go test
// run test with verbose => go test -v

func Say(n int) string {
	if n == 1 {
		return "1"
	}

	return "0"
}
