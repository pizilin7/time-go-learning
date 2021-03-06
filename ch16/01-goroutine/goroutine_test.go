package goroutine

import (
    "fmt"
    "testing"
    "time"
)

func TestGoroutine(t *testing.T) {
    for i := 0; i < 10; i ++ {
        // go func(i int) {
        // 	fmt.Printf("i = %v\n", i)
        // }(i)

        go func() {
            // 读取当前的i的数值，
            fmt.Printf("i = %v\n", i)
        }()
        fmt.Printf("x = %v\n", i)
    }
    time.Sleep(time.Second)
}