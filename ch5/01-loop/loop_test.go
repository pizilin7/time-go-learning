package loop_test

import "testing"

func TestWhileLoop(t *testing.T) {
	n := 0
	if n < 6 {
		t.Log(n)
		n ++
	}
}