package main

import "fmt"

func main() {
    // Обратный отсчет от 10 до 0
    for i := 10; i >= 0; i-- {
        fmt.Println(i)
        
        // После вывода 0 добавляем сообщение "Старт!"
        if i == 0 {
            fmt.Println("Старт!")
        }
    }
}






