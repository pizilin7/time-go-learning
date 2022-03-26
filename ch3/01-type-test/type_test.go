package type_test

import "testing"

type MyInt int32

// 不支持隐式转换
func TestImplicitConvert(t *testing.T) {
	var a int32 =1
	var b int64
	var c MyInt

	b = int64(a)
	c = MyInt(a)

	t.Log(b, a, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	t.Logf("%T %T", a, aPtr)
}

// 初始化的带有零值,而不是空值
func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
}