package main

import "fmt"

func main() {
	cards := deck{"Ace of Diamonds", newCard()} // a slice of type string

	// take the slice of cards & loop over it
	for i, card := range cards { // range is a keyword for iterating over every element in an array
		fmt.Println(i, card)
	}

}

func newCard() string { // must declare what type will be returned by function
	return "Five of Diamonds"
}

// ---- **Notes: General** ----
// vs code will auto import things when you save a file

// ---- **Notes: Variable Declarations** ----
// var deckSize int
// can initialize a variable outside of a function, we just can't assign a value to it.
//
// Variable Assignment & Reassignment:
// var card string = "Ace of Spades"
// card := "Ace of Spades"  // relies on Go compiler to figure out type (does infer types to a degree unlike Java/C++)
// card = "Five of Dimonds" // only need colon when you first initialize value, not for reassignment
// fundamental Go data types: bool, string, int, float64
// other Go data types: array, slice

// ---- **Notes: Functions & Return Types** ----
// func main() {
// 	card := newCard()
// 	fmt.Println(card)
// }
//
// func newCard() string { // must declare what type will be returned by function
// 	return "Five of Diamonds"
// }

// ---- **Notes: Slices & For Loops ----
// Files in the same package can freely call functions defined in other files.
// Two data structures for handling lists of records:
// 1. Array - fixed length list of things
// 2. Slice - an array that can grow or shrink;
// For both: each element in a slice/array must be of same type
//
//
// cards := []string{"Ace of Diamonds", newCard()} // a slice of type string
// fmt.Println(cards)
//
// add a new element to a slice
// cards = append(cards, "Six of Spades") // pass in existing slice, & new element that we want to append at the end of the slice
// append does not modify the existing slice, instead returns a new slice that we then assign back to the variable cards
//
// take the slice of cards & loop over it
// for i, card := range cards { // range is a keyword for iterating over every element in an array
// 	fmt.Println(i, card)
// }

// ---- **Quiz Q: Slices & For Loops** ----
// What represents a slice where each element in it is of type int?
// - []int{}
// Is the following code valid? `colors := []strings{"Red", "Yellow", "Blue"}`
// - No, "strings" is not a valid data type.
// How do we iterate through each element in a slice and print out its value & index position?
// - `colors :=[]string{"Red", "Yellow", "Blue"}`
// - `for index, color := range colors { fmt.Println(index, color) }`
// Would the following code compile successfully? `for index, card := range cards { fmt.Println(card) }`
// - No, because every variable we declare must be used in our code. Here, 'index' is not being used.
// Can a slice have both values of type 'int' and of type 'string' in it?
// - No, because a slice can only have one type of value in it.

// ---- **Notes: OO approach vs Go approuch** ----
// Go is not an object oriented language (no idea of classes within Go)
// However, can draw some of that object oriented approach over to Go

// OO approach:
// Use Deck Class to create a Deck instance.
// A deck instance might have a property called cards (which could be an array of strings).
// Could then have functions attached to the Deck instance (functions like print, shuffle, & saveToFile).
// Could then use these functions to manipulate that list of cards that are attached to the Deck instance.

// Go approach:
// To create the idea of a Deck inside of our Go program, we define a new type inside of Go (a Deck type).
// We want to extend a base data type & add extra functionality to it.
// So, we'll create an array of strings (type deck []string)...
// & add functions specifically for it (functions with 'deck' as a 'receiver').
// A function with a receiver is like a 'method' - a function that belongs to an 'instance'.
