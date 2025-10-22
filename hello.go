package main

import "fmt"

func main() {
    var code int
    fmt.Scan(&code)
    
    switch c := code; c {
    case 1:
        fmt.Println("Заказ создан")
    case 2:
        fmt.Println("В обработке")
    case 3:
        fmt.Println("Отправлен")
    case 4:
        fmt.Println("Доставлен")
    default:
        fmt.Println("Неизвестный статус")
    }
}






