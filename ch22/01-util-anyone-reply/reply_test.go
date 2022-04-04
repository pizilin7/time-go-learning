package replay_test

import (
    "fmt"
    "runtime"
    "testing"
    "time"
)

func runTask(id int) string {
    time.Sleep(10 * time.Millisecond)
    return fmt.Sprintf("the result is from %d", id)
}

func reply() string {
    n := 10
    ch := make(chan string)
    for i := 0; i < n; i ++ {
        go func(i int) {
            ret := runTask(i)
            ch <- ret
        }(i)
    }

    return <- ch
}

// 使用buffer channel
func reply1() string {
    n := 10
    ch := make(chan string, n)
    for i := 0; i < n; i ++ {
        go func(i int) {
            ret := runTask(i)
            ch <- ret
        }(i)
    }

    return <- ch
}

func TestReply(t *testing.T) {
    // t.Log("Before: ", runtime.NumGoroutine())
    // t.Log(reply())
    // time.Sleep(time.Second)
    // // 还有gorotine存在，造成协程泄露
    // t.Log("Afeter: ", runtime.NumGoroutine())

    t.Log("Before: ", runtime.NumGoroutine())
    t.Log(reply1())
    time.Sleep(time.Second)
    // gorutine不在等待接受者，不会造成泄露
    t.Log("Afeter: ", runtime.NumGoroutine())
}