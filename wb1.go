package main

import (
	"fmt"
)

type Human struct {
	Name string
	Age  int
}

func (h *Human) Speak() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", h.Name, h.Age)
}

func (h *Human) Walk() {
	fmt.Println("I'm walking...")
}

type Action struct {
	Human
	ActionType string
}

func (a *Action) DoAction() {
	fmt.Printf("%s is performing action: %s\n", a.Name, a.ActionType)
}

func main() {

	person := Action{
		Human:      Human{Name: "Alice", Age: 30},
		ActionType: "Running",
	}

	person.Speak()
	person.Walk()

	person.DoAction()
}
