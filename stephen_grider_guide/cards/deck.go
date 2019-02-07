package main

import "fmt"

// Create a custom type of 'deck' which is a slice of strings
type deck []string

// Create a receiver function, print(), for custom type 'deck'
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// ---- **Custom Type Declarations Notes** ----
// Create a new type of 'deck'
// which is a slice of strings
// type deck []string
//
// Now you can do:
// cards := deck{"Ace of Diamonds", newCard()}
// instead of:
// cards := []string{"Ace of Diamonds", newCard()}

// ---- **Receiver Functions Notes** ----
// Receiver: sets up methods on variables that we create
// any variable of type 'deck' now gets access to the 'print' method / can call this function on itself
//
// func (d deck) print() { // put receiver & type as arguments
// 	for i, card := range d { // for i in range d
// 		fmt.Println(i, card)
// 	}
// }
//
// d is the actual copy of the deck we're working with
//  (the instance of the deck variable that we're working with)
//   & is available in the function as a variable called 'd'.
//   Imagine that the variable 'cards' is passed into the print function as the variable 'd'.
//   'd' argument is also kind of like this or self, but called 1-3 letter abbreviation by convention
// deck is a reference to the type that we want to attach this method to

// ---- **Quiz Q: Receiver Functions** ----
// What would the following code print out?
// package main
//
// import "fmt"
//
// type book string
//
// func (b book) printTitle() {
//     fmt.Println(b)
// }
//
// func main() {
//     var b book = "Harry Potter"
//     b.printTitle()
// }
// - Answer: Harry Potter
// 
// By creating a new type with a function that has a receiver, we...
// - Are adding a 'method' to any value of that type. 
//
// In the following snippet, what does the variable 'ls' represent?
// type laptopSize float
// func (ls laptopSize) getSizeOfLaptop() { return ls }
// - A value of type 'laptopSize'
// 
// Is the following code valid? 
// type laptopSize float64
// func (this laptopSize) getSizeOfLaptop() laptopSize { return this }
// - Yes, it technically is & will compile, but it is breaking convention
// - Go avoids any mention of 'this' or 'self'