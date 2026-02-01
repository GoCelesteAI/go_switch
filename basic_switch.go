package main

import "fmt"

func main() {
  day := "Wednesday"

  switch day {
  case "Monday":
    fmt.Println("Start of the week")
  case "Wednesday":
    fmt.Println("Midweek")
  case "Friday":
    fmt.Println("Almost weekend!")
  default:
    fmt.Println("Regular day")
  }

  // Multiple values in one case
  switch day {
  case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
    fmt.Println("Weekday")
  case "Saturday", "Sunday":
    fmt.Println("Weekend")
  }
}
