package array_list_test

import "testing"

// 数组的初始化
// 使用下标获取数组的值
func TestArray(t *testing.T) {
	arr1 := [3]int{1,3,4}
	arr2 := [4]int{10,100,1000}
	arr3 := [...]int{0,3,23}

	t.Log(arr1[1], arr2[3],arr3[2])
}

// 数组的两种遍历方式
func TestArrayTravel(t *testing.T) {
	arr1 := [3]int{1,3,4}
	for i := 0; i < len(arr1); i ++ {
		t.Log(i, arr1[i])
	}

	for index, value := range(arr1) {
		t.Log(index, value)
	}
}

// 数组切片
func TestArraySlice(t *testing.T) {
	arr := [...]int{1,3,4,10,23,22}
	arr_slice := arr[:]
	arr_slice2 := arr[1:4]
	arr_slice3 := arr[3: len(arr)]

	t.Log(arr_slice)
	t.Log(arr_slice2)
	t.Log(arr_slice3)
}