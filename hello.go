package main

import (
    "fmt"
    "strconv"
)

func main() {
    var age int
    fmt.Scan(&age)
    
    ageStr := strconv.Itoa(age)
    result := "Возраст пользователя: " + ageStr + " лет"
    
    fmt.Println(result)
}






