package task_cancel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func isCancel(ch chan struct{}) bool {
	select {
	case <-ch:
		return true;
	default:
		return false;
	}
}

func cancel1(ch chan struct{}) {
	ch<-struct{}{}
}

func cancel2(ch chan struct{}) {
	close(ch)
}

func TestCancel(t *testing.T) {
	ch := make(chan struct{}, 1)
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int , ch chan struct{}) {
			for {
				if isCancel(ch){
					break;
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println("Cancelled i = ", i)
			wg.Done()
		}(i, ch)
	}

	// 接受者并不知道有多少个，也不知道发多少标识，接受者才能全部关闭
	// cancel1(ch)

	// 关闭通道
	cancel2(ch)
	wg.Wait()
	fmt.Println("success cancelled all")
}