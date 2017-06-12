package main

import ("fmt";"math/rand")

func main () {
    fmt.Println("")
    fmt.Println("Using math/rand library:")
    fmt.Println("------------------------")
    fmt.Println("A Random number is", rand.Intn(10))
}