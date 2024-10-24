package main

import (
    "fmt"
)

func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

func printPermutation(n, r int) {
    result := factorial(n) / factorial(n-r)
    fmt.Printf("Permutation(%d, %d) = %d\n", n, r, result)
}

func printCombination(n, r int) {
    result := factorial(n) / (factorial(r) * factorial(n-r))
    fmt.Printf("Combination(%d, %d) = %d\n", n, r, result)
}

func main() {
    var a, b, c, d int

    fmt.Print("Enter four integers (a, b, c, d): ")
    fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)

    if a >= c && b >= d {
        printPermutation(a, c)
        printCombination(a, c)
        printPermutation(b, d)
        printCombination(b, d)
    } else {
        fmt.Println("Invalid input")
    }
}
