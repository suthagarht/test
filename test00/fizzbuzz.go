package main

import (
  "fmt"
)

func createRange(min, max int) []int {
  a := make([]int, max-min+1)
  for i := range a {
    a[i] = min + i
  }
  return a
}

func fizzbuzz(a []int) {
  for i := range a {
    switch {
      case  a[i] % 3 == 0 && a[i] % 5 == 0:
        fmt.Printf(" %d  Can be divided by 3 and 5\n", a[i])
      case a[i] % 3 == 0:
        fmt.Printf(" %d can only be divided by 3.\n", a[i])
      case a[i] % 5 == 0 :
        fmt.Printf(" %d can only be divided by 5.\n", a[i])
    }
  }
}

func main() {
  a := createRange(0,100)
  fizzbuzz(a)
}
