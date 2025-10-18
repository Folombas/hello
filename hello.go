package main

import "fmt"

func main() {
    var a, b float64
    
    fmt.Scan(&a)
    fmt.Scan(&b)
    
    area := a * b
    
    fmt.Printf("Площадь: %.2f кв.м.\n", area)
}






