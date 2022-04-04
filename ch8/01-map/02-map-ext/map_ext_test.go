package map_ext

import "testing"

func TestMapWithFunvalue(t *testing.T) {
    // value可以是一个function
    m := map[int]func(op int)int {}
    m[1] = func(op int) int{return op}
    m[2] = func(op int) int{return op * op}
    m[3] = func(op int) int{return op * op * op}

    t.Log(m[1](10), m[2](10), m[3](10))
}

func TestMapForSet(t *testing.T) {
    mySet := map[int]bool{}
    mySet[1] = true
    n := 10
    if mySet[n] {
        t.Logf("%d is existing", n)
    } else {
        t.Logf("%d is not existing", n)
    }

    mySet[3] = true
    t.Log(len(mySet))
    delete(mySet, 1)
}