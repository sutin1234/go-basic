package fizzbuzz

import "strconv"

// RUN TEST
// filename fizzbuzz_test.go
// run test normal => go test
// run test with verbose => go test -v

func Say(n int) string {
	if n%15 == 0 {
		return "FizzBuzz"
	}
	if n%5 == 0 {
		return "Buzz"
	}
	if n%3 == 0 {
		return "Fizz"
	}
	return strconv.Itoa(n)
}
