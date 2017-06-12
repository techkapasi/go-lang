package main

import ("fmt")

func fibonacci(out, end chan int) {
	x, y := 0, 1
	for {
		select {
            case out <- x:
                x, y = y, x+y
            case <-end:
                return
		}
	}
}


func main () {
    fmt.Println("")
    fmt.Println("Concurrent - fiboanacci:")
    fmt.Println("------------------------")

    out := make(chan int)
    end := make(chan int)    
    var last_element int    
    var series_range int
    
    fmt.Print("Enter Series Range : ")
    fmt.Scanln(&series_range)
    
    go func() {
        for i := 0; i < series_range; i++ {
            if i >= (series_range - 1) {
                last_element = <-out
                fmt.Println(last_element)               
            } else {
                fmt.Println(<-out)
            }
		}
        end <- 0
	}()
    fibonacci(out, end)
    fmt.Println("Series Last Element : ", last_element)
}