package main

import "fmt"

func main() {
	fmt.Println("Let's Go world!")
}

// How do we run our code in our project?
//   - `go build` just compiles -> creates an executeble `main`; run via `./main`
//   - `go run` compiles & executes
// `go format` auto formats code inside files
// `go install` & `go get` handle dependencies inside projects
//  - `go install` compiles & installs a package
//  - `go get` downloads the raw source code of someone elses package
// `go test` runs any test associated with the project`
// What does 'package main' mean
// What does 'import "fmt"' mean?
