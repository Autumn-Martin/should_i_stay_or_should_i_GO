// Annotated code following parts 1-3

package main
// always use double quotes in Go
import ("fmt"
        "math"
        "math/rand") // slash notation for import instead of dot notation

func add(x, y float64) float64 {
  return x+y
}

// Here's a function that will return 2 things
func multiple(a,b string) (string, string) { // specify every return type (even if the same)
  return a,b
}

// main function will always run, like the init method in a class
func main() {
  // Casing: if 1st letter is capitalized, that function will be exported by Go (otherwise interpreted as internal)
  fmt.Println("Welcome to Go! Go your own way") // `$ godoc fmt Println` will give you documentation
  fmt.Println("--- PART 2 ---")
  fmt.Println("The square root of 625 is",math.Sqrt(625))
  fmt.Println(math.Sqrt(625), "minutes to Go")
  fmt.Println("A random number from 1-100 is", rand.Intn(101))

  fmt.Println("--- PART 3 ---")
  // var num1 float64 = 5.7
  // var num2 float64 = 7.8
  // var num1,num2 float64 = 5.7, 7.8 // condensed version of the 2 lines above

  // anything outside of the function, you would have to define using var and type
  // if inside the function, can use shorthand that does not include type
  // (but type cannot change once it's compiled -- not dynamic everytime the program comes to it, just one time)
  num1,num2 := 5.7, 7.8 // try running this, and change the add function types to float32
  // but if you want type to be explicit & not default, you should specify type

  fmt.Println(add(num1, num2))

  w1, w2 := "The show", "must Go on" // note that it will yell at you if you have any unused vars!
  fmt.Println(multiple(w1,w2))

  // converting types:
  var a int = 42
  var b float64 = float64(a)
  x := a // x will be type int
  fmt.Println(b, x)

  fmt.Println("--- PART 4 ---")
  // pointers: point to something using ampersand and the thing you want to point to
  o := 15
  p := &o // memory address
  fmt.Println(p) // prints memory address for where value of x is being stored
  // to read what's at a memory address, use an asterisk or a star
  fmt.Println(*p) // prints value stored at that memory address
  *p = 7
  fmt.Println(o)
  *p = *p**p
  fmt.Println(o)
  fmt.Println(*p)
}
