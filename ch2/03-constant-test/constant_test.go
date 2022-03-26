package constant_test

import "testing"

const (
	Monday = 1 + iota
	Tuesday
	Wenesday
)

const (
	Readable = 1 << iota
	Wirteable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Tuesday, Wenesday)
}

func TestConstantTry2(t *testing.T) {
	a := 7
	t.Log(a&Readable, a&Wirteable, a&Executable)
}