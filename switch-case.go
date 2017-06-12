package main

import ("fmt")

func main () {
    fmt.Println("")
    fmt.Println("Switch Case:")
    fmt.Println("------------")
    fmt.Print("Enter Number (between 1 to 5) : ")
    var number int ; fmt.Scanln(&number)
    switch number {
        case 1:
            fmt.Println("Entered Number is One")
        
        case 2:
            fmt.Println("Entered Number is Two")
        
        case 3:
            fmt.Println("Entered Number is Three")
        
        case 4:
            fmt.Println("Entered Number is Four")
        
        case 5:
            fmt.Println("Entered Number is Five")
        
        default:
            fmt.Println("Entered Number is Out of Range")
    }

}