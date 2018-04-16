package main

import "fmt"

type Animal struct {
	Name string
}

func (a *Animal) Speak() {
	fmt.Println("UGH!",
		"My name is", a.Name)
}

type Dog struct {
	Animal
	PackFactor int
}

func (d *Dog) Speak() {
	fmt.Println("Woof!",
		"My name is", d.Name,
		"I am a mammal with a pack factor of", d.PackFactor)
}

type Cat struct {
	Animal
	ClimbFactor int
}

func (c *Cat) Speak() {
	fmt.Println("Meow!",
		"My name is", c.Name,
		"I am a mammal with a climb factor of", c.ClimbFactor)
}

func main() {
	animals := []Animal{
		Dog{
			Animal: Animal{
				Name: "Fido",
			},
			PackFactor: 5,
		},
		Cat{
			Animal: Animal{
				Name: "Milo",
			},
			ClimbFactor: 4,
		},
	}

	for _, animal := range animals {
		animal.Speak()
	}
}

// =============================================================================

// NOTES:

// Smells:
// 	* The Animal type is providing an abstraction layer of reusable state.
// 	* The program never needs to create or solely use a value of type Animal.
// 	* The implementation of the Speak method for the Animal type is a generalization.
// 	* The Speak method for the Animal type is never going to be called.
