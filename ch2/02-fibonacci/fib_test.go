package fib_test

import "testing"

func TestFib(t *testing.T) {
	// 1 1 2 3 5 8 13
	var a int = 1
	var b int = 1

	for i := 0; i < 5; i ++ {
		temp := a
		a = b 
		b = a + temp 
		t.Log(b, " ")
	}
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	temp := a
	a = b
	b = temp

	t.Log("a = ", a, ", b = ", b)

	a1 := 100
	b1 := 200

	a1, b1 = b1, a1
	t.Log("a1 = ", a1, ", b1 = ", b1)
}