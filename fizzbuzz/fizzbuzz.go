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

	switch {
	case n%15 == 0:
		return "FizzBuzz"
	case n%5 == 0:
		return "Buzz"
	case n%3 == 0:
		return "Fizz"
	default:
		return strconv.Itoa(n)
	}
}
