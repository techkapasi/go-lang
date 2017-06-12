package main

import ("fmt")

func main () {
    fmt.Println("")
    fmt.Println("If Condition:")
    fmt.Println("--------------")    
    if x1,x2:= 1,1 ; x1 > 0 && x2 > 0  {
        fmt.Printf("x1 : %d , x2 : %d \n", x1, x2)
        fmt.Println("x1 & x2 both are greater than 0")
    } else if (x1 <= 0 && x2 <= 0) {
        fmt.Printf("x1 : %d , x2 : %d \n", x1, x2)
        fmt.Println("x1 & x2 both are less than or equals to 0")
    } else {
        fmt.Printf("x1 : %d , x2 : %d \n", x1, x2)
        fmt.Println("Either x1 or x2 is less than or equals to 0")
    }    
}