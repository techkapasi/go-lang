package main

import ("fmt";"math/cmplx")

var (
    Boolean bool = false
    MaxInteger uint64 = 1<<64 - 1
    ComplexNum complex128 = cmplx.Sqrt(-5 + 12i)
)

func main () {

    fmt.Println("")
    fmt.Println("Data Types:")
    fmt.Println("-----------")    
    const pattern_syntax = "%T(%v)\n"
    fmt.Printf(pattern_syntax, Boolean, Boolean)
    fmt.Printf(pattern_syntax, MaxInteger, MaxInteger)
    fmt.Printf(pattern_syntax, ComplexNum, ComplexNum)
    
    fmt.Println("")
    fmt.Println("Undefined Variable Data Types:")
    fmt.Println("------------------------------")    
    var integer int
    var float float64
    var bool_val bool
    var string_val string
    fmt.Printf("%v %v %v %q \n", integer, float, bool_val, string_val)
}