package main

import ("fmt";"time")

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main () {
    
    fmt.Println("")
    fmt.Println("Concurrency - GoRoutines:")
    fmt.Println("-------------------------")
    
//    go func () {
//        fmt.Println("world")
//    }()

//    go fmt.Println("world") 
//    fmt.Println("Hello")
//    time.Sleep(time.Second)    
//    //time.Sleep(time.Second/109500)

//    go say("World")
//    say("Hello")    
    
    for i := 0 ; i < 100 ; i++ {
        go func (i int) {
            fmt.Println("Hello", i)
        }(i)
    }
    time.Sleep(time.Second)

}