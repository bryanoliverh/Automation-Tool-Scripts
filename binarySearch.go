package main

import (
    "fmt"
    "flag"
    "strconv"
    "strings"
)

func binarySearch(slice []int, target int) int {
    high := len(slice) - 1
    low := 0
    for low <= high {
        mid := (high + low) / 2
        if slice[mid] == target {
            return mid
        }
        if slice[mid] < target {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return -1
}

func main() {
    // Define command-line arguments
    listStr := flag.String("L", "", "Comma-separated list of integers")
    targetStr := flag.String("T", "", "Target integer")
    flag.Parse()

    // Parse list and target arguments
    list := []int{}
    for _, s := range strings.Split(*listStr, ",") {
        n, err := strconv.Atoi(s)
        if err != nil {
            panic(err)
        }
        list = append(list, n)
    }
    target, err := strconv.Atoi(*targetStr)
    if err != nil {
        panic(err)
    }

    // Perform binary search
    index := binarySearch(list, target)
    if index == -1 {
        fmt.Println("Target not found in slice")
    } else {
        fmt.Printf("Target found at index %d\n", index)
    }
}
