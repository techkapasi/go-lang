package main

import ("fmt")

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main () {
    fmt.Println("")
    fmt.Println("Split:")
    fmt.Println("------")
    n1,n2:=split(17)
    fmt.Printf("Number %d split into %d and %d \n", 17, n1,n2)
//    fmt.Printf("Number %d split into %d and %d \n", 17, split(17))
//    fmt.Printf("Number %d split into %d and %d \n", 17, n1,n2:=split(17))
}