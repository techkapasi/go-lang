package main

import ("fmt")

func fibonacci() func() int {
    a,b := 0,1
    return func() int {
        temp:= a
        a, b = b, a + b
        return temp
    }
}

func main () {
    fmt.Println("")
    fmt.Println("Fibonacci Series via Closure Function:")
    fmt.Println("--------------------------------------")
    fib := fibonacci()
    for j:= 0 ; j < 10 ; j++ {
        fmt.Printf("%d ", fib())
    }
    fmt.Println("") 
}