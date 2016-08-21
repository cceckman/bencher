// 2016-08-21 cceckman <charles@cceckman.com>
package isprime

import (
	"testing"
)

func TestIsPrime(t *testing.T) {
	cases := map[int]bool{
		-1:    false,
		1:     false,
		2:     true,
		7:     true,
		8:     false,
		6343:  true,
		1181:  true,
		1180:  false,
		10000: false,
	}

	for in, expected := range cases {
		for name, impl := range Implementations {
			got := impl(in)
			if got != expected {
				t.Errorf("Error for input %v in function %v: got %v, expected %v",
					in, name, got, expected)
			}
		}
	}
}
