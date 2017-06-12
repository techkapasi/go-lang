package main

import ("fmt";"sync")
//import ("fmt";"sync";"time")

type Integer struct {
    value int
    m sync.Mutex
}

func (this *Integer) Inc() {
    this.m.Lock()
    this.value++
    this.m.Unlock()
}

func (this *Integer) Val() int { 
    return this.value
}

func say (n int) {
    fmt.Println("Hello # ",n)
}

func main () {
    
//    fmt.Println("")
//    fmt.Println("Concurrency - Sync - MuteX:")
//    fmt.Println("---------------------------")    
//    n := &Integer{}
//    for i := 0 ; i < 1000000 ; i++ {
//        go n.Inc()
//    }
//    time.Sleep(time.Second)
//    fmt.Println(n.value)
    
    fmt.Println("")
    fmt.Println("Concurrency - Sync - WorkGroup:")
    fmt.Println("-------------------------------")    
    var wg sync.WaitGroup
    wg.Add(10)
    
    for i:=0 ; i < 10 ; i++ {
        go func (i int) {
            defer wg.Done() 
            say (i)
        }(i)
    }
    wg.Wait()
}