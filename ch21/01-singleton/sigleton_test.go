package sigleton

import (
	"fmt"
	"sync"
	"testing"
)

type sigleton struct {}

var sigletonInstance *sigleton
var once sync.Once

func GetSigleton() *sigleton {
    // 保证只创建一个单例对象
    once.Do(func() {
        sigletonInstance = new(sigleton) 
        fmt.Println("create sigleon instance")
    })

    return sigletonInstance
}

func TestCreateSigleton(t *testing.T) {
    var wg sync.WaitGroup
    for i := 0; i < 10; i ++ {
        wg.Add(1)
        go func(i int) {
            instance := GetSigleton()
            fmt.Printf("i = %d, instance = %p\n", i, instance)
            wg.Done()
        }(i)
    }

    wg.Wait()
}