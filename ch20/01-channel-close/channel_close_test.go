package channel_close_test

import (
    "fmt"
    "sync"
    "testing"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {
    wg.Add(1)
    go func() {
        for i := 0; i < 10; i++ {
            ch <- i
        }
        // 通道关闭
        // 不能在通过通道发送数据
        // 从通道接受到零值
        close(ch)
        wg.Done()
    }()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
    wg.Add(1)
    go func() {
        for {
            if v, ok := <- ch; ok {
                fmt.Println("i = ", v)
            } else {
                break
            }
        }

        wg.Done()
    }()
}

func TestChannelClose(t *testing.T) {
    var wg sync.WaitGroup
    ch := make(chan int)
    dataProducer(ch, &wg)
    dataReceiver(ch, &wg)
    dataReceiver(ch, &wg)
    wg.Wait()
} 