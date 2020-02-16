package fizzbuzz

import "strconv"

// RUN TEST
// filename fizzbuzz_test.go
// run test normal => go test
// run test with verbose => go test -v

type FizzOOP int

func New(i int) FizzOOP {
	return FizzOOP(i)
}

func (f FizzOOP) String() string {
	return Say(int(f))
}

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
