package main

import "fmt"

func main() {
    for i := 1; i <= 50; i++ {
        // Проверяем на нечетность и выводим только нечетные числа
        if i%2 != 0 {
            fmt.Println(i)
        }
    }
}






