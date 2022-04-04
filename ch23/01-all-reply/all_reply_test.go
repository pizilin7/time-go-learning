package all_reply_test

import (
	"fmt"
	"testing"
	"time"
)

func runTask(id int) string {
    time.Sleep(10 * time.Millisecond)
    return fmt.Sprintf("the result is form %d", id)
}

func allReplay() string {
    n := 10
    ch := make(chan string)
    for i := 0; i < n; i ++ {
        go func(i int) {
            ret := runTask(i)
            ch <- ret
        }(i)
    }

    // var arr []string
	var str string;
    for i := 0; i < n; i ++ {
        str += <-ch + "\n"
    }

    return str
}

func TestAllReply(t *testing.T) {
	t.Logf("%v", allReplay())
}