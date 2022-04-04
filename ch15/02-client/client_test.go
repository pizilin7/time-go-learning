package client_test

import (
    "fmt"
    "testing"
    "time-go-learning/ch15/01-series"
)

func init() {
    fmt.Println("client test, init")
}

// init 依次执行
func init() {
    fmt.Println("client test, init 2")
}

func TestPackage(t *testing.T) {
    t.Log(series.GetFibonacci(10))
}