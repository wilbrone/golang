package main

import "fmt"

var x int

type person struct {
	Fname string
	Lname string
}

type secreteAgent struct {
	person
	lisenceToKill bool
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func (p person) speak() {
	fmt.Println(p.Fname, p.Lname, `is not the hero`)
}

func (sa secreteAgent) speak() {
	fmt.Println(sa.Fname, sa.Lname, `is the real hero`)
}

func main() {
	// x := 7
	// fmt.Println(x)
	// fmt.Printf("hello, world. and Welcome to my first Go\n")

	// xi := []int{2, 4, 7, 9, 42}
	// fmt.Println(xi)

	// m := map[string]int{
	// 	"Mark":   45,
	// 	"Philip": 42,
	// }
	// fmt.Println(m)

	p1 := person{
		"Wilbrone",
		"Baron",
	}
	fmt.Println(p1)
	// p1.speak()

	sa1 := secreteAgent{
		person{
			"Mark",
			"Philip",
		},
		true,
	}
	// sa1.speak()
	// sa1.person.speak()
	saySomething(p1)
	saySomething(sa1)
}
