//
// primebench.go
// Copyright (C) 2016 cceckman <charles@cceckman.com>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"fmt"
	"os"
	"github.com/cceckman/bencher"
	"github.com/cceckman/bencher/examples/isprime"
)

var(
	maybePrimes = []int {
		2, 3, 11, 16,
		7919, 7920, 7979,
		29443, 30253, 30257,
		106877, 106878, 106879,
	}
)

func main() {
	// Close the functions to bench over the inputs.
	cases := make(bencher.Cases)
	for name, impl := range isprime.Implementations {
		for _, n := range maybePrimes {
			caseName := fmt.Sprintf("%s(%d)", name, n)
			cases[caseName] = func() fmt.Stringer {
				return strungBool(impl(n))
			} // bencher.Runnable
		} // for each input
	} // for each implementation

	// cases has all our cases. Run all of them.
	if err := bencher.AutoBenchmark(cases); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// Implement fmt.Stringer for bool, since Go doesn't on its own.
type strungBool bool
func (b strungBool) String() string { return fmt.Sprintf("%t", b) }


