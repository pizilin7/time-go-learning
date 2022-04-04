// panic 用于不可恢复的错误，会执行defer
// os.Exist 不会执行defer，不会输出堆栈信息
package panic_recover_te_test

import (
    "errors"
    "fmt"
    // "os"
    "testing"
)

func TestPanicAndRecover(t *testing.T) {
    // panic下会执行
    defer fmt.Println("defer,1")	

    defer func() {
        // 错误信息被捕捉，并恢复
        // 不会抛出异常堆栈
        // 善用recover,僵尸进城，直接let it crash
        if error := recover(); error != nil {
            fmt.Println("error = ", error)
        }
    }()

    // os.Exit()使用
    // fmt.Println("Start")
    // os.Exit(-1)
    // fmt.Println("End")

    // fmt.Println("Start")
    error := errors.New("Something wrong")
    if error != nil {
        panic(error)
    }
    // fmt.Println("End")	

    // 不会执行
    defer fmt.Println("defer,2")	
}