package main

import ("fmt")

func adder() func(int) int{
    sum := 0    
    return func(x int) int {
            sum += x
            return sum
    }
}

func main () {
    fmt.Println("")
    fmt.Println("Adder via Closure Function:")
    fmt.Println("---------------------------")    
    pos,neg := adder(), adder()
    for i := 0 ; i < 10 ; i++ {
        fmt.Println(pos(i), neg(-1*i))
    }
}