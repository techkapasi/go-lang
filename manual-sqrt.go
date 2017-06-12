package main

import ("fmt";"math")

func Sqrt(x float64) float64 {    
    var x1,z float64 ; z = x
    for {
            x1 = (z - (  (z*z) - x )  / (2 * z) )
            z = x1

            if  (x1 * x1) > x   {
                    break
                }
        }
    return x1
}

func main () {
    fmt.Println("")
    fmt.Println("Square Root:")
    fmt.Println("------------")
    var sqrnum float64 = 2
    fmt.Println("Manual Square root of ", sqrnum, " is : ", Sqrt(sqrnum))
    fmt.Println("Auto Square root of ", sqrnum, " is : ",math.Sqrt(sqrnum))    

}