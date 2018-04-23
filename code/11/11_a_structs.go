package main

type Animal struct {
	Name     string
	IsMammal bool
}

func (a *Animal) Speak() {}

type Dog struct {
	Animal
	PackFactor int
}

func (d *Dog) Speak() {}

type Cat struct {
	Animal
	ClimbFactor int
}

func (c *Cat) Speak() {}
