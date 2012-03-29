package main

import "fmt"

type Movie struct {
	Rating int
	Name   string
}

type Critic struct {
	Name   string
	Movies []Movie
}

func main() {
	liza := Critic{
		Name: "Lisa Rose",
		Movies: []Movie{
			Movie{Name: "Lady in the Water", Rating: 2},
			Movie{Name: "Snakes on a Plane", Rating: 3},
		},
	}
	jack := Critic{
		Name: "Jack Matthews",
		Movies: []Movie{
			Movie{Name: "Lady in the Water", Rating: 3},
			Movie{Name: "Snakes on a Plane", Rating: 4},
		},
	}
	critics := []Critic{liza, jack}
	fmt.Println(critics)
}
