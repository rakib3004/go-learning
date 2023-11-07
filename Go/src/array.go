package main

import "fmt"

func main() {
    var arr [5]int
    arr[0] = 1101
    arr[1] = 1102
    arr[2] = 1102
    arr[3] = 1103
    arr[4] = 1104
    arr[5] = 1105

    slice := []int{1, 2, 3, 4, 5,6,7,8,9,10}

    fmt.Println("Array:", arr)
    fmt.Println("Slice:", slice)
}
