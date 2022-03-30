package error_test

import (
	"errors"
	"testing"
)

func GetFibonacci(n int) ([]int, error) {
	if n < 2 || n > 100 {
		return nil, errors.New("n should b [2,100]")
	}

	fibList := []int{1,1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i - 2] + fibList[i - 1])
	}

	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if v, error := GetFibonacci(-10); error != nil {
		t.Log(error)
	} else {
		t.Log(v)
	}
}