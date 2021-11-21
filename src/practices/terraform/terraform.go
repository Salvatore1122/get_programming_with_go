package terraform

import "fmt"

type planets []string

func (p planets) terraform() {
	for i := range p {
		p[i] = "New " + p[i]
	}
}

func Answer() {
	var planets planets = []string{"Mars", "Uranus", "Neptune"}
	planets.terraform()
	fmt.Println(planets)
}
