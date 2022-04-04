package select_test

import (
    "fmt"
    "testing"
    "time"
)

func service() string {
    time.Sleep(time.Millisecond * 500)
    return "Done"
}

func otherTask() {
    fmt.Println("working on somethind else")
    time.Sleep(time.Millisecond * 100)
    fmt.Println("Task is Done")
}

func AsyncService() chan string {
    // 等待接受者，自身协程会阻塞
    // retCh := make(chan string)

    // 创建buffer chan,会先往buffer塞消息，如果消息满了则会阻塞
    retCh := make(chan string, 10)
    go func(){
        ret := service()
        fmt.Println("return result.")
        retCh <- ret
        fmt.Println("service existed")
    }()

    return retCh
}

func TestSelect(t *testing.T) {
    select {
    case ret := <- AsyncService():
        t.Log(ret)
    case <- time.After(time.Millisecond):
        t.Error("time out")
    }
}