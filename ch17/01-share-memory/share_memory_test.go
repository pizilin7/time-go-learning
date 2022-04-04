package share_memory_test

import (
	"sync"
	"testing"
	"time"
)

//

// 使用锁
func TestCounterByMutex(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0;  i < 5000; i ++ {
		go func() {
			defer func(){
				mut.Unlock()
			}()

			mut.Lock()
			counter++
		}()
	}

	// 直接打印并不是最终的结果，有一些routine还在运行中
	// t.Logf("counter = %v", counter)
	time.Sleep(1 * time.Second)
	t.Logf("counter = %v", counter)
}

func TestMutexWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0;  i < 5000; i ++ {
		wg.Add(1)
		go func() {
			defer func(){
				mut.Unlock()
				wg.Done()
			}()

			mut.Lock()
			counter++
		}()
	}

	// 使用WaitGroup会更准确
	wg.Wait()
	t.Logf("counter = %v\n", counter)
}