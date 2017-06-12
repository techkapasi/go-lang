package main

import ("fmt";"math")

//func add(x int, y int) int {
func add(x, y int) int {    
	return x + y
}

func main () {
    fmt.Println("")
    fmt.Println("Using math library:")
    fmt.Println("-------------------")
    fmt.Printf("Square root of 7 is : %g. \n", math.Sqrt(7))
    fmt.Printf("Value of PI is : %g. \n", math.Pi)
    fmt.Printf("Addition of %d and %d => %d \n",42, 13, add(42, 13))
}