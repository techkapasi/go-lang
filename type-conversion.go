package main

import ("fmt";"math")

func main () {
    fmt.Println("")
    fmt.Println("Type Conversion:")
    fmt.Println("----------------")    
    var num1,num2 int = 3, 4
    //var float_num float64 = math.Sqrt((num1*num1 + num2*num2))
    var float_num float64 = math.Sqrt(float64(num1*num1 + num2*num2))
    var u_integer = uint(float_num)
    fmt.Println("Num 1        : ", num1)
    fmt.Println("Num 2        : ", num2)
    fmt.Println("Float Num    : ", float_num)
    fmt.Println("Unsigned Num : ", u_integer)

}