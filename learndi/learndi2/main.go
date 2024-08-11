package main

import(
	"fmt"
)

type Suzuki struct{}
type Honda struct{}
type Panasonic struct{}
type Shell struct{}

type CarTest1 struct {
    Honda // wheel
    Panasonic // gas
}
type CarTest2 struct {
    Suzuki // wheel
    Panasonic // gas
}
type CarTest3 struct {
    Suzuki // wheel
    Shell // gas
}

// ↑ どれもcarInterfaceを満たしている

type carInterface interface {
	Wheel()
	Gas()
}

type CarDi struct {
	carInterface
}

func (s *Suzuki) Wheel()  { fmt.Println("Wheel : Suzuki") }
func (s *Honda) Wheel()   { fmt.Println("Wheel : Honda") }
func (s *Shell) Gas()     { fmt.Println("Gas : Shell") }
func (s *Panasonic) Gas() { fmt.Println("Gas : Panasonic") }

func NewcarDi(i carInterface) *CarDi{
	return &CarDi{i}
}

// carInterfaceを満たす構造体ならなんでも受け入れるようになっている
func main() {
	car1 := &CarTest1{}
	di1 := NewcarDi(car1)
	di1.Gas()
	di1.Wheel()

	car2 := &CarTest2{}
	di2 := NewcarDi(car2)
	di2.Gas()
	di2.Wheel()

	car3 := &CarTest3{}
	di3 := NewcarDi(car3)
	di3.Gas()
	di3.Wheel()
}