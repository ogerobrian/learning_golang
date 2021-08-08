package main

import (
	"fmt"
)

// declare struct
type Actor struct {
	age       int
	actorName string
	coActors  []string // a slice
}

func main() {
	// initialize the struct
	myActor := Actor{
		age:       34,
		actorName: "Brian Kaka",
		coActors: []string{
			"Obed Kaka",
			"Hem Kaka",
			"Melvin Kaka",
		},
	}
	fmt.Println(myActor)

	// create an anonymous struct
	yActor := struct{ age int }{age: 28}
	copy_strc := &yActor
	copy_strc.age = 23
	fmt.Println(yActor)
	fmt.Println(copy_strc)

}
