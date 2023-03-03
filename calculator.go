package main

import "fmt"

func main() {
    for {
        var operator string
        var num1, num2 float64

        fmt.Print("Enter operator (+, -, *, /) or type 'exit': ")
        fmt.Scanln(&operator)

        if operator == "exit" {
            fmt.Println("Exiting calculator...")
            break
        }

        fmt.Print("Enter first number: ")
        fmt.Scanln(&num1)

        fmt.Print("Enter second number: ")
        fmt.Scanln(&num2)

        var result float64

        switch operator {
        case "+":
            result = num1 + num2
        case "-":
            result = num1 - num2
        case "*":
            result = num1 * num2
        case "/":
            result = num1 / num2
        default:
            fmt.Println("Invalid operator")
            continue
        }

        fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operator, num2, result)
    }
}
