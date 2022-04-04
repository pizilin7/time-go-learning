package context_test

import (
    "context"
    "fmt"
    "sync"
    "testing"
    "time"
)

func isCancel(ctx context.Context) bool {
    select {
    case <-ctx.Done():
        return true
    default:
        return false
    }
}

func TestContext(t *testing.T) {
    ctx, cancel := context.WithCancel(context.Background())
    wg := sync.WaitGroup{}
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(i int, ctx context.Context) {
            for {
                if isCancel(ctx) {
                    break
                }

                time.Sleep(time.Millisecond * 5)
            }
            fmt.Printf("i = %v is canceled\n", i)
            wg.Done()
        }(i, ctx)
    }

    cancel()
    wg.Wait()
}
