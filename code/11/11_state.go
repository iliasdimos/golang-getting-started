package main

func main() {
	// Create a list of Animals that know how to speak.
	animals := []Animal{
		Dog{
			Animal: Animal{
				Name:     "Fido",
				IsMammal: true,
			},
			PackFactor: 5,
		},
		Cat{
			Animal: Animal{
				Name:     "Milo",
				IsMammal: true,
			},
			ClimbFactor: 4,
		},
	}
	// Have the Animals speak.
	for _, animal := range animals {
		animal.Speak()
	}
}
