package main

import ("fmt"
        "math")

func foo(var1 float64) {
  fmt.Printf("Square root of %g is %g\n", var1, math.Sqrt(var1))
}

func main() {
  fmt.Printf("hello, world\n")
  
  foo(16)
  foo(4)
}
