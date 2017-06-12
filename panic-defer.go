package main

import ("fmt")

func main () {
    fmt.Println("")
    fmt.Println("Panic & Defer:")
    fmt.Println("--------------")
    fmt.Println("testing before calling panic")    
    defer func() {
        err := recover()
        fmt.Println("first defer  : ", err)       
        
        defer func() {
            err := recover()
            fmt.Println("second defer : ", err)
            //goto AfterPanic
        }()
        panic("panic2")
    }()
    //AfterPanic:
    panic("panic1")
    
    
    fmt.Println("testing after calling panic")    
//    panic("panic1 called")
//    recover()

}