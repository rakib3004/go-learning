package main

import "fmt"

func main() {
    for i := 1; i <= 5; i++ {
        fmt.Printf("Iteration %d\n", i)
    }

	for j := 10; j >= 5; j-- {
        fmt.Printf("Reverse Iteration %d\n", j)
    }
}
