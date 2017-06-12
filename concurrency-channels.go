package main

import ("fmt";"time")

func sum(s []int, c chan int) {
	// implement a method that sums all elements of `s` slice and
	// sends the result to `c` channel
}

func main () {
    fmt.Println("")
    fmt.Println("Concurrency - Channels:")
    fmt.Println("-----------------------")
//    c := make(chan int)
//    go func() {
//        time.Sleep(time.Second)
//        c <- 1
//        c <- 2
//        c <- 3        
//        close(c)
//    }()
//    fmt.Println("Run")
//    for v := range c {
//        fmt.Println(v)
//    }
//    
////    <- c
////    <- c
//    fmt.Println("Wait")
    

//    c := make(chan int)
//    go func () {
//        time.Sleep(time.Second)
//        c <- 1
//        c <- 2
//        c <- 3        
//        close(c)
//    }()
//    fmt.Println(len(c)) //Why the Length of Channel is 0 ? 

//    for v := range c {
////        for i := 0 ; i < len(c) ; i++) {
////            fmt.Println(c[i])
//        fmt.Println(v)        
//    }


//    c := make(chan int, 100)
//    c <- 1
//    c <- 2
//    c <- 3  
//    fmt.Println(len(c))

//	Sarray := []int{1, 5, -2, 7, 2, 8, -9, 4, 0}
//	//c := make(chan int)
//    
////    fmt.Println("Length => ", len(Sarray)/3)
//
//
//    //var SumArray [len(Sarray)/3] int
//    var SumArray [3] int
//    var SumArrayIndex int = 0 
//    
//    for i := 0 ; i < len(Sarray)-2 ; i+=3 {
//        SumArray[SumArrayIndex] = Sarray[i] + Sarray[i+1] + Sarray[i+2] ; SumArrayIndex++
//        var sliceArray [] int = Sarray[i:i+3]
//        fmt.Println(sliceArray)
//    }
//    
//    fmt.Println("SumArray:")
//    var tot_sum int = 0
//    for i := 0 ; i < len(SumArray) ; i++ {
//        fmt.Println(SumArray[i])
//        tot_sum += SumArray[i]
//    }
//    fmt.Println("total Sum : ", tot_sum)
    
	// sum `s` slice by taking three elements each time
	// print sum of each time and the total result

}