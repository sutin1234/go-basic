// go test custom tag
// +build integration
// go test -tags=integration

package fizzbuzz_test

import (
	"fmt"
	fizzbuzz "hello/fizzbuzz"
	"testing"
)

func TestFizzBuzz1To100(t *testing.T) {
	for i := 1; i <= 100; i++ {
		fmt.Println(fizzbuzz.Say(i))
	}
}
