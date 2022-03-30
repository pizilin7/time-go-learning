package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 返回多数值
func returnMultiValue() (int, int) {
	rand.Seed(time.Now().Unix())
	return rand.Intn(1000), rand.Intn(100)
}

// 函数运行花费的时间
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Printf("time spent: %v, ret = %v\n", time.Since(start).Seconds(), ret)
		return ret
	}
}

func slowFun(n int) int {
	time.Sleep(5)
	return n
}

func TestFn(t *testing.T) {
	a, b := returnMultiValue()
	t.Logf("a = %v, b = %v", a, b)

	// 函数式编程
	tsSF := timeSpent(slowFun)
	tsSF(100)
}