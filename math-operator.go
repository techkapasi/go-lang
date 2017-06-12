package main

import ("fmt")

func calc(fn func(int, int) int) int {
    return fn(4,2)
}

func Counter(count int) func() int{
    return func() int {
        count += 1
        return count
    }
}

func main () {
    fmt.Println("")
    fmt.Println("Math Operators:")
    fmt.Println("---------------")    

    add,sub,mul,div := 0,0,0,0
    
    add = calc (func(i int, j int) int {
        return i + j
    })

    sub = calc (func(i int, j int) int {
        return i - j
    })
    
    mul = calc (func(i int, j int) int {
        return i * j
    })
    
    div = calc (func(i int, j int) int {
        return i / j
    })
    fmt.Printf("Numbers       : %d and %d \n", 4,2)
    fmt.Printf("Addition      : %d \n", add)
    fmt.Printf("Subtraction   : %d \n", sub)
    fmt.Printf("Multiplicaton : %d \n", mul)    
    fmt.Printf("Division      : %d \n", div)
    
//    fmt.Println("")
//    fmt.Println("Math Operators:")
//    fmt.Println("---------------")    
//    count:=0
//    fmt.Println("Counter Called -- ", Counter(count))
//    fmt.Println("Counter Called -- ", Counter(count))
//    fmt.Println("Counter Called -- ", Counter(count)) 
//    fmt.Println("Counter Called -- ", Counter(count))    
    
}