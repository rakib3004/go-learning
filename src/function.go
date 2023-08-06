package main

import "fmt"

func add(a, b int) int {
    return a + b
}
func subtract(a, b int) int {
    return a - b
}

func main() {
    summation_result := add(5, 3)
	subtraction_restult := subtract(5, 3)

    fmt.Println("5 + 3 =", summation_result)
	fmt.Println("5 - 3 =", subtraction_restult)

}
