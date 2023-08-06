package main

import "fmt"

func main() {
    // For loop
    for i := 1; i <= 5; i++ {
        fmt.Printf("Iteration %d\n", i)
    }

    // While loop (Go has only for loop, but can be used as while loop)
    j := 1
    for j <= 5 {
        fmt.Printf("While loop iteration %d\n", j)
        j++
    }
}
