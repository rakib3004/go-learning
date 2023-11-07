package main

import "fmt"

func main() {
    number := 29
    if number > 0 {
        fmt.Println("Number is positive.")
    } else if number == 0 {
        fmt.Println("Number is zero.")
    } else {
        fmt.Println("Number is negative.")
    }
}
