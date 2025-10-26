package main

import "fmt"

func main() {
    var rating float64
    
    // Считываем оценку фильма
    fmt.Print("Введите среднюю оценку фильма (0.0 - 10.0): ")
    fmt.Scan(&rating)
    
    // Используем switch без выражения для проверки диапазонов
    switch {
    case rating >= 0.0 && rating <= 3.9:
        fmt.Println("Ужасно!")
    case rating >= 4.0 && rating <= 6.9:
        fmt.Println("Слабовато")
    case rating >= 7.0 && rating <= 8.4:
        fmt.Println("Неплохо")
    case rating >= 8.5 && rating <= 10.0:
        fmt.Println("Отлично!")
    default:
        fmt.Println("Оценка вне диапазона")
    }
}






