package main

import "fmt"

func main() {
    counter := 1
    for counter < 1000 { // Цикл выполняется, пока counter меньше 1000
        fmt.Println("Счетчик:", counter)
        counter *= 2 // Увеличиваем значение counter
    }
}






