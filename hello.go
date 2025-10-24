package main

import "fmt"

func main() {
    // Поиск числа 7 в последовательности от 1 до 10
    for i := 1; i <= 10; i++ {
        if i == 7 {
            fmt.Println("Число 7 найдено, выход из цикла.")
            break
        }
        fmt.Println(i)
    }
}






