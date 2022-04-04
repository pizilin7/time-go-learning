package remove
import "testing"
import "github.com/easierway/concurrent_map"

func TestConcurrentMap(t *testing.T) {
    m := concurrent_map.CreateConcurrentMap(99)
    m.Set(concurrent_map.StrKey("key"), 10)
    t.Log(m.Get(concurrent_map.StrKey("key")))
}

// 在go111module为off的情况
// 当前包下的vendor目录
// 向上级目录查找，直到找到src下的vendor目录
// 在GOPATH 下面查找依赖包
// 在GOROOT目录下查找