# Go Switch Statements

Code examples from **Go Tutorial for Beginners #7 - Control Flow: switch**

## Contents

- `basic_switch.go` - Basic switch statements with multiple case values
- `expressionless_switch.go` - Switch without expression (like if-else chains)
- `fallthrough.go` - Using the fallthrough keyword
- `type_switch.go` - Type switches for interface types

## Running the Examples

```bash
go run basic_switch.go
go run expressionless_switch.go
go run fallthrough.go
go run type_switch.go
```

## Key Concepts

### Basic Switch
```go
switch day {
case "Monday", "Tuesday":
    fmt.Println("Weekday")
case "Saturday", "Sunday":
    fmt.Println("Weekend")
default:
    fmt.Println("Unknown")
}
```

### Expressionless Switch (like if-else)
```go
switch {
case score >= 90:
    fmt.Println("Grade: A")
case score >= 80:
    fmt.Println("Grade: B")
default:
    fmt.Println("Grade: F")
}
```

### Fallthrough
```go
switch num {
case 1:
    fmt.Println("One")
    fallthrough
case 2:
    fmt.Println("Two")
}
```

### Type Switch
```go
switch v := i.(type) {
case int:
    fmt.Println("Integer:", v)
case string:
    fmt.Println("String:", v)
default:
    fmt.Printf("Unknown: %T\n", v)
}
```

## Watch the Tutorial

[Go Tutorial for Beginners #7 - Control Flow: switch](https://youtube.com/@CelesteAI)

## License

MIT
