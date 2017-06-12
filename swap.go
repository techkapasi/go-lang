package main

import ("fmt")

func swap(x, y string) (string, string) {
	return y, x
}

func main () {
    fmt.Println("")
    fmt.Println("Swapping:")
    fmt.Println("----------")  
    var a,b string ; a = "Hello" ; b = "World"
    fmt.Println("Value before Swapping: ", a, b)
    a, b = swap(a,b)
    //a, b := swap("hello", "world")
    fmt.Println("Value before Swapping: ", a, b)
}