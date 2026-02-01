package main

import "fmt"

func main() {
  num := 2

  fmt.Println("Without fallthrough:")
  switch num {
  case 1:
    fmt.Println("One")
  case 2:
    fmt.Println("Two")
  case 3:
    fmt.Println("Three")
  }

  fmt.Println("With fallthrough:")
  switch num {
  case 1:
    fmt.Println("One")
    fallthrough
  case 2:
    fmt.Println("Two")
    fallthrough
  case 3:
    fmt.Println("Three")
  }
}
