package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct {
	Name       string
	PackFactor int
}

func (d *Dog) Speak() {
	fmt.Println("Woof!",
		"My name is", d.Name,
		"I am a mammal with a pack factor of", d.PackFactor)
}

// Cat contains everything a Cat needs.
type Cat struct {
	Name        string
	ClimbFactor int
}

// Speak knows how to speak like a cat.
// This makes a Cat now part of a group of concrete
// types that know how to speak.
func (c *Cat) Speak() {
	fmt.Println("Meow!",
		"My name is", c.Name,
		"I am a mammal with a climb factor of", c.ClimbFactor)
}

func main() {

	// Create a list of Animals that know how to speak.
	speakers := []Speaker{

		// Create a Dog by initializing its Animal parts
		// and then its specific Dog attributes.
		&Dog{
			Name:       "Fido",
			PackFactor: 5,
		},

		// Create a Cat by initializing its Animal parts
		// and then its specific Cat attributes.
		&Cat{
			Name:        "Milo",
			ClimbFactor: 4,
		},
	}

	// Have the Animals speak.
	for _, spkr := range speakers {
		spkr.Speak()
	}
}

// =============================================================================

// NOTES:

// Here are some guidelines around declaring types:
// 	* Declare types that represent something new or unique.
// 	* Validate that a value of any type is created or used on its own.
// 	* Embed types to reuse existing behaviors you need to satisfy.
// 	* Question types that are an alias or abstraction for an existing type.
// 	* Question types whose sole purpose is to share common state.
