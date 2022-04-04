package string_func

import (
    "strconv"
    "strings"
    "testing"
)

func TestString(t *testing.T) {
    s := "A,B,C"
    s_arr := strings.Split(s, ",")
    t.Log("s arr = ", s_arr)

    t.Log("s arr = ", strings.Join(s_arr, "-"))
}

func TestConv( t *testing.T) {
    s := strconv.Itoa(100)
    t.Log("s = ", s)

    s1, err := strconv.Atoi("100")
    if err == nil {
        t.Log(s1)
    }
}


