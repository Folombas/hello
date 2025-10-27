package main

import "fmt"

func main() {
    result := 1
    
    for result <= 1000 {
        fmt.Println(result)
        result *= 2
    }
}
