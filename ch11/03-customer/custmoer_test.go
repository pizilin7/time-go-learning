package customer_test

import (
	"fmt"
	"testing"
	"time"
)

// 简化
type Inner func(op int) int

// 函数运行花费的时间
func timeSpent(inner Inner) Inner {
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
	// 函数式编程
	tsSF := timeSpent(slowFun)
	tsSF(100)
}