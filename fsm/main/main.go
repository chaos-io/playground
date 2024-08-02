package main

import (
	"fmt"

	"chaos-io/playground/fsm"
)

func main() {
	bp := fsm.New()
	bp.Start(0)
	bp.From(0).To(1)
	bp.From(1).To(3).Then(func(m *fsm.Machine) { fmt.Println("from 1 to 3") })
	bp.Print()
	bp.From(1).To(4).Then(func(m *fsm.Machine) { fmt.Println("from 1 to 4") })
	bp.From(3).To(5).Then(func(m *fsm.Machine) { fmt.Println("from 3 to 5") })
	bp.Print()

	fmt.Println("------")
	m := bp.Machine()
	m.State()        // => a
	err := m.Goto(1) // => error(nil)
	fmt.Printf("1:%v\n", err)
	m.State()       // => b
	err = m.Goto(4) // => error(nil)
	fmt.Printf("2:%v\n", err)
	// => hola!
	err = m.Goto(5) // => error, "Transition 2 to 1 not permitted."
	fmt.Printf("1:%v\n", err)
}
