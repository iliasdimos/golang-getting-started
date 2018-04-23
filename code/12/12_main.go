package main

func main() {
	// Create a list of Animals that know how to speak.
	speakers := []Speaker{
		// Create a Dog by initializing its Animal parts and specific Dog attributes
		&Dog{
			Name:       "Fido",
			IsMammal:   true,
			PackFactor: 5,
		},
		// Create a Cat by initializing its Animal parts and specific Cat attributes
		&Cat{
			Name:        "Milo",
			IsMammal:    true,
			ClimbFactor: 4,
		},
	}
	// Have the Animals speak.
	for _, spkr := range speakers {
		spkr.Speak()
	}
}
