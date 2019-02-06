package main

import "fmt"

func main() {
    fmt.Println("Let's Go world!")
}

// **Notes**
// How do we run our code in our project?
// - `go build` just compiles -> creates an executeble `main`; run via `./main`
// - `go run` compiles & executes
// `go format` auto formats code inside files
// `go install` & `go get` handle dependencies inside projects
// - `go install` compiles & installs a package
// - `go get` downloads the raw source code of someone elses package
// `go test` runs any test associated with the project`
// What does 'package main' mean?
// - package is a collection of common source code files
// - package == project == workspace
// - each file needs to declare which package it belongs to
// - 2 types of packages: executable (generates a file that we can run) & reusable (code used as helpers -- good place o put reusable logic, like libraries)
// - word `main` is used to build an executable type package
// - any other name (ex: `package calculator`) defines a package that can be used as a dependency (helper code)
// - note that any time we have a an executable package, it must have a function called `main` within it as well
// What does 'import "fmt"' mean?
// - gives our package access to code & functionality that's contained within another package
// - fmt is a standard library package that is included in Go by default
// - fmt library is used to print out info to terminal for debugging
// - by default, our package has no access to default libraries -> must spec link by using import
// - can also import packages that are authored by other engineers
// - golang.org/pkg for docs of default packages
// What is `func`?
// - declare functions with keyword `func`, name of function, set of parens with optional list of args, & curly brackets with body
// How is the main.go file organized?
// - same pattern: {top: package declaration, mid: import other packages that we need, bot: declare functions}

// **Quiz Q**
// purpose of packages :to group together code with a similar purpose
// What is the special name we use for a package to tell Go that we want it to be turned into a file that can be executed? `main`
// The one requirement of packages named main is that: we define a function named "main", which is ran automatically when the program runs
// Why do we use import statements? to give our package access to code written in another package
