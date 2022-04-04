package csp_test

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

// 串行输出
func TestService(t *testing.T) {
    fmt.Println(service())
    otherTask()
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

func TestAsynService(t *testing.T) {
    retCh := AsyncService()
    otherTask()
    fmt.Println(<-retCh)
}