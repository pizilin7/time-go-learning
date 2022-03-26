package map_test

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1 : 1, 2 : 4, 3 : 9}
	t.Logf("m1 = %v, len = %v", m1, len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2 = %v", len(m2))

	m3 := make(map[int]int, 10)
	t.Logf("len m3 = %v", len(m3))
}


func TestAccesNotExistStringKey(t *testing.T) {
	// 存在零值，如果没有存在的话返回是0
	m := map[int]int{}
	t.Log(m[1])

	m[2] = 0
	t.Log(m[2])

	if v, ok := m[1]; ok {
		t.Log(v)
	} else {
		t.Log("not exist")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1:3,10:2,11:90}
	for key, value := range(m1) {
		t.Logf("key = %v, value = %v", key, value)
	}
}