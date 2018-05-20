package main

import (
  "fmt"
)

func fibonacci(l int) {
  a := 0
  b := 1
  var r  []int
  r = append(r,0)
  for i := 0; i < l; i++ {
    a, b = b , a + b
    r = append(r,a)
  }
  fmt.Println(r)
}

func main() {
  fibonacci(10)
}
