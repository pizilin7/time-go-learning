package condition_test

import "testing"

// func TestIfMultiSec(t *testing.T) {
// 	if v, err := someFun(); err != nil {
// 		do something
// 	}
// }

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i ++ {
		switch i {
		case 0,2:
			t.Log("Event")
		case 1,3:
			t.Log("odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}