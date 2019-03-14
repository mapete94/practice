package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int, K int) []int {
    // write your code in Go 1.4
    if len(A) == 0 {
        return A
    }
    
    for len(A) < K  {
        K = K - len(A)
    }
    
    rotated := A[len(A)-K:]
    rotated = append(rotated, A[:len(A)-K]...)
    
    return rotated
}
