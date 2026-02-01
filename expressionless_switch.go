package main

import "fmt"

func main() {
  score := 85

  // Switch without expression
  switch {
  case score >= 90:
    fmt.Println("Grade: A - Excellent!")
  case score >= 80:
    fmt.Println("Grade: B - Good job!")
  case score >= 70:
    fmt.Println("Grade: C - Satisfactory")
  default:
    fmt.Println("Grade: F - Please study more")
  }

  hour := 14
  switch {
  case hour < 12:
    fmt.Println("Good morning!")
  case hour < 18:
    fmt.Println("Good afternoon!")
  default:
    fmt.Println("Good evening!")
  }
}
