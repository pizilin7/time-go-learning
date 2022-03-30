package interface_test

import (
	"fmt"
	"testing"
)

type Programmer interface {
	WirteHelloWorld() string
} 

type GoProgrammer struct {

}

func (g *GoProgrammer)WirteHelloWorld() string {
	return fmt.Sprintf("%p", g)
}

func TestClient(t *testing.T){
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WirteHelloWorld())
}