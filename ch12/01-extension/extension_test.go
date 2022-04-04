package extension_test

import (
    "fmt"
    "testing"
)

type Pet struct {

}

func (p *Pet) Speak() {
    fmt.Println("............")
}

func (p *Pet) SpeakTo(host string) {
    p.Speak()
    fmt.Println("Pet SpeakTo", host)
}

type PetPro struct {

}

func (p *PetPro) Talk() {
    fmt.Println("............")
}

func (p *PetPro) TalkTo(host string) {
    p.Talk()
    fmt.Println("Pet TalkTo", host)
}

type Dog struct {
    p *Pet
    *PetPro
}

func (d *Dog) Speak() {
    // d.p.Speak()
    fmt.Println("Dog:wang!wang!")		// SpeakTo方法依旧调用的是Pet的Speak方法，而不是Dog的Speak方法
}

func (d *Dog) SpeakTo(host string) {
    // d.p.SpeakTo(host)
    d.Speak()
    fmt.Println("Dog SpeakTo ", host)	// 自己实现的SpeakTo，不适用Pet的实现
}

// 是否可以重写
func (d *Dog) Talk() {
    fmt.Println("dog talk")
}

func TestDot(t *testing.T) {
    dog := new(Dog)
    dog.SpeakTo("Lin")

    // dog存在方法Talk
    // TalkTo依旧调用PetPro的Talk方法
    dog.TalkTo("Lin")
}