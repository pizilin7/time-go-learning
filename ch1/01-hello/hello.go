package main

import (
    "fmt"
    "os"
)

func main() {
    // 打印main函数带入的参数
    fmt.Println("args =", os.Args[1:])
    fmt.Println("hello world")

    // main函数返回
    os.Exit(-1)
}