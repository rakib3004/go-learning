package main

import "fmt"

func main() {
    // Arrays
    var arr [5]int
    arr[0] = 1
    arr[1] = 2
    arr[2] = 3

    // Slices
    slice := []int{1, 2, 3, 4, 5}

    fmt.Println("Array:", arr)
    fmt.Println("Slice:", slice)
}
