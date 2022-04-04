// 运算符
package operation_test

import "testing"

const (
    Readable = 1 << iota
    Wirteable
    Executable
)

func TestCompareArray(t *testing.T) {
    // a := [...]int{1, 2, 3, 4, 5}
    b := [...]int{1, 2, 3, 4}
    c := [...]int{1, 2, 4, 3}
    d := [...]int{1, 2, 3, 4}

    // 长度不相同的不能不计较，mismatched types [5]int and [4]int)
    // t.Log(a == b)
    t.Log(b == c)
    t.Log(b == d)
}

// 按位清运算符
// 1 &^ 0  1
// 1 &^ 1  0
// 0 &^ 1  0
// 0 &^ 0  0

func TestBitClear(t *testing.T) {
    a := 7
    a = a &^ Readable
    t.Log(a & Readable, a & Wirteable, a & Executable)
}
