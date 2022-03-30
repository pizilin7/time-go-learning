package empty_interace_test

import (
	"fmt"
	"testing"
)

func DoSomeThing(p interface{}) {
	switch v := p.(type){
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknow Type")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomeThing(10)
	DoSomeThing("1000")
}