// Annotated code following part 5
package main

import ("fmt"
        "net/http")


func home_handler(w http.ResponseWriter, r *http.Request) {
  // http.ResponseWriter is not a built-in Go type; it's a custom type
  // need a writer to put info on page and a reader for the request
  // want value of request so, use *http.Request

  // Fprintf will format based on what you specify (w)
  fmt.Fprintf(w, "Hey, Oh, Let's Go!")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "The Show Must Go On")
}


func main() {
  http.HandleFunc("/", home_handler) // handler that will take path and a function
  http.HandleFunc("/about/", about_handler)
  http.ListenAndServe(":8000", nil) // 8000 is the port, nil is Go equivalent of Python's None
}
