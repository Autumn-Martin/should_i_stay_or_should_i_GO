package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func (d deck) print() { // put receiver in argument
	for i, card := range d { // for i in range d
		fmt.Println(i, card)
	}
}
