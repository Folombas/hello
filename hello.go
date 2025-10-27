package main

import "fmt"

func main() {
    var a, b float64
    var operation string

    // Считываем первое число
    fmt.Scan(&a)

    // Считываем символ операции
    fmt.Scan(&operation)

    // Считываем второе число
    fmt.Scan(&b)

    // Обрабатываем операцию с помощью switch
    switch operation {
    case "+":
        result := a + b
        fmt.Printf("%.2f\n", result)
    case "-":
        result := a - b
        fmt.Printf("%.2f\n", result) // Убрал лишнюю строку с "x"
    case "*":
        result := a * b
        fmt.Printf("%.2f\n", result)
    case "/":
        if b == 0 {
            fmt.Println("Деление на ноль!")
        } else {
            result := a / b
            fmt.Printf("%.2f\n", result)
        }
    default:
        fmt.Println("Неизвестная операция")
    }
}
