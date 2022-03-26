package string_test

import "testing"

func TestString(t *testing.T) {
	// 初始化默认为“”
	var s string
	t.Log("s = ", s)
	s = "hello world"
	t.Log("s = ", s)

	// string是一个不可变的byte slice
	// canot assign s[0](value of byte)
	// s[0] = "H"

	// 可以存储任务二进制数据
	s = "\xE4\xB8\xA5"
	t.Log("s = ", s)
	t.Log("s len = ", len(s))

	// Unicode是一种字符集合
	// UTF8是Unicode的一种实现方式

	c := []rune(s)
	t.Log("c len = ", len(c))
	t.Logf("unicode = %x", c[0])
	t.Logf("utf8 = %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "你好世界"
	for _, c := range s {
		// c是rune类型，而不是byte
		t.Logf("%[1]c, %[1]x", c)
	}

}