package node

import "fmt"

type Speaker interface {
	SayHello()
	SayGoodbye(name string)
}

type Human struct {
	Greeting string
}

func (Human) SayHello() {
	fmt.Println("Hello")
}

func (Human) SayGoodbye(name string) {
	fmt.Println("goodbye " + name)
}