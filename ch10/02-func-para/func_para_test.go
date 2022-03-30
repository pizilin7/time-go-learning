package func_para

import (
	"fmt"
	"testing"
)

func sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret;
}

func TestVarParam(t *testing.T) {
	t.Log("sum = ", sum(1, 3, 5, 7, 9))
}

func Clear() {
	fmt.Println("clear")
}

func TestDefer(t *testing.T) {
	defer Clear()
	t.Log("test defer")

	// 依旧执行 defer
	panic("err")
}