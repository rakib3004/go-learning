package main

import "fmt"

func selectionSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        minIndex := i
        for j := i + 1; j < n; j++ {
            if arr[j] < arr[minIndex] {
                minIndex = j
            }
        }
        arr[i], arr[minIndex] = arr[minIndex], arr[i]
    }
}

func main() {
    // Create an array of 10 numbers
    arr := []int{64, 34, 25, 12, 22, 11, 90, 45, 67, 89}

    fmt.Println("Original array:", arr)

    // Sort the array using selection sort
    selectionSort(arr)

    fmt.Println("Sorted array:", arr)
}
