// A couple of "is X prime" solutions, of varying speed.
package isprime

import (
	"math"
)

type IsPrime func(int) bool

var (
	Implementations = map[string]IsPrime{
		"SimpleTestDiv": SimpleTestDiv,
		"BetterTestDiv": BetterTestDiv,
		"SieveErat":     SieveErat,
		"BetterErat":    BetterErat,
	}
)

func SimpleTestDiv(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func BetterTestDiv(n int) bool {
	switch {
	case n < 2:
		return false
	case n == 2:
		return true
	case n%2 == 0:
		return false
	}

	for i := 3; i < n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func SieveErat(n int) bool {
	switch {
	case n < 2:
		return false
	case n == 2:
		return true
	case n%2 == 0:
		return false
	}
	// Uses twice the memory needed... eh, whatever.
	composite := make([]bool, n+1)
	for i := 3; i < n; i += 2 {
		if composite[i] {
			continue
		}
		// i is prime; sieve
		for k := i * 2; k < len(composite); k += i {
			composite[k] = true
		}
	}
	return !composite[n]
}

func BetterErat(n int) bool {
	switch {
	case n < 2:
		return false
	case n == 2:
		return true
	case n%2 == 0:
		return false
	}
	// Optimize memory; skip evens.
	composite := make([]bool, n/2+1)
	sqrt := int(math.Ceil(math.Sqrt(float64(n))))

	for i := 3; i <= sqrt; i += 2 {
		if composite[(i-1)/2] {
			continue
		}
		// i is prime; hasn't yet been marked composite.

		for j := i; j <= n; j += (i + i) {
			composite[(j-1)/2] = true
		}
	}

	return !composite[(n-1)/2]
}
