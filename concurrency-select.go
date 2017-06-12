package main

import ("fmt";"time")

func main () {
    fmt.Println("")
    fmt.Println("Concurrency - Select:")
    fmt.Println("---------------------")    
    tick := time.Tick(100*time.Millisecond)
    boom := time.Tick(500*time.Millisecond)
    for
    {
        select {
            case <- tick:
                fmt.Println("tick")
            case <- boom:
                fmt.Println("boom") 
                goto EndSelect
            default:
            time.Sleep(100*time.Millisecond)
                fmt.Println(".") 
        }
    }
    EndSelect:
    
    fmt.Println("")
    fmt.Println("Test")   

}